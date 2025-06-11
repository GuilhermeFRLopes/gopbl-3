# Variáveis de configuração
SERVER_IP   ?= 127.0.0.1
SERVER_PORT ?= 8083

iniciarGeral:
    solc --abi --bin contratos/Contrato.sol -o contratos/build --overwrite --evm-version london
    abigen --abi contratos/build/Contrato.abi --bin contratos/build/Contrato.bin --pkg contratos --out contratos/contrato.go

run-server:
	@echo "Iniciando o servidor na porta ${SERVER_PORT}..."
	@go run servidor/servidor-22.go --port=${SERVER_PORT}

run-client-veiculo:
	@echo "Iniciando cliente de veículo, conectando a ${SERVER_IP}:${SERVER_PORT}..."
	@go run clienteVeiculo/cliente.go --ip=${SERVER_IP} --port=${SERVER_PORT}

run-client-posto:
	@echo "Iniciando cliente de posto, conectando a ${SERVER_IP}:${SERVER_PORT}..."
	@go run clientePosto/cliente-posto.go --ip=${SERVER_IP} --port=${SERVER_PORT}


# --- INSTRUÇÕES DE USO ---
# 1. iniciar ganache

# 2. Implante o contrato na blockchain (e.g., Ganache):
#    go run deploy_contract.go  (Lembre-se de atualizar a chave privada antes)
#
# 3. Inicie o servidor (pode especificar a porta):
#    make run-server
#    make run-server PORT=8084
#
# 4. Inicie os clientes (pode especificar IP e porta do servidor):
#    make run-client-veiculo
#    make run-client-veiculo IP=192.168.0.10 PORT=8084
#
#    make run-client-posto
#    make run-client-posto IP=192.168.0.10 PORT=8084