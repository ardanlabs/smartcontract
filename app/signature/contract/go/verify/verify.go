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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"MatchSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610610806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80639d4fa1b11461003b578063e51d691a1461006b575b600080fd5b610055600480360381019061005091906103a8565b61009b565b604051610062919061046c565b60405180910390f35b610085600480360381019061008091906103a8565b6100b3565b60405161009291906104a2565b60405180910390f35b60006100a98585858561010f565b9050949350505050565b6000806100c28686868661010f565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610101576001915050610107565b60009150505b949350505050565b6000808580519060200120905060006040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090506000818360405160200161016b92919061054f565b604051602081830303815290604052805190602001209050600181888888604051600081526020016040526040516101a69493929190610595565b6020604051602081039080840390855afa1580156101c8573d6000803e3d6000fd5b505050602060405103519350505050949350505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610246826101fd565b810181811067ffffffffffffffff821117156102655761026461020e565b5b80604052505050565b60006102786101df565b9050610284828261023d565b919050565b600067ffffffffffffffff8211156102a4576102a361020e565b5b6102ad826101fd565b9050602081019050919050565b82818337600083830152505050565b60006102dc6102d784610289565b61026e565b9050828152602081018484840111156102f8576102f76101f8565b5b6103038482856102ba565b509392505050565b600082601f8301126103205761031f6101f3565b5b81356103308482602086016102c9565b91505092915050565b600060ff82169050919050565b61034f81610339565b811461035a57600080fd5b50565b60008135905061036c81610346565b92915050565b6000819050919050565b61038581610372565b811461039057600080fd5b50565b6000813590506103a28161037c565b92915050565b600080600080608085870312156103c2576103c16101e9565b5b600085013567ffffffffffffffff8111156103e0576103df6101ee565b5b6103ec8782880161030b565b94505060206103fd8782880161035d565b935050604061040e87828801610393565b925050606061041f87828801610393565b91505092959194509250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104568261042b565b9050919050565b6104668161044b565b82525050565b6000602082019050610481600083018461045d565b92915050565b60008115159050919050565b61049c81610487565b82525050565b60006020820190506104b76000830184610493565b92915050565b600081519050919050565b600081905092915050565b60005b838110156104f15780820151818401526020810190506104d6565b60008484015250505050565b6000610508826104bd565b61051281856104c8565b93506105228185602086016104d3565b80840191505092915050565b6000819050919050565b61054961054482610372565b61052e565b82525050565b600061055b82856104fd565b91506105678284610538565b6020820191508190509392505050565b61058081610372565b82525050565b61058f81610339565b82525050565b60006080820190506105aa6000830187610577565b6105b76020830186610586565b6105c46040830185610577565b6105d16060830184610577565b9594505050505056fea2646970667358221220ead0e2f3eae5d716e87afe570acba589bbbe37c6be4ed8d9a8ea335baf89603e64736f6c63430008100033",
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

// Address is a free data retrieval call binding the contract method 0x9d4fa1b1.
//
// Solidity: function Address(bytes data, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Verify *VerifyCaller) Address(opts *bind.CallOpts, data []byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "Address", data, v, r, s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Address is a free data retrieval call binding the contract method 0x9d4fa1b1.
//
// Solidity: function Address(bytes data, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Verify *VerifySession) Address(data []byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Verify.Contract.Address(&_Verify.CallOpts, data, v, r, s)
}

// Address is a free data retrieval call binding the contract method 0x9d4fa1b1.
//
// Solidity: function Address(bytes data, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Verify *VerifyCallerSession) Address(data []byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Verify.Contract.Address(&_Verify.CallOpts, data, v, r, s)
}

// MatchSender is a free data retrieval call binding the contract method 0xe51d691a.
//
// Solidity: function MatchSender(bytes data, uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_Verify *VerifyCaller) MatchSender(opts *bind.CallOpts, data []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "MatchSender", data, v, r, s)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MatchSender is a free data retrieval call binding the contract method 0xe51d691a.
//
// Solidity: function MatchSender(bytes data, uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_Verify *VerifySession) MatchSender(data []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Verify.Contract.MatchSender(&_Verify.CallOpts, data, v, r, s)
}

// MatchSender is a free data retrieval call binding the contract method 0xe51d691a.
//
// Solidity: function MatchSender(bytes data, uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_Verify *VerifyCallerSession) MatchSender(data []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Verify.Contract.MatchSender(&_Verify.CallOpts, data, v, r, s)
}
