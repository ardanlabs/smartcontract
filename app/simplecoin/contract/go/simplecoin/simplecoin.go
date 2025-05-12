// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simplecoin

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

// SimplecoinMetaData contains all meta data concerning the Simplecoin contract.
var SimplecoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"EventFrozenAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EventTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CoinBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"freeze\",\"type\":\"bool\"}],\"name\":\"FreezeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FrozenAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516117ba3803806117ba83398181016040528101906100319190610210565b335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506100a05f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16826100a660201b60201c565b506102c3565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100fd575f5ffd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546101499190610268565b925050819055508173ffffffffffffffffffffffffffffffffffffffff165f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45836040516101cd91906102aa565b60405180910390a35050565b5f5ffd5b5f819050919050565b6101ef816101dd565b81146101f9575f5ffd5b50565b5f8151905061020a816101e6565b92915050565b5f60208284031215610225576102246101d9565b5b5f610232848285016101fc565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610272826101dd565b915061027d836101dd565b92508282019050808211156102955761029461023b565b5b92915050565b6102a4816101dd565b82525050565b5f6020820190506102bd5f83018461029b565b92915050565b6114ea806102d05f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80630f6798a51461006457806369ca02dd146100805780638bdf15341461009c578063ab31471e146100cc578063b4a99a4e146100fc578063d16a7a4b1461011a575b5f5ffd5b61007e60048036038101906100799190610d15565b610136565b005b61009a60048036038101906100959190610d15565b610269565b005b6100b660048036038101906100b19190610d53565b610443565b6040516100c39190610d8d565b60405180910390f35b6100e660048036038101906100e19190610d53565b610458565b6040516100f39190610dc0565b60405180910390f35b610104610475565b6040516101119190610de8565b60405180910390f35b610134600480360381019061012f9190610e2b565b610499565b005b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461018d575f5ffd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546101d99190610e96565b925050819055508173ffffffffffffffffffffffffffffffffffffffff165f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a458360405161025d9190610d8d565b60405180910390a35050565b5f610275338484610581565b9050805f0151156102c15780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b89190610f39565b60405180910390fd5b8160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461030d9190610f59565b925050819055508160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546103609190610e96565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610391836107fc565b61039a8561097a565b6103a33361097a565b6040516020016103b593929190611038565b6040516020818303038152906040526040516103d19190610f39565b60405180910390a18273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45846040516104369190610d8d565b60405180910390a3505050565b6001602052805f5260405f205f915090505481565b6002602052805f5260405f205f915054906101000a900460ff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104f0575f5ffd5b8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507f3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d8282604051610575929190611095565b60405180910390a15050565b610589610c69565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610601576105fa6040518060400160405280601f81526020017f63616e27742073656e64206d6f6e657920746f20616464726573732030783000815250610b33565b90506107f5565b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610710576107096040518060400160405280601781526020017f696e737566666963656e742066756e6473202062616c3a0000000000000000008152506106c560015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546107fc565b6040518060400160405280600a81526020017f2020616d6f756e743a2000000000000000000000000000000000000000000000815250610704866107fc565b610b5a565b90506107f5565b60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546107979190610e96565b116107ea576107e36040518060400160405280601081526020017f696e76616c696420616d6f756e743a20000000000000000000000000000000008152506107de846107fc565b610ba9565b90506107f5565b6107f2610bf2565b90505b9392505050565b60605f8203610842576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610975565b5f8290505f5b5f821461087157808061085a906110bc565b915050600a8261086a9190611130565b9150610848565b5f8167ffffffffffffffff81111561088c5761088b611160565b5b6040519080825280601f01601f1916602001820160405280156108be5781602001600182028036833780820191505090505b5090505f8290505b5f861461096d576001816108da9190610f59565b90505f600a80886108eb9190611130565b6108f5919061118d565b876109009190610f59565b603061090c91906111da565b90505f8160f81b9050808484815181106109295761092861120e565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a886109649190611130565b975050506108c6565b819450505050505b919050565b60605f602867ffffffffffffffff81111561099857610997611160565b5b6040519080825280601f01601f1916602001820160405280156109ca5781602001600182028036833780820191505090505b5090505f5f90505b6014811015610b29575f8160136109e99190610f59565b60086109f5919061118d565b6002610a01919061136a565b8573ffffffffffffffffffffffffffffffffffffffff16610a229190611130565b60f81b90505f60108260f81c610a3891906113b4565b60f81b90505f8160f81c6010610a4e91906113e4565b8360f81c610a5c9190611420565b60f81b9050610a6a82610c24565b85856002610a78919061118d565b81518110610a8957610a8861120e565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350610ac081610c24565b856001866002610ad0919061118d565b610ada9190610e96565b81518110610aeb57610aea61120e565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a90535050505080806001019150506109d2565b5080915050919050565b610b3b610c69565b6040518060400160405280600115158152602001838152509050919050565b610b62610c69565b604051806040016040528060011515815260200186868686604051602001610b8d9493929190611454565b6040516020818303038152906040528152509050949350505050565b610bb1610c69565b60405180604001604052806001151581526020018484604051602001610bd8929190611491565b604051602081830303815290604052815250905092915050565b610bfa610c69565b60405180604001604052805f1515815260200160405180602001604052805f815250815250905090565b5f600a8260f81c60ff161015610c4e5760308260f81c610c4491906111da565b60f81b9050610c64565b60578260f81c610c5e91906111da565b60f81b90505b919050565b60405180604001604052805f15158152602001606081525090565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610cb182610c88565b9050919050565b610cc181610ca7565b8114610ccb575f5ffd5b50565b5f81359050610cdc81610cb8565b92915050565b5f819050919050565b610cf481610ce2565b8114610cfe575f5ffd5b50565b5f81359050610d0f81610ceb565b92915050565b5f5f60408385031215610d2b57610d2a610c84565b5b5f610d3885828601610cce565b9250506020610d4985828601610d01565b9150509250929050565b5f60208284031215610d6857610d67610c84565b5b5f610d7584828501610cce565b91505092915050565b610d8781610ce2565b82525050565b5f602082019050610da05f830184610d7e565b92915050565b5f8115159050919050565b610dba81610da6565b82525050565b5f602082019050610dd35f830184610db1565b92915050565b610de281610ca7565b82525050565b5f602082019050610dfb5f830184610dd9565b92915050565b610e0a81610da6565b8114610e14575f5ffd5b50565b5f81359050610e2581610e01565b92915050565b5f5f60408385031215610e4157610e40610c84565b5b5f610e4e85828601610cce565b9250506020610e5f85828601610e17565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610ea082610ce2565b9150610eab83610ce2565b9250828201905080821115610ec357610ec2610e69565b5b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610f0b82610ec9565b610f158185610ed3565b9350610f25818560208601610ee3565b610f2e81610ef1565b840191505092915050565b5f6020820190508181035f830152610f518184610f01565b905092915050565b5f610f6382610ce2565b9150610f6e83610ce2565b9250828203905081811115610f8657610f85610e69565b5b92915050565b7f7472616e73666572656420000000000000000000000000000000000000000000815250565b5f81905092915050565b5f610fc682610ec9565b610fd08185610fb2565b9350610fe0818560208601610ee3565b80840191505092915050565b7f20746f2000000000000000000000000000000000000000000000000000000000815250565b7f2066726f6d200000000000000000000000000000000000000000000000000000815250565b5f61104282610f8c565b600b820191506110528286610fbc565b915061105d82610fec565b60048201915061106d8285610fbc565b915061107882611012565b6006820191506110888284610fbc565b9150819050949350505050565b5f6040820190506110a85f830185610dd9565b6110b56020830184610db1565b9392505050565b5f6110c682610ce2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110f8576110f7610e69565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61113a82610ce2565b915061114583610ce2565b92508261115557611154611103565b5b828204905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f61119782610ce2565b91506111a283610ce2565b92508282026111b081610ce2565b915082820484148315176111c7576111c6610e69565b5b5092915050565b5f60ff82169050919050565b5f6111e4826111ce565b91506111ef836111ce565b9250828201905060ff81111561120857611207610e69565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b60018511156112905780860481111561126c5761126b610e69565b5b600185161561127b5780820291505b80810290506112898561123b565b9450611250565b94509492505050565b5f826112a85760019050611363565b816112b5575f9050611363565b81600181146112cb57600281146112d557611304565b6001915050611363565b60ff8411156112e7576112e6610e69565b5b8360020a9150848211156112fe576112fd610e69565b5b50611363565b5060208310610133831016604e8410600b84101617156113395782820a90508381111561133457611333610e69565b5b611363565b6113468484846001611247565b9250905081840481111561135d5761135c610e69565b5b81810290505b9392505050565b5f61137482610ce2565b915061137f83610ce2565b92506113ac7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611299565b905092915050565b5f6113be826111ce565b91506113c9836111ce565b9250826113d9576113d8611103565b5b828204905092915050565b5f6113ee826111ce565b91506113f9836111ce565b9250828202611407816111ce565b915080821461141957611418610e69565b5b5092915050565b5f61142a826111ce565b9150611435836111ce565b9250828203905060ff81111561144e5761144d610e69565b5b92915050565b5f61145f8287610fbc565b915061146b8286610fbc565b91506114778285610fbc565b91506114838284610fbc565b915081905095945050505050565b5f61149c8285610fbc565b91506114a88284610fbc565b9150819050939250505056fea2646970667358221220cef80999912ba0f99bba8cacd25da277eb904dc38fa3a798c655e0b5f732e47d64736f6c634300081e0033",
}

// SimplecoinABI is the input ABI used to generate the binding from.
// Deprecated: Use SimplecoinMetaData.ABI instead.
var SimplecoinABI = SimplecoinMetaData.ABI

// SimplecoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimplecoinMetaData.Bin instead.
var SimplecoinBin = SimplecoinMetaData.Bin

// DeploySimplecoin deploys a new Ethereum contract, binding an instance of Simplecoin to it.
func DeploySimplecoin(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *Simplecoin, error) {
	parsed, err := SimplecoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimplecoinBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simplecoin{SimplecoinCaller: SimplecoinCaller{contract: contract}, SimplecoinTransactor: SimplecoinTransactor{contract: contract}, SimplecoinFilterer: SimplecoinFilterer{contract: contract}}, nil
}

// Simplecoin is an auto generated Go binding around an Ethereum contract.
type Simplecoin struct {
	SimplecoinCaller     // Read-only binding to the contract
	SimplecoinTransactor // Write-only binding to the contract
	SimplecoinFilterer   // Log filterer for contract events
}

// SimplecoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplecoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplecoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplecoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplecoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplecoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplecoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplecoinSession struct {
	Contract     *Simplecoin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimplecoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplecoinCallerSession struct {
	Contract *SimplecoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SimplecoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplecoinTransactorSession struct {
	Contract     *SimplecoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimplecoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplecoinRaw struct {
	Contract *Simplecoin // Generic contract binding to access the raw methods on
}

// SimplecoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplecoinCallerRaw struct {
	Contract *SimplecoinCaller // Generic read-only contract binding to access the raw methods on
}

// SimplecoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplecoinTransactorRaw struct {
	Contract *SimplecoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplecoin creates a new instance of Simplecoin, bound to a specific deployed contract.
func NewSimplecoin(address common.Address, backend bind.ContractBackend) (*Simplecoin, error) {
	contract, err := bindSimplecoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simplecoin{SimplecoinCaller: SimplecoinCaller{contract: contract}, SimplecoinTransactor: SimplecoinTransactor{contract: contract}, SimplecoinFilterer: SimplecoinFilterer{contract: contract}}, nil
}

// NewSimplecoinCaller creates a new read-only instance of Simplecoin, bound to a specific deployed contract.
func NewSimplecoinCaller(address common.Address, caller bind.ContractCaller) (*SimplecoinCaller, error) {
	contract, err := bindSimplecoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplecoinCaller{contract: contract}, nil
}

// NewSimplecoinTransactor creates a new write-only instance of Simplecoin, bound to a specific deployed contract.
func NewSimplecoinTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplecoinTransactor, error) {
	contract, err := bindSimplecoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplecoinTransactor{contract: contract}, nil
}

// NewSimplecoinFilterer creates a new log filterer instance of Simplecoin, bound to a specific deployed contract.
func NewSimplecoinFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplecoinFilterer, error) {
	contract, err := bindSimplecoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplecoinFilterer{contract: contract}, nil
}

// bindSimplecoin binds a generic wrapper to an already deployed contract.
func bindSimplecoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimplecoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplecoin *SimplecoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplecoin.Contract.SimplecoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplecoin *SimplecoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplecoin.Contract.SimplecoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplecoin *SimplecoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplecoin.Contract.SimplecoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplecoin *SimplecoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplecoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplecoin *SimplecoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplecoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplecoin *SimplecoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplecoin.Contract.contract.Transact(opts, method, params...)
}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_Simplecoin *SimplecoinCaller) CoinBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Simplecoin.contract.Call(opts, &out, "CoinBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_Simplecoin *SimplecoinSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _Simplecoin.Contract.CoinBalance(&_Simplecoin.CallOpts, arg0)
}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_Simplecoin *SimplecoinCallerSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _Simplecoin.Contract.CoinBalance(&_Simplecoin.CallOpts, arg0)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_Simplecoin *SimplecoinCaller) FrozenAccount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Simplecoin.contract.Call(opts, &out, "FrozenAccount", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_Simplecoin *SimplecoinSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _Simplecoin.Contract.FrozenAccount(&_Simplecoin.CallOpts, arg0)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_Simplecoin *SimplecoinCallerSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _Simplecoin.Contract.FrozenAccount(&_Simplecoin.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Simplecoin *SimplecoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simplecoin.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Simplecoin *SimplecoinSession) Owner() (common.Address, error) {
	return _Simplecoin.Contract.Owner(&_Simplecoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Simplecoin *SimplecoinCallerSession) Owner() (common.Address, error) {
	return _Simplecoin.Contract.Owner(&_Simplecoin.CallOpts)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_Simplecoin *SimplecoinTransactor) FreezeAccount(opts *bind.TransactOpts, target common.Address, freeze bool) (*types.Transaction, error) {
	return _Simplecoin.contract.Transact(opts, "FreezeAccount", target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_Simplecoin *SimplecoinSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _Simplecoin.Contract.FreezeAccount(&_Simplecoin.TransactOpts, target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_Simplecoin *SimplecoinTransactorSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _Simplecoin.Contract.FreezeAccount(&_Simplecoin.TransactOpts, target, freeze)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_Simplecoin *SimplecoinTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.contract.Transact(opts, "Mint", recipient, mintedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_Simplecoin *SimplecoinSession) Mint(recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.Mint(&_Simplecoin.TransactOpts, recipient, mintedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_Simplecoin *SimplecoinTransactorSession) Mint(recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.Mint(&_Simplecoin.TransactOpts, recipient, mintedAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_Simplecoin *SimplecoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.contract.Transact(opts, "Transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_Simplecoin *SimplecoinSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.Transfer(&_Simplecoin.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_Simplecoin *SimplecoinTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.Transfer(&_Simplecoin.TransactOpts, to, amount)
}

// SimplecoinEventFrozenAccountIterator is returned from FilterEventFrozenAccount and is used to iterate over the raw logs and unpacked data for EventFrozenAccount events raised by the Simplecoin contract.
type SimplecoinEventFrozenAccountIterator struct {
	Event *SimplecoinEventFrozenAccount // Event containing the contract specifics and raw log

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
func (it *SimplecoinEventFrozenAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplecoinEventFrozenAccount)
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
		it.Event = new(SimplecoinEventFrozenAccount)
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
func (it *SimplecoinEventFrozenAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplecoinEventFrozenAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplecoinEventFrozenAccount represents a EventFrozenAccount event raised by the Simplecoin contract.
type SimplecoinEventFrozenAccount struct {
	Target common.Address
	Frozen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventFrozenAccount is a free log retrieval operation binding the contract event 0x3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d.
//
// Solidity: event EventFrozenAccount(address target, bool frozen)
func (_Simplecoin *SimplecoinFilterer) FilterEventFrozenAccount(opts *bind.FilterOpts) (*SimplecoinEventFrozenAccountIterator, error) {

	logs, sub, err := _Simplecoin.contract.FilterLogs(opts, "EventFrozenAccount")
	if err != nil {
		return nil, err
	}
	return &SimplecoinEventFrozenAccountIterator{contract: _Simplecoin.contract, event: "EventFrozenAccount", logs: logs, sub: sub}, nil
}

// WatchEventFrozenAccount is a free log subscription operation binding the contract event 0x3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d.
//
// Solidity: event EventFrozenAccount(address target, bool frozen)
func (_Simplecoin *SimplecoinFilterer) WatchEventFrozenAccount(opts *bind.WatchOpts, sink chan<- *SimplecoinEventFrozenAccount) (event.Subscription, error) {

	logs, sub, err := _Simplecoin.contract.WatchLogs(opts, "EventFrozenAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplecoinEventFrozenAccount)
				if err := _Simplecoin.contract.UnpackLog(event, "EventFrozenAccount", log); err != nil {
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

// ParseEventFrozenAccount is a log parse operation binding the contract event 0x3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d.
//
// Solidity: event EventFrozenAccount(address target, bool frozen)
func (_Simplecoin *SimplecoinFilterer) ParseEventFrozenAccount(log types.Log) (*SimplecoinEventFrozenAccount, error) {
	event := new(SimplecoinEventFrozenAccount)
	if err := _Simplecoin.contract.UnpackLog(event, "EventFrozenAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimplecoinEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Simplecoin contract.
type SimplecoinEventLogIterator struct {
	Event *SimplecoinEventLog // Event containing the contract specifics and raw log

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
func (it *SimplecoinEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplecoinEventLog)
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
		it.Event = new(SimplecoinEventLog)
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
func (it *SimplecoinEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplecoinEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplecoinEventLog represents a EventLog event raised by the Simplecoin contract.
type SimplecoinEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Simplecoin *SimplecoinFilterer) FilterEventLog(opts *bind.FilterOpts) (*SimplecoinEventLogIterator, error) {

	logs, sub, err := _Simplecoin.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &SimplecoinEventLogIterator{contract: _Simplecoin.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Simplecoin *SimplecoinFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *SimplecoinEventLog) (event.Subscription, error) {

	logs, sub, err := _Simplecoin.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplecoinEventLog)
				if err := _Simplecoin.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_Simplecoin *SimplecoinFilterer) ParseEventLog(log types.Log) (*SimplecoinEventLog, error) {
	event := new(SimplecoinEventLog)
	if err := _Simplecoin.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimplecoinEventTransferIterator is returned from FilterEventTransfer and is used to iterate over the raw logs and unpacked data for EventTransfer events raised by the Simplecoin contract.
type SimplecoinEventTransferIterator struct {
	Event *SimplecoinEventTransfer // Event containing the contract specifics and raw log

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
func (it *SimplecoinEventTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplecoinEventTransfer)
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
		it.Event = new(SimplecoinEventTransfer)
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
func (it *SimplecoinEventTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplecoinEventTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplecoinEventTransfer represents a EventTransfer event raised by the Simplecoin contract.
type SimplecoinEventTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventTransfer is a free log retrieval operation binding the contract event 0x64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45.
//
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 amount)
func (_Simplecoin *SimplecoinFilterer) FilterEventTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimplecoinEventTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Simplecoin.contract.FilterLogs(opts, "EventTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SimplecoinEventTransferIterator{contract: _Simplecoin.contract, event: "EventTransfer", logs: logs, sub: sub}, nil
}

// WatchEventTransfer is a free log subscription operation binding the contract event 0x64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45.
//
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 amount)
func (_Simplecoin *SimplecoinFilterer) WatchEventTransfer(opts *bind.WatchOpts, sink chan<- *SimplecoinEventTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Simplecoin.contract.WatchLogs(opts, "EventTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplecoinEventTransfer)
				if err := _Simplecoin.contract.UnpackLog(event, "EventTransfer", log); err != nil {
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

// ParseEventTransfer is a log parse operation binding the contract event 0x64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45.
//
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 amount)
func (_Simplecoin *SimplecoinFilterer) ParseEventTransfer(log types.Log) (*SimplecoinEventTransfer, error) {
	event := new(SimplecoinEventTransfer)
	if err := _Simplecoin.contract.UnpackLog(event, "EventTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
