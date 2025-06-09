package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	// Adicione esta importação se não estiver lá
	abi_contracts "gopbl-3/contratos" // Ajuste para o caminho real do arquivo Go gerado pelo abigen

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// --- CONFIGURAÇÕES ---
	// privateKeyHex agora lê da variável de ambiente "ETHEREUM_PRIVATE_KEY"
	privateKeyHex := os.Getenv("ETHEREUM_PRIVATE_KEY")
	if privateKeyHex == "" {
		// Se a variável de ambiente não estiver definida, use uma chave privada de uma conta do Ganache diretamente AQUI para teste
		// EXTREMAMENTE IMPORTANTE: NUNCA FAÇA ISSO COM CHAVES DE PRODUÇÃO.
		// Para teste local, é aceitável para conveniência.
		privateKeyHex = "4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d" // Ex: 4f3edf983ad256c709e6a10f1723ad2eeaec49c9fda382894563704043b075ff
		log.Println("Usando chave privada diretamente no código para teste local. Use variáveis de ambiente em produção!")
	}

	// infuraURL agora lê da variável de ambiente "INFURA_SEPOLIA_URL"
	infuraURL := os.Getenv("INFURA_SEPOLIA_URL") // <-- CORRIGIDO AQUI!
	if infuraURL == "" {
		infuraURL = "HTTP://127.0.0.1:8545"
		//log.Fatal("Variável de ambiente INFURA_SEPOLIA_URL não configurada. Por favor, defina-a no terminal antes de executar (ex: set INFURA_SEPOLIA_URL=\"https://...\").")
	}

	chainID := big.NewInt(1337) // Chain ID padrão do Ganache

	// --- CONEXÃO ETHEREUM ---
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao Infura: %v", err)
	}
	defer client.Close()
	fmt.Println("Conectado ao Infura (Sepolia)!")

	// --- CARREGAR CHAVE PRIVADA ---
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Falha ao carregar chave privada. Verifique se ETHEREUM_PRIVATE_KEY está correta e não inclui o '0x': %v", err) // Adicionado dica
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Não é uma chave pública ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("Endereço da conta de implantação: %s\n", fromAddress.Hex()) // LINHA IMPORTANTE AQUI!

	// Verificar saldo
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Printf("Aviso: Falha ao obter saldo da conta: %v", err)
	} else {
		fmt.Printf("Saldo da conta: %s ETH (para gás)\n", new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18)))
		if balance.Cmp(big.NewInt(0)) == 0 {
			log.Fatal("Sua conta não tem ETH na rede Sepolia. Obtenha ETH de teste de um Sepolia Faucet.")
		}
	}

	// --- PREPARAR TRANSACTOR ---
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Falha ao obter nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Falha ao obter preço do gás: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Falha ao criar transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // Não enviando Ether junto com a implantação
	auth.GasLimit = uint64(6000000) // Limite de gás para implantação (pode precisar ajustar)
	auth.GasPrice = gasPrice

	fmt.Println("Preparando implantação do contrato...")

	// --- IMPLANTAR CONTRATO ---
	// O construtor do ChargingStationManager não aceita argumentos.
	address, tx, _, err := abi_contracts.DeployContratos(auth, client) // Se o nome do contrato for diferente, substitua "DeployChargingStationManager" pelo nome correto.

	if err != nil {
		log.Fatalf("Falha ao implantar contrato: %v", err)
	}

	fmt.Printf("Transação de implantação enviada. Hash: %s\n", tx.Hash().Hex())
	fmt.Printf("Aguardando confirmação do contrato... (pode levar alguns segundos)\n")

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Erro ao aguardar a mineração da transação de implantação: %v", err)
	}

	if receipt.Status != 1 {
		log.Fatalf("Implantação do contrato falhou na blockchain. Status: %d\n", receipt.Status)
	}

	fmt.Printf("Contrato ChargingStationManager implantado com sucesso em: %s\n", address.Hex())
	fmt.Printf("Hash da transação de implantação: %s\n", tx.Hash().Hex())

	// Opcional: Salvar o endereço do contrato em um arquivo para uso posterior
	err = ioutil.WriteFile("contract_address.txt", []byte(address.Hex()), 0644)
	if err != nil {
		log.Printf("Aviso: Falha ao salvar o endereço do contrato em arquivo: %v", err)
	}
}

// Pequena função auxiliar para obter variáveis de ambiente (esta função não será usada diretamente no main agora, mas é um bom exemplo)
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
