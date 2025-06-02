package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"

	"gopbl-3/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Lista de chaves privadas comuns do Ganache (sem o prefixo 0x)
var ganachePrivateKeys = []string{
	"4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d", // Chave correta da saída do Ganache
	"ae6ae8e5ccbfb04590405997ee2d52d2b330726137b875053c36d94e974d162f", // Conta 1
	"0dbbe8e4ae425a6d2687f1a7e3ba17bc98c673636790f1b8ad91193c05875ef1", // Conta 2
	"c88b703fb08cbea894b6aeff5a544fb92e78a18e19814cd85da83b71f772aa6c", // Conta 3
	"388c684f0ba1ef5017716adb5d21a053ea8e90277d0868337519f97bede61418", // Conta 4
}

func main() {
	log.Println("Aguardando 5 segundos para o Ganache inicializar...")
	time.Sleep(5 * time.Second)

	var privateKeyHex string
	envKey := os.Getenv("ETHEREUM_PRIVATE_KEY")
	if envKey != "" {
		privateKeyHex = envKey
		log.Println("Usando chave privada da variável de ambiente ETHEREUM_PRIVATE_KEY.")
	}

	infuraURL := os.Getenv("INFURA_SEPOLIA_URL")
	rpcURL := "http://127.0.0.1:8545"
	if infuraURL != "" {
		rpcURL = infuraURL
		log.Println("Usando URL da variável de ambiente INFURA_SEPOLIA_URL:", rpcURL)
	} else {
		log.Println("Variável de ambiente INFURA_SEPOLIA_URL não configurada. Usando RPC local do Ganache:", rpcURL)
	}

	chainID := big.NewInt(1337)

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao Ethereum via RPC (%s): %v", rpcURL, err)
	}
	defer client.Close()
	fmt.Printf("Conectado ao Ethereum via RPC (%s)!\n", rpcURL)

	var deployedAddress common.Address
	var deployedTx *types.Transaction
	var deploymentError error

	keysToTry := ganachePrivateKeys
	if privateKeyHex != "" { // Se uma chave foi fornecida via ENV, tenta ela primeiro
		keysToTry = append([]string{privateKeyHex}, ganachePrivateKeys...)
	}

	for i, pkHex := range keysToTry {
		log.Printf("Tentando implantar com a chave %d...\n", i)
		privateKey, err := crypto.HexToECDSA(pkHex)
		if err != nil {
			log.Printf("Falha ao carregar chave privada %s: %v. Tentando próxima.", pkHex, err)
			deploymentError = err
			continue
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Println("Não é uma chave pública ECDSA. Tentando próxima.")
			deploymentError = fmt.Errorf("não é uma chave pública ECDSA")
			continue
		}
		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		fmt.Printf("Endereço da conta de implantação: %s\n", fromAddress.Hex())

		balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
		if err != nil {
			log.Printf("Aviso: Falha ao obter saldo da conta %s: %v", fromAddress.Hex(), err)
			deploymentError = err
			// Não continuar se não conseguir nem checar o saldo
			continue
		}
		fmt.Printf("Saldo da conta: %s ETH (para gás)\n", new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18)))
		if balance.Cmp(big.NewInt(0)) == 0 {
			log.Printf("Aviso: Conta %s não tem ETH. Tentando próxima.", fromAddress.Hex())
			deploymentError = fmt.Errorf("conta %s sem fundos", fromAddress.Hex())
			continue
		}

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Printf("Falha ao obter nonce para %s: %v. Tentando próxima.", fromAddress.Hex(), err)
			deploymentError = err
			continue
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Printf("Falha ao obter preço do gás: %v. Tentando próxima.", err)
			deploymentError = err
			continue
		}

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			log.Printf("Falha ao criar transactor para %s: %v. Tentando próxima.", fromAddress.Hex(), err)
			deploymentError = err
			continue
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(5000000)
		auth.GasPrice = gasPrice

		fmt.Println("Preparando implantação do contrato...")
		address, tx, _, errDeploy := contracts.DeployContracts(auth, client)
		if errDeploy != nil {
			log.Printf("Falha ao implantar contrato com a conta %s: %v. Tentando próxima.", fromAddress.Hex(), errDeploy)
			deploymentError = errDeploy
			continue
		}

		fmt.Printf("Transação de implantação enviada. Hash: %s\n", tx.Hash().Hex())
		fmt.Printf("Aguardando confirmação do contrato...\n")

		receipt, errMine := bind.WaitMined(context.Background(), client, tx)
		if errMine != nil {
			log.Fatalf("Erro ao aguardar a mineração da transação de implantação: %v", errMine)
			deploymentError = errMine
			continue // Embora fatal, vamos tentar a próxima chave por via das dúvidas
		}

		if receipt.Status != 1 {
			log.Printf("Implantação do contrato falhou na blockchain. Status: %d. Tentando próxima.", receipt.Status)
			deploymentError = fmt.Errorf("implantação falhou com status %d", receipt.Status)
			continue
		}

		// Sucesso!
		deployedAddress = address
		deployedTx = tx
		deploymentError = nil // Limpa o erro anterior
		log.Printf("Contrato GerenciadorPostos implantado com sucesso com a conta %s!\n", fromAddress.Hex())
		break // Sai do loop se o deploy for bem-sucedido
	}

	if deploymentError != nil || deployedTx == nil {
		log.Fatalf("Falha ao implantar o contrato após tentar todas as chaves: %v", deploymentError)
	}

	fmt.Printf("Contrato GerenciadorPostos implantado com sucesso em: %s\n", deployedAddress.Hex())
	fmt.Printf("Hash da transação de implantação: %s\n", deployedTx.Hash().Hex())

	err = ioutil.WriteFile("contract_address.txt", []byte(deployedAddress.Hex()), 0644)
	if err != nil {
		log.Printf("Aviso: Falha ao salvar o endereço do contrato em arquivo: %v", err)
	}
}

// Pequena função auxiliar para obter variáveis de ambiente (não utilizada no momento, mantida para referência)
// func getEnv(key, defaultValue string) string {
// 	value := os.Getenv(key)
// 	if value == "" {
// 		return defaultValue
// 	}
// 	return value
// }
