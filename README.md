# Sistema de Recarga de VEs Baseado em Blockchain

Este projeto implementa um sistema descentralizado para o gerenciamento de postos de recarga de veículos elétricos (VEs), utilizando a tecnologia blockchain para garantir a segurança, transparência e auditabilidade das transações. A solução foi desenvolvida como parte do Problema 3 da disciplina TEC502.

O sistema visa resolver desafios enfrentados por usuários e empresas do setor de recarga de VEs, como a garantia de registros corretos de reservas e pagamentos, a prevenção de fraudes e a integração confiável entre diferentes operadores de postos de recarga.

## 📍 Visão Geral do Projeto

O objetivo central é criar um sistema que utiliza um livro-razão distribuído (blockchain) para registrar todas as operações de forma segura e transparente. As principais operações incluem o cadastro de novos postos de recarga, a reserva de vagas por usuários de VEs, o registro do histórico de recargas e a gestão de pagamentos entre os envolvidos.

A arquitetura é genuinamente descentralizada, evitando o uso de soluções centralizadas para assegurar maior transparência e confiabilidade, conforme as restrições do problema.

## 🚀 Funcionalidades Principais

O sistema é composto por três componentes principais:

### Smart Contract (Contrato Inteligente)

O núcleo da lógica de negócios, implementado em Solidity e implantado na blockchain Ethereum. Ele é responsável por:

- **Registrar Postos**: Adicionar novos postos de recarga à rede.
- **Gerenciar Reservas**: Controlar a reserva e a liberação de vagas nos postos.
- **Consultar Dados**: Fornecer acesso público ao histórico de transações e ao estado atual dos postos.

### Servidor Backend (Go)

Atua como uma interface entre os clientes e a blockchain. Suas responsabilidades são:

- Expor uma API REST para que os clientes possam interagir com o sistema.
- Conectar-se a um nó Ethereum (como o Ganache) para enviar transações e consultar o estado do smart contract.
- Orquestrar as chamadas para as funções do contrato, como `registrarPosto` e `reservarVagaPostos`.

### Clientes (Go)

Aplicações de linha de comando que simulam a interação de usuários e empresas com o sistema:

- **Cliente de Posto (cliente-posto)**: Permite que empresas de recarga cadastrem novos postos na rede.
- **Cliente de Veículo (cliente)**: Permite que usuários de VEs consultem postos disponíveis, montem rotas e reservem vagas para recarga.

## 📁 Estrutura do Projeto

```
/
|-- clientePosto/
|   |-- cliente-posto.go     # Aplicação cliente para cadastrar postos de recarga.
|-- clienteVeiculo/
|   |-- cliente.go           # Aplicação cliente para usuários de VEs (consultar, reservar).
|-- contratos/
|   |-- Contrato.sol         # Código-fonte do Smart Contract em Solidity (arquivo original).
|   |-- contrato.go          # Binding do contrato gerado pelo Abigen para uso em Go.
|   |-- build/
|       |-- Contrato.abi     # ABI (Application Binary Interface) do contrato.
|       |-- Contrato.bin     # Bytecode compilado do contrato.
|-- servidor/
|   |-- servidor-22.go       # Implementação do servidor backend em Go (API REST).
|-- debug/
|   |-- debug_tx.go          # Script para depurar transações na blockchain.
|-- modelo/                  # Structs de dados compartilhadas (Posto, Veiculo, etc.).
|-- deploy_contract.go       # Script para implantar o smart contract na blockchain.
|-- makefile                 # Comandos para compilar e gerar o binding do contrato.
|-- contract_address.txt     # Armazena o endereço do contrato implantado na rede.
`-- comandos.txt             # Histórico de comandos útis para desenvolvimento e execução.
```

## 🛠️ Como Executar (Ambiente Local sem Docker)

### Pré-requisitos

- Go (versão 1.18 ou superior).
- Node.js e npm para instalar o Ganache.
- Ganache CLI: `npm install -g ganache`
- Solidity Compiler (`solc`) (versão 0.8.30 ou compatível).
- Abigen, uma ferramenta do Go-Ethereum.

### 1. Iniciar a Rede Blockchain Local

```bash
ganache
```

Anote uma das chaves privadas exibidas. O Ganache estará rodando em `http://127.0.0.1:7545`, é necessário ajustar o ganache para rodar em `http://0.0.0.0:7545` no laboratorio. Lembrando que o endereço `http://0.0.0.0:7545`, em caso de rodar um servidor na mesma máquina onde o ganache está rodando. Caso deseje rodar o servidor em outra máquina, a qual o ganache não está presente, é necessário ajustar o endereço para `http://(endereço de ip da máquina onde o ganache está rodando):7545`.

Por exemplo, se o ganache for iniciado na máquina com endereço `172.16.201.11`, então os ip's de destino do ganhace dos servidores que estão rodando em outras máquinas devem ser 
`http://127.0.0.1:7545`. Isso é definido na função 'conectarEthereum' que os servidores contém.

### 2. Compilar o Smart Contract e Gerar o Binding

```bash
solc --abi --bin contratos/Contrato.sol -o contratos/build --overwrite --evm-version london
abigen --abi contratos/build/Contrato.abi --bin contratos/build/Contrato.bin --pkg contratos --out contratos/contrato.go
```

### 3. Configurar e Implantar o Smart Contract

Edite o ficheiro `deploy_contract.go`:

- Substitua `privateKeyHex` pela chave privada copiada do Ganache.
- Verifique se `infuraURL` está como `HTTP://127.0.0.1:7545` (ou a porta e ip corretos).

Execute:

```bash
go run deploy_contract.go
```

O endereço do contrato será mostrado e salvo em `contract_address.txt`.

### 4. Configurar e Iniciar o Servidor

Edite `servidor/servidor-22.go`:

- Cole o endereço do contrato na variável `contractAddress`.
- Use uma chave privada diferente da implantação para `privateKeyHex`.
- Confirme o `ethClient.Dial` com `HTTP://127.0.0.1:7545` ou o adequado.

Execute:

```bash
go run servidor/servidor-22.go
```

Servidor na porta `:8083`.

para rodar outros servidores é necessario mudar a porta, 8083,8084,8085 etc.

De outra maneira, os servidores também podem ser executados pelo makefile:
```bash
# Para que isso funcione corretamente, as variáveis KEY1, KEY2 e KEY3,
# além de PORT1, PORT2 e PORT3 devem ser alteradas no próprio makefile,
# as quais devem ser especificadas chaves privadas e portas diferentes.
# Feito isso, os seguintes comandos podem ser executados:

    make run-server1
    make run-server2
    make run-server3
```

### 5. Configurar e Executar os Clientes

Verifique `servidorBackendURL` em:

- `clientePosto/cliente-posto.go`
- `clienteVeiculo/cliente.go`

Execute os clientes:

```bash
# Cliente de Posto
go run clientePosto/cliente-posto.go


# Cliente de Veículo
go run clienteVeiculo/cliente.go
```
 De outra maneira, os clientes também podem ser executados pelo makefile:
```bash
# Deve ser especificado o IP e a porta do servidor que deseja se conectar.
make run-client-veiculo IP=192.168.0.10 PORT=8084 


make run-client-posto IP=192.168.0.10 PORT=8084
```

---



