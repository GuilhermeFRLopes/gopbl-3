package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	blockchain_contracts "gopbl-3/contratos"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	//"os"
)

// Modelo de dados para o posto
type Posto struct {
	ID             string `json:"id"`
	Latitude       int64  `json:"latitude"`
	Longitude      int64  `json:"longitude"`
	Ocupado        bool   `json:"ocupado"`
	ServidorOrigem string `json:"servidorOrigem"`
	Cidade         string `json:"cidade"`
	Reservador     string `json:"reservador"`
	ValorPagamento uint64 `json:"valorPagamento"`
}

// Modelo de dados para a requisição de reserva
type ReservaData struct {
	IDPostos       []string `json:"idPostos"`
	Reservar       bool     `json:"reservar"`
	ClientID       string   `json:"clientId"` // O ID do veículo cliente (será o reservador no contrato)
	ValorPagamento uint64   `json:"valorPagamento"`
}

// Variáveis para interação com Ethereum
var ethClient *ethclient.Client

var contractInstance *blockchain_contracts.Contratos
var contractAddress common.Address
var privateKey *ecdsa.PrivateKey
var fromAddress common.Address

func main() {
	if err := conectarEthereum(); err != nil {
		log.Fatalf("Erro fatal ao conectar ao Ethereum: %v", err)
	}
	fmt.Println("Conectado ao Ethereum com sucesso!")

	//numero do contrato
	contractAddress = common.HexToAddress("Ef8F65612B109E927f23fc654f8EF3FE8601DC72")
	// Se usar variáveis de ambiente
	// contractAddress = common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	var err error
	contractInstance, err = blockchain_contracts.NewContratos(contractAddress, ethClient)
	if err != nil {
		log.Fatalf("Falha ao instanciar o contrato: %v", err)
	}
	fmt.Printf("Instância do contrato 'Contrato' carregada do endereço: %s\n", contractAddress.Hex())

	//Configurar Servidor Gin HTTP
	router := gin.Default()
	router.GET("/postosDisponiveis", postosDisponiveisHandler)
	router.POST("/cadastrar", cadastrarPostoHandler)
	router.PUT("/reservar", editarPostoHandler)

	//Iniciar Servidor HTTP
	fmt.Println("Servidor Gin HTTP iniciado na porta :8083")
	router.Run(":8083")
}

func conectarEthereum() error {
	var err error
	// Conecta ao Ganache local
	ethClient, err = ethclient.Dial("HTTP://127.0.0.1:7545") //endereço do ganache
	if err != nil {
		return fmt.Errorf("falha ao conectar ao Ethereum (Ganache): %v", err)
	}

	// Carregar a chave privada da conta do Ganache
	// Lembre-se de que essa chave deve ser de uma das contas pré-financiadas do SEU Ganache
	//tem que mudar a chave toda hora que abrir o ganache
	privateKeyHex := "4d158b3c2a6dbeff6fa06b8e4c8c3f4125f41fdfc7601b5459bdad77f9c0b677"
	privateKey, err = crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("falha ao carregar chave privada: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("não é uma chave pública ECDSA")
	}
	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("Usando endereço Ethereum para transações: %s\n", fromAddress.Hex())

	balance, err := ethClient.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Printf("Aviso: Falha ao obter saldo da conta: %v", err)
	} else {
		fmt.Printf("Saldo da conta %s: %s Wei\n", fromAddress.Hex(), balance.String())
	}

	return nil
}

// getAuthTransactor prepara as opções para enviar uma transação (nonce, gas price, gas limit)
func getAuthTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := ethClient.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter nonce: %v", err)
	}

	gasPrice, err := ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter preço do gás: %v", err)
	}

	// ChainID para o Ganache é 1337
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // <-- Mantido 1337 para Ganache
	if err != nil {
		return nil, fmt.Errorf("falha ao criar transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(6000000) // Ajustado para 6 milhões para Ganache (limite de bloco)
	auth.GasPrice = gasPrice

	return auth, nil
}

// postosDisponiveisHandler consulta os postos diretamente da blockchain
func postosDisponiveisHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Obter todos os IDs de postos do contrato Ethereum
	//chama a funçao GetAllPostos do contrato, ela retorna a lista com todos os postos
	idsPostos, err := contractInstance.GetAllPostos(&bind.CallOpts{Context: ctx})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar IDs de postos do contrato: %v", err)})
		return
	}

	var todosPostos []Posto // Usando a struct Posto definida neste arquivo
	for _, id := range idsPostos {

		stationData, err := contractInstance.PostosRecarga(&bind.CallOpts{Context: ctx}, id)
		if err != nil {
			log.Printf("Aviso: Erro ao obter detalhes completos do posto %s: %v", id, err)
			continue
		}

		posto := Posto{
			ID:             stationData.Id,
			Latitude:       stationData.Latitude.Int64(),
			Longitude:      stationData.Longitude.Int64(),
			Ocupado:        stationData.Ocupado,
			ServidorOrigem: stationData.ServidorOrigem,
			Cidade:         stationData.Cidade,
			Reservador:     stationData.Reservador,
			ValorPagamento: stationData.ValorPagamento.Uint64(),
		}
		todosPostos = append(todosPostos, posto)
	}

	c.JSON(http.StatusOK, todosPostos)
}

// cadastrarPostoHandler registra um novo posto na blockchain (ADAPTADO)
func cadastrarPostoHandler(c *gin.Context) {
	var novoPosto Posto
	if erro := c.ShouldBindJSON(&novoPosto); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON de posto inválido"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	auth, err := getAuthTransactor(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao preparar transação Ethereum: %v", err)})
		return
	}

	// Chamar a nova função 'registrarPosto' com os parâmetros corretos
	tx, err := contractInstance.RegistrarPosto(auth, novoPosto.ID, big.NewInt(novoPosto.Latitude), big.NewInt(novoPosto.Longitude), novoPosto.ServidorOrigem, novoPosto.Cidade) // <-- ADAPTADO PARA 'RegistrarPosto'
	if err != nil {
		if strings.Contains(err.Error(), "Posto de recarga ja cadastrado") { // <-- Mensagem de erro do NOVO contrato
			c.JSON(http.StatusConflict, gin.H{"error": "Posto já cadastrado na blockchain."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao chamar contrato para registrar posto: %v", err)})
		return
	}

	log.Printf("Transação de registro de posto enviada. Hash: %s\n", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, ethClient, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao aguardar confirmação da transação de registro: %v", err)})
		return
	}
	if receipt.Status != 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Transação de registro falhou na blockchain. Status: %d", receipt.Status)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Posto cadastrado com sucesso na blockchain!", "txHash": tx.Hash().Hex()})
}

// editarPostoHandler para reservar/liberar postos na blockchain
func editarPostoHandler(c *gin.Context) {
	var data ReservaData
	if erro := c.ShouldBindJSON(&data); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON de reserva inválido."})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	reservadorString := data.ClientID

	auth, err := getAuthTransactor(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao preparar transação: %v", err)})
		return
	}

	var txHash string
	if data.Reservar {
		// Lógica para reservar postos
		valorPagamentoBigInt := new(big.Int).SetUint64(data.ValorPagamento)
		tx, err := contractInstance.ReservarVagaPostos(auth, data.IDPostos, true, reservadorString, valorPagamentoBigInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao reservar postos: %v", err)})
			return
		}
		txHash = tx.Hash().Hex()
		log.Printf("Transação de RESERVA enviada. Hash: %s\n", txHash)

	} else {
		// Lógica para liberar postos
		tx, err := contractInstance.ReservarVagaPostos(auth, data.IDPostos, false, reservadorString, big.NewInt(0))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao liberar postos: %v", err)})
			return
		}
		txHash = tx.Hash().Hex()
		log.Printf("Transação de LIBERAÇÃO enviada. Hash: %s\n", txHash)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operação de %s enviada com sucesso.", func() string {
			if data.Reservar {
				return "reserva"
			}
			return "liberação"
		}()),
		"txHash": txHash,
	})
}