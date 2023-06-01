// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verify

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

// VerifyMetaData contains all meta data concerning the Verify contract.
var VerifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MatchSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610a8f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063406b55231461003b578063abbddf3b1461006b575b600080fd5b61005560048036038101906100509190610672565b61009b565b6040516100629190610731565b60405180910390f35b61008560048036038101906100809190610672565b610178565b604051610092919061075b565b60405180910390f35b6000808686866040516020016100b393929190610804565b6040516020818303038152906040528051906020012090506000806100d9838787610214565b915091508060000151156101285780602001516040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161011f9190610842565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610167576001935050505061016f565b600093505050505b95945050505050565b60008086868660405160200161019093929190610804565b6040516020818303038152906040528051906020012090506000806101b6838787610214565b915091508060000151156102055780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fc9190610842565b60405180910390fd5b81935050505095945050505050565b600061021e610408565b604184849050146102725760006102696040518060400160405280601881526020017f696e76616c6964207369676e6174757265206c656e67746800000000000000008152506103ad565b915091506103a5565b60006040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081876040516020016102c19291906108d6565b604051602081830303815290604052805190602001209050600086866000906020926102ef93929190610908565b906102fa919061095b565b90506000878760209060409261031293929190610908565b9061031d919061095b565b9050600088886040818110610335576103346109ba565b5b9050013560f81c60f81b60f81c9050600184828585604051600081526020016040526040516103679493929190610a14565b6020604051602081039080840390855afa158015610389573d6000803e3d6000fd5b5050506020604051035161039b6103d4565b9650965050505050505b935093915050565b6103b5610408565b6040518060400160405280600115158152602001838152509050919050565b6103dc610408565b604051806040016040528060001515815260200160405180602001604052806000815250815250905090565b6040518060400160405280600015158152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61048b82610442565b810181811067ffffffffffffffff821117156104aa576104a9610453565b5b80604052505050565b60006104bd610424565b90506104c98282610482565b919050565b600067ffffffffffffffff8211156104e9576104e8610453565b5b6104f282610442565b9050602081019050919050565b82818337600083830152505050565b600061052161051c846104ce565b6104b3565b90508281526020810184848401111561053d5761053c61043d565b5b6105488482856104ff565b509392505050565b600082601f83011261056557610564610438565b5b813561057584826020860161050e565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105a98261057e565b9050919050565b6105b98161059e565b81146105c457600080fd5b50565b6000813590506105d6816105b0565b92915050565b6000819050919050565b6105ef816105dc565b81146105fa57600080fd5b50565b60008135905061060c816105e6565b92915050565b600080fd5b600080fd5b60008083601f84011261063257610631610438565b5b8235905067ffffffffffffffff81111561064f5761064e610612565b5b60208301915083600182028301111561066b5761066a610617565b5b9250929050565b60008060008060006080868803121561068e5761068d61042e565b5b600086013567ffffffffffffffff8111156106ac576106ab610433565b5b6106b888828901610550565b95505060206106c9888289016105c7565b94505060406106da888289016105fd565b935050606086013567ffffffffffffffff8111156106fb576106fa610433565b5b6107078882890161061c565b92509250509295509295909350565b60008115159050919050565b61072b81610716565b82525050565b60006020820190506107466000830184610722565b92915050565b6107558161059e565b82525050565b6000602082019050610770600083018461074c565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156107b0578082015181840152602081019050610795565b60008484015250505050565b60006107c782610776565b6107d18185610781565b93506107e1818560208601610792565b6107ea81610442565b840191505092915050565b6107fe816105dc565b82525050565b6000606082019050818103600083015261081e81866107bc565b905061082d602083018561074c565b61083a60408301846107f5565b949350505050565b6000602082019050818103600083015261085c81846107bc565b905092915050565b600081519050919050565b600081905092915050565b600061088582610864565b61088f818561086f565b935061089f818560208601610792565b80840191505092915050565b6000819050919050565b6000819050919050565b6108d06108cb826108ab565b6108b5565b82525050565b60006108e2828561087a565b91506108ee82846108bf565b6020820191508190509392505050565b600080fd5b600080fd5b6000808585111561091c5761091b6108fe565b5b8386111561092d5761092c610903565b5b6001850283019150848603905094509492505050565b600082905092915050565b600082821b905092915050565b60006109678383610943565b8261097281356108ab565b925060208210156109b2576109ad7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8360200360080261094e565b831692505b505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6109f2816108ab565b82525050565b600060ff82169050919050565b610a0e816109f8565b82525050565b6000608082019050610a2960008301876109e9565b610a366020830186610a05565b610a4360408301856109e9565b610a5060608301846109e9565b9594505050505056fea2646970667358221220bce35f54c8db7e00248de91973547d8b94da4a123ac6533c87389c43c18c605d64736f6c63430008140033",
}

// VerifyABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyMetaData.ABI instead.
var VerifyABI = VerifyMetaData.ABI

// VerifyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifyMetaData.Bin instead.
var VerifyBin = VerifyMetaData.Bin

// DeployVerify deploys a new Ethereum contract, binding an instance of Verify to it.
func DeployVerify(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verify, error) {
	parsed, err := VerifyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verify{VerifyCaller: VerifyCaller{contract: contract}, VerifyTransactor: VerifyTransactor{contract: contract}, VerifyFilterer: VerifyFilterer{contract: contract}}, nil
}

// Verify is an auto generated Go binding around an Ethereum contract.
type Verify struct {
	VerifyCaller     // Read-only binding to the contract
	VerifyTransactor // Write-only binding to the contract
	VerifyFilterer   // Log filterer for contract events
}

// VerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifySession struct {
	Contract     *Verify           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyCallerSession struct {
	Contract *VerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyTransactorSession struct {
	Contract     *VerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyRaw struct {
	Contract *Verify // Generic contract binding to access the raw methods on
}

// VerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyCallerRaw struct {
	Contract *VerifyCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyTransactorRaw struct {
	Contract *VerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerify creates a new instance of Verify, bound to a specific deployed contract.
func NewVerify(address common.Address, backend bind.ContractBackend) (*Verify, error) {
	contract, err := bindVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verify{VerifyCaller: VerifyCaller{contract: contract}, VerifyTransactor: VerifyTransactor{contract: contract}, VerifyFilterer: VerifyFilterer{contract: contract}}, nil
}

// NewVerifyCaller creates a new read-only instance of Verify, bound to a specific deployed contract.
func NewVerifyCaller(address common.Address, caller bind.ContractCaller) (*VerifyCaller, error) {
	contract, err := bindVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyCaller{contract: contract}, nil
}

// NewVerifyTransactor creates a new write-only instance of Verify, bound to a specific deployed contract.
func NewVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyTransactor, error) {
	contract, err := bindVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyTransactor{contract: contract}, nil
}

// NewVerifyFilterer creates a new log filterer instance of Verify, bound to a specific deployed contract.
func NewVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyFilterer, error) {
	contract, err := bindVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyFilterer{contract: contract}, nil
}

// bindVerify binds a generic wrapper to an already deployed contract.
func bindVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verify *VerifyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verify.Contract.VerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verify *VerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verify.Contract.VerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verify *VerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verify.Contract.VerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verify *VerifyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verify *VerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verify *VerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verify.Contract.contract.Transact(opts, method, params...)
}

// Address is a free data retrieval call binding the contract method 0xabbddf3b.
//
// Solidity: function Address(string id, address participant, uint256 nonce, bytes sig) pure returns(address)
func (_Verify *VerifyCaller) Address(opts *bind.CallOpts, id string, participant common.Address, nonce *big.Int, sig []byte) (common.Address, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "Address", id, participant, nonce, sig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Address is a free data retrieval call binding the contract method 0xabbddf3b.
//
// Solidity: function Address(string id, address participant, uint256 nonce, bytes sig) pure returns(address)
func (_Verify *VerifySession) Address(id string, participant common.Address, nonce *big.Int, sig []byte) (common.Address, error) {
	return _Verify.Contract.Address(&_Verify.CallOpts, id, participant, nonce, sig)
}

// Address is a free data retrieval call binding the contract method 0xabbddf3b.
//
// Solidity: function Address(string id, address participant, uint256 nonce, bytes sig) pure returns(address)
func (_Verify *VerifyCallerSession) Address(id string, participant common.Address, nonce *big.Int, sig []byte) (common.Address, error) {
	return _Verify.Contract.Address(&_Verify.CallOpts, id, participant, nonce, sig)
}

// MatchSender is a free data retrieval call binding the contract method 0x406b5523.
//
// Solidity: function MatchSender(string id, address participant, uint256 nonce, bytes sig) view returns(bool)
func (_Verify *VerifyCaller) MatchSender(opts *bind.CallOpts, id string, participant common.Address, nonce *big.Int, sig []byte) (bool, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "MatchSender", id, participant, nonce, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MatchSender is a free data retrieval call binding the contract method 0x406b5523.
//
// Solidity: function MatchSender(string id, address participant, uint256 nonce, bytes sig) view returns(bool)
func (_Verify *VerifySession) MatchSender(id string, participant common.Address, nonce *big.Int, sig []byte) (bool, error) {
	return _Verify.Contract.MatchSender(&_Verify.CallOpts, id, participant, nonce, sig)
}

// MatchSender is a free data retrieval call binding the contract method 0x406b5523.
//
// Solidity: function MatchSender(string id, address participant, uint256 nonce, bytes sig) view returns(bool)
func (_Verify *VerifyCallerSession) MatchSender(id string, participant common.Address, nonce *big.Int, sig []byte) (bool, error) {
	return _Verify.Contract.MatchSender(&_Verify.CallOpts, id, participant, nonce, sig)
}
