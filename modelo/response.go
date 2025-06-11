package modelo

type RecomendadoResponse struct {
	ID_posto  string  `json:"id_posto"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PagamentoJson struct {
	Veiculo  Veiculo `json:"id_veiculo"`
	Valor    float64 `json:"valor"`
	ID_posto string  `json:"id_posto"`
}

type Pagamento struct {
	Veiculo  string `json:"id_veiculo"`
	Valor    float64 `json:"valor"`
	ID_posto string  `json:"id_posto"`
}

type ReqAtomica struct {
	Veiculo Veiculo `json:"veiculo"`
	Posto1 Posto   `json:"posto-1"`
	Posto2 Posto  `json:"posto-2"`
	Posto3 Posto `json:"posto-3"`
}

// type estruturaVeiculo struct {
// 	Latitude            float64 `json:"latitude"`
// 	Longitude           float64 `json:"longitude"`
// 	IsDeslocandoAoPosto bool    `json:"isDeslocandoAoPosto"`
// 	IsCarregando        bool    `json:"isCarregando"`
// }

type ReservarVagaJson struct {
	Veiculo Veiculo `json:"veiculo"`
}

type RetornarVagaJson struct {
	Posto Posto `json:"posto"`
	ID_veiculo string `json:"id_veiculo"`
}

type AtualizarPosicaoNaFila struct {
	Veiculo  Veiculo `json:"veiculo"`
	ID_posto string  `json:"id_posto"`
}

type RetornarAtualizarPosicaoFila struct {
	Veiculo Veiculo `json:"veiculo"`
	Posto   Posto   `json:"posto"`
}