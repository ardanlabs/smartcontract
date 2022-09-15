// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bankproxy

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
)

// BankproxyMetaData contains all meta data concerning the Bankproxy contract.
var BankproxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"api\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"API\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"api\",\"type\":\"address\"}],\"name\":\"UpgradeAPI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200195538038062001955833981810160405281019062000037919062000129565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200015b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000f182620000c4565b9050919050565b6200010381620000e4565b81146200010f57600080fd5b50565b6000815190506200012381620000f8565b92915050565b600060208284031215620001425762000141620000bf565b5b6000620001528482850162000112565b91505092915050565b6117ea806200016b6000396000f3fe60806040526004361061007b5760003560e01c8063c32ab72e1161004e578063c32ab72e1461010b578063e63f341f14610134578063ed21248c14610171578063fa84fd8e1461017b5761007b565b80630ef678871461008057806357ea89b6146100ab5780637d7b0099146100b5578063b4a99a4e146100e0575b600080fd5b34801561008c57600080fd5b506100956101a4565b6040516100a29190610ca0565b60405180910390f35b6100b36101eb565b005b3480156100c157600080fd5b506100ca6103ae565b6040516100d79190610cfc565b60405180910390f35b3480156100ec57600080fd5b506100f56103d4565b6040516101029190610cfc565b60405180910390f35b34801561011757600080fd5b50610132600480360381019061012d9190610d4d565b6103f8565b005b34801561014057600080fd5b5061015b60048036038101906101569190610d4d565b6104f2565b6040516101689190610ca0565b60405180910390f35b610179610594565b005b34801561018757600080fd5b506101a2600480360381019061019d9190610e0b565b610693565b005b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003390506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403610272576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026990610ef0565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156102fc573d6000803e3d6000fd5b506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61036c33610873565b61037583610a36565b604051602001610386929190610ff3565b6040516020818303038152906040526040516103a2919061108e565b60405180910390a15050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461045057600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6104bb82610873565b6040516020016104cb91906110d6565b6040516020818303038152906040526040516104e7919061108e565b60405180910390a150565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461054d57600080fd5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b34600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105e3919061113a565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61061433610873565b61065c600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610a36565b60405160200161066d9291906111ba565b604051602081830303815290604052604051610689919061108e565b60405180910390a1565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106eb57600080fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16868686868660405160240161073f9594939291906112ce565b6040516020818303038152906040527f3b197bd8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516107c99190611363565b600060405180830381855af49150503d8060008114610804576040519150601f19603f3d011682016040523d82523d6000602084013e610809565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61083782610bbe565b60405160200161084791906113a0565b604051602081830303815290604052604051610863919061108e565b60405180910390a1505050505050565b60606000602867ffffffffffffffff811115610892576108916113d5565b5b6040519080825280601f01601f1916602001820160405280156108c45781602001600182028036833780820191505090505b50905060005b6014811015610a2c5760008160136108e29190611404565b60086108ee9190611438565b60026108fa91906115ad565b8573ffffffffffffffffffffffffffffffffffffffff1661091b9190611627565b60f81b9050600060108260f81c6109329190611665565b60f81b905060008160f81c60106109499190611696565b8360f81c61095791906116d3565b60f81b905061096582610c41565b858560026109739190611438565b8151811061098457610983611708565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506109bc81610c41565b8560018660026109cc9190611438565b6109d6919061113a565b815181106109e7576109e6611708565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610a2490611737565b9150506108ca565b5080915050919050565b606060008203610a7d576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610bb9565b600082905060005b60008214610aaf578080610a9890611737565b915050600a82610aa89190611627565b9150610a85565b60008167ffffffffffffffff811115610acb57610aca6113d5565b5b6040519080825280601f01601f191660200182016040528015610afd5781602001600182028036833780820191505090505b50905060008290505b60008614610bb157600181610b1b9190611404565b90506000600a8088610b2d9190611627565b610b379190611438565b87610b429190611404565b6030610b4e919061177f565b905060008160f81b905080848481518110610b6c57610b6b611708565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610ba89190611627565b97505050610b06565b819450505050505b919050565b60608115610c03576040518060400160405280600481526020017f74727565000000000000000000000000000000000000000000000000000000008152509050610c3c565b6040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525090505b919050565b6000600a8260f81c60ff161015610c6c5760308260f81c610c62919061177f565b60f81b9050610c82565b60578260f81c610c7c919061177f565b60f81b90505b919050565b6000819050919050565b610c9a81610c87565b82525050565b6000602082019050610cb56000830184610c91565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ce682610cbb565b9050919050565b610cf681610cdb565b82525050565b6000602082019050610d116000830184610ced565b92915050565b600080fd5b600080fd5b610d2a81610cdb565b8114610d3557600080fd5b50565b600081359050610d4781610d21565b92915050565b600060208284031215610d6357610d62610d17565b5b6000610d7184828501610d38565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610d9f57610d9e610d7a565b5b8235905067ffffffffffffffff811115610dbc57610dbb610d7f565b5b602083019150836020820283011115610dd857610dd7610d84565b5b9250929050565b610de881610c87565b8114610df357600080fd5b50565b600081359050610e0581610ddf565b92915050565b600080600080600060808688031215610e2757610e26610d17565b5b6000610e3588828901610d38565b955050602086013567ffffffffffffffff811115610e5657610e55610d1c565b5b610e6288828901610d89565b94509450506040610e7588828901610df6565b9250506060610e8688828901610df6565b9150509295509295909350565b600082825260208201905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000610eda601283610e93565b9150610ee582610ea4565b602082019050919050565b60006020820190508181036000830152610f0981610ecd565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b83811015610f6a578082015181840152602081019050610f4f565b60008484015250505050565b6000610f8182610f36565b610f8b8185610f41565b9350610f9b818560208601610f4c565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b6000610ffe82610f10565b60098201915061100e8285610f76565b915061101982610fa7565b6009820191506110298284610f76565b915061103482610fcd565b6001820191508190509392505050565b6000601f19601f8301169050919050565b600061106082610f36565b61106a8185610e93565b935061107a818560208601610f4c565b61108381611044565b840191505092915050565b600060208201905081810360008301526110a88184611055565b905092915050565b7f757067726164655b000000000000000000000000000000000000000000000000815250565b60006110e1826110b0565b6008820191506110f18284610f76565b91506110fc82610fcd565b60018201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061114582610c87565b915061115083610c87565b92508282019050808211156111685761116761110b565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b60006111c58261116e565b6008820191506111d58285610f76565b91506111e082611194565b600a820191506111f08284610f76565b91506111fb82610fcd565b6001820191508190509392505050565b600082825260208201905092915050565b6000819050919050565b61122f81610cdb565b82525050565b60006112418383611226565b60208301905092915050565b600061125c6020840184610d38565b905092915050565b6000602082019050919050565b600061127d838561120b565b93506112888261121c565b8060005b858110156112c15761129e828461124d565b6112a88882611235565b97506112b383611264565b92505060018101905061128c565b5085925050509392505050565b60006080820190506112e36000830188610ced565b81810360208301526112f6818688611271565b90506113056040830185610c91565b6113126060830184610c91565b9695505050505050565b600081519050919050565b600081905092915050565b600061133d8261131c565b6113478185611327565b9350611357818560208601610f4c565b80840191505092915050565b600061136f8284611332565b915081905092915050565b7f7265636f6e63696c655b00000000000000000000000000000000000000000000815250565b60006113ab8261137a565b600a820191506113bb8284610f76565b91506113c682610fcd565b60018201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600061140f82610c87565b915061141a83610c87565b92508282039050818111156114325761143161110b565b5b92915050565b600061144382610c87565b915061144e83610c87565b925082820261145c81610c87565b915082820484148315176114735761147261110b565b5b5092915050565b60008160011c9050919050565b6000808291508390505b60018511156114d1578086048111156114ad576114ac61110b565b5b60018516156114bc5780820291505b80810290506114ca8561147a565b9450611491565b94509492505050565b6000826114ea57600190506115a6565b816114f857600090506115a6565b816001811461150e576002811461151857611547565b60019150506115a6565b60ff84111561152a5761152961110b565b5b8360020a9150848211156115415761154061110b565b5b506115a6565b5060208310610133831016604e8410600b841016171561157c5782820a9050838111156115775761157661110b565b5b6115a6565b6115898484846001611487565b925090508184048111156115a05761159f61110b565b5b81810290505b9392505050565b60006115b882610c87565b91506115c383610c87565b92506115f07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846114da565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061163282610c87565b915061163d83610c87565b92508261164d5761164c6115f8565b5b828204905092915050565b600060ff82169050919050565b600061167082611658565b915061167b83611658565b92508261168b5761168a6115f8565b5b828204905092915050565b60006116a182611658565b91506116ac83611658565b92508282026116ba81611658565b91508082146116cc576116cb61110b565b5b5092915050565b60006116de82611658565b91506116e983611658565b9250828203905060ff8111156117025761170161110b565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061174282610c87565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117745761177361110b565b5b600182019050919050565b600061178a82611658565b915061179583611658565b9250828201905060ff8111156117ae576117ad61110b565b5b9291505056fea264697066735822122013be7cee9e27c6e351199ae23cbbb7907d82ac09d63f3ed17df4064c55f5de4f64736f6c63430008110033",
}

// BankproxyABI is the input ABI used to generate the binding from.
// Deprecated: Use BankproxyMetaData.ABI instead.
var BankproxyABI = BankproxyMetaData.ABI

// BankproxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankproxyMetaData.Bin instead.
var BankproxyBin = BankproxyMetaData.Bin

// DeployBankproxy deploys a new Ethereum contract, binding an instance of Bankproxy to it.
func DeployBankproxy(auth *bind.TransactOpts, backend bind.ContractBackend, api common.Address) (common.Address, *types.Transaction, *Bankproxy, error) {
	parsed, err := BankproxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankproxyBin), backend, api)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bankproxy{BankproxyCaller: BankproxyCaller{contract: contract}, BankproxyTransactor: BankproxyTransactor{contract: contract}, BankproxyFilterer: BankproxyFilterer{contract: contract}}, nil
}

// Bankproxy is an auto generated Go binding around an Ethereum contract.
type Bankproxy struct {
	BankproxyCaller     // Read-only binding to the contract
	BankproxyTransactor // Write-only binding to the contract
	BankproxyFilterer   // Log filterer for contract events
}

// BankproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankproxySession struct {
	Contract     *Bankproxy        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankproxyCallerSession struct {
	Contract *BankproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BankproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankproxyTransactorSession struct {
	Contract     *BankproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BankproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankproxyRaw struct {
	Contract *Bankproxy // Generic contract binding to access the raw methods on
}

// BankproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankproxyCallerRaw struct {
	Contract *BankproxyCaller // Generic read-only contract binding to access the raw methods on
}

// BankproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankproxyTransactorRaw struct {
	Contract *BankproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBankproxy creates a new instance of Bankproxy, bound to a specific deployed contract.
func NewBankproxy(address common.Address, backend bind.ContractBackend) (*Bankproxy, error) {
	contract, err := bindBankproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bankproxy{BankproxyCaller: BankproxyCaller{contract: contract}, BankproxyTransactor: BankproxyTransactor{contract: contract}, BankproxyFilterer: BankproxyFilterer{contract: contract}}, nil
}

// NewBankproxyCaller creates a new read-only instance of Bankproxy, bound to a specific deployed contract.
func NewBankproxyCaller(address common.Address, caller bind.ContractCaller) (*BankproxyCaller, error) {
	contract, err := bindBankproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankproxyCaller{contract: contract}, nil
}

// NewBankproxyTransactor creates a new write-only instance of Bankproxy, bound to a specific deployed contract.
func NewBankproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*BankproxyTransactor, error) {
	contract, err := bindBankproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankproxyTransactor{contract: contract}, nil
}

// NewBankproxyFilterer creates a new log filterer instance of Bankproxy, bound to a specific deployed contract.
func NewBankproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*BankproxyFilterer, error) {
	contract, err := bindBankproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankproxyFilterer{contract: contract}, nil
}

// bindBankproxy binds a generic wrapper to an already deployed contract.
func bindBankproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankproxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankproxy *BankproxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankproxy.Contract.BankproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankproxy *BankproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankproxy.Contract.BankproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankproxy *BankproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankproxy.Contract.BankproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankproxy *BankproxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankproxy *BankproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankproxy *BankproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankproxy.Contract.contract.Transact(opts, method, params...)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankproxy *BankproxyCaller) API(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankproxy.contract.Call(opts, &out, "API")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankproxy *BankproxySession) API() (common.Address, error) {
	return _Bankproxy.Contract.API(&_Bankproxy.CallOpts)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankproxy *BankproxyCallerSession) API() (common.Address, error) {
	return _Bankproxy.Contract.API(&_Bankproxy.CallOpts)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bankproxy *BankproxyCaller) AccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bankproxy.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bankproxy *BankproxySession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bankproxy.Contract.AccountBalance(&_Bankproxy.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bankproxy *BankproxyCallerSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bankproxy.Contract.AccountBalance(&_Bankproxy.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bankproxy *BankproxyCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bankproxy.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bankproxy *BankproxySession) Balance() (*big.Int, error) {
	return _Bankproxy.Contract.Balance(&_Bankproxy.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bankproxy *BankproxyCallerSession) Balance() (*big.Int, error) {
	return _Bankproxy.Contract.Balance(&_Bankproxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankproxy *BankproxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankproxy.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankproxy *BankproxySession) Owner() (common.Address, error) {
	return _Bankproxy.Contract.Owner(&_Bankproxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankproxy *BankproxyCallerSession) Owner() (common.Address, error) {
	return _Bankproxy.Contract.Owner(&_Bankproxy.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankproxy *BankproxyTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankproxy.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankproxy *BankproxySession) Deposit() (*types.Transaction, error) {
	return _Bankproxy.Contract.Deposit(&_Bankproxy.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankproxy *BankproxyTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bankproxy.Contract.Deposit(&_Bankproxy.TransactOpts)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bankproxy *BankproxyTransactor) Reconcile(opts *bind.TransactOpts, winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankproxy.contract.Transact(opts, "Reconcile", winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bankproxy *BankproxySession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankproxy.Contract.Reconcile(&_Bankproxy.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bankproxy *BankproxyTransactorSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankproxy.Contract.Reconcile(&_Bankproxy.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// UpgradeAPI is a paid mutator transaction binding the contract method 0xc32ab72e.
//
// Solidity: function UpgradeAPI(address api) returns()
func (_Bankproxy *BankproxyTransactor) UpgradeAPI(opts *bind.TransactOpts, api common.Address) (*types.Transaction, error) {
	return _Bankproxy.contract.Transact(opts, "UpgradeAPI", api)
}

// UpgradeAPI is a paid mutator transaction binding the contract method 0xc32ab72e.
//
// Solidity: function UpgradeAPI(address api) returns()
func (_Bankproxy *BankproxySession) UpgradeAPI(api common.Address) (*types.Transaction, error) {
	return _Bankproxy.Contract.UpgradeAPI(&_Bankproxy.TransactOpts, api)
}

// UpgradeAPI is a paid mutator transaction binding the contract method 0xc32ab72e.
//
// Solidity: function UpgradeAPI(address api) returns()
func (_Bankproxy *BankproxyTransactorSession) UpgradeAPI(api common.Address) (*types.Transaction, error) {
	return _Bankproxy.Contract.UpgradeAPI(&_Bankproxy.TransactOpts, api)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankproxy *BankproxyTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankproxy.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankproxy *BankproxySession) Withdraw() (*types.Transaction, error) {
	return _Bankproxy.Contract.Withdraw(&_Bankproxy.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankproxy *BankproxyTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bankproxy.Contract.Withdraw(&_Bankproxy.TransactOpts)
}

// BankproxyEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bankproxy contract.
type BankproxyEventLogIterator struct {
	Event *BankproxyEventLog // Event containing the contract specifics and raw log

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
func (it *BankproxyEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankproxyEventLog)
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
		it.Event = new(BankproxyEventLog)
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
func (it *BankproxyEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankproxyEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankproxyEventLog represents a EventLog event raised by the Bankproxy contract.
type BankproxyEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankproxy *BankproxyFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankproxyEventLogIterator, error) {

	logs, sub, err := _Bankproxy.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankproxyEventLogIterator{contract: _Bankproxy.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankproxy *BankproxyFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankproxyEventLog) (event.Subscription, error) {

	logs, sub, err := _Bankproxy.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankproxyEventLog)
				if err := _Bankproxy.contract.UnpackLog(event, "EventLog", log); err != nil {
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

// ParseEventLog is a log parse operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankproxy *BankproxyFilterer) ParseEventLog(log types.Log) (*BankproxyEventLog, error) {
	event := new(BankproxyEventLog)
	if err := _Bankproxy.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
