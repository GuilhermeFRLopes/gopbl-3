ganache:
	ganache --deterministic --accounts 10 --miner.blockTime 1

iniciarGeral:
	solc --abi --bin contratos/Contrato.sol -o contratos/build --overwrite --evm-version london
	abigen --abi contratos/build/Contrato.abi --bin contratos/build/Contrato.bin --pkg contratos --out contratos/contrato.go


# 1- iniciar geral
# 2- altere a chave privada do deploy_contract e servidor
# 3- go run deploy_contract
# 4- go run servidor