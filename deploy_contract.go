package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings" // Adicione esta importação se não estiver lá

	blockchain_contracts "gopbl-3/contracts" // Ajuste para o seu caminho real

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// --- CONFIGURAÇÕES ---
	// SUBSTITUA PELA SUA CHAVE PRIVADA (da conta com ETH de teste na Sepolia)
	privateKeyHex := os.Getenv("0x34009605e23e913ddcd731063e3aa8b37a014a6494441a43e64f549cab704589")
	if privateKeyHex == "" {
		log.Fatal("Variável de ambiente ETHEREUM_PRIVATE_KEY não configurada.")
	}

	// SUBSTITUA PELO SEU ENDPOINT INFURA DA REDE SEPOLIA
	infuraURL := os.Getenv("https://sepolia.infura.io/v3/8ce7834b4e0f4f9c8eb42d2012df9ec3")
	if infuraURL == "" {
		log.Fatal("Variável de ambiente INFURA_SEPOLIA_URL não configurada.")
	}

	// Chain ID da rede Sepolia
	chainID := big.NewInt(11155111)

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
		log.Fatalf("Falha ao carregar chave privada: %v", err)
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
			log.Fatal("Sua conta não tem ETH. Obtenha ETH de teste de um Sepolia Faucet.")
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
	auth.Value = big.NewInt(0)     // Não enviando Ether junto com a implantação
	auth.GasLimit = uint64(5000000) // Limite de gás para implantação (pode precisar ajustar)
	auth.GasPrice = gasPrice

	fmt.Println("Preparando implantação do contrato...")

	// --- IMPLANTAR CONTRATO ---
	// O construtor do ChargingStationManager não aceita argumentos.
	address, tx, instance, err := blockchain_contracts.DeployChargingStationManager(auth, client)
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

// Pequena função auxiliar para obter variáveis de ambiente
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}