package db

import (
	"context"
	"encoding/json"
	"fmt"
	"gopbl-3/modelo"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	TopicPrefix          = "sincronizacao/postos/"
	TopicPostoAtualizado = TopicPrefix + "atualizado"
	TopicPostoRemovido   = TopicPrefix + "removido"
	TopicSolicitacao     = TopicPrefix + "solicitacao"
)

// SincronizadorMQTT gerencia a sincronização via MQTT entre diferentes bases de dados
type SincronizadorMQTT struct {
	Conexao       *ConexaoServidorDB
	Cliente       mqtt.Client
	ServidorID    string
	IntervaloSync time.Duration
}

// NovoSincronizadorMQTT cria um novo sincronizador baseado em MQTT
func NovoSincronizadorMQTT(conexao *ConexaoServidorDB, brokerURL, servidorID string, intervalo time.Duration) (*SincronizadorMQTT, error) {
	// Configurar opções do cliente MQTT
	opts := mqtt.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID(fmt.Sprintf("sincronizador-%s", servidorID)).
		SetKeepAlive(60 * time.Second).
		SetPingTimeout(1 * time.Second).
		SetCleanSession(false).
		SetAutoReconnect(true).
		SetMaxReconnectInterval(30 * time.Second)

	// Criar cliente MQTT
	cliente := mqtt.NewClient(opts)
	token := cliente.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("erro ao conectar ao broker MQTT: %v", token.Error())
	}

	// Criar sincronizador
	sincronizador := &SincronizadorMQTT{
		Conexao:       conexao,
		Cliente:       cliente,
		ServidorID:    servidorID,
		IntervaloSync: intervalo,
	}

	// Configurar handlers para os tópicos
	sincronizador.configurarHandlers()

	return sincronizador, nil
}

// configurarHandlers configura os callbacks para os tópicos MQTT
func (s *SincronizadorMQTT) configurarHandlers() {
	// Handler para postos atualizados
	s.Cliente.Subscribe(TopicPostoAtualizado, 1, func(client mqtt.Client, msg mqtt.Message) {
		var posto modelo.Posto
		if err := json.Unmarshal(msg.Payload(), &posto); err != nil {
			log.Printf("Erro ao decodificar posto atualizado: %v", err)
			return
		}
		s.atualizarPostoLocal(&posto)
	})

	// Handler para solicitações de sincronização completa
	s.Cliente.Subscribe(TopicSolicitacao, 1, func(client mqtt.Client, msg mqtt.Message) {
		var solicitacao struct {
			ServidorID string `json:"servidorID"`
		}
		if err := json.Unmarshal(msg.Payload(), &solicitacao); err != nil {
			log.Printf("Erro ao decodificar solicitação: %v", err)
			return
		}

		// Evitar responder à própria solicitação
		if solicitacao.ServidorID != s.ServidorID {
			s.enviarTodosPostos()
		}
	})
}

// IniciarSincronizacao inicia o processo de sincronização periódica via MQTT
func (s *SincronizadorMQTT) IniciarSincronizacao() {
	// Solicitar sincronização imediata
	s.SolicitarSincronizacao()

	// Configurar timer para sincronização periódica
	go func() {
		ticker := time.NewTicker(s.IntervaloSync)
		defer ticker.Stop()

		for range ticker.C {
			s.SolicitarSincronizacao()
		}
	}()
}

// SolicitarSincronizacao solicita a todos os servidores que enviem seus postos
func (s *SincronizadorMQTT) SolicitarSincronizacao() {
	solicitacao := struct {
		ServidorID string `json:"servidorID"`
		Timestamp  int64  `json:"timestamp"`
	}{
		ServidorID: s.ServidorID,
		Timestamp:  time.Now().UnixNano(),
	}

	// Serializar a solicitação
	payload, err := json.Marshal(solicitacao)
	if err != nil {
		log.Printf("Erro ao serializar solicitação: %v", err)
		return
	}

	// Publicar no tópico de solicitação
	token := s.Cliente.Publish(TopicSolicitacao, 1, false, payload)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Erro ao publicar solicitação: %v", token.Error())
	}
}

// NotificarAlteracao envia atualizações de um posto via MQTT
func (s *SincronizadorMQTT) NotificarAlteracao(posto *modelo.Posto) {
	// Serializar o posto
	payload, err := json.Marshal(posto)
	if err != nil {
		log.Printf("Erro ao serializar posto: %v", err)
		return
	}

	// Publicar no tópico de atualização
	token := s.Cliente.Publish(TopicPostoAtualizado, 1, false, payload)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Erro ao publicar atualização: %v", token.Error())
	}
}

// atualizarPostoLocal atualiza um posto no banco de dados local
func (s *SincronizadorMQTT) atualizarPostoLocal(posto *modelo.Posto) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Definir o filtro para buscar o posto
	filtro := bson.M{"id": posto.ID}

	// Obter o posto existente, se houver
	var postoExistente modelo.Posto
	err := s.Conexao.PostosCollection.FindOne(ctx, filtro).Decode(&postoExistente)

	// Se o posto não existe ou se a atualização é mais recente
	if err != nil || posto.UltimaAtualizacao.After(postoExistente.UltimaAtualizacao) {
		// Definir opções (upsert)
		opcoes := options.Update().SetUpsert(true)

		// Definir atualização
		update := bson.M{"$set": posto}

		// Atualizar posto
		_, err := s.Conexao.PostosCollection.UpdateOne(ctx, filtro, update, opcoes)
		if err != nil {
			log.Printf("Erro ao atualizar posto local: %v", err)
		} else {
			log.Printf("Posto %s atualizado localmente", posto.ID)
		}
	}
}

// enviarTodosPostos envia todos os postos locais para sincronização
func (s *SincronizadorMQTT) enviarTodosPostos() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Buscar todos os postos
	cursor, err := s.Conexao.PostosCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Erro ao buscar postos locais: %v", err)
		return
	}
	defer cursor.Close(ctx)

	// Processar cada posto
	var postos []modelo.Posto
	if err := cursor.All(ctx, &postos); err != nil {
		log.Printf("Erro ao decodificar postos locais: %v", err)
		return
	}

	// Enviar cada posto
	for i := range postos {
		s.NotificarAlteracao(&postos[i])
		// Pequeno delay para evitar sobrecarregar o broker
		time.Sleep(50 * time.Millisecond)
	}
}

// Fechar encerra a conexão com o broker MQTT
func (s *SincronizadorMQTT) Fechar() {
	if s.Cliente.IsConnected() {
		s.Cliente.Unsubscribe(TopicPostoAtualizado, TopicSolicitacao)
		s.Cliente.Disconnect(250)
	}
}
