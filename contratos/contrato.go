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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"idPosto\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disponivel\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"servidorOrigem\",\"type\":\"string\"}],\"name\":\"AtualDisponibilidadePosto\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"idPosto\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cidade\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"servidorOrigem\",\"type\":\"string\"}],\"name\":\"PostoRegistrado\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAllPostos\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"idsPostos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"postosRecarga\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"latitude\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"longitude\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"ocupado\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"servidorOrigem\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cidade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reservador\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_latitude\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_longitude\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_servidorOrigem\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cidade\",\"type\":\"string\"}],\"name\":\"registrarPosto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_idPostos\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"_reservar\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_reservador\",\"type\":\"string\"}],\"name\":\"reservarVagaPostos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_idPosto\",\"type\":\"string\"}],\"name\":\"verificarDisponibilidade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50611a9e8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063527ce6f31461006757806355bb5cbe146100975780635dbafe29146100c75780636dba52dc146100e35780638f8aa6f1146100ff57806395ec3afb14610135575b600080fd5b610081600480360381019061007c9190610b1e565b610153565b60405161008e9190610bdb565b60405180910390f35b6100b160048036038101906100ac9190610d32565b6101ff565b6040516100be9190610d96565b60405180910390f35b6100e160048036038101906100dc9190610de7565b61023c565b005b6100fd60048036038101906100f89190610fc8565b610428565b005b61011960048036038101906101149190610d32565b610776565b60405161012c9796959493929190611062565b60405180910390f35b61013d6109fb565b60405161014a91906111f9565b60405180910390f35b6001818154811061016357600080fd5b90600052602060002001600091509050805461017e9061124a565b80601f01602080910402602001604051908101604052809291908181526020018280546101aa9061124a565b80156101f75780601f106101cc576101008083540402835291602001916101f7565b820191906000526020600020905b8154815290600101906020018083116101da57829003601f168201915b505050505081565b60008060008360405161021291906112b7565b908152602001604051809103902090508060030160009054906101000a900460ff16915050919050565b6000808660405161024d91906112b7565b908152602001604051809103902060000180546102699061124a565b9050146102ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a29061131a565b60405180910390fd5b6040518060e001604052808681526020018581526020018481526020016000151581526020018381526020018281526020016040518060200160405280600081525081525060008660405161030091906112b7565b9081526020016040518091039020600082015181600001908161032391906114e6565b50602082015181600101556040820151816002015560608201518160030160006101000a81548160ff021916908315150217905550608082015181600401908161036d91906114e6565b5060a082015181600501908161038391906114e6565b5060c082015181600601908161039991906114e6565b509050506001859080600181540180825580915050600190039060005260206000200160009091909190915090816103d191906114e6565b50846040516103e091906112b7565b60405180910390207f8c9754c7fc6b0df199ea5d8a036af296a20873a52888d9fbc75ff5ac9aaa244282846040516104199291906115b8565b60405180910390a25050505050565b600083511161046c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104639061163b565b60405180910390fd5b60005b835181101561077057600084828151811061048d5761048c61165b565b5b60200260200101519050600080826040516104a891906112b7565b9081526020016040518091039020905060008160000180546104c99061124a565b90501161050b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610502906116d6565b60405180910390fd5b841561060357600015158160030160009054906101000a900460ff16151514610569576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056090611742565b60405180910390fd5b60018160030160006101000a81548160ff0219169083151502179055508381600601908161059791906114e6565b50806000016040516105a991906117e5565b60405180910390207faf99a1509360a139a5b5aab49ac6571796571db04782badc6574f3aa7014b98a8260030160009054906101000a900460ff16836006016040516105f6929190611880565b60405180910390a2610761565b600115158160030160009054906101000a900460ff1615151461065b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610652906118fc565b60405180910390fd5b83805190602001208160060160405161067491906119bf565b6040518091039020146106bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b390611a48565b60405180910390fd5b60008160030160006101000a81548160ff021916908315150217905550604051806020016040528060008152508160060190816106f991906114e6565b508060000160405161070b91906117e5565b60405180910390207faf99a1509360a139a5b5aab49ac6571796571db04782badc6574f3aa7014b98a8260030160009054906101000a900460ff1683600601604051610758929190611880565b60405180910390a25b5050808060010191505061046f565b50505050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546107af9061124a565b80601f01602080910402602001604051908101604052809291908181526020018280546107db9061124a565b80156108285780601f106107fd57610100808354040283529160200191610828565b820191906000526020600020905b81548152906001019060200180831161080b57829003601f168201915b5050505050908060010154908060020154908060030160009054906101000a900460ff169080600401805461085c9061124a565b80601f01602080910402602001604051908101604052809291908181526020018280546108889061124a565b80156108d55780601f106108aa576101008083540402835291602001916108d5565b820191906000526020600020905b8154815290600101906020018083116108b857829003601f168201915b5050505050908060050180546108ea9061124a565b80601f01602080910402602001604051908101604052809291908181526020018280546109169061124a565b80156109635780601f1061093857610100808354040283529160200191610963565b820191906000526020600020905b81548152906001019060200180831161094657829003601f168201915b5050505050908060060180546109789061124a565b80601f01602080910402602001604051908101604052809291908181526020018280546109a49061124a565b80156109f15780601f106109c6576101008083540402835291602001916109f1565b820191906000526020600020905b8154815290600101906020018083116109d457829003601f168201915b5050505050905087565b60606001805480602002602001604051908101604052809291908181526020016000905b82821015610acb578382906000526020600020018054610a3e9061124a565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6a9061124a565b8015610ab75780601f10610a8c57610100808354040283529160200191610ab7565b820191906000526020600020905b815481529060010190602001808311610a9a57829003601f168201915b505050505081526020019060010190610a1f565b50505050905090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610afb81610ae8565b8114610b0657600080fd5b50565b600081359050610b1881610af2565b92915050565b600060208284031215610b3457610b33610ade565b5b6000610b4284828501610b09565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b85578082015181840152602081019050610b6a565b60008484015250505050565b6000601f19601f8301169050919050565b6000610bad82610b4b565b610bb78185610b56565b9350610bc7818560208601610b67565b610bd081610b91565b840191505092915050565b60006020820190508181036000830152610bf58184610ba2565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610c3f82610b91565b810181811067ffffffffffffffff82111715610c5e57610c5d610c07565b5b80604052505050565b6000610c71610ad4565b9050610c7d8282610c36565b919050565b600067ffffffffffffffff821115610c9d57610c9c610c07565b5b610ca682610b91565b9050602081019050919050565b82818337600083830152505050565b6000610cd5610cd084610c82565b610c67565b905082815260208101848484011115610cf157610cf0610c02565b5b610cfc848285610cb3565b509392505050565b600082601f830112610d1957610d18610bfd565b5b8135610d29848260208601610cc2565b91505092915050565b600060208284031215610d4857610d47610ade565b5b600082013567ffffffffffffffff811115610d6657610d65610ae3565b5b610d7284828501610d04565b91505092915050565b60008115159050919050565b610d9081610d7b565b82525050565b6000602082019050610dab6000830184610d87565b92915050565b6000819050919050565b610dc481610db1565b8114610dcf57600080fd5b50565b600081359050610de181610dbb565b92915050565b600080600080600060a08688031215610e0357610e02610ade565b5b600086013567ffffffffffffffff811115610e2157610e20610ae3565b5b610e2d88828901610d04565b9550506020610e3e88828901610dd2565b9450506040610e4f88828901610dd2565b935050606086013567ffffffffffffffff811115610e7057610e6f610ae3565b5b610e7c88828901610d04565b925050608086013567ffffffffffffffff811115610e9d57610e9c610ae3565b5b610ea988828901610d04565b9150509295509295909350565b600067ffffffffffffffff821115610ed157610ed0610c07565b5b602082029050602081019050919050565b600080fd5b6000610efa610ef584610eb6565b610c67565b90508083825260208201905060208402830185811115610f1d57610f1c610ee2565b5b835b81811015610f6457803567ffffffffffffffff811115610f4257610f41610bfd565b5b808601610f4f8982610d04565b85526020850194505050602081019050610f1f565b5050509392505050565b600082601f830112610f8357610f82610bfd565b5b8135610f93848260208601610ee7565b91505092915050565b610fa581610d7b565b8114610fb057600080fd5b50565b600081359050610fc281610f9c565b92915050565b600080600060608486031215610fe157610fe0610ade565b5b600084013567ffffffffffffffff811115610fff57610ffe610ae3565b5b61100b86828701610f6e565b935050602061101c86828701610fb3565b925050604084013567ffffffffffffffff81111561103d5761103c610ae3565b5b61104986828701610d04565b9150509250925092565b61105c81610db1565b82525050565b600060e082019050818103600083015261107c818a610ba2565b905061108b6020830189611053565b6110986040830188611053565b6110a56060830187610d87565b81810360808301526110b78186610ba2565b905081810360a08301526110cb8185610ba2565b905081810360c08301526110df8184610ba2565b905098975050505050505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b600061113582610b4b565b61113f8185611119565b935061114f818560208601610b67565b61115881610b91565b840191505092915050565b600061116f838361112a565b905092915050565b6000602082019050919050565b600061118f826110ed565b61119981856110f8565b9350836020820285016111ab85611109565b8060005b858110156111e757848403895281516111c88582611163565b94506111d383611177565b925060208a019950506001810190506111af565b50829750879550505050505092915050565b600060208201905081810360008301526112138184611184565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061126257607f821691505b6020821081036112755761127461121b565b5b50919050565b600081905092915050565b600061129182610b4b565b61129b818561127b565b93506112ab818560208601610b67565b80840191505092915050565b60006112c38284611286565b915081905092915050565b7f506f73746f2064652072656361726761206a61206361646173747261646f0000600082015250565b6000611304601e83610b56565b915061130f826112ce565b602082019050919050565b60006020820190508181036000830152611333816112f7565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261139c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261135f565b6113a6868361135f565b95508019841693508086168417925050509392505050565b6000819050919050565b60006113e36113de6113d984610ae8565b6113be565b610ae8565b9050919050565b6000819050919050565b6113fd836113c8565b611411611409826113ea565b84845461136c565b825550505050565b600090565b611426611419565b6114318184846113f4565b505050565b5b818110156114555761144a60008261141e565b600181019050611437565b5050565b601f82111561149a5761146b8161133a565b6114748461134f565b81016020851015611483578190505b61149761148f8561134f565b830182611436565b50505b505050565b600082821c905092915050565b60006114bd6000198460080261149f565b1980831691505092915050565b60006114d683836114ac565b9150826002028217905092915050565b6114ef82610b4b565b67ffffffffffffffff81111561150857611507610c07565b5b611512825461124a565b61151d828285611459565b600060209050601f831160018114611550576000841561153e578287015190505b61154885826114ca565b8655506115b0565b601f19841661155e8661133a565b60005b8281101561158657848901518255600182019150602085019450602081019050611561565b868310156115a3578489015161159f601f8916826114ac565b8355505b6001600288020188555050505b505050505050565b600060408201905081810360008301526115d28185610ba2565b905081810360208301526115e68184610ba2565b90509392505050565b7f4e656e68756d20494420646520706f73746f20666f726e656369646f2e000000600082015250565b6000611625601d83610b56565b9150611630826115ef565b602082019050919050565b6000602082019050818103600083015261165481611618565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f506f73746f206e616f206578697374652e000000000000000000000000000000600082015250565b60006116c0601183610b56565b91506116cb8261168a565b602082019050919050565b600060208201905081810360008301526116ef816116b3565b9050919050565b7f506f73746f206a612065737461206f63757061646f2e00000000000000000000600082015250565b600061172c601683610b56565b9150611737826116f6565b602082019050919050565b6000602082019050818103600083015261175b8161171f565b9050919050565b6000815461176f8161124a565b611779818661127b565b9450600182166000811461179457600181146117a9576117dc565b60ff19831686528115158202860193506117dc565b6117b28561133a565b60005b838110156117d4578154818901526001820191506020810190506117b5565b838801955050505b50505092915050565b60006117f18284611762565b915081905092915050565b600081546118098161124a565b6118138186610b56565b9450600182166000811461182e576001811461184457611877565b60ff198316865281151560200286019350611877565b61184d8561133a565b60005b8381101561186f57815481890152600182019150602081019050611850565b808801955050505b50505092915050565b60006040820190506118956000830185610d87565b81810360208301526118a781846117fc565b90509392505050565b7f506f73746f206a61206573746120646973706f6e6976656c2e00000000000000600082015250565b60006118e6601983610b56565b91506118f1826118b0565b602082019050919050565b60006020820190508181036000830152611915816118d9565b9050919050565b600081905092915050565b60008190508160005260206000209050919050565b600081546119498161124a565b611953818661191c565b9450600182166000811461196e5760018114611983576119b6565b60ff19831686528115158202860193506119b6565b61198c85611927565b60005b838110156119ae5781548189015260018201915060208101905061198f565b838801955050505b50505092915050565b60006119cb828461193c565b915081905092915050565b7f4170656e6173206f2072657365727661646f7220706f6465206c69626572617260008201527f206573746520706f73746f2e0000000000000000000000000000000000000000602082015250565b6000611a32602c83610b56565b9150611a3d826119d6565b604082019050919050565b60006020820190508181036000830152611a6181611a25565b905091905056fea26469706673582212205053be243f2a4575fe2b4320f14e4bc7cec586fcb8ac48cd5b717fce9f7ee7ef64736f6c634300081e0033",
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
// Solidity: function postosRecarga(string ) view returns(string id, int256 latitude, int256 longitude, bool ocupado, string servidorOrigem, string cidade, string reservador)
func (_Contratos *ContratosCaller) PostosRecarga(opts *bind.CallOpts, arg0 string) (struct {
	Id             string
	Latitude       *big.Int
	Longitude      *big.Int
	Ocupado        bool
	ServidorOrigem string
	Cidade         string
	Reservador     string
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

	return *outstruct, err

}

// PostosRecarga is a free data retrieval call binding the contract method 0x8f8aa6f1.
//
// Solidity: function postosRecarga(string ) view returns(string id, int256 latitude, int256 longitude, bool ocupado, string servidorOrigem, string cidade, string reservador)
func (_Contratos *ContratosSession) PostosRecarga(arg0 string) (struct {
	Id             string
	Latitude       *big.Int
	Longitude      *big.Int
	Ocupado        bool
	ServidorOrigem string
	Cidade         string
	Reservador     string
}, error) {
	return _Contratos.Contract.PostosRecarga(&_Contratos.CallOpts, arg0)
}

// PostosRecarga is a free data retrieval call binding the contract method 0x8f8aa6f1.
//
// Solidity: function postosRecarga(string ) view returns(string id, int256 latitude, int256 longitude, bool ocupado, string servidorOrigem, string cidade, string reservador)
func (_Contratos *ContratosCallerSession) PostosRecarga(arg0 string) (struct {
	Id             string
	Latitude       *big.Int
	Longitude      *big.Int
	Ocupado        bool
	ServidorOrigem string
	Cidade         string
	Reservador     string
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

// ReservarVagaPostos is a paid mutator transaction binding the contract method 0x6dba52dc.
//
// Solidity: function reservarVagaPostos(string[] _idPostos, bool _reservar, string _reservador) returns()
func (_Contratos *ContratosTransactor) ReservarVagaPostos(opts *bind.TransactOpts, _idPostos []string, _reservar bool, _reservador string) (*types.Transaction, error) {
	return _Contratos.contract.Transact(opts, "reservarVagaPostos", _idPostos, _reservar, _reservador)
}

// ReservarVagaPostos is a paid mutator transaction binding the contract method 0x6dba52dc.
//
// Solidity: function reservarVagaPostos(string[] _idPostos, bool _reservar, string _reservador) returns()
func (_Contratos *ContratosSession) ReservarVagaPostos(_idPostos []string, _reservar bool, _reservador string) (*types.Transaction, error) {
	return _Contratos.Contract.ReservarVagaPostos(&_Contratos.TransactOpts, _idPostos, _reservar, _reservador)
}

// ReservarVagaPostos is a paid mutator transaction binding the contract method 0x6dba52dc.
//
// Solidity: function reservarVagaPostos(string[] _idPostos, bool _reservar, string _reservador) returns()
func (_Contratos *ContratosTransactorSession) ReservarVagaPostos(_idPostos []string, _reservar bool, _reservador string) (*types.Transaction, error) {
	return _Contratos.Contract.ReservarVagaPostos(&_Contratos.TransactOpts, _idPostos, _reservar, _reservador)
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
