// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"stationId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newAvailability\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reservor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StationAvailabilityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"stationId\",\"type\":\"string\"}],\"name\":\"StationDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"stationId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"city\",\"type\":\"string\"}],\"name\":\"StationRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"chargingStations\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nome\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cidade\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proprietario\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"disponivel\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"ultimaAtualizacao\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"veiculoReservador\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"deleteStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllStationIds\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getStationAvailability\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nome\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cidade\",\"type\":\"string\"}],\"name\":\"registerStation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stationIds\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_stationIds\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"_reservar\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_veiculoId\",\"type\":\"address\"}],\"name\":\"updateStationAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50611dfb8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c80639781e91b116100595780639781e91b146100e9578063a836e3b214610107578063e391daf91461013d578063e701768e1461016d5761007b565b806337fa26561461007f5780635d7e2e301461009b5780636a4e3b04146100cd575b5f5ffd5b61009960048036038101906100949190610fbd565b610189565b005b6100b560048036038101906100b09190611061565b610409565b6040516100c493929190611119565b60405180910390f35b6100e760048036038101906100e29190611061565b6104c6565b005b6100f161077f565b6040516100fe9190611269565b60405180910390f35b610121600480360381019061011c9190611061565b610853565b60405161013497969594939291906112d1565b60405180910390f35b6101576004803603810190610152919061137d565b610a85565b60405161016491906113a8565b60405180910390f35b610187600480360381019061018291906114fe565b610b2b565b005b5f5f8460405161019991906115a4565b90815260200160405180910390205f0180546101b4906115e7565b9050146101f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ed90611687565b60405180910390fd5b600183908060018154018082558091505060019003905f5260205f20015f9091909190915090816102279190611845565b506040518060e001604052808481526020018381526020018281526020013373ffffffffffffffffffffffffffffffffffffffff1681526020016001151581526020014281526020015f73ffffffffffffffffffffffffffffffffffffffff168152505f8460405161029991906115a4565b90815260200160405180910390205f820151815f0190816102ba9190611845565b5060208201518160010190816102d09190611845565b5060408201518160020190816102e69190611845565b506060820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160146101000a81548160ff02191690831515021790555060a0820151816004015560c0820151816005015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff16836040516103c591906115a4565b60405180910390207fed098a22e0ff25594835cdcd04fe2ec2bb711e3b7cb28d8fc4f02aa153459f2b836040516103fc91906113a8565b60405180910390a3505050565b5f5f5f5f5f8560405161041c91906115a4565b908152602001604051809103902090505f815f01805461043b906115e7565b90501161047d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104749061195e565b60405180910390fd5b8060030160149054906101000a900460ff16816005015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260040154935093509350509193909250565b5f816040516104d591906115a4565b90815260200160405180910390206003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610573576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056a906119ec565b60405180910390fd5b5f5f90505b600180549050811015610686578160405160200161059691906115a4565b60405160208183030381529060405280519060200120600182815481106105c0576105bf611a0a565b5b905f5260205f20016040516020016105d89190611ab7565b60405160208183030381529060405280519060200120036106795760018080805490506106059190611afa565b8154811061061657610615611a0a565b5b905f5260205f20016001828154811061063257610631611a0a565b5b905f5260205f200190816106469190611b54565b50600180548061065957610658611c39565b5b600190038181905f5260205f20015f6106729190610e18565b9055610686565b8080600101915050610578565b505f8160405161069691906115a4565b90815260200160405180910390205f5f82015f6106b39190610e18565b600182015f6106c29190610e18565b600282015f6106d19190610e18565b600382015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160146101000a81549060ff0219169055600482015f9055600582015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550508060405161074891906115a4565b60405180910390207fd7468e363ec0597d4d2d75dcac9670a851a25bfec8c52379cbbeedd59b73c6c160405160405180910390a250565b60606001805480602002602001604051908101604052809291908181526020015f905b8282101561084a578382905f5260205f200180546107bf906115e7565b80601f01602080910402602001604051908101604052809291908181526020018280546107eb906115e7565b80156108365780601f1061080d57610100808354040283529160200191610836565b820191905f5260205f20905b81548152906001019060200180831161081957829003601f168201915b5050505050815260200190600101906107a2565b50505050905090565b5f818051602081018201805184825260208301602085012081835280955050505050505f91509050805f018054610889906115e7565b80601f01602080910402602001604051908101604052809291908181526020018280546108b5906115e7565b80156109005780601f106108d757610100808354040283529160200191610900565b820191905f5260205f20905b8154815290600101906020018083116108e357829003601f168201915b505050505090806001018054610915906115e7565b80601f0160208091040260200160405190810160405280929190818152602001828054610941906115e7565b801561098c5780601f106109635761010080835404028352916020019161098c565b820191905f5260205f20905b81548152906001019060200180831161096f57829003601f168201915b5050505050908060020180546109a1906115e7565b80601f01602080910402602001604051908101604052809291908181526020018280546109cd906115e7565b8015610a185780601f106109ef57610100808354040283529160200191610a18565b820191905f5260205f20905b8154815290600101906020018083116109fb57829003601f168201915b505050505090806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160149054906101000a900460ff1690806004015490806005015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905087565b60018181548110610a94575f80fd5b905f5260205f20015f915090508054610aac906115e7565b80601f0160208091040260200160405190810160405280929190818152602001828054610ad8906115e7565b8015610b235780601f10610afa57610100808354040283529160200191610b23565b820191905f5260205f20905b815481529060010190602001808311610b0657829003601f168201915b505050505081565b5f835111610b6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6590611cb0565b60405180910390fd5b5f5f90505b8351811015610e12575f848281518110610b9057610b8f611a0a565b5b602002602001015190505f5f82604051610baa91906115a4565b908152602001604051809103902090505f815f018054610bc9906115e7565b905011610c0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c029061195e565b60405180910390fd5b8415610ccc57600115158160030160149054906101000a900460ff16151514610c69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6090611d18565b60405180910390fd5b5f8160030160146101000a81548160ff02191690831515021790555083816005015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610d83565b5f15158160030160149054906101000a900460ff16151514610d23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d1a90611d80565b60405180910390fd5b60018160030160146101000a81548160ff0219169083151502179055505f816005015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b4281600401819055508373ffffffffffffffffffffffffffffffffffffffff1682604051610db191906115a4565b60405180910390207f9f45fd480a7696877da95e871642a241d6b0a68fb0fe585a562ca074de2c21458360030160149054906101000a900460ff1642604051610dfb929190611d9e565b60405180910390a350508080600101915050610b73565b50505050565b508054610e24906115e7565b5f825580601f10610e355750610e52565b601f0160209004905f5260205f2090810190610e519190610e55565b5b50565b5b80821115610e6c575f815f905550600101610e56565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610ecf82610e89565b810181811067ffffffffffffffff82111715610eee57610eed610e99565b5b80604052505050565b5f610f00610e70565b9050610f0c8282610ec6565b919050565b5f67ffffffffffffffff821115610f2b57610f2a610e99565b5b610f3482610e89565b9050602081019050919050565b828183375f83830152505050565b5f610f61610f5c84610f11565b610ef7565b905082815260208101848484011115610f7d57610f7c610e85565b5b610f88848285610f41565b509392505050565b5f82601f830112610fa457610fa3610e81565b5b8135610fb4848260208601610f4f565b91505092915050565b5f5f5f60608486031215610fd457610fd3610e79565b5b5f84013567ffffffffffffffff811115610ff157610ff0610e7d565b5b610ffd86828701610f90565b935050602084013567ffffffffffffffff81111561101e5761101d610e7d565b5b61102a86828701610f90565b925050604084013567ffffffffffffffff81111561104b5761104a610e7d565b5b61105786828701610f90565b9150509250925092565b5f6020828403121561107657611075610e79565b5b5f82013567ffffffffffffffff81111561109357611092610e7d565b5b61109f84828501610f90565b91505092915050565b5f8115159050919050565b6110bc816110a8565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6110eb826110c2565b9050919050565b6110fb816110e1565b82525050565b5f819050919050565b61111381611101565b82525050565b5f60608201905061112c5f8301866110b3565b61113960208301856110f2565b611146604083018461110a565b949350505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6111a982611177565b6111b38185611181565b93506111c3818560208601611191565b6111cc81610e89565b840191505092915050565b5f6111e2838361119f565b905092915050565b5f602082019050919050565b5f6112008261114e565b61120a8185611158565b93508360208202850161121c85611168565b805f5b85811015611257578484038952815161123885826111d7565b9450611243836111ea565b925060208a0199505060018101905061121f565b50829750879550505050505092915050565b5f6020820190508181035f83015261128181846111f6565b905092915050565b5f82825260208201905092915050565b5f6112a382611177565b6112ad8185611289565b93506112bd818560208601611191565b6112c681610e89565b840191505092915050565b5f60e0820190508181035f8301526112e9818a611299565b905081810360208301526112fd8189611299565b905081810360408301526113118188611299565b905061132060608301876110f2565b61132d60808301866110b3565b61133a60a083018561110a565b61134760c08301846110f2565b98975050505050505050565b61135c81611101565b8114611366575f5ffd5b50565b5f8135905061137781611353565b92915050565b5f6020828403121561139257611391610e79565b5b5f61139f84828501611369565b91505092915050565b5f6020820190508181035f8301526113c08184611299565b905092915050565b5f67ffffffffffffffff8211156113e2576113e1610e99565b5b602082029050602081019050919050565b5f5ffd5b5f611409611404846113c8565b610ef7565b9050808382526020820190506020840283018581111561142c5761142b6113f3565b5b835b8181101561147357803567ffffffffffffffff81111561145157611450610e81565b5b80860161145e8982610f90565b8552602085019450505060208101905061142e565b5050509392505050565b5f82601f83011261149157611490610e81565b5b81356114a18482602086016113f7565b91505092915050565b6114b3816110a8565b81146114bd575f5ffd5b50565b5f813590506114ce816114aa565b92915050565b6114dd816110e1565b81146114e7575f5ffd5b50565b5f813590506114f8816114d4565b92915050565b5f5f5f6060848603121561151557611514610e79565b5b5f84013567ffffffffffffffff81111561153257611531610e7d565b5b61153e8682870161147d565b935050602061154f868287016114c0565b9250506040611560868287016114ea565b9150509250925092565b5f81905092915050565b5f61157e82611177565b611588818561156a565b9350611598818560208601611191565b80840191505092915050565b5f6115af8284611574565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806115fe57607f821691505b602082108103611611576116106115ba565b5b50919050565b7f4368617267696e672073746174696f6e20616c726561647920726567697374655f8201527f7265642e00000000000000000000000000000000000000000000000000000000602082015250565b5f611671602483611289565b915061167c82611617565b604082019050919050565b5f6020820190508181035f83015261169e81611665565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026117017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826116c6565b61170b86836116c6565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61174661174161173c84611101565b611723565b611101565b9050919050565b5f819050919050565b61175f8361172c565b61177361176b8261174d565b8484546116d2565b825550505050565b5f5f905090565b61178a61177b565b611795818484611756565b505050565b5b818110156117b8576117ad5f82611782565b60018101905061179b565b5050565b601f8211156117fd576117ce816116a5565b6117d7846116b7565b810160208510156117e6578190505b6117fa6117f2856116b7565b83018261179a565b50505b505050565b5f82821c905092915050565b5f61181d5f1984600802611802565b1980831691505092915050565b5f611835838361180e565b9150826002028217905092915050565b61184e82611177565b67ffffffffffffffff81111561186757611866610e99565b5b61187182546115e7565b61187c8282856117bc565b5f60209050601f8311600181146118ad575f841561189b578287015190505b6118a5858261182a565b86555061190c565b601f1984166118bb866116a5565b5f5b828110156118e2578489015182556001820191506020850194506020810190506118bd565b868310156118ff57848901516118fb601f89168261180e565b8355505b6001600288020188555050505b505050505050565b7f53746174696f6e20646f6573206e6f742065786973742e0000000000000000005f82015250565b5f611948601783611289565b915061195382611914565b602082019050919050565b5f6020820190508181035f8301526119758161193c565b9050919050565b7f4f6e6c7920746865206f776e65722063616e2064656c657465207468697320735f8201527f746174696f6e2e00000000000000000000000000000000000000000000000000602082015250565b5f6119d6602783611289565b91506119e18261197c565b604082019050919050565b5f6020820190508181035f830152611a03816119ca565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8154611a43816115e7565b611a4d818661156a565b9450600182165f8114611a675760018114611a7c57611aae565b60ff1983168652811515820286019350611aae565b611a85856116a5565b5f5b83811015611aa657815481890152600182019150602081019050611a87565b838801955050505b50505092915050565b5f611ac28284611a37565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611b0482611101565b9150611b0f83611101565b9250828203905081811115611b2757611b26611acd565b5b92915050565b5f81549050611b3b816115e7565b9050919050565b5f819050815f5260205f209050919050565b818103611b62575050611c37565b611b6b82611b2d565b67ffffffffffffffff811115611b8457611b83610e99565b5b611b8e82546115e7565b611b998282856117bc565b5f601f831160018114611bc6575f8415611bb4578287015490505b611bbe858261182a565b865550611c30565b601f198416611bd487611b42565b9650611bdf866116a5565b5f5b82811015611c0657848901548255600182019150600185019450602081019050611be1565b86831015611c235784890154611c1f601f89168261180e565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e6f2073746174696f6e204944732070726f76696465642e00000000000000005f82015250565b5f611c9a601883611289565b9150611ca582611c66565b602082019050919050565b5f6020820190508181035f830152611cc781611c8e565b9050919050565b7f53746174696f6e20697320616c72656164792072657365727665642e000000005f82015250565b5f611d02601c83611289565b9150611d0d82611cce565b602082019050919050565b5f6020820190508181035f830152611d2f81611cf6565b9050919050565b7f53746174696f6e20697320616c726561647920617661696c61626c652e0000005f82015250565b5f611d6a601d83611289565b9150611d7582611d36565b602082019050919050565b5f6020820190508181035f830152611d9781611d5e565b9050919050565b5f604082019050611db15f8301856110b3565b611dbe602083018461110a565b939250505056fea2646970667358221220a1949b22d79d084375fad675ec3786048b9260a812e442d6bd47eeb2012c6c9b64736f6c634300081e0033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// ChargingStations is a free data retrieval call binding the contract method 0xa836e3b2.
//
// Solidity: function chargingStations(string ) view returns(string id, string nome, string cidade, address proprietario, bool disponivel, uint256 ultimaAtualizacao, address veiculoReservador)
func (_Contracts *ContractsCaller) ChargingStations(opts *bind.CallOpts, arg0 string) (struct {
	Id                string
	Nome              string
	Cidade            string
	Proprietario      common.Address
	Disponivel        bool
	UltimaAtualizacao *big.Int
	VeiculoReservador common.Address
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "chargingStations", arg0)

	outstruct := new(struct {
		Id                string
		Nome              string
		Cidade            string
		Proprietario      common.Address
		Disponivel        bool
		UltimaAtualizacao *big.Int
		VeiculoReservador common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Nome = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Cidade = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Proprietario = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Disponivel = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.UltimaAtualizacao = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.VeiculoReservador = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ChargingStations is a free data retrieval call binding the contract method 0xa836e3b2.
//
// Solidity: function chargingStations(string ) view returns(string id, string nome, string cidade, address proprietario, bool disponivel, uint256 ultimaAtualizacao, address veiculoReservador)
func (_Contracts *ContractsSession) ChargingStations(arg0 string) (struct {
	Id                string
	Nome              string
	Cidade            string
	Proprietario      common.Address
	Disponivel        bool
	UltimaAtualizacao *big.Int
	VeiculoReservador common.Address
}, error) {
	return _Contracts.Contract.ChargingStations(&_Contracts.CallOpts, arg0)
}

// ChargingStations is a free data retrieval call binding the contract method 0xa836e3b2.
//
// Solidity: function chargingStations(string ) view returns(string id, string nome, string cidade, address proprietario, bool disponivel, uint256 ultimaAtualizacao, address veiculoReservador)
func (_Contracts *ContractsCallerSession) ChargingStations(arg0 string) (struct {
	Id                string
	Nome              string
	Cidade            string
	Proprietario      common.Address
	Disponivel        bool
	UltimaAtualizacao *big.Int
	VeiculoReservador common.Address
}, error) {
	return _Contracts.Contract.ChargingStations(&_Contracts.CallOpts, arg0)
}

// GetAllStationIds is a free data retrieval call binding the contract method 0x9781e91b.
//
// Solidity: function getAllStationIds() view returns(string[])
func (_Contracts *ContractsCaller) GetAllStationIds(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllStationIds")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllStationIds is a free data retrieval call binding the contract method 0x9781e91b.
//
// Solidity: function getAllStationIds() view returns(string[])
func (_Contracts *ContractsSession) GetAllStationIds() ([]string, error) {
	return _Contracts.Contract.GetAllStationIds(&_Contracts.CallOpts)
}

// GetAllStationIds is a free data retrieval call binding the contract method 0x9781e91b.
//
// Solidity: function getAllStationIds() view returns(string[])
func (_Contracts *ContractsCallerSession) GetAllStationIds() ([]string, error) {
	return _Contracts.Contract.GetAllStationIds(&_Contracts.CallOpts)
}

// GetStationAvailability is a free data retrieval call binding the contract method 0x5d7e2e30.
//
// Solidity: function getStationAvailability(string _id) view returns(bool, address, uint256)
func (_Contracts *ContractsCaller) GetStationAvailability(opts *bind.CallOpts, _id string) (bool, common.Address, *big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getStationAvailability", _id)

	if err != nil {
		return *new(bool), *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStationAvailability is a free data retrieval call binding the contract method 0x5d7e2e30.
//
// Solidity: function getStationAvailability(string _id) view returns(bool, address, uint256)
func (_Contracts *ContractsSession) GetStationAvailability(_id string) (bool, common.Address, *big.Int, error) {
	return _Contracts.Contract.GetStationAvailability(&_Contracts.CallOpts, _id)
}

// GetStationAvailability is a free data retrieval call binding the contract method 0x5d7e2e30.
//
// Solidity: function getStationAvailability(string _id) view returns(bool, address, uint256)
func (_Contracts *ContractsCallerSession) GetStationAvailability(_id string) (bool, common.Address, *big.Int, error) {
	return _Contracts.Contract.GetStationAvailability(&_Contracts.CallOpts, _id)
}

// StationIds is a free data retrieval call binding the contract method 0xe391daf9.
//
// Solidity: function stationIds(uint256 ) view returns(string)
func (_Contracts *ContractsCaller) StationIds(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "stationIds", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StationIds is a free data retrieval call binding the contract method 0xe391daf9.
//
// Solidity: function stationIds(uint256 ) view returns(string)
func (_Contracts *ContractsSession) StationIds(arg0 *big.Int) (string, error) {
	return _Contracts.Contract.StationIds(&_Contracts.CallOpts, arg0)
}

// StationIds is a free data retrieval call binding the contract method 0xe391daf9.
//
// Solidity: function stationIds(uint256 ) view returns(string)
func (_Contracts *ContractsCallerSession) StationIds(arg0 *big.Int) (string, error) {
	return _Contracts.Contract.StationIds(&_Contracts.CallOpts, arg0)
}

// DeleteStation is a paid mutator transaction binding the contract method 0x6a4e3b04.
//
// Solidity: function deleteStation(string _id) returns()
func (_Contracts *ContractsTransactor) DeleteStation(opts *bind.TransactOpts, _id string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deleteStation", _id)
}

// DeleteStation is a paid mutator transaction binding the contract method 0x6a4e3b04.
//
// Solidity: function deleteStation(string _id) returns()
func (_Contracts *ContractsSession) DeleteStation(_id string) (*types.Transaction, error) {
	return _Contracts.Contract.DeleteStation(&_Contracts.TransactOpts, _id)
}

// DeleteStation is a paid mutator transaction binding the contract method 0x6a4e3b04.
//
// Solidity: function deleteStation(string _id) returns()
func (_Contracts *ContractsTransactorSession) DeleteStation(_id string) (*types.Transaction, error) {
	return _Contracts.Contract.DeleteStation(&_Contracts.TransactOpts, _id)
}

// RegisterStation is a paid mutator transaction binding the contract method 0x37fa2656.
//
// Solidity: function registerStation(string _id, string _nome, string _cidade) returns()
func (_Contracts *ContractsTransactor) RegisterStation(opts *bind.TransactOpts, _id string, _nome string, _cidade string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerStation", _id, _nome, _cidade)
}

// RegisterStation is a paid mutator transaction binding the contract method 0x37fa2656.
//
// Solidity: function registerStation(string _id, string _nome, string _cidade) returns()
func (_Contracts *ContractsSession) RegisterStation(_id string, _nome string, _cidade string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterStation(&_Contracts.TransactOpts, _id, _nome, _cidade)
}

// RegisterStation is a paid mutator transaction binding the contract method 0x37fa2656.
//
// Solidity: function registerStation(string _id, string _nome, string _cidade) returns()
func (_Contracts *ContractsTransactorSession) RegisterStation(_id string, _nome string, _cidade string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterStation(&_Contracts.TransactOpts, _id, _nome, _cidade)
}

// UpdateStationAvailability is a paid mutator transaction binding the contract method 0xe701768e.
//
// Solidity: function updateStationAvailability(string[] _stationIds, bool _reservar, address _veiculoId) returns()
func (_Contracts *ContractsTransactor) UpdateStationAvailability(opts *bind.TransactOpts, _stationIds []string, _reservar bool, _veiculoId common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateStationAvailability", _stationIds, _reservar, _veiculoId)
}

// UpdateStationAvailability is a paid mutator transaction binding the contract method 0xe701768e.
//
// Solidity: function updateStationAvailability(string[] _stationIds, bool _reservar, address _veiculoId) returns()
func (_Contracts *ContractsSession) UpdateStationAvailability(_stationIds []string, _reservar bool, _veiculoId common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateStationAvailability(&_Contracts.TransactOpts, _stationIds, _reservar, _veiculoId)
}

// UpdateStationAvailability is a paid mutator transaction binding the contract method 0xe701768e.
//
// Solidity: function updateStationAvailability(string[] _stationIds, bool _reservar, address _veiculoId) returns()
func (_Contracts *ContractsTransactorSession) UpdateStationAvailability(_stationIds []string, _reservar bool, _veiculoId common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateStationAvailability(&_Contracts.TransactOpts, _stationIds, _reservar, _veiculoId)
}

// ContractsStationAvailabilityUpdatedIterator is returned from FilterStationAvailabilityUpdated and is used to iterate over the raw logs and unpacked data for StationAvailabilityUpdated events raised by the Contracts contract.
type ContractsStationAvailabilityUpdatedIterator struct {
	Event *ContractsStationAvailabilityUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsStationAvailabilityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsStationAvailabilityUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsStationAvailabilityUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsStationAvailabilityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsStationAvailabilityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsStationAvailabilityUpdated represents a StationAvailabilityUpdated event raised by the Contracts contract.
type ContractsStationAvailabilityUpdated struct {
	StationId       common.Hash
	NewAvailability bool
	Reservor        common.Address
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStationAvailabilityUpdated is a free log retrieval operation binding the contract event 0x9f45fd480a7696877da95e871642a241d6b0a68fb0fe585a562ca074de2c2145.
//
// Solidity: event StationAvailabilityUpdated(string indexed stationId, bool newAvailability, address indexed reservor, uint256 timestamp)
func (_Contracts *ContractsFilterer) FilterStationAvailabilityUpdated(opts *bind.FilterOpts, stationId []string, reservor []common.Address) (*ContractsStationAvailabilityUpdatedIterator, error) {

	var stationIdRule []interface{}
	for _, stationIdItem := range stationId {
		stationIdRule = append(stationIdRule, stationIdItem)
	}

	var reservorRule []interface{}
	for _, reservorItem := range reservor {
		reservorRule = append(reservorRule, reservorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "StationAvailabilityUpdated", stationIdRule, reservorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsStationAvailabilityUpdatedIterator{contract: _Contracts.contract, event: "StationAvailabilityUpdated", logs: logs, sub: sub}, nil
}

// WatchStationAvailabilityUpdated is a free log subscription operation binding the contract event 0x9f45fd480a7696877da95e871642a241d6b0a68fb0fe585a562ca074de2c2145.
//
// Solidity: event StationAvailabilityUpdated(string indexed stationId, bool newAvailability, address indexed reservor, uint256 timestamp)
func (_Contracts *ContractsFilterer) WatchStationAvailabilityUpdated(opts *bind.WatchOpts, sink chan<- *ContractsStationAvailabilityUpdated, stationId []string, reservor []common.Address) (event.Subscription, error) {

	var stationIdRule []interface{}
	for _, stationIdItem := range stationId {
		stationIdRule = append(stationIdRule, stationIdItem)
	}

	var reservorRule []interface{}
	for _, reservorItem := range reservor {
		reservorRule = append(reservorRule, reservorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "StationAvailabilityUpdated", stationIdRule, reservorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsStationAvailabilityUpdated)
				if err := _Contracts.contract.UnpackLog(event, "StationAvailabilityUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStationAvailabilityUpdated is a log parse operation binding the contract event 0x9f45fd480a7696877da95e871642a241d6b0a68fb0fe585a562ca074de2c2145.
//
// Solidity: event StationAvailabilityUpdated(string indexed stationId, bool newAvailability, address indexed reservor, uint256 timestamp)
func (_Contracts *ContractsFilterer) ParseStationAvailabilityUpdated(log types.Log) (*ContractsStationAvailabilityUpdated, error) {
	event := new(ContractsStationAvailabilityUpdated)
	if err := _Contracts.contract.UnpackLog(event, "StationAvailabilityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsStationDeletedIterator is returned from FilterStationDeleted and is used to iterate over the raw logs and unpacked data for StationDeleted events raised by the Contracts contract.
type ContractsStationDeletedIterator struct {
	Event *ContractsStationDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsStationDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsStationDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsStationDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsStationDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsStationDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsStationDeleted represents a StationDeleted event raised by the Contracts contract.
type ContractsStationDeleted struct {
	StationId common.Hash
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStationDeleted is a free log retrieval operation binding the contract event 0xd7468e363ec0597d4d2d75dcac9670a851a25bfec8c52379cbbeedd59b73c6c1.
//
// Solidity: event StationDeleted(string indexed stationId)
func (_Contracts *ContractsFilterer) FilterStationDeleted(opts *bind.FilterOpts, stationId []string) (*ContractsStationDeletedIterator, error) {

	var stationIdRule []interface{}
	for _, stationIdItem := range stationId {
		stationIdRule = append(stationIdRule, stationIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "StationDeleted", stationIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsStationDeletedIterator{contract: _Contracts.contract, event: "StationDeleted", logs: logs, sub: sub}, nil
}

// WatchStationDeleted is a free log subscription operation binding the contract event 0xd7468e363ec0597d4d2d75dcac9670a851a25bfec8c52379cbbeedd59b73c6c1.
//
// Solidity: event StationDeleted(string indexed stationId)
func (_Contracts *ContractsFilterer) WatchStationDeleted(opts *bind.WatchOpts, sink chan<- *ContractsStationDeleted, stationId []string) (event.Subscription, error) {

	var stationIdRule []interface{}
	for _, stationIdItem := range stationId {
		stationIdRule = append(stationIdRule, stationIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "StationDeleted", stationIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsStationDeleted)
				if err := _Contracts.contract.UnpackLog(event, "StationDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStationDeleted is a log parse operation binding the contract event 0xd7468e363ec0597d4d2d75dcac9670a851a25bfec8c52379cbbeedd59b73c6c1.
//
// Solidity: event StationDeleted(string indexed stationId)
func (_Contracts *ContractsFilterer) ParseStationDeleted(log types.Log) (*ContractsStationDeleted, error) {
	event := new(ContractsStationDeleted)
	if err := _Contracts.contract.UnpackLog(event, "StationDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsStationRegisteredIterator is returned from FilterStationRegistered and is used to iterate over the raw logs and unpacked data for StationRegistered events raised by the Contracts contract.
type ContractsStationRegisteredIterator struct {
	Event *ContractsStationRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsStationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsStationRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsStationRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsStationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsStationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsStationRegistered represents a StationRegistered event raised by the Contracts contract.
type ContractsStationRegistered struct {
	StationId common.Hash
	Owner     common.Address
	City      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStationRegistered is a free log retrieval operation binding the contract event 0xed098a22e0ff25594835cdcd04fe2ec2bb711e3b7cb28d8fc4f02aa153459f2b.
//
// Solidity: event StationRegistered(string indexed stationId, address indexed owner, string city)
func (_Contracts *ContractsFilterer) FilterStationRegistered(opts *bind.FilterOpts, stationId []string, owner []common.Address) (*ContractsStationRegisteredIterator, error) {

	var stationIdRule []interface{}
	for _, stationIdItem := range stationId {
		stationIdRule = append(stationIdRule, stationIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "StationRegistered", stationIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsStationRegisteredIterator{contract: _Contracts.contract, event: "StationRegistered", logs: logs, sub: sub}, nil
}

// WatchStationRegistered is a free log subscription operation binding the contract event 0xed098a22e0ff25594835cdcd04fe2ec2bb711e3b7cb28d8fc4f02aa153459f2b.
//
// Solidity: event StationRegistered(string indexed stationId, address indexed owner, string city)
func (_Contracts *ContractsFilterer) WatchStationRegistered(opts *bind.WatchOpts, sink chan<- *ContractsStationRegistered, stationId []string, owner []common.Address) (event.Subscription, error) {

	var stationIdRule []interface{}
	for _, stationIdItem := range stationId {
		stationIdRule = append(stationIdRule, stationIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "StationRegistered", stationIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsStationRegistered)
				if err := _Contracts.contract.UnpackLog(event, "StationRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStationRegistered is a log parse operation binding the contract event 0xed098a22e0ff25594835cdcd04fe2ec2bb711e3b7cb28d8fc4f02aa153459f2b.
//
// Solidity: event StationRegistered(string indexed stationId, address indexed owner, string city)
func (_Contracts *ContractsFilterer) ParseStationRegistered(log types.Log) (*ContractsStationRegistered, error) {
	event := new(ContractsStationRegistered)
	if err := _Contracts.contract.UnpackLog(event, "StationRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
