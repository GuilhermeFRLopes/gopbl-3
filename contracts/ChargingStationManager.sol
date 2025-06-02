// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract GerenciadorPostos {
    struct Posto {
        string id;
        string nome;
        string cidade;
        address proprietario;
        bool disponivel;
        uint256 ultimaAtualizacao;
        address veiculoReservador;
    }

    struct Reserva {
        string idPosto;
        address veiculo;
        uint256 timestamp;
    }

    struct Recarga {
        string idPosto;
        address veiculo;
        uint256 kWh;
        uint256 valor;
        uint256 timestamp;
    }

    struct Pagamento {
        string idPosto;
        address veiculo;
        uint256 valor;
        uint256 timestamp;
    }

    // --- Disputa descentralizada entre empresas ---
    enum StatusDisputa { Nenhuma, Aberta, Resolvida }
    struct Disputa {
        string idPosto;
        address empresaRequerente;
        string motivo;
        uint256 votosAFavor;
        uint256 votosContra;
        StatusDisputa status;
        mapping(address => bool) votou;
        address[] votantes;
        address vencedora; // empresa vencedora da disputa
    }

    mapping(string => Posto) public postos;
    string[] public listaIdsPostos;
    mapping(address => string[]) public postosPorEmpresa;

    Reserva[] public historicoReservas;
    Recarga[] public historicoRecargas;
    Pagamento[] public historicoPagamentos;

    mapping(string => Disputa) public disputas; // disputa por id do posto
    string[] public listaDisputas;

    event PostoCadastrado(string idPosto, address proprietario, string cidade);
    event ReservaEfetuada(string idPosto, address veiculo, uint256 timestamp);
    event RecargaEfetuada(string idPosto, address veiculo, uint256 kWh, uint256 valor, uint256 timestamp);
    event PagamentoEfetuado(string idPosto, address veiculo, uint256 valor, uint256 timestamp);
    event DisputaAberta(string idPosto, address empresa, string motivo);
    event VotoRegistrado(string idPosto, address empresa, bool aFavor);
    event DisputaResolvida(string idPosto, address vencedora);

    function cadastrarPosto(string memory _id, string memory _nome, string memory _cidade) public {
        require(bytes(postos[_id].id).length == 0, "Posto ja cadastrado.");
        postos[_id] = Posto(_id, _nome, _cidade, msg.sender, true, block.timestamp, address(0));
        listaIdsPostos.push(_id);
        postosPorEmpresa[msg.sender].push(_id);
        emit PostoCadastrado(_id, msg.sender, _cidade);
    }

    function reservarPosto(string memory _id) public {
        require(bytes(postos[_id].id).length > 0, "Posto inexistente.");
        require(postos[_id].disponivel, "Posto ja reservado.");
        postos[_id].disponivel = false;
        postos[_id].veiculoReservador = msg.sender;
        postos[_id].ultimaAtualizacao = block.timestamp;
        historicoReservas.push(Reserva(_id, msg.sender, block.timestamp));
        emit ReservaEfetuada(_id, msg.sender, block.timestamp);
    }

    function liberarPosto(string memory _id) public {
        require(bytes(postos[_id].id).length > 0, "Posto inexistente.");
        require(!postos[_id].disponivel, "Posto ja esta disponivel.");
        require(postos[_id].veiculoReservador == msg.sender, "Apenas o veiculo que reservou pode liberar.");
        postos[_id].disponivel = true;
        postos[_id].veiculoReservador = address(0);
        postos[_id].ultimaAtualizacao = block.timestamp;
    }

    function registrarRecarga(string memory _id, uint256 _kWh, uint256 _valor) public {
        require(bytes(postos[_id].id).length > 0, "Posto inexistente.");
        require(!postos[_id].disponivel, "Posto deve estar reservado.");
        require(postos[_id].veiculoReservador == msg.sender, "Apenas o veiculo que reservou pode registrar recarga.");
        historicoRecargas.push(Recarga(_id, msg.sender, _kWh, _valor, block.timestamp));
        emit RecargaEfetuada(_id, msg.sender, _kWh, _valor, block.timestamp);
    }

    function registrarPagamento(string memory _id, uint256 _valor) public {
        require(bytes(postos[_id].id).length > 0, "Posto inexistente.");
        require(!postos[_id].disponivel, "Posto deve estar reservado.");
        require(postos[_id].veiculoReservador == msg.sender, "Apenas o veiculo que reservou pode pagar.");
        historicoPagamentos.push(Pagamento(_id, msg.sender, _valor, block.timestamp));
        emit PagamentoEfetuado(_id, msg.sender, _valor, block.timestamp);
    }

    function consultarPostosPorEmpresa(address _empresa) public view returns (string[] memory) {
        return postosPorEmpresa[_empresa];
    }

    function consultarHistoricoReservas() public view returns (Reserva[] memory) {
        return historicoReservas;
    }

    function consultarHistoricoRecargas() public view returns (Recarga[] memory) {
        return historicoRecargas;
    }

    function consultarHistoricoPagamentos() public view returns (Pagamento[] memory) {
        return historicoPagamentos;
    }

    function consultarPosto(string memory _id) public view returns (Posto memory) {
        require(bytes(postos[_id].id).length > 0, "Posto inexistente.");
        return postos[_id];
    }

    function consultarTodosPostos() public view returns (string[] memory) {
        return listaIdsPostos;
    }

    // Abrir disputa por um posto
    function abrirDisputa(string memory _idPosto, string memory _motivo) public {
        require(bytes(postos[_idPosto].id).length > 0, "Posto inexistente.");
        require(disputas[_idPosto].status != StatusDisputa.Aberta, "Ja existe disputa aberta para este posto.");
        Disputa storage d = disputas[_idPosto];
        d.idPosto = _idPosto;
        d.empresaRequerente = msg.sender;
        d.motivo = _motivo;
        d.votosAFavor = 0;
        d.votosContra = 0;
        d.status = StatusDisputa.Aberta;
        d.vencedora = address(0);
        delete d.votantes;
        listaDisputas.push(_idPosto);
        emit DisputaAberta(_idPosto, msg.sender, _motivo);
    }

    // Votar em uma disputa (apenas empresas que possuem pelo menos um posto)
    function votarDisputa(string memory _idPosto, bool aFavor) public {
        require(disputas[_idPosto].status == StatusDisputa.Aberta, "Disputa nao esta aberta.");
        require(possuiPosto(msg.sender), "So empresas com posto podem votar.");
        Disputa storage d = disputas[_idPosto];
        require(!d.votou[msg.sender], "Empresa ja votou nesta disputa.");
        d.votou[msg.sender] = true;
        d.votantes.push(msg.sender);
        if (aFavor) {
            d.votosAFavor++;
        } else {
            d.votosContra++;
        }
        emit VotoRegistrado(_idPosto, msg.sender, aFavor);
    }

    // Resolver disputa (qualquer um pode chamar, resultado por maioria simples)
    function resolverDisputa(string memory _idPosto) public {
        Disputa storage d = disputas[_idPosto];
        require(d.status == StatusDisputa.Aberta, "Disputa nao esta aberta.");
        require(d.votantes.length > 0, "Nenhum voto registrado.");
        if (d.votosAFavor > d.votosContra) {
            d.vencedora = d.empresaRequerente;
        } else {
            d.vencedora = postos[_idPosto].proprietario;
        }
        d.status = StatusDisputa.Resolvida;
        emit DisputaResolvida(_idPosto, d.vencedora);
    }

    // Consultar status da disputa
    function consultarDisputa(string memory _idPosto) public view returns (
        string memory idPosto,
        address empresaRequerente,
        string memory motivo,
        uint256 votosAFavor,
        uint256 votosContra,
        StatusDisputa status,
        address vencedora,
        address[] memory votantes
    ) {
        Disputa storage d = disputas[_idPosto];
        return (
            d.idPosto,
            d.empresaRequerente,
            d.motivo,
            d.votosAFavor,
            d.votosContra,
            d.status,
            d.vencedora,
            d.votantes
        );
    }

    // Verifica se o endereço possui pelo menos um posto
    function possuiPosto(address empresa) internal view returns (bool) {
        return postosPorEmpresa[empresa].length > 0;
    }
}