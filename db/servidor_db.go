package db

import (
	"context"
	"fmt"
	"gopbl-3/modelo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConexaoServidorDB struct {
	Cliente          *mongo.Client
	PostosCollection *mongo.Collection
	Nome             string
}

func NovaConexaoDB(nomeServidor, hostDB string, porta int) (*ConexaoServidorDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // timeout para operações com o banco de dados
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%d", hostDB, porta) // URI de conexão específica para este servidor

	clienteOpts := options.Client().ApplyURI(uri)
	cliente, erro := mongo.Connect(ctx, clienteOpts)
	if erro != nil {
		return nil, fmt.Errorf("erro ao conectar ao MongoDB para %s: %v", nomeServidor, erro)
	}

	// testar a conexão
	if erro = cliente.Ping(ctx, nil); erro != nil {
		return nil, fmt.Errorf("erro ao verificar conexão com MongoDB para %s: %v", nomeServidor, erro)
	}

	// cada servidor tem seu database
	dbName := fmt.Sprintf("reservas_%s", nomeServidor)

	collection := cliente.Database(dbName).Collection("postos")
	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("erro ao limpar coleção 'postos': %v", err)
	}

	return &ConexaoServidorDB{
		Cliente:          cliente,
		PostosCollection: cliente.Database(dbName).Collection("postos"),
		Nome:             nomeServidor,
	}, nil
}

// encerra a conexão com o banco de dados
func (c *ConexaoServidorDB) Fechar() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Cliente.Disconnect(ctx)
}

// consulta a disponibilidade dos postos em todos os servidores
func (c *ConexaoServidorDB) ConsultarPostosEmTodosServidores(idsPostos []string) (map[string]bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// buscar postos locais
	filtro := bson.M{"id": bson.M{"$in": idsPostos}}
	cursor, err := c.PostosCollection.Find(ctx, filtro)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar postos locais: %v", err)
	}
	defer cursor.Close(ctx)

	disponibilidade := make(map[string]bool) // para armazenar a disponibilidade dos postos

	// processa postos locais
	var postos []modelo.Posto
	if err := cursor.All(ctx, &postos); err != nil {
		return nil, fmt.Errorf("erro ao decodificar postos locais: %v", err)
	}

	for _, posto := range postos {
		disponibilidade[posto.ID] = posto.Disponivel
	}

	// buscar postos em outros servidores
	filtroOutros := bson.M{
		"id":             bson.M{"$in": idsPostos},
		"servidorOrigem": bson.M{"$ne": c.Nome},
	}
	cursorOutros, err := c.PostosCollection.Find(ctx, filtroOutros)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar postos de outros servidores: %v", err)
	}
	defer cursorOutros.Close(ctx)

	// processa postos de outros servidores
	var postosOutros []modelo.Posto
	if err := cursorOutros.All(ctx, &postosOutros); err != nil {
		return nil, fmt.Errorf("erro ao decodificar postos de outros servidores: %v", err)
	}

	for _, posto := range postosOutros {
		disponibilidade[posto.ID] = posto.Disponivel
	}

	return disponibilidade, nil
}
