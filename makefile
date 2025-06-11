# Variáveis de configuração
IP   ?= 127.0.0.1
PORT ?= 8083

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