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
)

// VerifyMetaData contains all meta data concerning the Verify contract.
var VerifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MatchSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610732806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637079bd281461003b578063ad6265571461006b575b600080fd5b610055600480360381019061005091906103fd565b61009b565b6040516100629190610478565b60405180910390f35b610085600480360381019061008091906103fd565b610146565b60405161009291906104d4565b60405180910390f35b60008060006100ab8686866101b1565b915091508060000151156100fa5780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100f1919061057f565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036101385760019250505061013f565b6000925050505b9392505050565b60008060006101568686866101b1565b915091508060000151156101a55780602001516040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019c919061057f565b60405180910390fd5b81925050509392505050565b60006101bb61033c565b6041848490501461020f5760006102066040518060400160405280601881526020017f696e76616c6964207369676e6174757265206c656e67746800000000000000008152506102e1565b915091506102d9565b60008484600090602092610225939291906105ab565b9061023091906105fe565b905060008585602090604092610248939291906105ab565b9061025391906105fe565b905060008686604081811061026b5761026a61065d565b5b9050013560f81c60f81b60f81c90506001888285856040516000815260200160405260405161029d94939291906106b7565b6020604051602081039080840390855afa1580156102bf573d6000803e3d6000fd5b505050602060405103516102d1610308565b945094505050505b935093915050565b6102e961033c565b6040518060400160405280600115158152602001838152509050919050565b61031061033c565b604051806040016040528060001515815260200160405180602001604052806000815250815250905090565b6040518060400160405280600015158152602001606081525090565b600080fd5b600080fd5b6000819050919050565b61037581610362565b811461038057600080fd5b50565b6000813590506103928161036c565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126103bd576103bc610398565b5b8235905067ffffffffffffffff8111156103da576103d961039d565b5b6020830191508360018202830111156103f6576103f56103a2565b5b9250929050565b60008060006040848603121561041657610415610358565b5b600061042486828701610383565b935050602084013567ffffffffffffffff8111156104455761044461035d565b5b610451868287016103a7565b92509250509250925092565b60008115159050919050565b6104728161045d565b82525050565b600060208201905061048d6000830184610469565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104be82610493565b9050919050565b6104ce816104b3565b82525050565b60006020820190506104e960008301846104c5565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561052957808201518184015260208101905061050e565b60008484015250505050565b6000601f19601f8301169050919050565b6000610551826104ef565b61055b81856104fa565b935061056b81856020860161050b565b61057481610535565b840191505092915050565b600060208201905081810360008301526105998184610546565b905092915050565b600080fd5b600080fd5b600080858511156105bf576105be6105a1565b5b838611156105d0576105cf6105a6565b5b6001850283019150848603905094509492505050565b600082905092915050565b600082821b905092915050565b600061060a83836105e6565b826106158135610362565b92506020821015610655576106507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff836020036008026105f1565b831692505b505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b61069581610362565b82525050565b600060ff82169050919050565b6106b18161069b565b82525050565b60006080820190506106cc600083018761068c565b6106d960208301866106a8565b6106e6604083018561068c565b6106f3606083018461068c565b9594505050505056fea2646970667358221220b8a741fb199de3ec8ffea46631cd5712746dae4385e2a1c5726381300a59a52264736f6c63430008100033",
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
	parsed, err := abi.JSON(strings.NewReader(VerifyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// Address is a free data retrieval call binding the contract method 0xad626557.
//
// Solidity: function Address(bytes32 data, bytes sig) pure returns(address)
func (_Verify *VerifyCaller) Address(opts *bind.CallOpts, data [32]byte, sig []byte) (common.Address, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "Address", data, sig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Address is a free data retrieval call binding the contract method 0xad626557.
//
// Solidity: function Address(bytes32 data, bytes sig) pure returns(address)
func (_Verify *VerifySession) Address(data [32]byte, sig []byte) (common.Address, error) {
	return _Verify.Contract.Address(&_Verify.CallOpts, data, sig)
}

// Address is a free data retrieval call binding the contract method 0xad626557.
//
// Solidity: function Address(bytes32 data, bytes sig) pure returns(address)
func (_Verify *VerifyCallerSession) Address(data [32]byte, sig []byte) (common.Address, error) {
	return _Verify.Contract.Address(&_Verify.CallOpts, data, sig)
}

// MatchSender is a free data retrieval call binding the contract method 0x7079bd28.
//
// Solidity: function MatchSender(bytes32 data, bytes sig) view returns(bool)
func (_Verify *VerifyCaller) MatchSender(opts *bind.CallOpts, data [32]byte, sig []byte) (bool, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "MatchSender", data, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MatchSender is a free data retrieval call binding the contract method 0x7079bd28.
//
// Solidity: function MatchSender(bytes32 data, bytes sig) view returns(bool)
func (_Verify *VerifySession) MatchSender(data [32]byte, sig []byte) (bool, error) {
	return _Verify.Contract.MatchSender(&_Verify.CallOpts, data, sig)
}

// MatchSender is a free data retrieval call binding the contract method 0x7079bd28.
//
// Solidity: function MatchSender(bytes32 data, bytes sig) view returns(bool)
func (_Verify *VerifyCallerSession) MatchSender(data [32]byte, sig []byte) (bool, error) {
	return _Verify.Contract.MatchSender(&_Verify.CallOpts, data, sig)
}
