package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"

	//"log"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"sort"

	//"sync"
	"time"
)

type Posto struct {
	ID                string    `json:"id"`
	Nome              string    `json:"nome"`
	Cidade            string    `json:"cidade"`
	Latitude          float64   `json:"latitude"`
	Longitude         float64   `json:"longitude"`
	Ocupado           bool      `json:"ocupado"`
	UltimaAtualizacao time.Time `json:"ultimaAtualizacao"`
	ServidorOrigem    string    `json:"servidorOrigem"`
	VeiculoReservador string    `json:"veiculoReservador"`
	ValorPagamento    uint64    `json:"valorPagamento"`
}

// Modelo de dados para a requisição de reserva
type ReservaData struct {
	IDPostos       []string `json:"idPostos"`
	Reservar       bool     `json:"reservar"`
	ClientID       string   `json:"clientId"`
	ValorPagamento uint64   `json:"valorPagamento"`
}

// Modelo para o veículo
type Veiculo struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var veiculo Veiculo
var cadastrado bool = false
var idPostosReservados []string

var servidorBackendURL string // URL do seu servidor Go

// var servidorBackendURL string = "http://172.16.201.11:8083" // URL do seu servidor Go
// 192.168.0.110

func main() {
	// Definição dos flags para o endereço do servidor
	serverIP := flag.String("ip", "127.0.0.1", "IP do servidor backend")
	serverPort := flag.String("port", "8083", "Porta do servidor backend")
	flag.Parse()

	// Constrói a URL do backend
	servidorBackendURL = fmt.Sprintf("http://%s:%s", *serverIP, *serverPort)
	fmt.Printf("Conectando ao servidor backend em: %s\n", servidorBackendURL)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("\nRecebido sinal de interrupção. Encerrando cliente...")

		if len(idPostosReservados) > 0 && veiculo.ID != "" {
			fmt.Println("Tentando liberar postos reservados...")
			onSubmit(false, 0) // Tenta liberar os postos antes de sair
		}
		fmt.Println("Cliente encerrado.")
		os.Exit(0)
	}()

	menu()
}

func cadastrarVeiculo() {
	fmt.Print("Digite o ID do veículo : ")
	fmt.Scanln(&veiculo.ID)
	fmt.Print("Digite a Latitude do veículo: ")
	fmt.Scanln(&veiculo.Latitude)

	fmt.Print("Digite a Longitude do veículo: ")
	fmt.Scanln(&veiculo.Longitude)

	cadastrado = true
	fmt.Println("Veículo cadastrado:", veiculo)
}

func onSubmit(reservar bool, preco float64) {
	data := ReservaData{
		IDPostos:       idPostosReservados,
		Reservar:       reservar,
		ClientID:       veiculo.ID,
		ValorPagamento: uint64(preco * 100), // Converter para centavos
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
		if reservar {
			return "reserva"
		} else {
			return "liberação"
		}
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

	fmt.Println("\n--- Postos Disponíveis ---")
	for _, posto := range postos {
		estado := "Disponível"
		if posto.Ocupado {
			estado = fmt.Sprintf("Reservado por %s (Valor: R$ %.2f)", posto.VeiculoReservador, float64(posto.ValorPagamento)/100.0)
		}
		fmt.Printf("ID: %s, Cidade: %s, Estado: %s\n", posto.ID, posto.Cidade, estado)
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
		if !posto.Ocupado {
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
	// Calcular a distância total da rota
	distanciaTotal := 0.0
	pontoAnterior := veiculo // Começa do veículo
	for _, posto := range rotaEscolhida {
		distanciaTotal += distanciaHaversine(pontoAnterior.Latitude, pontoAnterior.Longitude, posto.Latitude, posto.Longitude)
		pontoAnterior.Latitude = posto.Latitude
		pontoAnterior.Longitude = posto.Longitude
	}

	precoPorKm := 0.5 // Exemplo de preço por KM
	precoTotal := distanciaTotal * precoPorKm

	fmt.Printf("\n--- Detalhes da Rota e Pagamento ---\n")
	fmt.Printf("Rota Selecionada: ")
	for i, p := range rotaEscolhida {
		fmt.Printf("%s", p.ID)
		if i < len(rotaEscolhida)-1 {
			fmt.Printf(" -> ")
		}
	}
	fmt.Println()
	fmt.Printf("Distância Total: %.2f km\n", distanciaTotal)
	fmt.Printf("Preço Total: R$ %.2f\n", precoTotal)

	fmt.Println("\nDeseja confirmar a reserva e o pagamento? (s/n)")
	var confirmacao string
	fmt.Scanln(&confirmacao)

	if confirmacao == "s" || confirmacao == "S" {
		idPostosReservados = []string{}
		for _, p := range rotaEscolhida {
			idPostosReservados = append(idPostosReservados, p.ID)
		}

		if len(idPostosReservados) > 0 {
			fmt.Println("Confirmando reserva para os postos:", idPostosReservados)
			onSubmit(true, precoTotal) // Tenta reservar
		} else {
			fmt.Println("Nenhum posto selecionado.")
		}
	} else {
		fmt.Println("Reserva cancelada.")
	}
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func distanciaHaversine(lat1, lon1, lat2, lon2 float64) float64 {
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6371 // Raio da Terra em KM

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)
	distancia := 2 * r * math.Asin(math.Sqrt(h))
	return distancia
}

func alterarServidorBackend() {
	var novoIP, novaPorta string

	fmt.Printf("Servidor backend atual: %s\n", servidorBackendURL)
	fmt.Print("Digite o novo IP do servidor: ")
	fmt.Scanln(&novoIP)
	fmt.Print("Digite a nova porta do servidor: ")
	fmt.Scanln(&novaPorta)

	if novoIP != "" && novaPorta != "" {
		servidorBackendURL = fmt.Sprintf("http://%s:%s", novoIP, novaPorta)
		fmt.Printf("Servidor backend alterado para: %s\n", servidorBackendURL)
	} else {
		fmt.Println("Entrada inválida. O endereço do servidor não foi alterado.")
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
		fmt.Println("6 - Alterar servidor backend")
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
				onSubmit(false, 0) // Tenta finalizar a reserva
			} else {
				fmt.Println("Você não possui reservas ativas para finalizar.")
			}
		case 6:
			alterarServidorBackend()
		//case 6:
		// Lógica de encerramento já está na goroutine do sinal
		//return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
