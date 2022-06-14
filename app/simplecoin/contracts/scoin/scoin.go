// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scoin

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

// ScoinMetaData contains all meta data concerning the Scoin contract.
var ScoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizedAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowanceAmount\",\"type\":\"uint256\"}],\"name\":\"authorize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"coinBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001309380380620013098339818101604052810190620000379190620000c1565b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050620000f3565b600080fd5b6000819050919050565b6200009b8162000086565b8114620000a757600080fd5b50565b600081519050620000bb8162000090565b92915050565b600060208284031215620000da57620000d962000081565b5b6000620000ea84828501620000aa565b91505092915050565b61120680620001036000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806323b872dd1461005c578063a9059cbb14610078578063c1dbd9b214610094578063dd62ed3e146100c4578063fabde80c146100f4575b600080fd5b61007660048036038101906100719190610be2565b610124565b005b610092600480360381019061008d9190610c35565b610390565b005b6100ae60048036038101906100a99190610c35565b610568565b6040516100bb9190610c90565b60405180910390f35b6100de60048036038101906100d99190610cab565b6105f5565b6040516100eb9190610cfa565b60405180910390f35b61010e60048036038101906101099190610d15565b61061a565b60405161011b9190610cfa565b60405180910390f35b6000610131848484610632565b905080600001511561017e5780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101759190610ddb565b60405180910390fd5b7fcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab6040516101ab90610e49565b60405180910390a1816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102019190610e98565b92505081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102569190610ecc565b9250508190555081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102e99190610e98565b925050819055507fcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab60405161031d90610f6e565b60405180910390a18273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516103829190610cfa565b60405180910390a350505050565b600061039d338484610632565b90508060000151156103ea5780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e19190610ddb565b60405180910390fd5b7fcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab60405161041790610e49565b60405180910390a1816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461046d9190610e98565b92505081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104c29190610ecc565b925050819055507fcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab6040516104f690610f6e565b60405180910390a18273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161055b9190610cfa565b60405180910390a3505050565b600081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001905092915050565b6001602052816000526040600020602052806000526040600020600091509150505481565b60006020528060005260406000206000915090505481565b61063a610b2d565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036106b3576106ac6040518060400160405280601f81526020017f63616e27742073656e64206d6f6e657920746f206164647265737320307830008152506108b2565b90506108ab565b816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156107c4576107bd6040518060400160405280601781526020017f696e737566666963656e742066756e6473202062616c3a0000000000000000008152506107796000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546108d9565b6040518060400160405280600a81526020017f2020616d6f756e743a20000000000000000000000000000000000000000000008152506107b8866108d9565b610a61565b90506108ab565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461084d9190610ecc565b116108a0576108996040518060400160405280601081526020017f696e76616c696420616d6f756e743a2000000000000000000000000000000000815250610894846108d9565b610ab0565b90506108ab565b6108a8610af9565b90505b9392505050565b6108ba610b2d565b6040518060400160405280600115158152602001838152509050919050565b606060008203610920576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610a5c565b600082905060005b6000821461095257808061093b90610f8e565b915050600a8261094b9190611005565b9150610928565b60008167ffffffffffffffff81111561096e5761096d611036565b5b6040519080825280601f01601f1916602001820160405280156109a05781602001600182028036833780820191505090505b50905060008290505b60008614610a54576001816109be9190610e98565b90506000600a80886109d09190611005565b6109da9190611065565b876109e59190610e98565b60306109f191906110cc565b905060008160f81b905080848481518110610a0f57610a0e611103565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610a4b9190611005565b975050506109a9565b819450505050505b919050565b610a69610b2d565b604051806040016040528060011515815260200186868686604051602001610a94949392919061116e565b6040516020818303038152906040528152509050949350505050565b610ab8610b2d565b60405180604001604052806001151581526020018484604051602001610adf9291906111ac565b604051602081830303815290604052815250905092915050565b610b01610b2d565b604051806040016040528060001515815260200160405180602001604052806000815250815250905090565b6040518060400160405280600015158152602001606081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b7982610b4e565b9050919050565b610b8981610b6e565b8114610b9457600080fd5b50565b600081359050610ba681610b80565b92915050565b6000819050919050565b610bbf81610bac565b8114610bca57600080fd5b50565b600081359050610bdc81610bb6565b92915050565b600080600060608486031215610bfb57610bfa610b49565b5b6000610c0986828701610b97565b9350506020610c1a86828701610b97565b9250506040610c2b86828701610bcd565b9150509250925092565b60008060408385031215610c4c57610c4b610b49565b5b6000610c5a85828601610b97565b9250506020610c6b85828601610bcd565b9150509250929050565b60008115159050919050565b610c8a81610c75565b82525050565b6000602082019050610ca56000830184610c81565b92915050565b60008060408385031215610cc257610cc1610b49565b5b6000610cd085828601610b97565b9250506020610ce185828601610b97565b9150509250929050565b610cf481610bac565b82525050565b6000602082019050610d0f6000830184610ceb565b92915050565b600060208284031215610d2b57610d2a610b49565b5b6000610d3984828501610b97565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610d7c578082015181840152602081019050610d61565b83811115610d8b576000848401525b50505050565b6000601f19601f8301169050919050565b6000610dad82610d42565b610db78185610d4d565b9350610dc7818560208601610d5e565b610dd081610d91565b840191505092915050565b60006020820190508181036000830152610df58184610da2565b905092915050565b7f7374617274696e67207472616e73666572000000000000000000000000000000600082015250565b6000610e33601183610d4d565b9150610e3e82610dfd565b602082019050919050565b60006020820190508181036000830152610e6281610e26565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ea382610bac565b9150610eae83610bac565b925082821015610ec157610ec0610e69565b5b828203905092915050565b6000610ed782610bac565b9150610ee283610bac565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610f1757610f16610e69565b5b828201905092915050565b7f656e64696e67207472616e736665720000000000000000000000000000000000600082015250565b6000610f58600f83610d4d565b9150610f6382610f22565b602082019050919050565b60006020820190508181036000830152610f8781610f4b565b9050919050565b6000610f9982610bac565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fcb57610fca610e69565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061101082610bac565b915061101b83610bac565b92508261102b5761102a610fd6565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600061107082610bac565b915061107b83610bac565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156110b4576110b3610e69565b5b828202905092915050565b600060ff82169050919050565b60006110d7826110bf565b91506110e2836110bf565b92508260ff038211156110f8576110f7610e69565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081905092915050565b600061114882610d42565b6111528185611132565b9350611162818560208601610d5e565b80840191505092915050565b600061117a828761113d565b9150611186828661113d565b9150611192828561113d565b915061119e828461113d565b915081905095945050505050565b60006111b8828561113d565b91506111c4828461113d565b9150819050939250505056fea26469706673582212209ad6d68da413dd93d67faaa3dd4a72dcbdbec703d5cabcee33ff79c4575ce3d164736f6c634300080e0033",
}

// ScoinABI is the input ABI used to generate the binding from.
// Deprecated: Use ScoinMetaData.ABI instead.
var ScoinABI = ScoinMetaData.ABI

// ScoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ScoinMetaData.Bin instead.
var ScoinBin = ScoinMetaData.Bin

// DeployScoin deploys a new Ethereum contract, binding an instance of Scoin to it.
func DeployScoin(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *Scoin, error) {
	parsed, err := ScoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScoinBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Scoin{ScoinCaller: ScoinCaller{contract: contract}, ScoinTransactor: ScoinTransactor{contract: contract}, ScoinFilterer: ScoinFilterer{contract: contract}}, nil
}

// Scoin is an auto generated Go binding around an Ethereum contract.
type Scoin struct {
	ScoinCaller     // Read-only binding to the contract
	ScoinTransactor // Write-only binding to the contract
	ScoinFilterer   // Log filterer for contract events
}

// ScoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScoinSession struct {
	Contract     *Scoin            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScoinCallerSession struct {
	Contract *ScoinCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ScoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScoinTransactorSession struct {
	Contract     *ScoinTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScoinRaw struct {
	Contract *Scoin // Generic contract binding to access the raw methods on
}

// ScoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScoinCallerRaw struct {
	Contract *ScoinCaller // Generic read-only contract binding to access the raw methods on
}

// ScoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScoinTransactorRaw struct {
	Contract *ScoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScoin creates a new instance of Scoin, bound to a specific deployed contract.
func NewScoin(address common.Address, backend bind.ContractBackend) (*Scoin, error) {
	contract, err := bindScoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Scoin{ScoinCaller: ScoinCaller{contract: contract}, ScoinTransactor: ScoinTransactor{contract: contract}, ScoinFilterer: ScoinFilterer{contract: contract}}, nil
}

// NewScoinCaller creates a new read-only instance of Scoin, bound to a specific deployed contract.
func NewScoinCaller(address common.Address, caller bind.ContractCaller) (*ScoinCaller, error) {
	contract, err := bindScoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScoinCaller{contract: contract}, nil
}

// NewScoinTransactor creates a new write-only instance of Scoin, bound to a specific deployed contract.
func NewScoinTransactor(address common.Address, transactor bind.ContractTransactor) (*ScoinTransactor, error) {
	contract, err := bindScoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScoinTransactor{contract: contract}, nil
}

// NewScoinFilterer creates a new log filterer instance of Scoin, bound to a specific deployed contract.
func NewScoinFilterer(address common.Address, filterer bind.ContractFilterer) (*ScoinFilterer, error) {
	contract, err := bindScoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScoinFilterer{contract: contract}, nil
}

// bindScoin binds a generic wrapper to an already deployed contract.
func bindScoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scoin *ScoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scoin.Contract.ScoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scoin *ScoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scoin.Contract.ScoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scoin *ScoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scoin.Contract.ScoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scoin *ScoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scoin *ScoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scoin *ScoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Scoin *ScoinCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scoin.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Scoin *ScoinSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Scoin.Contract.Allowance(&_Scoin.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Scoin *ScoinCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Scoin.Contract.Allowance(&_Scoin.CallOpts, arg0, arg1)
}

// CoinBalance is a free data retrieval call binding the contract method 0xfabde80c.
//
// Solidity: function coinBalance(address ) view returns(uint256)
func (_Scoin *ScoinCaller) CoinBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scoin.contract.Call(opts, &out, "coinBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinBalance is a free data retrieval call binding the contract method 0xfabde80c.
//
// Solidity: function coinBalance(address ) view returns(uint256)
func (_Scoin *ScoinSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _Scoin.Contract.CoinBalance(&_Scoin.CallOpts, arg0)
}

// CoinBalance is a free data retrieval call binding the contract method 0xfabde80c.
//
// Solidity: function coinBalance(address ) view returns(uint256)
func (_Scoin *ScoinCallerSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _Scoin.Contract.CoinBalance(&_Scoin.CallOpts, arg0)
}

// Authorize is a paid mutator transaction binding the contract method 0xc1dbd9b2.
//
// Solidity: function authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_Scoin *ScoinTransactor) Authorize(opts *bind.TransactOpts, authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _Scoin.contract.Transact(opts, "authorize", authorizedAccount, allowanceAmount)
}

// Authorize is a paid mutator transaction binding the contract method 0xc1dbd9b2.
//
// Solidity: function authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_Scoin *ScoinSession) Authorize(authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Authorize(&_Scoin.TransactOpts, authorizedAccount, allowanceAmount)
}

// Authorize is a paid mutator transaction binding the contract method 0xc1dbd9b2.
//
// Solidity: function authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_Scoin *ScoinTransactorSession) Authorize(authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Authorize(&_Scoin.TransactOpts, authorizedAccount, allowanceAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Scoin *ScoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Scoin *ScoinSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Transfer(&_Scoin.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Scoin *ScoinTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Transfer(&_Scoin.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns()
func (_Scoin *ScoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns()
func (_Scoin *ScoinSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.TransferFrom(&_Scoin.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns()
func (_Scoin *ScoinTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.TransferFrom(&_Scoin.TransactOpts, from, to, amount)
}

// ScoinLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the Scoin contract.
type ScoinLogIterator struct {
	Event *ScoinLog // Event containing the contract specifics and raw log

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
func (it *ScoinLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoinLog)
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
		it.Event = new(ScoinLog)
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
func (it *ScoinLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoinLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoinLog represents a Log event raised by the Scoin contract.
type ScoinLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string value)
func (_Scoin *ScoinFilterer) FilterLog(opts *bind.FilterOpts) (*ScoinLogIterator, error) {

	logs, sub, err := _Scoin.contract.FilterLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return &ScoinLogIterator{contract: _Scoin.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string value)
func (_Scoin *ScoinFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *ScoinLog) (event.Subscription, error) {

	logs, sub, err := _Scoin.contract.WatchLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoinLog)
				if err := _Scoin.contract.UnpackLog(event, "Log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string value)
func (_Scoin *ScoinFilterer) ParseLog(log types.Log) (*ScoinLog, error) {
	event := new(ScoinLog)
	if err := _Scoin.contract.UnpackLog(event, "Log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Scoin contract.
type ScoinTransferIterator struct {
	Event *ScoinTransfer // Event containing the contract specifics and raw log

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
func (it *ScoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoinTransfer)
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
		it.Event = new(ScoinTransfer)
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
func (it *ScoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoinTransfer represents a Transfer event raised by the Scoin contract.
type ScoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scoin *ScoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Scoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScoinTransferIterator{contract: _Scoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scoin *ScoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ScoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Scoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoinTransfer)
				if err := _Scoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scoin *ScoinFilterer) ParseTransfer(log types.Log) (*ScoinTransfer, error) {
	event := new(ScoinTransfer)
	if err := _Scoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
