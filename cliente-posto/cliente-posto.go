package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"time"

	"gopbl-3/contracts" // Importe o pacote do contrato gerado
	// Para as structs de modelo, se necessário
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Variáveis para interação com Ethereum
var ethClient *ethclient.Client
var contratoGerenciador *contracts.Contracts // Instância do seu contrato - CORRIGIDO
var enderecoContrato common.Address          // Endereço do contrato implantado
var chavePrivada *ecdsa.PrivateKey           // Chave privada para assinar transações
var enderecoDaConta common.Address           // Endereço da conta que envia transações

// --- Configurações da Rede Ethereum (ajuste conforme sua rede de teste) ---
const (
	rpcURL  = "http://127.0.0.1:8545" // Ou "https://sepolia.infura.io/v3/SUA_CHAVE_INFURA"
	chainID = 1337                    // 1337 para Geth --dev, 11155111 para Sepolia
	// Chave privada para uma conta com saldo na rede de teste
	// ATENÇÃO: NUNCA USE CHAVES PRIVADAS EM PRODUÇÃO DIRETAMENTE NO CÓDIGO!
	chavePrivadaHex     = "59c6995e998f97a5a004496608237ee6b68eef6727284b578e8dc038161f0340" // Exemplo: Geth --dev account 0
	enderecoContratoHex = "0x762f01C9503923d142A9f697523f6e91122a6136"                       // SUBSTITUA PELO SEU ENDEREÇO DE CONTRATO IMPLANTADO
)

func main() {
	if err := conectarEthereum(); err != nil {
		log.Fatalf("Erro fatal ao conectar ao Ethereum: %v", err)
	}
	fmt.Println("Conectado ao Ethereum com sucesso!")

	menuPrincipal()
}

func conectarEthereum() error {
	var err error
	ethClient, err = ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("falha ao conectar ao Ethereum: %v", err)
	}

	chavePrivada, err = crypto.HexToECDSA(chavePrivadaHex)
	if err != nil {
		return fmt.Errorf("falha ao carregar chave privada: %v", err)
	}

	publicKey := chavePrivada.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("não é uma chave pública ECDSA")
	}
	enderecoDaConta = crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("Usando endereço Ethereum para transações: %s\n", enderecoDaConta.Hex())

	balance, err := ethClient.BalanceAt(context.Background(), enderecoDaConta, nil)
	if err != nil {
		log.Printf("Aviso: Falha ao obter saldo da conta: %v", err)
	} else {
		fmt.Printf("Saldo da conta %s: %s Ether\n", enderecoDaConta.Hex(), new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18)).String())
	}

	enderecoContrato = common.HexToAddress(enderecoContratoHex)
	contratoGerenciador, err = contracts.NewContracts(enderecoContrato, ethClient) // CORRIGIDO
	if err != nil {
		return fmt.Errorf("falha ao instanciar o contrato GerenciadorPostos: %v", err)
	}
	fmt.Printf("Instância do contrato GerenciadorPostos carregada do endereço: %s\n", enderecoContrato.Hex())

	return nil
}

func getOpcoesTransacao(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := ethClient.PendingNonceAt(ctx, enderecoDaConta)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter nonce: %v", err)
	}

	gasPrice, err := ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter preço do gás: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(chavePrivada, big.NewInt(chainID))
	if err != nil {
		return nil, fmt.Errorf("falha ao criar transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(5000000) // Limite de gás padrão, ajuste se necessário
	auth.GasPrice = gasPrice

	return auth, nil
}

func aguardarTransacao(ctx context.Context, tx *types.Transaction) {
	fmt.Printf("Transação enviada. Hash: %s\n", tx.Hash().Hex())
	fmt.Println("Aguardando confirmação da transação... (pode levar alguns segundos)")

	receipt, err := bind.WaitMined(ctx, ethClient, tx)
	if err != nil {
		log.Printf("Erro ao aguardar a mineração da transação: %v", err)
		return
	}
	if receipt.Status != 1 {
		log.Printf("Transação falhou na blockchain. Status: %d", receipt.Status)
	} else {
		fmt.Println("Transação confirmada com sucesso!")
	}
}

func menuPrincipal() {
	for {
		fmt.Println("\n--- Menu Cliente Empresa (Posto) ---")
		fmt.Println("1. Cadastrar Novo Posto")
		fmt.Println("2. Abrir Disputa")
		fmt.Println("3. Votar em Disputa")
		fmt.Println("4. Resolver Disputa")
		fmt.Println("5. Consultar Disputa")
		fmt.Println("6. Consultar Meus Postos")
		fmt.Println("7. Sair")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			cadastrarNovoPosto()
		case 2:
			abrirDisputa()
		case 3:
			votarDisputa()
		case 4:
			resolverDisputa()
		case 5:
			consultarDisputa()
		case 6:
			consultarMeusPostos()
		case 7:
			fmt.Println("Encerrando cliente.")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func cadastrarNovoPosto() {
	var id, nome, cidade string
	fmt.Println("\n--- Cadastro de Novo Posto ---")
	fmt.Print("ID do Posto: ")
	fmt.Scanln(&id)
	fmt.Print("Nome do Posto: ")
	fmt.Scanln(&nome)
	fmt.Print("Cidade do Posto: ")
	fmt.Scanln(&cidade)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	auth, err := getOpcoesTransacao(ctx)
	if err != nil {
		log.Printf("Erro ao preparar transação: %v", err)
		return
	}

	tx, err := contratoGerenciador.RegisterStation(auth, id, nome, cidade) // CORRIGIDO: nome da função no contrato
	if err != nil {
		fmt.Printf("Erro ao cadastrar posto no contrato: %v\n", err)
		return
	}
	aguardarTransacao(ctx, tx)
}

func abrirDisputa() {
	var idPosto, motivo string
	fmt.Println("\n--- Abrir Disputa ---")
	fmt.Print("ID do Posto para disputa: ")
	fmt.Scanln(&idPosto)
	fmt.Print("Motivo da disputa: ")
	fmt.Scanln(&motivo)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	auth, err := getOpcoesTransacao(ctx)
	if err != nil {
		log.Printf("Erro ao preparar transação: %v", err)
		return
	}

	tx, err := contratoGerenciador.AbrirDisputa(auth, idPosto, motivo)
	if err != nil {
		fmt.Printf("Erro ao abrir disputa no contrato: %v\n", err)
		return
	}
	aguardarTransacao(ctx, tx)
}

func votarDisputa() {
	var idPosto string
	var votoOpcao int
	fmt.Println("\n--- Votar em Disputa ---")
	fmt.Print("ID do Posto da disputa: ")
	fmt.Scanln(&idPosto)
	fmt.Println("Seu voto:")
	fmt.Println("1. A favor da empresa requerente")
	fmt.Println("2. Contra a empresa requerente (a favor do proprietário atual)")
	fmt.Print("Escolha uma opção: ")
	fmt.Scanln(&votoOpcao)

	aFavor := false
	if votoOpcao == 1 {
		aFavor = true
	} else if votoOpcao != 2 {
		fmt.Println("Opção de voto inválida. Voto cancelado.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	auth, err := getOpcoesTransacao(ctx)
	if err != nil {
		log.Printf("Erro ao preparar transação: %v", err)
		return
	}

	tx, err := contratoGerenciador.VotarDisputa(auth, idPosto, aFavor)
	if err != nil {
		fmt.Printf("Erro ao votar na disputa: %v\n", err)
		return
	}
	aguardarTransacao(ctx, tx)
}

func resolverDisputa() {
	var idPosto string
	fmt.Println("\n--- Resolver Disputa ---")
	fmt.Print("ID do Posto da disputa a ser resolvida: ")
	fmt.Scanln(&idPosto)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	auth, err := getOpcoesTransacao(ctx)
	if err != nil {
		log.Printf("Erro ao preparar transação: %v", err)
		return
	}

	tx, err := contratoGerenciador.ResolverDisputa(auth, idPosto)
	if err != nil {
		fmt.Printf("Erro ao resolver disputa: %v\n", err)
		return
	}
	aguardarTransacao(ctx, tx)
}

func consultarDisputa() {
	var idPosto string
	fmt.Println("\n--- Consultar Disputa ---")
	fmt.Print("ID do Posto da disputa para consulta: ")
	fmt.Scanln(&idPosto)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// A função gerada ConsultarDisputa retorna uma struct com todos os campos da disputa
	disputa, err := contratoGerenciador.ConsultarDisputa(&bind.CallOpts{Context: ctx}, idPosto)
	if err != nil {
		fmt.Printf("Erro ao consultar disputa: %v\n", err)
		return
	}

	fmt.Println("\n--- Detalhes da Disputa ---")
	fmt.Printf("ID do Posto: %s\n", disputa.IdPosto)
	fmt.Printf("Empresa Requerente: %s\n", disputa.EmpresaRequerente.Hex())
	fmt.Printf("Motivo: %s\n", disputa.Motivo)
	fmt.Printf("Votos a Favor: %d\n", disputa.VotosAFavor)
	fmt.Printf("Votos Contra: %d\n", disputa.VotosContra)
	// Convertendo o enum StatusDisputa para string
	statusDisputaStr := ""
	switch disputa.Status {
	case 0: // Nenhuma
		statusDisputaStr = "Nenhuma"
	case 1: // Aberta
		statusDisputaStr = "Aberta"
	case 2: // Resolvida
		statusDisputaStr = "Resolvida"
	default:
		statusDisputaStr = "Desconhecido"
	}
	fmt.Printf("Status: %s\n", statusDisputaStr)
	fmt.Printf("Vencedora: %s\n", disputa.Vencedora.Hex())
	fmt.Printf("Votantes: %v\n", disputa.Votantes)
}

func consultarMeusPostos() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// O contrato armazena os IDs dos postos por empresa
	idsPostos, err := contratoGerenciador.PostosPorEmpresa(&bind.CallOpts{Context: ctx}, enderecoDaConta)
	if err != nil {
		fmt.Printf("Erro ao consultar postos da sua empresa: %v\n", err)
		return
	}

	if len(idsPostos) == 0 {
		fmt.Println("\nSua empresa não possui postos cadastrados.")
		return
	}

	fmt.Println("\n--- Seus Postos Cadastrados ---")
	for _, id := range idsPostos {
		// A função ConsultarPosto do contrato Solidity retorna a struct completa do posto
		posto, err := contratoGerenciador.ChargingStations(&bind.CallOpts{Context: ctx}, id) // CORRIGIDO
		if err != nil {
			fmt.Printf("Erro ao consultar detalhes do posto %s: %v\n", id, err)
			continue
		}
		status := "Disponível"
		if !posto.Disponivel {
			status = fmt.Sprintf("Reservado por: %s", posto.VeiculoReservador.Hex())
		}
		fmt.Printf("ID: %s, Nome: %s, Cidade: %s, Status: %s, Última Atualização: %s\n",
			posto.Id, posto.Nome, posto.Cidade, status, time.Unix(int64(posto.UltimaAtualizacao.Uint64()), 0).Format(time.RFC3339))
	}
}

// A função init() é executada antes da main(), mas aqui é apenas para capturar Ctrl+C
// e não é um "loop" ou parte principal do fluxo de execução do cliente.
func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("\nRecebido sinal de interrupção. Encerrando cliente...")
		os.Exit(0)
	}()
}
