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
	Bin: "0x6080604052348015600f57600080fd5b50611ed98061001f6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80639781e91b1161005b5780639781e91b146100ec578063a836e3b21461010a578063e391daf914610140578063e701768e146101705761007d565b806337fa2656146100825780635d7e2e301461009e5780636a4e3b04146100d0575b600080fd5b61009c60048036038101906100979190611012565b61018c565b005b6100b860048036038101906100b391906110b9565b610417565b6040516100c793929190611177565b60405180910390f35b6100ea60048036038101906100e591906110b9565b6104da565b005b6100f46107a4565b60405161010191906112ef565b60405180910390f35b610124600480360381019061011f91906110b9565b61087d565b604051610137979695949392919061135b565b60405180910390f35b61015a6004803603810190610155919061140b565b610aba565b6040516101679190611438565b60405180910390f35b61018a60048036038101906101859190611598565b610b66565b005b6000808460405161019d9190611643565b908152602001604051809103902060000180546101b990611689565b9050146101fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f29061172c565b60405180910390fd5b60018390806001815401808255809150506001900390600052602060002001600090919091909150908161022f91906118f8565b506040518060e001604052808481526020018381526020018281526020013373ffffffffffffffffffffffffffffffffffffffff168152602001600115158152602001428152602001600073ffffffffffffffffffffffffffffffffffffffff168152506000846040516102a39190611643565b908152602001604051809103902060008201518160000190816102c691906118f8565b5060208201518160010190816102dc91906118f8565b5060408201518160020190816102f291906118f8565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160146101000a81548160ff02191690831515021790555060a0820151816004015560c08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff16836040516103d39190611643565b60405180910390207fed098a22e0ff25594835cdcd04fe2ec2bb711e3b7cb28d8fc4f02aa153459f2b8360405161040a9190611438565b60405180910390a3505050565b60008060008060008560405161042d9190611643565b90815260200160405180910390209050600081600001805461044e90611689565b905011610490576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048790611a16565b60405180910390fd5b8060030160149054906101000a900460ff168160050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260040154935093509350509193909250565b6000816040516104ea9190611643565b908152602001604051809103902060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610589576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058090611aa8565b60405180910390fd5b60005b6001805490508110156106a357816040516020016105aa9190611643565b60405160208183030381529060405280519060200120600182815481106105d4576105d3611ac8565b5b906000526020600020016040516020016105ee9190611b7a565b604051602081830303815290604052805190602001200361069657600180808054905061061b9190611bc0565b8154811061062c5761062b611ac8565b5b906000526020600020016001828154811061064a57610649611ac8565b5b9060005260206000200190816106609190611c1f565b50600180548061067357610672611d07565b5b60019003818190600052602060002001600061068f9190610e5b565b90556106a3565b808060010191505061058c565b506000816040516106b49190611643565b9081526020016040518091039020600080820160006106d39190610e5b565b6001820160006106e39190610e5b565b6002820160006106f39190610e5b565b6003820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160146101000a81549060ff021916905560048201600090556005820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550508060405161076d9190611643565b60405180910390207fd7468e363ec0597d4d2d75dcac9670a851a25bfec8c52379cbbeedd59b73c6c160405160405180910390a250565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156108745783829060005260206000200180546107e790611689565b80601f016020809104026020016040519081016040528092919081815260200182805461081390611689565b80156108605780601f1061083557610100808354040283529160200191610860565b820191906000526020600020905b81548152906001019060200180831161084357829003601f168201915b5050505050815260200190600101906107c8565b50505050905090565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546108b690611689565b80601f01602080910402602001604051908101604052809291908181526020018280546108e290611689565b801561092f5780601f106109045761010080835404028352916020019161092f565b820191906000526020600020905b81548152906001019060200180831161091257829003601f168201915b50505050509080600101805461094490611689565b80601f016020809104026020016040519081016040528092919081815260200182805461097090611689565b80156109bd5780601f10610992576101008083540402835291602001916109bd565b820191906000526020600020905b8154815290600101906020018083116109a057829003601f168201915b5050505050908060020180546109d290611689565b80601f01602080910402602001604051908101604052809291908181526020018280546109fe90611689565b8015610a4b5780601f10610a2057610100808354040283529160200191610a4b565b820191906000526020600020905b815481529060010190602001808311610a2e57829003601f168201915b5050505050908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160149054906101000a900460ff16908060040154908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905087565b60018181548110610aca57600080fd5b906000526020600020016000915090508054610ae590611689565b80601f0160208091040260200160405190810160405280929190818152602001828054610b1190611689565b8015610b5e5780601f10610b3357610100808354040283529160200191610b5e565b820191906000526020600020905b815481529060010190602001808311610b4157829003601f168201915b505050505081565b6000835111610baa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ba190611d82565b60405180910390fd5b60005b8351811015610e55576000848281518110610bcb57610bca611ac8565b5b6020026020010151905060008082604051610be69190611643565b908152602001604051809103902090506000816000018054610c0790611689565b905011610c49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4090611a16565b60405180910390fd5b8415610d0c57600115158160030160149054906101000a900460ff16151514610ca7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9e90611dee565b60405180910390fd5b60008160030160146101000a81548160ff021916908315150217905550838160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610dc6565b600015158160030160149054906101000a900460ff16151514610d64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5b90611e5a565b60405180910390fd5b60018160030160146101000a81548160ff02191690831515021790555060008160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b4281600401819055508373ffffffffffffffffffffffffffffffffffffffff1682604051610df49190611643565b60405180910390207f9f45fd480a7696877da95e871642a241d6b0a68fb0fe585a562ca074de2c21458360030160149054906101000a900460ff1642604051610e3e929190611e7a565b60405180910390a350508080600101915050610bad565b50505050565b508054610e6790611689565b6000825580601f10610e795750610e98565b601f016020900490600052602060002090810190610e979190610e9b565b5b50565b5b80821115610eb4576000816000905550600101610e9c565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f1f82610ed6565b810181811067ffffffffffffffff82111715610f3e57610f3d610ee7565b5b80604052505050565b6000610f51610eb8565b9050610f5d8282610f16565b919050565b600067ffffffffffffffff821115610f7d57610f7c610ee7565b5b610f8682610ed6565b9050602081019050919050565b82818337600083830152505050565b6000610fb5610fb084610f62565b610f47565b905082815260208101848484011115610fd157610fd0610ed1565b5b610fdc848285610f93565b509392505050565b600082601f830112610ff957610ff8610ecc565b5b8135611009848260208601610fa2565b91505092915050565b60008060006060848603121561102b5761102a610ec2565b5b600084013567ffffffffffffffff81111561104957611048610ec7565b5b61105586828701610fe4565b935050602084013567ffffffffffffffff81111561107657611075610ec7565b5b61108286828701610fe4565b925050604084013567ffffffffffffffff8111156110a3576110a2610ec7565b5b6110af86828701610fe4565b9150509250925092565b6000602082840312156110cf576110ce610ec2565b5b600082013567ffffffffffffffff8111156110ed576110ec610ec7565b5b6110f984828501610fe4565b91505092915050565b60008115159050919050565b61111781611102565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006111488261111d565b9050919050565b6111588161113d565b82525050565b6000819050919050565b6111718161115e565b82525050565b600060608201905061118c600083018661110e565b611199602083018561114f565b6111a66040830184611168565b949350505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156112145780820151818401526020810190506111f9565b60008484015250505050565b600061122b826111da565b61123581856111e5565b93506112458185602086016111f6565b61124e81610ed6565b840191505092915050565b60006112658383611220565b905092915050565b6000602082019050919050565b6000611285826111ae565b61128f81856111b9565b9350836020820285016112a1856111ca565b8060005b858110156112dd57848403895281516112be8582611259565b94506112c98361126d565b925060208a019950506001810190506112a5565b50829750879550505050505092915050565b60006020820190508181036000830152611309818461127a565b905092915050565b600082825260208201905092915050565b600061132d826111da565b6113378185611311565b93506113478185602086016111f6565b61135081610ed6565b840191505092915050565b600060e0820190508181036000830152611375818a611322565b905081810360208301526113898189611322565b9050818103604083015261139d8188611322565b90506113ac606083018761114f565b6113b9608083018661110e565b6113c660a0830185611168565b6113d360c083018461114f565b98975050505050505050565b6113e88161115e565b81146113f357600080fd5b50565b600081359050611405816113df565b92915050565b60006020828403121561142157611420610ec2565b5b600061142f848285016113f6565b91505092915050565b600060208201905081810360008301526114528184611322565b905092915050565b600067ffffffffffffffff82111561147557611474610ee7565b5b602082029050602081019050919050565b600080fd5b600061149e6114998461145a565b610f47565b905080838252602082019050602084028301858111156114c1576114c0611486565b5b835b8181101561150857803567ffffffffffffffff8111156114e6576114e5610ecc565b5b8086016114f38982610fe4565b855260208501945050506020810190506114c3565b5050509392505050565b600082601f83011261152757611526610ecc565b5b813561153784826020860161148b565b91505092915050565b61154981611102565b811461155457600080fd5b50565b60008135905061156681611540565b92915050565b6115758161113d565b811461158057600080fd5b50565b6000813590506115928161156c565b92915050565b6000806000606084860312156115b1576115b0610ec2565b5b600084013567ffffffffffffffff8111156115cf576115ce610ec7565b5b6115db86828701611512565b93505060206115ec86828701611557565b92505060406115fd86828701611583565b9150509250925092565b600081905092915050565b600061161d826111da565b6116278185611607565b93506116378185602086016111f6565b80840191505092915050565b600061164f8284611612565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806116a157607f821691505b6020821081036116b4576116b361165a565b5b50919050565b7f4368617267696e672073746174696f6e20616c7265616479207265676973746560008201527f7265642e00000000000000000000000000000000000000000000000000000000602082015250565b6000611716602483611311565b9150611721826116ba565b604082019050919050565b6000602082019050818103600083015261174581611709565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026117ae7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611771565b6117b88683611771565b95508019841693508086168417925050509392505050565b6000819050919050565b60006117f56117f06117eb8461115e565b6117d0565b61115e565b9050919050565b6000819050919050565b61180f836117da565b61182361181b826117fc565b84845461177e565b825550505050565b600090565b61183861182b565b611843818484611806565b505050565b5b818110156118675761185c600082611830565b600181019050611849565b5050565b601f8211156118ac5761187d8161174c565b61188684611761565b81016020851015611895578190505b6118a96118a185611761565b830182611848565b50505b505050565b600082821c905092915050565b60006118cf600019846008026118b1565b1980831691505092915050565b60006118e883836118be565b9150826002028217905092915050565b611901826111da565b67ffffffffffffffff81111561191a57611919610ee7565b5b6119248254611689565b61192f82828561186b565b600060209050601f8311600181146119625760008415611950578287015190505b61195a85826118dc565b8655506119c2565b601f1984166119708661174c565b60005b8281101561199857848901518255600182019150602085019450602081019050611973565b868310156119b557848901516119b1601f8916826118be565b8355505b6001600288020188555050505b505050505050565b7f53746174696f6e20646f6573206e6f742065786973742e000000000000000000600082015250565b6000611a00601783611311565b9150611a0b826119ca565b602082019050919050565b60006020820190508181036000830152611a2f816119f3565b9050919050565b7f4f6e6c7920746865206f776e65722063616e2064656c6574652074686973207360008201527f746174696f6e2e00000000000000000000000000000000000000000000000000602082015250565b6000611a92602783611311565b9150611a9d82611a36565b604082019050919050565b60006020820190508181036000830152611ac181611a85565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008154611b0481611689565b611b0e8186611607565b94506001821660008114611b295760018114611b3e57611b71565b60ff1983168652811515820286019350611b71565b611b478561174c565b60005b83811015611b6957815481890152600182019150602081019050611b4a565b838801955050505b50505092915050565b6000611b868284611af7565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611bcb8261115e565b9150611bd68361115e565b9250828203905081811115611bee57611bed611b91565b5b92915050565b600081549050611c0381611689565b9050919050565b60008190508160005260206000209050919050565b818103611c2d575050611d05565b611c3682611bf4565b67ffffffffffffffff811115611c4f57611c4e610ee7565b5b611c598254611689565b611c6482828561186b565b6000601f831160018114611c935760008415611c81578287015490505b611c8b85826118dc565b865550611cfe565b601f198416611ca187611c0a565b9650611cac8661174c565b60005b82811015611cd457848901548255600182019150600185019450602081019050611caf565b86831015611cf15784890154611ced601f8916826118be565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e6f2073746174696f6e204944732070726f76696465642e0000000000000000600082015250565b6000611d6c601883611311565b9150611d7782611d36565b602082019050919050565b60006020820190508181036000830152611d9b81611d5f565b9050919050565b7f53746174696f6e20697320616c72656164792072657365727665642e00000000600082015250565b6000611dd8601c83611311565b9150611de382611da2565b602082019050919050565b60006020820190508181036000830152611e0781611dcb565b9050919050565b7f53746174696f6e20697320616c726561647920617661696c61626c652e000000600082015250565b6000611e44601d83611311565b9150611e4f82611e0e565b602082019050919050565b60006020820190508181036000830152611e7381611e37565b9050919050565b6000604082019050611e8f600083018561110e565b611e9c6020830184611168565b939250505056fea264697066735822122068ce3d54783379507a678246310540fc638641b634bb7c6d75c5df1c879b6a2264736f6c634300081e0033",
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
