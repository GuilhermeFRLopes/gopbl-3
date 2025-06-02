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
	Bin: "0x6080604052348015600f57600080fd5b5061188d8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063527ce6f31461006757806355bb5cbe146100975780635dbafe29146100c75780636dba52dc146100e35780638f8aa6f1146100ff57806395ec3afb14610135575b600080fd5b610081600480360381019061007c9190610ae3565b610153565b60405161008e9190610ba0565b60405180910390f35b6100b160048036038101906100ac9190610cf7565b6101ff565b6040516100be9190610d5b565b60405180910390f35b6100e160048036038101906100dc9190610dac565b61023c565b005b6100fd60048036038101906100f89190610f8d565b610428565b005b61011960048036038101906101149190610cf7565b61073b565b60405161012c9796959493929190611027565b60405180910390f35b61013d6109c0565b60405161014a91906111be565b60405180910390f35b6001818154811061016357600080fd5b90600052602060002001600091509050805461017e9061120f565b80601f01602080910402602001604051908101604052809291908181526020018280546101aa9061120f565b80156101f75780601f106101cc576101008083540402835291602001916101f7565b820191906000526020600020905b8154815290600101906020018083116101da57829003601f168201915b505050505081565b600080600083604051610212919061127c565b908152602001604051809103902090508060030160009054906101000a900460ff16915050919050565b6000808660405161024d919061127c565b908152602001604051809103902060000180546102699061120f565b9050146102ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a2906112df565b60405180910390fd5b6040518060e0016040528086815260200185815260200184815260200160001515815260200183815260200182815260200160405180602001604052806000815250815250600086604051610300919061127c565b9081526020016040518091039020600082015181600001908161032391906114ab565b50602082015181600101556040820151816002015560608201518160030160006101000a81548160ff021916908315150217905550608082015181600401908161036d91906114ab565b5060a082015181600501908161038391906114ab565b5060c082015181600601908161039991906114ab565b509050506001859080600181540180825580915050600190039060005260206000200160009091909190915090816103d191906114ab565b50846040516103e0919061127c565b60405180910390207f8c9754c7fc6b0df199ea5d8a036af296a20873a52888d9fbc75ff5ac9aaa2442828460405161041992919061157d565b60405180910390a25050505050565b816105f25760005b6001805490508110156105ec57600060018281548110610453576104526115b4565b5b9060005260206000200180546104689061120f565b80601f01602080910402602001604051908101604052809291908181526020018280546104949061120f565b80156104e15780601f106104b6576101008083540402835291602001916104e1565b820191906000526020600020905b8154815290600101906020018083116104c457829003601f168201915b50505050509050600080826040516104f9919061127c565b908152602001604051809103902090508060030160009054906101000a900460ff16156105dd578380519060200120816006016040516105399190611686565b6040518091039020036105dc5760008160030160006101000a81548160ff0219169083151502179055508381600601908161057491906114ab565b50806000016040516105869190611720565b60405180910390207faf99a1509360a139a5b5aab49ac6571796571db04782badc6574f3aa7014b98a8260030160009054906101000a900460ff16836004016040516105d39291906117bb565b60405180910390a25b5b50508080600101915050610430565b50610736565b60005b835181101561073457600080858381518110610614576106136115b4565b5b6020026020010151604051610629919061127c565b90815260200160405180910390209050600115158160030160009054906101000a900460ff16151514610691576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068890611837565b60405180910390fd5b60018160030160006101000a81548160ff021916908315150217905550828160060190816106bf91906114ab565b50806000016040516106d19190611720565b60405180910390207faf99a1509360a139a5b5aab49ac6571796571db04782badc6574f3aa7014b98a8260030160009054906101000a900460ff168360040160405161071e9291906117bb565b60405180910390a25080806001019150506105f5565b505b505050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546107749061120f565b80601f01602080910402602001604051908101604052809291908181526020018280546107a09061120f565b80156107ed5780601f106107c2576101008083540402835291602001916107ed565b820191906000526020600020905b8154815290600101906020018083116107d057829003601f168201915b5050505050908060010154908060020154908060030160009054906101000a900460ff16908060040180546108219061120f565b80601f016020809104026020016040519081016040528092919081815260200182805461084d9061120f565b801561089a5780601f1061086f5761010080835404028352916020019161089a565b820191906000526020600020905b81548152906001019060200180831161087d57829003601f168201915b5050505050908060050180546108af9061120f565b80601f01602080910402602001604051908101604052809291908181526020018280546108db9061120f565b80156109285780601f106108fd57610100808354040283529160200191610928565b820191906000526020600020905b81548152906001019060200180831161090b57829003601f168201915b50505050509080600601805461093d9061120f565b80601f01602080910402602001604051908101604052809291908181526020018280546109699061120f565b80156109b65780601f1061098b576101008083540402835291602001916109b6565b820191906000526020600020905b81548152906001019060200180831161099957829003601f168201915b5050505050905087565b60606001805480602002602001604051908101604052809291908181526020016000905b82821015610a90578382906000526020600020018054610a039061120f565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2f9061120f565b8015610a7c5780601f10610a5157610100808354040283529160200191610a7c565b820191906000526020600020905b815481529060010190602001808311610a5f57829003601f168201915b5050505050815260200190600101906109e4565b50505050905090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610ac081610aad565b8114610acb57600080fd5b50565b600081359050610add81610ab7565b92915050565b600060208284031215610af957610af8610aa3565b5b6000610b0784828501610ace565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b4a578082015181840152602081019050610b2f565b60008484015250505050565b6000601f19601f8301169050919050565b6000610b7282610b10565b610b7c8185610b1b565b9350610b8c818560208601610b2c565b610b9581610b56565b840191505092915050565b60006020820190508181036000830152610bba8184610b67565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610c0482610b56565b810181811067ffffffffffffffff82111715610c2357610c22610bcc565b5b80604052505050565b6000610c36610a99565b9050610c428282610bfb565b919050565b600067ffffffffffffffff821115610c6257610c61610bcc565b5b610c6b82610b56565b9050602081019050919050565b82818337600083830152505050565b6000610c9a610c9584610c47565b610c2c565b905082815260208101848484011115610cb657610cb5610bc7565b5b610cc1848285610c78565b509392505050565b600082601f830112610cde57610cdd610bc2565b5b8135610cee848260208601610c87565b91505092915050565b600060208284031215610d0d57610d0c610aa3565b5b600082013567ffffffffffffffff811115610d2b57610d2a610aa8565b5b610d3784828501610cc9565b91505092915050565b60008115159050919050565b610d5581610d40565b82525050565b6000602082019050610d706000830184610d4c565b92915050565b6000819050919050565b610d8981610d76565b8114610d9457600080fd5b50565b600081359050610da681610d80565b92915050565b600080600080600060a08688031215610dc857610dc7610aa3565b5b600086013567ffffffffffffffff811115610de657610de5610aa8565b5b610df288828901610cc9565b9550506020610e0388828901610d97565b9450506040610e1488828901610d97565b935050606086013567ffffffffffffffff811115610e3557610e34610aa8565b5b610e4188828901610cc9565b925050608086013567ffffffffffffffff811115610e6257610e61610aa8565b5b610e6e88828901610cc9565b9150509295509295909350565b600067ffffffffffffffff821115610e9657610e95610bcc565b5b602082029050602081019050919050565b600080fd5b6000610ebf610eba84610e7b565b610c2c565b90508083825260208201905060208402830185811115610ee257610ee1610ea7565b5b835b81811015610f2957803567ffffffffffffffff811115610f0757610f06610bc2565b5b808601610f148982610cc9565b85526020850194505050602081019050610ee4565b5050509392505050565b600082601f830112610f4857610f47610bc2565b5b8135610f58848260208601610eac565b91505092915050565b610f6a81610d40565b8114610f7557600080fd5b50565b600081359050610f8781610f61565b92915050565b600080600060608486031215610fa657610fa5610aa3565b5b600084013567ffffffffffffffff811115610fc457610fc3610aa8565b5b610fd086828701610f33565b9350506020610fe186828701610f78565b925050604084013567ffffffffffffffff81111561100257611001610aa8565b5b61100e86828701610cc9565b9150509250925092565b61102181610d76565b82525050565b600060e0820190508181036000830152611041818a610b67565b90506110506020830189611018565b61105d6040830188611018565b61106a6060830187610d4c565b818103608083015261107c8186610b67565b905081810360a08301526110908185610b67565b905081810360c08301526110a48184610b67565b905098975050505050505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b60006110fa82610b10565b61110481856110de565b9350611114818560208601610b2c565b61111d81610b56565b840191505092915050565b600061113483836110ef565b905092915050565b6000602082019050919050565b6000611154826110b2565b61115e81856110bd565b935083602082028501611170856110ce565b8060005b858110156111ac578484038952815161118d8582611128565b94506111988361113c565b925060208a01995050600181019050611174565b50829750879550505050505092915050565b600060208201905081810360008301526111d88184611149565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061122757607f821691505b60208210810361123a576112396111e0565b5b50919050565b600081905092915050565b600061125682610b10565b6112608185611240565b9350611270818560208601610b2c565b80840191505092915050565b6000611288828461124b565b915081905092915050565b7f506f73746f2064652072656361726761206a61206361646173747261646f0000600082015250565b60006112c9601e83610b1b565b91506112d482611293565b602082019050919050565b600060208201905081810360008301526112f8816112bc565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026113617fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611324565b61136b8683611324565b95508019841693508086168417925050509392505050565b6000819050919050565b60006113a86113a361139e84610aad565b611383565b610aad565b9050919050565b6000819050919050565b6113c28361138d565b6113d66113ce826113af565b848454611331565b825550505050565b600090565b6113eb6113de565b6113f68184846113b9565b505050565b5b8181101561141a5761140f6000826113e3565b6001810190506113fc565b5050565b601f82111561145f57611430816112ff565b61143984611314565b81016020851015611448578190505b61145c61145485611314565b8301826113fb565b50505b505050565b600082821c905092915050565b600061148260001984600802611464565b1980831691505092915050565b600061149b8383611471565b9150826002028217905092915050565b6114b482610b10565b67ffffffffffffffff8111156114cd576114cc610bcc565b5b6114d7825461120f565b6114e282828561141e565b600060209050601f8311600181146115155760008415611503578287015190505b61150d858261148f565b865550611575565b601f198416611523866112ff565b60005b8281101561154b57848901518255600182019150602085019450602081019050611526565b868310156115685784890151611564601f891682611471565b8355505b6001600288020188555050505b505050505050565b600060408201905081810360008301526115978185610b67565b905081810360208301526115ab8184610b67565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081905092915050565b60008190508160005260206000209050919050565b600081546116108161120f565b61161a81866115e3565b94506001821660008114611635576001811461164a5761167d565b60ff198316865281151582028601935061167d565b611653856115ee565b60005b8381101561167557815481890152600182019150602081019050611656565b838801955050505b50505092915050565b60006116928284611603565b915081905092915050565b600081546116aa8161120f565b6116b48186611240565b945060018216600081146116cf57600181146116e457611717565b60ff1983168652811515820286019350611717565b6116ed856112ff565b60005b8381101561170f578154818901526001820191506020810190506116f0565b838801955050505b50505092915050565b600061172c828461169d565b915081905092915050565b600081546117448161120f565b61174e8186610b1b565b94506001821660008114611769576001811461177f576117b2565b60ff1983168652811515602002860193506117b2565b611788856112ff565b60005b838110156117aa5781548189015260018201915060208101905061178b565b808801955050505b50505092915050565b60006040820190506117d06000830185610d4c565b81810360208301526117e28184611737565b90509392505050565b7f506f73746f206a6120657374612072657365727661646f000000000000000000600082015250565b6000611821601783610b1b565b915061182c826117eb565b602082019050919050565b6000602082019050818103600083015261185081611814565b905091905056fea2646970667358221220e4a9a2f236a00aede70040328e50d8dc6c7c9f4de6fd4c0acf814c1796dffdb964736f6c634300081e0033",
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
