// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ChargingStationManager {
    // Estrutura para representar um Posto de Recarga
    struct ChargingStation {
        string id;             // ID único do posto (ex: "posto-shell-001")
        string nome;           // Nome/descrição
        string cidade;         // Cidade do posto
        address proprietario;  // Endereço da conta que cadastrou/gerencia o posto
        bool disponivel;       // true se disponível, false se reservado/ocupado
        uint256 ultimaAtualizacao; // Timestamp da última atualização (Unix)
        address veiculoReservador; // Endereço do veículo que reservou o posto (address(0) se não reservado)
    }

    // Mapeamento de ID do posto para sua estrutura ChargingStation
    mapping(string => ChargingStation) public chargingStations;

    // Array para armazenar todos os IDs de postos cadastrados (para iteração)
    string[] public stationIds;

    // Eventos para notificar sobre mudanças no estado da blockchain
    event StationRegistered(string indexed stationId, address indexed owner, string city);
    event StationAvailabilityUpdated(string indexed stationId, bool newAvailability, address indexed reservor, uint256 timestamp);
    event StationDeleted(string indexed stationId);

    // Constructor (opcional, para inicialização)
    constructor() {
        // Inicializações podem ser feitas aqui, se necessário
    }

    // Função para cadastrar um novo posto de recarga
    function registerStation(string memory _id, string memory _nome, string memory _cidade) public {
        require(bytes(chargingStations[_id].id).length == 0, "Charging station already registered."); // Verifica se o ID já existe
        
        stationIds.push(_id); // Adiciona o ID ao array de IDs
        chargingStations[_id] = ChargingStation(
            _id,
            _nome,
            _cidade,
            msg.sender, // Quem chamou a função é o proprietário
            true,       // Inicialmente disponível
            block.timestamp,
            address(0) // Ninguém reservou ainda
        );

        emit StationRegistered(_id, msg.sender, _cidade);
    }

    // Função para reservar/liberar um posto de recarga
    function updateStationAvailability(string[] memory _stationIds, bool _reservar, address _veiculoId) public {
        require(_stationIds.length > 0, "No station IDs provided.");

        for (uint i = 0; i < _stationIds.length; i++) {
            string memory stationId = _stationIds[i];
            ChargingStation storage station = chargingStations[stationId];

            require(bytes(station.id).length > 0, "Station does not exist."); // Verifica se o posto existe

            if (_reservar) {
                // Lógica para reservar
                require(station.disponivel == true, "Station is already reserved.");
                station.disponivel = false;
                station.veiculoReservador = _veiculoId;
            } else {
                // Lógica para liberar
                require(station.disponivel == false, "Station is already available.");
                // Opcional: require(station.veiculoReservador == _veiculoId, "Only the reserving vehicle can release.");
                station.disponivel = true;
                station.veiculoReservador = address(0);
            }
            station.ultimaAtualizacao = block.timestamp;

            emit StationAvailabilityUpdated(stationId, station.disponivel, _veiculoId, block.timestamp);
        }
    }

    // Função para consultar a disponibilidade de um posto (view function, não custa gás)
    function getStationAvailability(string memory _id) public view returns (bool, address, uint256) {
        ChargingStation storage station = chargingStations[_id];
        require(bytes(station.id).length > 0, "Station does not exist.");
        return (station.disponivel, station.veiculoReservador, station.ultimaAtualizacao);
    }

    // Função para obter todos os IDs de postos (para facilitar a consulta off-chain)
    function getAllStationIds() public view returns (string[] memory) {
        return stationIds;
    }

    // Função para deletar um posto (opcional, com controle de acesso)
    function deleteStation(string memory _id) public {
        // Exemplo: Somente o proprietário ou um admin pode deletar
        require(msg.sender == chargingStations[_id].proprietario, "Only the owner can delete this station.");
        
        // Remove do array de IDs
        for (uint i = 0; i < stationIds.length; i++) {
            if (keccak256(abi.encodePacked(stationIds[i])) == keccak256(abi.encodePacked(_id))) {
                stationIds[i] = stationIds[stationIds.length - 1];
                stationIds.pop();
                break;
            }
        }
        delete chargingStations[_id]; // Remove do mapeamento
        emit StationDeleted(_id);
    }
}