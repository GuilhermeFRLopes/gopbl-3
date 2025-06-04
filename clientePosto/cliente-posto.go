package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log" // Adicionado para logs
	"net/http"
	"os"
	"os/signal"
	"strconv" // Para strconv.ParseFloat
	"time"
	//"gopbl-3/modelo"
)

type Posto struct {
	ID             string  `json:"id"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Ocupado        bool    `json:"ocupado"` // Corresponde ao contrato Solidity
	ServidorOrigem string  `json:"servidorOrigem"`
	Cidade         string  `json:"cidade"`
	Reservador     string  `json:"reservador"`
}

var posto_criado Posto
var servidorBackendURL string = "http://172.17.0.1:8083"

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Goroutine para tratar o sinal de interrupção (Ctrl+C)
	go func() {
		<-c
		fmt.Println("\nRecebido sinal de interrupção. Encerrando cliente-posto...")

		if posto_criado.ID != "" && posto_criado.ServidorOrigem != "" {
			//logica pra apagar o posto
			fmt.Println("O posto cadastrado precisaria ser deletado da blockchain. Considere implementar um endpoint DELETE no servidor para isso.")
		}

		fmt.Println("Cliente-posto desconectado.")
		os.Exit(0)
	}()

	// Gerar dados do posto (a partir de variáveis de ambiente ou input manual)
	//gerarPostoDados()
	//input manual:
	cadastrarPostoManual()

	fmt.Printf("Dados do posto configurados: ID=%s, Cidade=%s, Servidor=%s\n", posto_criado.ID, posto_criado.Cidade, posto_criado.ServidorOrigem)

	// Enviar solicitação de cadastro
	enviarCadastroPosto()

	// O cliente permanecerá ativo para exibir a resposta.
	// Pressione Ctrl+C para sair.
	select {} // Mantém o programa rodando até o sinal de interrupção
}

// gerarPostoDados lê as informações do posto a partir de variáveis de ambiente
func gerarPostoDados() {
	idEnv := os.Getenv("POSTO_ID")
	if idEnv == "" {
		log.Fatal("Variável de ambiente POSTO_ID não definida. Ex: set POSTO_ID=posto-FSA-001")
	}

	latEnvStr := os.Getenv("POSTO_LAT")
	latEnv, err := strconv.ParseFloat(latEnvStr, 64)
	if err != nil {
		log.Fatalf("Variável de ambiente POSTO_LAT inválida: %v", err)
	}

	longEnvStr := os.Getenv("POSTO_LONG")
	longEnv, err := strconv.ParseFloat(longEnvStr, 64)
	if err != nil {
		log.Fatalf("Variável de ambiente POSTO_LONG inválida: %v", err)
	}

	servidorEnv := os.Getenv("POSTO_SERVIDOR")
	if servidorEnv == "" {
		log.Fatal("Variável de ambiente POSTO_SERVIDOR não definida. Ex: set POSTO_SERVIDOR=Ipiranga")
	}

	cidadeEnv := os.Getenv("POSTO_CIDADE")
	if cidadeEnv == "" {
		log.Fatal("Variável de ambiente POSTO_CIDADE não definida. Ex: set POSTO_CIDADE=\"Feira de Santana\"")
	}

	// Cria a struct Posto, que deve corresponder à struct Posto no seu main.go
	posto_criado = Posto{
		ID:             idEnv,
		Latitude:       latEnv,
		Longitude:      longEnv,
		ServidorOrigem: servidorEnv,
		Cidade:         cidadeEnv,
		Ocupado:        false, // Contrato inicializa como falso (disponível)
		Reservador:     "",    // Contrato inicializa como vazio
	}
}

// cadastrarPostoManual permite input manual dos dados do posto (opcional)
func cadastrarPostoManual() {
	var id string
	var lat, long float64
	var servidor, cidade string

	fmt.Println("--- Cadastro de Posto de Recarga ---")
	fmt.Print("Digite o ID do posto (ex: posto-FSA-001): ")
	fmt.Scanln(&id)
	fmt.Print("Digite a latitude (ex: -12.23): ")
	fmt.Scanln(&lat)
	fmt.Print("Digite a longitude (ex: -38.96): ")
	fmt.Scanln(&long)

	// Seleção de servidor e cidade (mantendo a lógica original)
	fmt.Println("\nSelecione o servidor:")
	fmt.Println("1 - Ipiranga")
	fmt.Println("2 - 22")
	fmt.Println("3 - Shell")
	var opcaoServidor int
	fmt.Print("Escolha uma opção: ")
	fmt.Scanln(&opcaoServidor)
	switch opcaoServidor {
	case 1:
		servidor = "Ipiranga"
	case 2:
		servidor = "22"
	case 3:
		servidor = "Shell"
	default:
		fmt.Println("Opção inválida. Utilizando Ipiranga como padrão.")
		servidor = "Ipiranga"
	}

	fmt.Println("\nInforme a cidade do posto:")
	fmt.Println("1 - Feira de Santana")
	fmt.Println("2 - São Gonçalo")
	fmt.Println("3 - Serrinha")
	// Adicione mais cidades conforme a sua lógica de switches para outros servidores
	// Simplificado para este exemplo:
	var opcaoCidade int
	fmt.Print("Escolha uma opção: ")
	fmt.Scanln(&opcaoCidade)
	switch opcaoCidade {
	case 1:
		cidade = "Feira de Santana"
	case 2:
		cidade = "São Gonçalo"
	case 3:
		cidade = "Serrinha"
	default:
		fmt.Println("Opção inválida. Cidade padrão: Feira de Santana")
		cidade = "Feira de Santana"
	}

	posto_criado = Posto{
		ID:             id,
		Latitude:       lat,
		Longitude:      long,
		ServidorOrigem: servidor,
		Cidade:         cidade,
		Ocupado:        false,
		Reservador:     "",
	}
}

// enviarCadastroPosto envia a requisição HTTP POST para o servidor
func enviarCadastroPosto() {
	payload, err := json.Marshal(posto_criado)
	if err != nil {
		log.Fatalf("Erro ao codificar JSON do posto: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, servidorBackendURL+"/cadastrar", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalf("Erro ao criar requisição HTTP POST: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second} // Timeout para operações de blockchain
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erro ao enviar requisição HTTP POST: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Erro ao decodificar resposta do servidor: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Falha ao cadastrar posto (Status: %d): %s", resp.StatusCode, result["error"])
	}

	fmt.Printf("Posto '%s' cadastrado com sucesso! Mensagem do servidor: %s (Tx Hash: %s)\n",
		posto_criado.ID, result["message"], result["txHash"])
}
