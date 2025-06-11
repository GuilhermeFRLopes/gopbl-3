// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contratos

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

// ContratosMetaData contains all meta data concerning the Contratos contract.
var ContratosMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"idPosto\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disponivel\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"servidorOrigem\",\"type\":\"string\"}],\"name\":\"AtualDisponibilidadePosto\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"idPosto\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cidade\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"servidorOrigem\",\"type\":\"string\"}],\"name\":\"PostoRegistrado\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAllPostos\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"idsPostos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"postosRecarga\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"latitude\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"longitude\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"ocupado\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"servidorOrigem\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cidade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reservador\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"valorPagamento\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_latitude\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_longitude\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_servidorOrigem\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cidade\",\"type\":\"string\"}],\"name\":\"registrarPosto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_idPostos\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"_reservar\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_reservador\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_valorPagamento\",\"type\":\"uint256\"}],\"name\":\"reservarVagaPostos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_idPosto\",\"type\":\"string\"}],\"name\":\"verificarDisponibilidade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50611afd8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063527ce6f31461006757806355bb5cbe146100975780635dbafe29146100c75780638f8aa6f1146100e357806395ec3afb1461011a578063f7877e3a14610138575b600080fd5b610081600480360381019061007c9190610b4b565b610154565b60405161008e9190610c08565b60405180910390f35b6100b160048036038101906100ac9190610d5f565b610200565b6040516100be9190610dc3565b60405180910390f35b6100e160048036038101906100dc9190610e14565b61023d565b005b6100fd60048036038101906100f89190610d5f565b61043b565b604051610111989796959493929190610f01565b60405180910390f35b6101226106c6565b60405161012f91906110a7565b60405180910390f35b610152600480360381019061014d91906111db565b61079f565b005b6001818154811061016457600080fd5b90600052602060002001600091509050805461017f906112a9565b80601f01602080910402602001604051908101604052809291908181526020018280546101ab906112a9565b80156101f85780601f106101cd576101008083540402835291602001916101f8565b820191906000526020600020905b8154815290600101906020018083116101db57829003601f168201915b505050505081565b6000806000836040516102139190611316565b908152602001604051809103902090508060030160009054906101000a900460ff16915050919050565b6000808660405161024e9190611316565b9081526020016040518091039020600001805461026a906112a9565b9050146102ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a390611379565b60405180910390fd5b60405180610100016040528086815260200185815260200184815260200160001515815260200183815260200182815260200160405180602001604052806000815250815260200160008152506000866040516103099190611316565b9081526020016040518091039020600082015181600001908161032c9190611545565b50602082015181600101556040820151816002015560608201518160030160006101000a81548160ff02191690831515021790555060808201518160040190816103769190611545565b5060a082015181600501908161038c9190611545565b5060c08201518160060190816103a29190611545565b5060e082015181600701559050506001859080600181540180825580915050600190039060005260206000200160009091909190915090816103e49190611545565b50846040516103f39190611316565b60405180910390207f8c9754c7fc6b0df199ea5d8a036af296a20873a52888d9fbc75ff5ac9aaa2442828460405161042c929190611617565b60405180910390a25050505050565b600081805160208101820180518482526020830160208501208183528095505050505050600091509050806000018054610474906112a9565b80601f01602080910402602001604051908101604052809291908181526020018280546104a0906112a9565b80156104ed5780601f106104c2576101008083540402835291602001916104ed565b820191906000526020600020905b8154815290600101906020018083116104d057829003601f168201915b5050505050908060010154908060020154908060030160009054906101000a900460ff1690806004018054610521906112a9565b80601f016020809104026020016040519081016040528092919081815260200182805461054d906112a9565b801561059a5780601f1061056f5761010080835404028352916020019161059a565b820191906000526020600020905b81548152906001019060200180831161057d57829003601f168201915b5050505050908060050180546105af906112a9565b80601f01602080910402602001604051908101604052809291908181526020018280546105db906112a9565b80156106285780601f106105fd57610100808354040283529160200191610628565b820191906000526020600020905b81548152906001019060200180831161060b57829003601f168201915b50505050509080600601805461063d906112a9565b80601f0160208091040260200160405190810160405280929190818152602001828054610669906112a9565b80156106b65780601f1061068b576101008083540402835291602001916106b6565b820191906000526020600020905b81548152906001019060200180831161069957829003601f168201915b5050505050908060070154905088565b60606001805480602002602001604051908101604052809291908181526020016000905b82821015610796578382906000526020600020018054610709906112a9565b80601f0160208091040260200160405190810160405280929190818152602001828054610735906112a9565b80156107825780601f1061075757610100808354040283529160200191610782565b820191906000526020600020905b81548152906001019060200180831161076557829003601f168201915b5050505050815260200190600101906106ea565b50505050905090565b60008451116107e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107da9061169a565b60405180910390fd5b60005b8451811015610afa576000858281518110610804576108036116ba565b5b602002602001015190506000808260405161081f9190611316565b908152602001604051809103902090506000816000018054610840906112a9565b905011610882576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087990611735565b60405180910390fd5b851561098357600015158160030160009054906101000a900460ff161515146108e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d7906117a1565b60405180910390fd5b60018160030160006101000a81548160ff0219169083151502179055508481600601908161090e9190611545565b50838160070181905550806000016040516109299190611844565b60405180910390207faf99a1509360a139a5b5aab49ac6571796571db04782badc6574f3aa7014b98a8260030160009054906101000a900460ff16836006016040516109769291906118df565b60405180910390a2610aeb565b600115158160030160009054906101000a900460ff161515146109db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d29061195b565b60405180910390fd5b8480519060200120816006016040516109f49190611a1e565b604051809103902014610a3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3390611aa7565b60405180910390fd5b60008160030160006101000a81548160ff02191690831515021790555060405180602001604052806000815250816006019081610a799190611545565b506000816007018190555080600001604051610a959190611844565b60405180910390207faf99a1509360a139a5b5aab49ac6571796571db04782badc6574f3aa7014b98a8260030160009054906101000a900460ff1683600601604051610ae29291906118df565b60405180910390a25b505080806001019150506107e6565b5050505050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610b2881610b15565b8114610b3357600080fd5b50565b600081359050610b4581610b1f565b92915050565b600060208284031215610b6157610b60610b0b565b5b6000610b6f84828501610b36565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610bb2578082015181840152602081019050610b97565b60008484015250505050565b6000601f19601f8301169050919050565b6000610bda82610b78565b610be48185610b83565b9350610bf4818560208601610b94565b610bfd81610bbe565b840191505092915050565b60006020820190508181036000830152610c228184610bcf565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610c6c82610bbe565b810181811067ffffffffffffffff82111715610c8b57610c8a610c34565b5b80604052505050565b6000610c9e610b01565b9050610caa8282610c63565b919050565b600067ffffffffffffffff821115610cca57610cc9610c34565b5b610cd382610bbe565b9050602081019050919050565b82818337600083830152505050565b6000610d02610cfd84610caf565b610c94565b905082815260208101848484011115610d1e57610d1d610c2f565b5b610d29848285610ce0565b509392505050565b600082601f830112610d4657610d45610c2a565b5b8135610d56848260208601610cef565b91505092915050565b600060208284031215610d7557610d74610b0b565b5b600082013567ffffffffffffffff811115610d9357610d92610b10565b5b610d9f84828501610d31565b91505092915050565b60008115159050919050565b610dbd81610da8565b82525050565b6000602082019050610dd86000830184610db4565b92915050565b6000819050919050565b610df181610dde565b8114610dfc57600080fd5b50565b600081359050610e0e81610de8565b92915050565b600080600080600060a08688031215610e3057610e2f610b0b565b5b600086013567ffffffffffffffff811115610e4e57610e4d610b10565b5b610e5a88828901610d31565b9550506020610e6b88828901610dff565b9450506040610e7c88828901610dff565b935050606086013567ffffffffffffffff811115610e9d57610e9c610b10565b5b610ea988828901610d31565b925050608086013567ffffffffffffffff811115610eca57610ec9610b10565b5b610ed688828901610d31565b9150509295509295909350565b610eec81610dde565b82525050565b610efb81610b15565b82525050565b6000610100820190508181036000830152610f1c818b610bcf565b9050610f2b602083018a610ee3565b610f386040830189610ee3565b610f456060830188610db4565b8181036080830152610f578187610bcf565b905081810360a0830152610f6b8186610bcf565b905081810360c0830152610f7f8185610bcf565b9050610f8e60e0830184610ef2565b9998505050505050505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b6000610fe382610b78565b610fed8185610fc7565b9350610ffd818560208601610b94565b61100681610bbe565b840191505092915050565b600061101d8383610fd8565b905092915050565b6000602082019050919050565b600061103d82610f9b565b6110478185610fa6565b93508360208202850161105985610fb7565b8060005b8581101561109557848403895281516110768582611011565b945061108183611025565b925060208a0199505060018101905061105d565b50829750879550505050505092915050565b600060208201905081810360008301526110c18184611032565b905092915050565b600067ffffffffffffffff8211156110e4576110e3610c34565b5b602082029050602081019050919050565b600080fd5b600061110d611108846110c9565b610c94565b905080838252602082019050602084028301858111156111305761112f6110f5565b5b835b8181101561117757803567ffffffffffffffff81111561115557611154610c2a565b5b8086016111628982610d31565b85526020850194505050602081019050611132565b5050509392505050565b600082601f83011261119657611195610c2a565b5b81356111a68482602086016110fa565b91505092915050565b6111b881610da8565b81146111c357600080fd5b50565b6000813590506111d5816111af565b92915050565b600080600080608085870312156111f5576111f4610b0b565b5b600085013567ffffffffffffffff81111561121357611212610b10565b5b61121f87828801611181565b9450506020611230878288016111c6565b935050604085013567ffffffffffffffff81111561125157611250610b10565b5b61125d87828801610d31565b925050606061126e87828801610b36565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806112c157607f821691505b6020821081036112d4576112d361127a565b5b50919050565b600081905092915050565b60006112f082610b78565b6112fa81856112da565b935061130a818560208601610b94565b80840191505092915050565b600061132282846112e5565b915081905092915050565b7f506f73746f2064652072656361726761206a61206361646173747261646f0000600082015250565b6000611363601e83610b83565b915061136e8261132d565b602082019050919050565b6000602082019050818103600083015261139281611356565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026113fb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826113be565b61140586836113be565b95508019841693508086168417925050509392505050565b6000819050919050565b600061144261143d61143884610b15565b61141d565b610b15565b9050919050565b6000819050919050565b61145c83611427565b61147061146882611449565b8484546113cb565b825550505050565b600090565b611485611478565b611490818484611453565b505050565b5b818110156114b4576114a960008261147d565b600181019050611496565b5050565b601f8211156114f9576114ca81611399565b6114d3846113ae565b810160208510156114e2578190505b6114f66114ee856113ae565b830182611495565b50505b505050565b600082821c905092915050565b600061151c600019846008026114fe565b1980831691505092915050565b6000611535838361150b565b9150826002028217905092915050565b61154e82610b78565b67ffffffffffffffff81111561156757611566610c34565b5b61157182546112a9565b61157c8282856114b8565b600060209050601f8311600181146115af576000841561159d578287015190505b6115a78582611529565b86555061160f565b601f1984166115bd86611399565b60005b828110156115e5578489015182556001820191506020850194506020810190506115c0565b8683101561160257848901516115fe601f89168261150b565b8355505b6001600288020188555050505b505050505050565b600060408201905081810360008301526116318185610bcf565b905081810360208301526116458184610bcf565b90509392505050565b7f4e656e68756d20494420646520706f73746f20666f726e656369646f2e000000600082015250565b6000611684601d83610b83565b915061168f8261164e565b602082019050919050565b600060208201905081810360008301526116b381611677565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f506f73746f206e616f206578697374652e000000000000000000000000000000600082015250565b600061171f601183610b83565b915061172a826116e9565b602082019050919050565b6000602082019050818103600083015261174e81611712565b9050919050565b7f506f73746f206a612065737461206f63757061646f2e00000000000000000000600082015250565b600061178b601683610b83565b915061179682611755565b602082019050919050565b600060208201905081810360008301526117ba8161177e565b9050919050565b600081546117ce816112a9565b6117d881866112da565b945060018216600081146117f357600181146118085761183b565b60ff198316865281151582028601935061183b565b61181185611399565b60005b8381101561183357815481890152600182019150602081019050611814565b838801955050505b50505092915050565b600061185082846117c1565b915081905092915050565b60008154611868816112a9565b6118728186610b83565b9450600182166000811461188d57600181146118a3576118d6565b60ff1983168652811515602002860193506118d6565b6118ac85611399565b60005b838110156118ce578154818901526001820191506020810190506118af565b808801955050505b50505092915050565b60006040820190506118f46000830185610db4565b8181036020830152611906818461185b565b90509392505050565b7f506f73746f206a61206573746120646973706f6e6976656c2e00000000000000600082015250565b6000611945601983610b83565b91506119508261190f565b602082019050919050565b6000602082019050818103600083015261197481611938565b9050919050565b600081905092915050565b60008190508160005260206000209050919050565b600081546119a8816112a9565b6119b2818661197b565b945060018216600081146119cd57600181146119e257611a15565b60ff1983168652811515820286019350611a15565b6119eb85611986565b60005b83811015611a0d578154818901526001820191506020810190506119ee565b838801955050505b50505092915050565b6000611a2a828461199b565b915081905092915050565b7f4170656e6173206f2072657365727661646f7220706f6465206c69626572617260008201527f206573746520706f73746f2e0000000000000000000000000000000000000000602082015250565b6000611a91602c83610b83565b9150611a9c82611a35565b604082019050919050565b60006020820190508181036000830152611ac081611a84565b905091905056fea26469706673582212204e07b91f5eee123a7aa04fc6d1e6e1c88aa0a3d820f76d8b185a08ae1bef935664736f6c634300081e0033",
}

// ContratosABI is the input ABI used to generate the binding from.
// Deprecated: Use ContratosMetaData.ABI instead.
var ContratosABI = ContratosMetaData.ABI

// ContratosBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContratosMetaData.Bin instead.
var ContratosBin = ContratosMetaData.Bin

// DeployContratos deploys a new Ethereum contract, binding an instance of Contratos to it.
func DeployContratos(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contratos, error) {
	parsed, err := ContratosMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContratosBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contratos{ContratosCaller: ContratosCaller{contract: contract}, ContratosTransactor: ContratosTransactor{contract: contract}, ContratosFilterer: ContratosFilterer{contract: contract}}, nil
}

// Contratos is an auto generated Go binding around an Ethereum contract.
type Contratos struct {
	ContratosCaller     // Read-only binding to the contract
	ContratosTransactor // Write-only binding to the contract
	ContratosFilterer   // Log filterer for contract events
}

// ContratosCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContratosCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContratosTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContratosTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContratosFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContratosFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContratosSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContratosSession struct {
	Contract     *Contratos        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContratosCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContratosCallerSession struct {
	Contract *ContratosCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContratosTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContratosTransactorSession struct {
	Contract     *ContratosTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContratosRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContratosRaw struct {
	Contract *Contratos // Generic contract binding to access the raw methods on
}

// ContratosCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContratosCallerRaw struct {
	Contract *ContratosCaller // Generic read-only contract binding to access the raw methods on
}

// ContratosTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContratosTransactorRaw struct {
	Contract *ContratosTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContratos creates a new instance of Contratos, bound to a specific deployed contract.
func NewContratos(address common.Address, backend bind.ContractBackend) (*Contratos, error) {
	contract, err := bindContratos(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contratos{ContratosCaller: ContratosCaller{contract: contract}, ContratosTransactor: ContratosTransactor{contract: contract}, ContratosFilterer: ContratosFilterer{contract: contract}}, nil
}

// NewContratosCaller creates a new read-only instance of Contratos, bound to a specific deployed contract.
func NewContratosCaller(address common.Address, caller bind.ContractCaller) (*ContratosCaller, error) {
	contract, err := bindContratos(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContratosCaller{contract: contract}, nil
}

// NewContratosTransactor creates a new write-only instance of Contratos, bound to a specific deployed contract.
func NewContratosTransactor(address common.Address, transactor bind.ContractTransactor) (*ContratosTransactor, error) {
	contract, err := bindContratos(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContratosTransactor{contract: contract}, nil
}

// NewContratosFilterer creates a new log filterer instance of Contratos, bound to a specific deployed contract.
func NewContratosFilterer(address common.Address, filterer bind.ContractFilterer) (*ContratosFilterer, error) {
	contract, err := bindContratos(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContratosFilterer{contract: contract}, nil
}

// bindContratos binds a generic wrapper to an already deployed contract.
func bindContratos(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContratosMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contratos *ContratosRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contratos.Contract.ContratosCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contratos *ContratosRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contratos.Contract.ContratosTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contratos *ContratosRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contratos.Contract.ContratosTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contratos *ContratosCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contratos.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contratos *ContratosTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contratos.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contratos *ContratosTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contratos.Contract.contract.Transact(opts, method, params...)
}

// GetAllPostos is a free data retrieval call binding the contract method 0x95ec3afb.
//
// Solidity: function getAllPostos() view returns(string[])
func (_Contratos *ContratosCaller) GetAllPostos(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Contratos.contract.Call(opts, &out, "getAllPostos")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllPostos is a free data retrieval call binding the contract method 0x95ec3afb.
//
// Solidity: function getAllPostos() view returns(string[])
func (_Contratos *ContratosSession) GetAllPostos() ([]string, error) {
	return _Contratos.Contract.GetAllPostos(&_Contratos.CallOpts)
}

// GetAllPostos is a free data retrieval call binding the contract method 0x95ec3afb.
//
// Solidity: function getAllPostos() view returns(string[])
func (_Contratos *ContratosCallerSession) GetAllPostos() ([]string, error) {
	return _Contratos.Contract.GetAllPostos(&_Contratos.CallOpts)
}

// IdsPostos is a free data retrieval call binding the contract method 0x527ce6f3.
//
// Solidity: function idsPostos(uint256 ) view returns(string)
func (_Contratos *ContratosCaller) IdsPostos(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Contratos.contract.Call(opts, &out, "idsPostos", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IdsPostos is a free data retrieval call binding the contract method 0x527ce6f3.
//
// Solidity: function idsPostos(uint256 ) view returns(string)
func (_Contratos *ContratosSession) IdsPostos(arg0 *big.Int) (string, error) {
	return _Contratos.Contract.IdsPostos(&_Contratos.CallOpts, arg0)
}

// IdsPostos is a free data retrieval call binding the contract method 0x527ce6f3.
//
// Solidity: function idsPostos(uint256 ) view returns(string)
func (_Contratos *ContratosCallerSession) IdsPostos(arg0 *big.Int) (string, error) {
	return _Contratos.Contract.IdsPostos(&_Contratos.CallOpts, arg0)
}

// PostosRecarga is a free data retrieval call binding the contract method 0x8f8aa6f1.
//
// Solidity: function postosRecarga(string ) view returns(string id, int256 latitude, int256 longitude, bool ocupado, string servidorOrigem, string cidade, string reservador, uint256 valorPagamento)
func (_Contratos *ContratosCaller) PostosRecarga(opts *bind.CallOpts, arg0 string) (struct {
	Id             string
	Latitude       *big.Int
	Longitude      *big.Int
	Ocupado        bool
	ServidorOrigem string
	Cidade         string
	Reservador     string
	ValorPagamento *big.Int
}, error) {
	var out []interface{}
	err := _Contratos.contract.Call(opts, &out, "postosRecarga", arg0)

	outstruct := new(struct {
		Id             string
		Latitude       *big.Int
		Longitude      *big.Int
		Ocupado        bool
		ServidorOrigem string
		Cidade         string
		Reservador     string
		ValorPagamento *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Latitude = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Longitude = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Ocupado = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.ServidorOrigem = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Cidade = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Reservador = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.ValorPagamento = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PostosRecarga is a free data retrieval call binding the contract method 0x8f8aa6f1.
//
// Solidity: function postosRecarga(string ) view returns(string id, int256 latitude, int256 longitude, bool ocupado, string servidorOrigem, string cidade, string reservador, uint256 valorPagamento)
func (_Contratos *ContratosSession) PostosRecarga(arg0 string) (struct {
	Id             string
	Latitude       *big.Int
	Longitude      *big.Int
	Ocupado        bool
	ServidorOrigem string
	Cidade         string
	Reservador     string
	ValorPagamento *big.Int
}, error) {
	return _Contratos.Contract.PostosRecarga(&_Contratos.CallOpts, arg0)
}

// PostosRecarga is a free data retrieval call binding the contract method 0x8f8aa6f1.
//
// Solidity: function postosRecarga(string ) view returns(string id, int256 latitude, int256 longitude, bool ocupado, string servidorOrigem, string cidade, string reservador, uint256 valorPagamento)
func (_Contratos *ContratosCallerSession) PostosRecarga(arg0 string) (struct {
	Id             string
	Latitude       *big.Int
	Longitude      *big.Int
	Ocupado        bool
	ServidorOrigem string
	Cidade         string
	Reservador     string
	ValorPagamento *big.Int
}, error) {
	return _Contratos.Contract.PostosRecarga(&_Contratos.CallOpts, arg0)
}

// VerificarDisponibilidade is a free data retrieval call binding the contract method 0x55bb5cbe.
//
// Solidity: function verificarDisponibilidade(string _idPosto) view returns(bool)
func (_Contratos *ContratosCaller) VerificarDisponibilidade(opts *bind.CallOpts, _idPosto string) (bool, error) {
	var out []interface{}
	err := _Contratos.contract.Call(opts, &out, "verificarDisponibilidade", _idPosto)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerificarDisponibilidade is a free data retrieval call binding the contract method 0x55bb5cbe.
//
// Solidity: function verificarDisponibilidade(string _idPosto) view returns(bool)
func (_Contratos *ContratosSession) VerificarDisponibilidade(_idPosto string) (bool, error) {
	return _Contratos.Contract.VerificarDisponibilidade(&_Contratos.CallOpts, _idPosto)
}

// VerificarDisponibilidade is a free data retrieval call binding the contract method 0x55bb5cbe.
//
// Solidity: function verificarDisponibilidade(string _idPosto) view returns(bool)
func (_Contratos *ContratosCallerSession) VerificarDisponibilidade(_idPosto string) (bool, error) {
	return _Contratos.Contract.VerificarDisponibilidade(&_Contratos.CallOpts, _idPosto)
}

// RegistrarPosto is a paid mutator transaction binding the contract method 0x5dbafe29.
//
// Solidity: function registrarPosto(string _id, int256 _latitude, int256 _longitude, string _servidorOrigem, string _cidade) returns()
func (_Contratos *ContratosTransactor) RegistrarPosto(opts *bind.TransactOpts, _id string, _latitude *big.Int, _longitude *big.Int, _servidorOrigem string, _cidade string) (*types.Transaction, error) {
	return _Contratos.contract.Transact(opts, "registrarPosto", _id, _latitude, _longitude, _servidorOrigem, _cidade)
}

// RegistrarPosto is a paid mutator transaction binding the contract method 0x5dbafe29.
//
// Solidity: function registrarPosto(string _id, int256 _latitude, int256 _longitude, string _servidorOrigem, string _cidade) returns()
func (_Contratos *ContratosSession) RegistrarPosto(_id string, _latitude *big.Int, _longitude *big.Int, _servidorOrigem string, _cidade string) (*types.Transaction, error) {
	return _Contratos.Contract.RegistrarPosto(&_Contratos.TransactOpts, _id, _latitude, _longitude, _servidorOrigem, _cidade)
}

// RegistrarPosto is a paid mutator transaction binding the contract method 0x5dbafe29.
//
// Solidity: function registrarPosto(string _id, int256 _latitude, int256 _longitude, string _servidorOrigem, string _cidade) returns()
func (_Contratos *ContratosTransactorSession) RegistrarPosto(_id string, _latitude *big.Int, _longitude *big.Int, _servidorOrigem string, _cidade string) (*types.Transaction, error) {
	return _Contratos.Contract.RegistrarPosto(&_Contratos.TransactOpts, _id, _latitude, _longitude, _servidorOrigem, _cidade)
}

// ReservarVagaPostos is a paid mutator transaction binding the contract method 0xf7877e3a.
//
// Solidity: function reservarVagaPostos(string[] _idPostos, bool _reservar, string _reservador, uint256 _valorPagamento) returns()
func (_Contratos *ContratosTransactor) ReservarVagaPostos(opts *bind.TransactOpts, _idPostos []string, _reservar bool, _reservador string, _valorPagamento *big.Int) (*types.Transaction, error) {
	return _Contratos.contract.Transact(opts, "reservarVagaPostos", _idPostos, _reservar, _reservador, _valorPagamento)
}

// ReservarVagaPostos is a paid mutator transaction binding the contract method 0xf7877e3a.
//
// Solidity: function reservarVagaPostos(string[] _idPostos, bool _reservar, string _reservador, uint256 _valorPagamento) returns()
func (_Contratos *ContratosSession) ReservarVagaPostos(_idPostos []string, _reservar bool, _reservador string, _valorPagamento *big.Int) (*types.Transaction, error) {
	return _Contratos.Contract.ReservarVagaPostos(&_Contratos.TransactOpts, _idPostos, _reservar, _reservador, _valorPagamento)
}

// ReservarVagaPostos is a paid mutator transaction binding the contract method 0xf7877e3a.
//
// Solidity: function reservarVagaPostos(string[] _idPostos, bool _reservar, string _reservador, uint256 _valorPagamento) returns()
func (_Contratos *ContratosTransactorSession) ReservarVagaPostos(_idPostos []string, _reservar bool, _reservador string, _valorPagamento *big.Int) (*types.Transaction, error) {
	return _Contratos.Contract.ReservarVagaPostos(&_Contratos.TransactOpts, _idPostos, _reservar, _reservador, _valorPagamento)
}

// ContratosAtualDisponibilidadePostoIterator is returned from FilterAtualDisponibilidadePosto and is used to iterate over the raw logs and unpacked data for AtualDisponibilidadePosto events raised by the Contratos contract.
type ContratosAtualDisponibilidadePostoIterator struct {
	Event *ContratosAtualDisponibilidadePosto // Event containing the contract specifics and raw log

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
func (it *ContratosAtualDisponibilidadePostoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContratosAtualDisponibilidadePosto)
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
		it.Event = new(ContratosAtualDisponibilidadePosto)
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
func (it *ContratosAtualDisponibilidadePostoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContratosAtualDisponibilidadePostoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContratosAtualDisponibilidadePosto represents a AtualDisponibilidadePosto event raised by the Contratos contract.
type ContratosAtualDisponibilidadePosto struct {
	IdPosto        common.Hash
	Disponivel     bool
	ServidorOrigem string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAtualDisponibilidadePosto is a free log retrieval operation binding the contract event 0xaf99a1509360a139a5b5aab49ac6571796571db04782badc6574f3aa7014b98a.
//
// Solidity: event AtualDisponibilidadePosto(string indexed idPosto, bool disponivel, string servidorOrigem)
func (_Contratos *ContratosFilterer) FilterAtualDisponibilidadePosto(opts *bind.FilterOpts, idPosto []string) (*ContratosAtualDisponibilidadePostoIterator, error) {

	var idPostoRule []interface{}
	for _, idPostoItem := range idPosto {
		idPostoRule = append(idPostoRule, idPostoItem)
	}

	logs, sub, err := _Contratos.contract.FilterLogs(opts, "AtualDisponibilidadePosto", idPostoRule)
	if err != nil {
		return nil, err
	}
	return &ContratosAtualDisponibilidadePostoIterator{contract: _Contratos.contract, event: "AtualDisponibilidadePosto", logs: logs, sub: sub}, nil
}

// WatchAtualDisponibilidadePosto is a free log subscription operation binding the contract event 0xaf99a1509360a139a5b5aab49ac6571796571db04782badc6574f3aa7014b98a.
//
// Solidity: event AtualDisponibilidadePosto(string indexed idPosto, bool disponivel, string servidorOrigem)
func (_Contratos *ContratosFilterer) WatchAtualDisponibilidadePosto(opts *bind.WatchOpts, sink chan<- *ContratosAtualDisponibilidadePosto, idPosto []string) (event.Subscription, error) {

	var idPostoRule []interface{}
	for _, idPostoItem := range idPosto {
		idPostoRule = append(idPostoRule, idPostoItem)
	}

	logs, sub, err := _Contratos.contract.WatchLogs(opts, "AtualDisponibilidadePosto", idPostoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContratosAtualDisponibilidadePosto)
				if err := _Contratos.contract.UnpackLog(event, "AtualDisponibilidadePosto", log); err != nil {
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

// ParseAtualDisponibilidadePosto is a log parse operation binding the contract event 0xaf99a1509360a139a5b5aab49ac6571796571db04782badc6574f3aa7014b98a.
//
// Solidity: event AtualDisponibilidadePosto(string indexed idPosto, bool disponivel, string servidorOrigem)
func (_Contratos *ContratosFilterer) ParseAtualDisponibilidadePosto(log types.Log) (*ContratosAtualDisponibilidadePosto, error) {
	event := new(ContratosAtualDisponibilidadePosto)
	if err := _Contratos.contract.UnpackLog(event, "AtualDisponibilidadePosto", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContratosPostoRegistradoIterator is returned from FilterPostoRegistrado and is used to iterate over the raw logs and unpacked data for PostoRegistrado events raised by the Contratos contract.
type ContratosPostoRegistradoIterator struct {
	Event *ContratosPostoRegistrado // Event containing the contract specifics and raw log

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
func (it *ContratosPostoRegistradoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContratosPostoRegistrado)
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
		it.Event = new(ContratosPostoRegistrado)
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
func (it *ContratosPostoRegistradoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContratosPostoRegistradoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContratosPostoRegistrado represents a PostoRegistrado event raised by the Contratos contract.
type ContratosPostoRegistrado struct {
	IdPosto        common.Hash
	Cidade         string
	ServidorOrigem string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPostoRegistrado is a free log retrieval operation binding the contract event 0x8c9754c7fc6b0df199ea5d8a036af296a20873a52888d9fbc75ff5ac9aaa2442.
//
// Solidity: event PostoRegistrado(string indexed idPosto, string cidade, string servidorOrigem)
func (_Contratos *ContratosFilterer) FilterPostoRegistrado(opts *bind.FilterOpts, idPosto []string) (*ContratosPostoRegistradoIterator, error) {

	var idPostoRule []interface{}
	for _, idPostoItem := range idPosto {
		idPostoRule = append(idPostoRule, idPostoItem)
	}

	logs, sub, err := _Contratos.contract.FilterLogs(opts, "PostoRegistrado", idPostoRule)
	if err != nil {
		return nil, err
	}
	return &ContratosPostoRegistradoIterator{contract: _Contratos.contract, event: "PostoRegistrado", logs: logs, sub: sub}, nil
}

// WatchPostoRegistrado is a free log subscription operation binding the contract event 0x8c9754c7fc6b0df199ea5d8a036af296a20873a52888d9fbc75ff5ac9aaa2442.
//
// Solidity: event PostoRegistrado(string indexed idPosto, string cidade, string servidorOrigem)
func (_Contratos *ContratosFilterer) WatchPostoRegistrado(opts *bind.WatchOpts, sink chan<- *ContratosPostoRegistrado, idPosto []string) (event.Subscription, error) {

	var idPostoRule []interface{}
	for _, idPostoItem := range idPosto {
		idPostoRule = append(idPostoRule, idPostoItem)
	}

	logs, sub, err := _Contratos.contract.WatchLogs(opts, "PostoRegistrado", idPostoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContratosPostoRegistrado)
				if err := _Contratos.contract.UnpackLog(event, "PostoRegistrado", log); err != nil {
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

// ParsePostoRegistrado is a log parse operation binding the contract event 0x8c9754c7fc6b0df199ea5d8a036af296a20873a52888d9fbc75ff5ac9aaa2442.
//
// Solidity: event PostoRegistrado(string indexed idPosto, string cidade, string servidorOrigem)
func (_Contratos *ContratosFilterer) ParsePostoRegistrado(log types.Log) (*ContratosPostoRegistrado, error) {
	event := new(ContratosPostoRegistrado)
	if err := _Contratos.contract.UnpackLog(event, "PostoRegistrado", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
