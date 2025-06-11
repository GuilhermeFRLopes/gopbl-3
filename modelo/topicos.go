package modelo

const (
	// Tópicos para comunicação entre cliente e servidor
	TopicPostosDisponiveis      = "postos/disponiveis"
	TopicCadastrarPosto         = "postos/cadastrar"
	TopicDeletarPosto           = "postos/deletar"
	TopicReservarPosto          = "postos/reservar"
	TopicResposta               = "postos/resposta"
	//TopicReservaIntermediador   = "postos/reserva/intermediador"
	TopicReservaBloqueio        = "postos/reserva/bloqueio"
	TopicReservaEscutarBloqueio = "postos/reserva/escutarBloqueio"

	// Novos tópicos para comunicação específica por servidor
	// O formato será "postos/servidor/{nome-servidor}/{operacao}"
	TopicServidorPrefix = "postos/servidor/"
)

// Função para gerar tópico específico para um servidor
func GetTopicServidor(nomeServidor, operacao string) string {
	return TopicServidorPrefix + nomeServidor + "/" + operacao
}
