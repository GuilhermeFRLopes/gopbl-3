package modelo

// Estrutura padrão para Posto
// Pode ser usada em scripts Go para interação com o contrato

type Posto struct {
	ID                string // Identificador único
	Nome              string // Nome do posto
	Cidade            string // Cidade onde está localizado
	Proprietario      string // Endereço Ethereum do proprietário
	Disponivel        bool   // Disponibilidade
	UltimaAtualizacao int64  // Timestamp da última atualização
	VeiculoReservador string // Endereço do veículo que reservou
}
