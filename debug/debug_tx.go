package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	//"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	
	contratos_binding "gopbl-3/contratos" 
)

func main() {
	// --- CONFIGURAÇÕES ---
	ganacheURL := "http://127.0.0.1:7545" // IP e porta do Ganache 
	// usando Docker Ganache e rodando este script no host: http://127.0.0.1:8545
	// em outro PC na rede: http://IP_DO_PC_PRINCIPAL:8545

	// Hash da transação
	// Exemplo de hash de transação de reserva (adapte conforme suas transações)
	// VOCÊ PRECISA ATUALIZAR ESTE HASH COM O HASH REAL DE UMA TRANSAÇÃO DE RESERVA OU CADASTRO!
	txHash := common.HexToHash("0x982cf6250599da24a8af20eacbdd08add84bc4ae7a713c1a88bbfef4063fbfb7") // Hash de exemplo

	// Endereço do contrato com o qual a transação interagiu
	// É o "TO CONTRACT ADDRESS" da transação (o endereço do seu contrato implantado)
	contractAddr := common.HexToAddress("0xEf8F65612B109E927f23fc654F8eF3FE8601DC72") // Endereço de contrato de exemplo

	// --- CONEXÃO ETHEREUM ---
	client, err := ethclient.Dial(ganacheURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao Ganache: %v", err)
	}
	defer client.Close()
	fmt.Println("Conectado ao Ganache!")

	// --- CARREGAR ABI DO CONTRATO ---
	// Carrega a ABI diretamente do metadata gerado pelo abigen
	contractABI, err := abi.JSON(strings.NewReader(contratos_binding.ContratosMetaData.ABI))
	if err != nil {
		log.Fatalf("Falha ao parsear ABI: %v", err)
	}

	// --- OBTER E DECODIFICAR A TRANSAÇÃO ---
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatalf("Falha ao obter transação pelo hash %s: %v", txHash.Hex(), err)
	}

	if isPending {
		fmt.Println("Transação ainda pendente. Aguarde a mineração.")
		return
	}

	// Recupera o remetente da transação
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Falha ao obter ChainID da rede: %v", err)
	}
	signer := types.LatestSignerForChainID(chainID)

	senderAddr, err := signer.Sender(tx)
	if err != nil {
		log.Printf("Aviso: Falha ao recuperar remetente da transação: %v. Pode ser uma transação de criação de contrato antiga ou formatada incorretamente.", err)
		senderAddr = common.Address{} // Definir como zero address em caso de falha
	}

	fmt.Printf("\n--- Detalhes da Transação %s ---\n", txHash.Hex())
	fmt.Printf("De: %s\n", senderAddr.Hex())
	fmt.Printf("Para (Contrato): %s\n", tx.To().Hex())
	fmt.Printf("Valor (ETH): %s\n", new(big.Float).Quo(new(big.Float).SetInt(tx.Value()), big.NewFloat(1e18)).String())
	fmt.Printf("Gás Limite: %d\n", tx.Gas())
	fmt.Printf("Preço do Gás: %s Wei\n", tx.GasPrice().String())
	fmt.Printf("Nonce: %d\n", tx.Nonce())
	fmt.Printf("Dados de Entrada (Raw Input Data): %x\n", tx.Data())

	// Tenta decodificar os dados de entrada como chamada de função
	if tx.To() != nil && *tx.To() == contractAddr && len(tx.Data()) >= 4 {
		methodID := tx.Data()[:4]
		method, err := contractABI.MethodById(methodID)
		if err != nil {
			fmt.Printf("Aviso: Não foi possível encontrar método da ABI pelo seletor %x (pode não ser uma chamada de função ou a ABI está incompleta): %v\n", methodID, err)
		} else {
			fmt.Printf("\nFunção Chamada: %s\n", method.Name)
			fmt.Println("Parâmetros:")

			inputs := make(map[string]interface{})
			err = method.Inputs.UnpackIntoMap(inputs, tx.Data()[4:])
			if err != nil {
				fmt.Printf("Aviso: Falha ao decodificar parâmetros: %v\n", err)
			} else {
				for name, value := range inputs {
					// Impressão mais robusta para diferentes tipos
					if val, ok := value.(*big.Int); ok {
						fmt.Printf("  %s: %s (big.Int - uint256/int256)\n", name, val.String())
					} else if val, ok := value.([]string); ok {
						fmt.Printf("  %s: %v (string array)\n", name, val)
					} else if val, ok := value.(common.Address); ok {
						fmt.Printf("  %s: %s (address)\n", name, val.Hex())
					} else if val, ok := value.(bool); ok {
						fmt.Printf("  %s: %t (bool)\n", name, val)
					} else if val, ok := value.(string); ok {
						fmt.Printf("  %s: \"%s\" (string)\n", name, val)
					} else {
						fmt.Printf("  %s: %v (tipo desconhecido)\n", name, value)
					}
				}
			}
		}
	} else if tx.To() == nil {
		fmt.Println("\nEsta transação é uma transação de criação de contrato (To é nil).")
	} else {
		fmt.Println("\nEsta transação não é uma chamada para o contrato especificado, ou é uma transação simples de ETH.")
	}

	// --- Opcional: Obter o Recibo da Transação e Eventos ---
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Printf("Aviso: Falha ao obter recibo da transação: %v", err)
	} else {
		fmt.Printf("\nStatus do Recibo: %d (1 = Sucesso, 0 = Falha)\n", receipt.Status)
		if len(receipt.Logs) > 0 {
			fmt.Println("--- Eventos Emitidos ---")
			for _, logEntry := range receipt.Logs {
				fmt.Printf("  Log do Endereço: %s\n", logEntry.Address.Hex())
				fmt.Printf("  Tópicos (%d): %v\n", len(logEntry.Topics), logEntry.Topics)
				fmt.Printf("  Dados: %x\n", logEntry.Data)

				if len(logEntry.Topics) > 0 {
					eventID := logEntry.Topics[0]
					event, err := contractABI.EventByID(eventID)
					if err == nil {
						fmt.Printf("  Evento: %s\n", event.Name)
						decodedEventData := make(map[string]interface{})
						// Descompacta os dados não indexados (do logEntry.Data)
						err = contractABI.UnpackIntoMap(decodedEventData, event.Name, logEntry.Data)
						if err == nil {
							fmt.Println("  Dados do Evento Decodificados (dados não-indexados):")
							for k, v := range decodedEventData {
								fmt.Printf("    %s: %v\n", k, v)
							}
						} else {
							fmt.Printf("  Aviso: Falha ao descompactar dados NÃO indexados do evento: %v\n", err)
						}

						// Decodifica os dados indexados (dos tópicos)
						fmt.Println("  Dados do Evento Decodificados (dados indexados):")
						// O primeiro tópico (Topics[0]) é o hash do evento. Os tópicos seguintes são os parâmetros 'indexed'.
						for i, input := range event.Inputs {
							if input.Indexed {
								if i+1 < len(logEntry.Topics) { // +1 porque Topics[0] é o hash do evento
									topic := logEntry.Topics[i+1]
									if input.Type.String() == "string" {
										// Strings indexed são armazenadas como keccak256 hash no tópico
										fmt.Printf("    %s (indexed, string): %s (hash keccak256)\n", input.Name, topic.Hex())
									} else if input.Type.String() == "address" {
										fmt.Printf("    %s (indexed, address): %s\n", input.Name, common.BytesToAddress(topic.Bytes()).Hex())
									} else if input.Type.String() == "bool" {
										fmt.Printf("    %s (indexed, bool): %t\n", input.Name, topic.Big().Cmp(big.NewInt(0)) != 0) // Booleans são 0 ou 1
									} else if input.Type.String() == "uint256" || input.Type.String() == "int256" {
										fmt.Printf("    %s (indexed, %s): %s\n", input.Name, input.Type.String(), topic.Big().String())
									} else {
										fmt.Printf("    %s (indexed, %s): %s\n", input.Name, input.Type.String(), topic.Hex())
									}
								} else {
									fmt.Printf("    Aviso: Tópico indexado esperado para '%s' mas não encontrado.\n", input.Name)
								}
							}
						}
					} else {
						fmt.Printf("  Aviso: Não foi possível encontrar a definição do evento na ABI para o tópico %s: %v\n", eventID.Hex(), err)
					}
				}
				fmt.Println("---")
			}
		} else {
			fmt.Println("Nenhum evento emitido nesta transação.") // Confirma o "NO EVENTS"
		}
	}
}