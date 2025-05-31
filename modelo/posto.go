package modelo

import (
	"fmt"
	"time"
	//"math"
	//"sort"
	//"sync"
)

type Posto struct {
	ID                string    `bson:"id" json:"id"`
	Latitude          float64   `bson:"latitude" json:"latitude"`
	Longitude         float64   `bson:"longitude" json:"longitude"`
	Disponivel        bool      `bson:"disponivel" json:"disponivel"`
	UltimaAtualizacao time.Time `bson:"ultimaAtualizacao" json:"ultimaAtualizacao"`
	ServidorOrigem    string    `bson:"servidorOrigem" json:"servidorOrigem"`
	Cidade            string    `bson:"cidade" json:"cidade"`
}

func NovoPosto(id string, lat float64, long float64, servidorOrigem string, cidade string) Posto {
	fmt.Printf("Posto %s criado na localização (%.6f, %.6f) no servidor %s",
		id, lat, long, servidorOrigem)

	return Posto{
		ID:                id,
		Latitude:          lat,
		Longitude:         long,
		Disponivel:        true,
		UltimaAtualizacao: time.Now(),
		ServidorOrigem:    servidorOrigem,
		Cidade:            cidade,
	}
}

func String(p Posto) string {
	status := "disponível"
	if !p.Disponivel {
		status = "reservado"
	}
	return fmt.Sprintf("Posto id:%s\n Cidade: %s\n Posto localizado em (%.4f, %.4f)\n Status: %s\n Última atualização: %s",
		p.ID, p.Cidade, p.Latitude, p.Longitude, status, p.UltimaAtualizacao.Format(time.RFC3339))
}

// atualiza o status de disponibilidade e o timestamp
func AtualizarDisponibilidade(p *Posto, disponivel bool) {
	p.Disponivel = disponivel
	p.UltimaAtualizacao = time.Now()
}
