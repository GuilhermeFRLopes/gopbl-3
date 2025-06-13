# Variáveis de configuração
IP   ?= 172.16.201.11
PORT ?= 8083

# Criar servidores diferentes
KEY1   ?= 90e07a66a636fb06099e14d6ba1aa7d424d6285cdab07e6337e16e3afc2a010d
KEY2   ?= 5d9313fa2f3fbcba45c25571f269613fc089f16c35bd754f872fa2e5cb4449b0
KEY3   ?= 895618e042a2397e5136e373e2038ff8e02331f23a241af925e1ad1f2a192f97

PORT1 ?= 8083
PORT2 ?= 8084
PORT3 ?= 8085

ganache:
	ganache --deterministic --accounts 10 --miner.blockTime 1

iniciarGeral:
	@echo "Compilando contrato Solidity e gerando Go bindings..."
	@solc --abi --bin contratos/Contrato.sol -o contratos/build --overwrite --evm-version london
	@abigen --abi contratos/build/Contrato.abi --bin contratos/build/Contrato.bin --pkg contratos --out contratos/contrato.go
	@echo "Contrato compilado com sucesso!"

run-server:
	@echo "Iniciando o servidor na porta ${PORT}..."
	@go run servidor/servidor-22.go --port=${PORT}

run-client-veiculo:
	@echo "Iniciando cliente de veículo, conectando a ${IP}:${PORT}..."
	@go run clienteVeiculo/cliente.go --ip=${IP} --port=${PORT}

run-client-posto:
	@echo "Iniciando cliente de posto, conectando a ${IP}:${PORT}..."
	@go run clientePosto/cliente-posto.go --ip=${IP} --port=${PORT}

deploy-contract:
	@echo "Implantando o contrato..."
	@go run deploy_contract.go


run-server1:
	@echo "Iniciando o servidor na porta ${PORT1}..."
	@go run servidor/servidor-22.go --port=${PORT1} --key=${KEY1}

run-server2:
	@echo "Iniciando o servidor na porta ${PORT2}..."
	@go run servidor/servidor-22.go --port=${PORT2} --key=${KEY2}

run-server3:
	@echo "Iniciando o servidor na porta ${PORT3}..."
	@go run servidor/servidor-22.go --port=${PORT3} --key=${KEY3}


# --- INSTRUÇÕES DE USO ---
# 1. Inicie sua blockchain local (ex: Ganache)
#
# 2. Compile o contrato:
#    make iniciarGeral
#
# 3. Implante o contrato na blockchain:
#    make deploy-contract
#    (Lembre-se de atualizar a chave privada em deploy_contract.go e servidor-22.go)
#
# 4. Inicie o servidor (pode especificar a porta):
#    make run-server
#    make run-server PORT=8084
#
# 5. Inicie os clientes (pode especificar IP e porta do servidor):
#    make run-client-veiculo IP=192.168.0.10 PORT=8084
#
#    make run-client-posto IP=192.168.0.10 PORT=8084