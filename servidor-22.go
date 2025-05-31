package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	// Importar o wrapper do contrato gerado pelo abigen
	blockchain_contracts "gopbl-3/contracts" // Ajuste para o seu caminho real

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// Modelo de dados simplificado para o posto (semelhante ao seu modelo.Posto)
type Posto struct {
	ID                 string    `json:"id"`
	Nome               string    `json:"nome"`
	Cidade             string    `json:"cidade"`
	Latitude           float64   `json:"latitude"`
	Longitude          float64   `json:"longitude"`
	Disponivel         bool      `json:"disponivel"`
	UltimaAtualizacao  time.Time `json:"ultimaAtualizacao"`
	ServidorOrigem     string    `json:"servidorOrigem"` // Pode ser o endereço do contrato ou do servidor que registrou
	VeiculoReservador string    `json:"veiculoReservador"` // Endereço ETH do veículo que reservou
}

// Modelo de dados para a requisição de reserva
type ReservaData struct {
	IDPostos []string `json:"idPostos"`
	Reservar bool     `json:"reservar"`
	ClientID string   `json:"clientId"` // O ID do veículo cliente
}

// Variáveis para interação com Ethereum
var ethClient *ethclient.Client
var contractInstance *blockchain_contracts.ChargingStationManager // Instância do seu contrato
var contractAddress common.Address                              // Endereço do contrato implantado
var privateKey *ecdsa.PrivateKey                                // Chave privada para assinar transações
var fromAddress common.Address                                  // Endereço da conta que envia transações

func main() {
	// 1. Configurar Conexão Ethereum
	if err := conectarEthereum(); err != nil {
		log.Fatalf("Erro fatal ao conectar ao Ethereum: %v", err)
	}
	fmt.Println("Conectado ao Ethereum com sucesso!")

	// 2. Carregar o Contrato Inteligente
	// Em um ambiente real, você obteria este endereço de uma variável de ambiente ou configuração.
	// Para este exemplo, use um endereço de contrato que você implantou manualmente
	// na sua rede Geth --dev.
	// Ex: Execute o comando de deploy uma vez e cole o endereço aqui.
	// Exemplo de deploy manual:
	// geth attach http://127.0.0.1:8545
	// var abi = <conteúdo do ChargingStationManager.abi>
	// var bytecode = <conteúdo do ChargingStationManager.bin>
	// var contract = web3.eth.contract(abi);
	// var contractData = contract.new.getData(bytecode, {data: bytecode});
	// var deployedContract = contract.new({from: web3.eth.accounts[0], data: contractData, gas: '4700000'});
	// console.log(deployedContract.address); // Copie este endereço!
	// Substitua pelo endereço real do seu contrato implantado!
	// ATENÇÃO: SUBSTITUA PELO ENDEREÇO REAL DO SEU CONTRATO DEPLOYADO!
	contractAddress = common.HexToAddress("0x762f01C9503923d142A9f697523f6e91122a6136") 
	
	var err error
	contractInstance, err = blockchain_contracts.NewChargingStationManager(contractAddress, ethClient)
	if err != nil {
		log.Fatalf("Falha ao instanciar o contrato: %v", err)
	}
	fmt.Printf("Instância do contrato ChargingStationManager carregada do endereço: %s\n", contractAddress.Hex())

	// 3. Configurar Servidor Gin HTTP
	router := gin.Default()
	router.GET("/postosDisponiveis", postosDisponiveisHandler)
	router.POST("/cadastrar", cadastrarPostoHandler)
	router.PUT("/reservar", editarPostoHandler) // Agora irá interagir com o contrato

	// 4. Iniciar Servidor HTTP
	fmt.Println("Servidor Gin HTTP iniciado na porta :8083")
	router.Run(":8083")
}

func conectarEthereum() error {
	var err error
	// Conecta ao seu nó Geth. Se estiver usando --dev, será 127.0.0.1:8545
	// Se for uma rede de teste como Sepolia via Infura/Alchemy, use o URL fornecido.
	ethClient, err = ethclient.Dial("https://sepolia.infura.io/v3/8ce7834b4e0f4f9c8eb42d2012df9ec3")
	if err != nil {
		return fmt.Errorf("falha ao conectar ao Ethereum: %v", err)
	}

	// Carregar a chave privada da conta que irá assinar as transações.
	// ESTE É UM PONTO CRÍTICO DE SEGURANÇA. Em produção, use um KMS (Key Management System)
	// ou outras soluções seguras. Não hardcode chaves privadas!
	// Esta é a chave privada da primeira conta do Geth em modo --dev (0x70997970C51812dc3A010837D6298BD35Cf24626)
	privateKeyHex := "59c6995e998f97a5a004496608237ee6b68eef6727284b578e8dc038161f0340"
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

	// Opcional: Verificar saldo da conta
	balance, err := ethClient.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Printf("Aviso: Falha ao obter saldo da conta: %v", err)
	} else {
		fmt.Printf("Saldo da conta %s: %s Wei\n", fromAddress.Hex(), balance.String())
		// Lembre-se que você precisa de ETH para pagar as taxas de gás!
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

	// ChainID para redes de desenvolvimento Geth geralmente é 1337
	// Para Mainnet, Sepolia, etc., use o ChainID correto (e.g., 1 para Mainnet, 11155111 para Sepolia)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return nil, fmt.Errorf("falha ao criar transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // Não enviando Ether junto com a chamada do contrato
	auth.GasLimit = uint64(5000000) // Limite de gás (ajuste se suas operações forem mais complexas)
	auth.GasPrice = gasPrice

	return auth, nil
}

// postosDisponiveisHandler consulta os postos diretamente da blockchain
func postosDisponiveisHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Obter todos os IDs de postos do contrato Ethereum
	stationIds, err := contractInstance.GetAllStationIds(&bind.CallOpts{Context: ctx})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao buscar IDs de postos do contrato: %v", err)})
		return
	}

	var todosPostos []Posto // Usando a struct Posto definida neste arquivo
	for _, id := range stationIds {
		// Para cada ID, consulte o estado na blockchain
		disponivel, reservador, ultimaAtualizacaoUnix, err := contractInstance.GetStationAvailability(&bind.CallOpts{Context: ctx}, id)
		if err != nil {
			log.Printf("Aviso: Erro ao obter detalhes do posto %s do contrato: %v", id, err)
			continue // Pular para o próximo posto se houver erro
		}

		// Obter nome e cidade (se não estiverem no contrato, você precisaria de um cache local ou metadados)
		// Para este exemplo, vamos simplificar e assumir que são strings vazias ou derivá-las do ID se não houver um sistema de metadados.
		// Em um sistema real, você talvez tenha um mapeamento ID -> Detalhes no seu backend ou outro contrato.
		// Ou, idealmente, você incluiria nome e cidade no struct ChargingStation do Solidity, como já fizemos.
		
		// Recupere os detalhes completos do posto do contrato
		// Note: A função 'chargingStations' no contrato é pública e gera um getter,
		// mas retorna a struct completa, o que é mais eficiente do que chamar várias 'getters'
		stationData, err := contractInstance.ChargingStations(&bind.CallOpts{Context: ctx}, id)
		if err != nil {
			log.Printf("Aviso: Erro ao obter detalhes completos do posto %s: %v", id, err)
			continue
		}


		posto := Posto{
			ID:                 stationData.Id,
			Nome:               stationData.Nome,
			Cidade:             stationData.Cidade,
			Disponivel:         disponivel,
			UltimaAtualizacao:  time.Unix(int64(ultimaAtualizacaoUnix.Uint64()), 0),
			VeiculoReservador:  reservador.Hex(),
			ServidorOrigem:     "blockchain", // Indica que a info veio da blockchain
		}
		todosPostos = append(todosPostos, posto)
	}

	c.JSON(http.StatusOK, todosPostos)
}

// cadastrarPostoHandler registra um novo posto na blockchain
func cadastrarPostoHandler(c *gin.Context) {
	var novoPosto Posto
	if erro := c.ShouldBindJSON(&novoPosto); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON de posto inválido"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Tempo para a transação ser minerada
	defer cancel()

	// 1. Preparar opções para a transação
	auth, err := getAuthTransactor(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao preparar transação Ethereum: %v", err)})
		return
	}

	// 2. Chamar a função `registerStation` do contrato inteligente
	tx, err := contractInstance.RegisterStation(auth, novoPosto.ID, novoPosto.Nome, novoPosto.Cidade)
	if err != nil {
		// O erro pode ser de "station already registered" se o contrato exigir isso
		if strings.Contains(err.Error(), "Charging station already registered.") {
            c.JSON(http.StatusConflict, gin.H{"error": "Posto já cadastrado na blockchain."})
            return
        }
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao chamar contrato para registrar posto: %v", err)})
		return
	}

	log.Printf("Transação de registro de posto enviada. Hash: %s\n", tx.Hash().Hex())

	// 3. Aguardar a confirmação da transação na blockchain
	receipt, err := bind.WaitMined(ctx, ethClient, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao aguardar confirmação da transação de registro: %v", err)})
		return
	}
	if receipt.Status != 1 { // 1 significa sucesso
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Transação de registro falhou na blockchain. Status: %d", receipt.Status)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Posto cadastrado com sucesso na blockchain!", "txHash": tx.Hash().Hex()})
}

// editarPostoHandler para reservar/liberar postos na blockchain
func editarPostoHandler(c *gin.Context) {
	var data ReservaData // Assume que ReservaData tem IDPostos []string e Reservar bool
	if erro := c.ShouldBindJSON(&data); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON de reserva inválido."})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second) // Mais tempo para a transação
	defer cancel()

	// Simulação do endereço Ethereum do veículo (idealmente viria do cliente via assinatura)
	// Para este exemplo, vou gerar um endereço a partir do ClientID (não é um endereço real, mas serve para o contrato)
	// Em um sistema real, o cliente teria uma carteira Ethereum e passaria seu endereço, ou o servidor faria isso por ele.
	veiculoEthAddress := common.HexToAddress(crypto.Keccak256Hash([]byte(data.ClientID)).Hex()[len(crypto.Keccak256Hash([]byte(data.ClientID)).Hex())-40:])


	// 1. Preparar opções para a transação
	auth, err := getAuthTransactor(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao preparar transação Ethereum: %v", err)})
		return
	}

	// 2. Chamar a função `updateStationAvailability` no contrato Ethereum
	tx, err := contractInstance.UpdateStationAvailability(auth, data.IDPostos, data.Reservar, veiculoEthAddress)
	if err != nil {
		// Tratamento de erros específicos do contrato
		if strings.Contains(err.Error(), "Station is already reserved.") {
            c.JSON(http.StatusConflict, gin.H{"error": "Um ou mais postos já estão reservados."})
            return
        }
        if strings.Contains(err.Error(), "Station is already available.") {
            c.JSON(http.StatusConflict, gin.H{"error": "Um ou mais postos já estão disponíveis para liberação."})
            return
        }
        if strings.Contains(err.Error(), "Station does not exist.") {
            c.JSON(http.StatusNotFound, gin.H{"error": "Um ou mais postos não foram encontrados na blockchain."})
            return
        }
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao chamar contrato para %s posto(s): %v", func() string {
			if data.Reservar { return "reservar" } else { return "liberar" }
		}(), err)})
		return
	}

	log.Printf("Transação de %s posto(s) enviada. Hash: %s\n", func() string {
		if data.Reservar { return "reserva" } else { return "liberação" }
	}(), tx.Hash().Hex())

	// 3. Aguardar a confirmação da transação
	receipt, err := bind.WaitMined(ctx, ethClient, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao aguardar confirmação da transação de %s: %v", func() string {
			if data.Reservar { return "reserva" } else { return "liberação" }
		}(), err)})
		return
	}
	if receipt.Status != 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Transação de %s falhou na blockchain. Status: %d", func() string {
			if data.Reservar { return "reserva" } else { return "liberação" }
		}(), receipt.Status)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Postos %s com sucesso na blockchain!", func() string {
		if data.Reservar { return "reservados" } else { return "liberados" }
	}()), "txHash": tx.Hash().Hex()})
}

// Funções utilitárias (getEnv, etc.) não são mais necessárias para esta implementação direta.
// A lógica de MQTT e MongoDB foi removida.