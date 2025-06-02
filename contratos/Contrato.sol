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
    function reservarVagaPostos(string[] memory _idPostos, bool _reservar, string memory _reservador) public{
        if (!_reservar){
            //logica para liberar os postos
            for (uint i = 0; i < idsPostos.length; i++){
                string memory p = idsPostos[i];
                PostoRecarga storage posto = postosRecarga[p];
                if (posto.ocupado){
                    if (keccak256(bytes(posto.reservador)) == keccak256(bytes(_reservador))){
                        posto.ocupado = false;
                        posto.reservador = _reservador;
                        //liberou o posto
                        emit AtualDisponibilidadePosto(posto.id, posto.ocupado, posto.servidorOrigem);
                    }
                } 
                // emit AtualDisponibilidadePosto(idPosto, disponivel, servidorOrigem);
            }
        } else {
            //logica para reservar
            for (uint i = 0; i < _idPostos.length; i++){
                PostoRecarga storage posto = postosRecarga[(_idPostos[i])];
                require(posto.ocupado == true, "Posto ja esta reservado");
                posto.ocupado = true;
                posto.reservador =  _reservador;
                //reservou o posto
                emit AtualDisponibilidadePosto(posto.id, posto.ocupado, posto.servidorOrigem);
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