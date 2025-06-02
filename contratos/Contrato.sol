// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;
//estrutura do contrato
contract Contrato {
    //estrutura do posto
    struct PostoRecarga{
        string id;
        int256 latitude;
        int256 longitude;
        bool ocupado;
        string servidorOrigem;
        string cidade; 
        string reservador;       
    }

    constructor(){
        //inicializar
    }

    //dicionario dos postos, chave é o id
    mapping(string => PostoRecarga) public postosRecarga;

    //lista dos ids dos postos
    string[] public idsPostos;

    //eventos
    event PostoRegistrado(string indexed idPosto, string cidade, string servidorOrigem);
    event AtualDisponibilidadePosto(string indexed idPosto, bool disponivel, string servidorOrigem);
    

    function registrarPosto(string memory _id, int256 _latitude, int256 _longitude, string memory _servidorOrigem, string memory _cidade) public {
        //verifica se ja nao tem posto com esse id
        require(bytes(postosRecarga[_id].id).length == 0, "Posto de recarga ja cadastrado");
        postosRecarga[_id] = PostoRecarga(_id,_latitude,_longitude,false,_servidorOrigem,_cidade,"");
        idsPostos.push(_id);
        emit PostoRegistrado(_id, _cidade, _servidorOrigem);
    }

    // function getReservador(string memory _idP) public view returns (string memory){
    //     PostoRecarga storage p = idsPostos[_idP];        
    //     return  p.reservador;
    // }

    //funçao que reserva ou libera os postos
   function reservarVagaPostos(string[] memory _idPostos, bool _reservar, string memory _reservador) public {
        require(_idPostos.length > 0, "Nenhum ID de posto fornecido.");

        for (uint i = 0; i < _idPostos.length; i++) {
            string memory idDoPostoAtual = _idPostos[i]; // Pega o ID da lista que veio na chamada
            PostoRecarga storage posto = postosRecarga[idDoPostoAtual]; // Acessa o posto pelo ID

            require(bytes(posto.id).length > 0, "Posto nao existe."); // Garante que o posto existe

            if (_reservar) {
                // Logica para reservar:
                // O posto DEVE estar NAO OCUPADO para ser reservado.
                require(posto.ocupado == false, "Posto ja esta ocupado.");
                posto.ocupado = true;
                posto.reservador = _reservador;
                emit AtualDisponibilidadePosto(posto.id, posto.ocupado, posto.reservador); // Emite evento

            } else {
                // Logica para liberar:
                // O posto DEVE estar OCUPADO para ser liberado.
                require(posto.ocupado == true, "Posto ja esta disponivel.");                
                require(keccak256(bytes(posto.reservador)) == keccak256(bytes(_reservador)), "Apenas o reservador pode liberar este posto.");
                posto.ocupado = false;
                posto.reservador = ""; // Limpa o reservador
                emit AtualDisponibilidadePosto(posto.id, posto.ocupado, posto.reservador); // Emite evento
            }
        }
    }

    function verificarDisponibilidade(string memory _idPosto) public view returns (bool){
        PostoRecarga storage posto = postosRecarga[_idPosto];
        return posto.ocupado;
    }

    function getAllPostos() public view returns (string[] memory){
        return idsPostos;
    }


}