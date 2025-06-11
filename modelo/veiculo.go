package modelo

// definir rotas para: consulta de pontos de recarga disponíveis;
// reserva de pontos de recarga;
// registro de recargas realizadas.

import (
	//"fmt"
	"fmt"
	"math/rand"
)

type Veiculo struct {
	ID                  string
	Latitude            float64
	Longitude           float64
	Bateria             float64
	descarregamento 	int
	isDeslocando        bool
}

func NovoVeiculo(id string, inicialLat float64, inicialLong float64) Veiculo {
	return Veiculo{
		ID:                  id,
		Latitude:            inicialLat,
		Longitude:           inicialLong,
		Bateria:             100.0, // começa com bateria cheia
		isDeslocando:        false,
		descarregamento: 	 rand.Intn(3), // 0: lenta, 1: normal, 2: rapida
	}
}

func (v Veiculo) String() string {
	return fmt.Sprintf("Veiculo id:%s\n Veiculo na posicao (%.4f, %.4f)\n Bateria: %.2f\n Deslocando: %t\n Descarregamento: %d (0 - lento, 1 - normal, 2 - rapido)\n", v.ID, v.Latitude, v.Longitude, v.Bateria, v.isDeslocando, v.descarregamento)
}

