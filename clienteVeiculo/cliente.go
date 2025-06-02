package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"log"
	"net/http"
	"os"
	"os/signal"
	"sort"
	//"sync"
	"time"
)

// Modelo de dados simplificado para o posto (replicando a struct do servidor)
type Posto struct {
	ID                 string    `json:"id"`
	Nome               string    `json:"nome"`
	Cidade             string    `json:"cidade"`
	Latitude           float64   `json:"latitude"`
	Longitude          float64   `json:"longitude"`
	Disponivel         bool      `json:"disponivel"`
	UltimaAtualizacao  time.Time `json:"ultimaAtualizacao"`
	ServidorOrigem     string    `json:"servidorOrigem"`
	VeiculoReservador string    `json:"veiculoReservador"`
}

// Modelo de dados para a requisição de reserva
type ReservaData struct {
	IDPostos []string `json:"idPostos"`
	Reservar bool     `json:"reservar"`
	ClientID string   `json:"clientId"`
}

// Modelo para o veículo
type Veiculo struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var veiculo Veiculo
var cadastrado bool = false
var idPostosReservados []string // Mudando o nome para maior clareza
var servidorBackendURL string = "http://127.0.0.1:8083" // URL do seu servidor Go

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("\nRecebido sinal de interrupção. Encerrando cliente...")

		if len(idPostosReservados) > 0 && veiculo.ID != "" {
			fmt.Println("Tentando liberar postos reservados...")
			onSubmit(false) // Tenta liberar os postos antes de sair
		}
		fmt.Println("Cliente encerrado.")
		os.Exit(0)
	}()

	menu()
}

func cadastrarVeiculo() {
	fmt.Print("Digite o ID do veículo (deve ser um endereço ETH válido para teste): ")
	fmt.Scanln(&veiculo.ID) // Em um cenário real, o ID do veículo seria um endereço ETH gerado pelo cliente.

	fmt.Print("Digite a Latitude do veículo: ")
	fmt.Scanln(&veiculo.Latitude)

	fmt.Print("Digite a Longitude do veículo: ")
	fmt.Scanln(&veiculo.Longitude)

	cadastrado = true
	fmt.Println("Veículo cadastrado:", veiculo)
}

func onSubmit(reservar bool) {
	data := ReservaData{
		IDPostos: idPostosReservados,
		Reservar: reservar,
		ClientID: veiculo.ID,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erro ao codificar JSON:", err)
		return
	}

	req, err := http.NewRequest(http.MethodPut, servidorBackendURL+"/reservar", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second} // Aumentar timeout para operações de blockchain
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar requisição HTTP:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Erro ao decodificar resposta:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Erro do servidor (%d): %s\n", resp.StatusCode, result["error"])
		if reservar {
			idPostosReservados = []string{} // Limpa se a reserva falhou
		}
		return
	}

	fmt.Printf("Operação de %s concluída: %s (Tx Hash: %s)\n", func() string {
		if reservar { return "reserva" } else { return "liberação" }
	}(), result["message"], result["txHash"])

	if !reservar { // Se a operação foi de liberar, limpa a lista de postos reservados
		idPostosReservados = []string{}
	}
}

func listarPostos() []Posto {
	resp, err := http.Get(servidorBackendURL + "/postosDisponiveis")
	if err != nil {
		fmt.Println("Erro ao consultar postos disponíveis:", err)
		return []Posto{}
	}
	defer resp.Body.Close()

	var postos []Posto
	if err := json.NewDecoder(resp.Body).Decode(&postos); err != nil {
		fmt.Println("Erro ao decodificar resposta:", err)
		return []Posto{}
	}

	rotasGeradas := montarRotas(postos)

	if len(rotasGeradas) > 0 {
		var chaves []int
		for k := range rotasGeradas {
			chaves = append(chaves, k)
		}
		sort.Ints(chaves)
		fmt.Println("\n--- Rotas Disponíveis ---")
		for _, idx := range chaves {
			rota := rotasGeradas[idx]
			fmt.Printf("Rota %d: ", idx)
			for _, posto := range rota {
				fmt.Printf("%s (%s) ", posto.ID, posto.Cidade)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Nenhuma rota disponível encontrada.")
	}

	return postos
}

func montarRotas(postos []Posto) map[int][]Posto {
    var postosFSA []Posto
    var postosSonga []Posto
    var postosSerrinha []Posto
    var rotas = make(map[int][]Posto)
    var quantidadeRotas int = 0

    // Filtra apenas postos disponíveis para montar rotas
    for _, posto := range postos {
        if posto.Disponivel {
            switch posto.Cidade {
            case "Feira de Santana":
                postosFSA = append(postosFSA, posto)
            case "Serrinha":
                postosSerrinha = append(postosSerrinha, posto)
            case "São Gonçalo":
                postosSonga = append(postosSonga, posto)
            }
        }
    }

    var todosPostosDisponiveis []Posto
    todosPostosDisponiveis = append(todosPostosDisponiveis, postosFSA...)
    todosPostosDisponiveis = append(todosPostosDisponiveis, postosSonga...)
    todosPostosDisponiveis = append(todosPostosDisponiveis, postosSerrinha...)

    // Rotas com 1 posto
    for _, p := range todosPostosDisponiveis {
        rotas[quantidadeRotas] = []Posto{p}
        quantidadeRotas++
    }

    // Rotas com 2 postos (ordem importa, cidades diferentes)
    for i := 0; i < len(todosPostosDisponiveis); i++ {
        for j := 0; j < len(todosPostosDisponiveis); j++ {
            if i != j && todosPostosDisponiveis[i].Cidade != todosPostosDisponiveis[j].Cidade {
                rotas[quantidadeRotas] = []Posto{todosPostosDisponiveis[i], todosPostosDisponiveis[j]}
                quantidadeRotas++
            }
        }
    }

    // Rotas com 3 postos (ordem importa, cidades diferentes)
    for i := 0; i < len(todosPostosDisponiveis); i++ {
        for j := 0; j < len(todosPostosDisponiveis); j++ {
            for k := 0; k < len(todosPostosDisponiveis); k++ {
                if i != j && i != k && j != k {
                    // Verifica se as 3 cidades são distintas
                    ci := todosPostosDisponiveis[i].Cidade
                    cj := todosPostosDisponiveis[j].Cidade
                    ck := todosPostosDisponiveis[k].Cidade
                    if ci != cj && ci != ck && cj != ck {
                        rotas[quantidadeRotas] = []Posto{todosPostosDisponiveis[i], todosPostosDisponiveis[j], todosPostosDisponiveis[k]}
                        quantidadeRotas++
                    }
                }
            }
        }
    }

    return rotas
}

func procurarPostosParaReserva(rotas map[int][]Posto) {
	var escolha int

	fmt.Println("\nDigite o número da rota que deseja reservar, ou -1 para sair:")
	fmt.Scanln(&escolha)

	if escolha == -1 {
		fmt.Println("Operação cancelada.")
		return
	}

	rotaEscolhida, existe := rotas[escolha]
	if !existe {
		fmt.Println("Rota inválida.")
		return
	}
	idPostosReservados = []string{}
	for _, p := range rotaEscolhida {
		idPostosReservados = append(idPostosReservados, p.ID)
	}

	if len(idPostosReservados) > 0 {
		fmt.Println("Postos selecionados para reserva:", idPostosReservados)
		onSubmit(true) // Tenta reservar
	} else {
		fmt.Println("Nenhum posto selecionado.")
	}
}

func menu() {
	for {
		fmt.Println("\nMenu de Ações:")
		fmt.Println("1 - Cadastrar veículo")
		fmt.Println("2 - Atualizar posição do veículo")
		fmt.Println("3 - Consultar postos disponíveis")
		fmt.Println("4 - Reservar posto")
		fmt.Println("5 - Finalizar viagem")
		fmt.Println("6 - Sair") // Removida opção de servidor preferido, pois não é mais relevante no cliente para este cenário
		var opcao int
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			fmt.Println("Cadastrar veículo")
			cadastrarVeiculo()
		case 2:
			if !cadastrado {
				fmt.Println("Veículo não cadastrado!")
				continue
			}
			fmt.Println("Digite a nova Latitude do veículo:")
			var lat float64
			fmt.Scanln(&lat)
			fmt.Println("Digite a nova Longitude do veículo:")
			var long float64
			fmt.Scanln(&long)
			veiculo.Latitude = lat
			veiculo.Longitude = long
			fmt.Println("Veículo atualizado:", veiculo)
		case 3:
			listarPostos()
		case 4:
			if len(idPostosReservados) == 0 { // Só permite reservar se não houver reserva ativa
				postos := listarPostos() // Lista para exibir as rotas disponíveis
				if len(postos) > 0 {
					procurarPostosParaReserva(montarRotas(postos))
				} else {
					fmt.Println("Nenhum posto disponível para reserva.")
				}
			} else {
				fmt.Println("Você já possui reservas em andamento. Finalize a viagem para reservar novamente.")
			}
		case 5:
			if len(idPostosReservados) > 0 {
				onSubmit(false) // Tenta finalizar a reserva
			} else {
				fmt.Println("Você não possui reservas ativas para finalizar.")
			}
		case 6:
			// Lógica de encerramento já está na goroutine do sinal
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}