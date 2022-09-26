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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MatchSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610765806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063311b23ed1461003b578063f96ee26d1461006b575b600080fd5b61005560048036038101906100509190610402565b61009b565b60405161006291906104bf565b60405180910390f35b61008560048036038101906100809190610402565b6100b1565b60405161009291906104f5565b60405180910390f35b60006100a884848461010b565b90509392505050565b6000806100bf85858561010b565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100fe576001915050610104565b60009150505b9392505050565b6000808480519060200120905060006040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081836040516020016101679291906105ac565b60405160208183030381529060405280519060200120905060008686600090602092610195939291906105de565b906101a09190610631565b9050600087876020906040926101b8939291906105de565b906101c39190610631565b90506000888860408181106101db576101da610690565b5b9050013560f81c60f81b60f81c90506001848285856040516000815260200160405260405161020d94939291906106ea565b6020604051602081039080840390855afa15801561022f573d6000803e3d6000fd5b5050506020604051035196505050505050509392505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102af82610266565b810181811067ffffffffffffffff821117156102ce576102cd610277565b5b80604052505050565b60006102e1610248565b90506102ed82826102a6565b919050565b600067ffffffffffffffff82111561030d5761030c610277565b5b61031682610266565b9050602081019050919050565b82818337600083830152505050565b6000610345610340846102f2565b6102d7565b90508281526020810184848401111561036157610360610261565b5b61036c848285610323565b509392505050565b600082601f8301126103895761038861025c565b5b8135610399848260208601610332565b91505092915050565b600080fd5b600080fd5b60008083601f8401126103c2576103c161025c565b5b8235905067ffffffffffffffff8111156103df576103de6103a2565b5b6020830191508360018202830111156103fb576103fa6103a7565b5b9250929050565b60008060006040848603121561041b5761041a610252565b5b600084013567ffffffffffffffff81111561043957610438610257565b5b61044586828701610374565b935050602084013567ffffffffffffffff81111561046657610465610257565b5b610472868287016103ac565b92509250509250925092565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104a98261047e565b9050919050565b6104b98161049e565b82525050565b60006020820190506104d460008301846104b0565b92915050565b60008115159050919050565b6104ef816104da565b82525050565b600060208201905061050a60008301846104e6565b92915050565b600081519050919050565b600081905092915050565b60005b83811015610544578082015181840152602081019050610529565b60008484015250505050565b600061055b82610510565b610565818561051b565b9350610575818560208601610526565b80840191505092915050565b6000819050919050565b6000819050919050565b6105a66105a182610581565b61058b565b82525050565b60006105b88285610550565b91506105c48284610595565b6020820191508190509392505050565b600080fd5b600080fd5b600080858511156105f2576105f16105d4565b5b83861115610603576106026105d9565b5b6001850283019150848603905094509492505050565b600082905092915050565b600082821b905092915050565b600061063d8383610619565b826106488135610581565b92506020821015610688576106837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83602003600802610624565b831692505b505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6106c881610581565b82525050565b600060ff82169050919050565b6106e4816106ce565b82525050565b60006080820190506106ff60008301876106bf565b61070c60208301866106db565b61071960408301856106bf565b61072660608301846106bf565b9594505050505056fea2646970667358221220e2ed1962c60712cd7f1d0890b47d1768e3dc1ffa67610e1508fed96024a59d7c64736f6c63430008100033",
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

// Address is a free data retrieval call binding the contract method 0x311b23ed.
//
// Solidity: function Address(bytes data, bytes sig) pure returns(address)
func (_Verify *VerifyCaller) Address(opts *bind.CallOpts, data []byte, sig []byte) (common.Address, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "Address", data, sig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Address is a free data retrieval call binding the contract method 0x311b23ed.
//
// Solidity: function Address(bytes data, bytes sig) pure returns(address)
func (_Verify *VerifySession) Address(data []byte, sig []byte) (common.Address, error) {
	return _Verify.Contract.Address(&_Verify.CallOpts, data, sig)
}

// Address is a free data retrieval call binding the contract method 0x311b23ed.
//
// Solidity: function Address(bytes data, bytes sig) pure returns(address)
func (_Verify *VerifyCallerSession) Address(data []byte, sig []byte) (common.Address, error) {
	return _Verify.Contract.Address(&_Verify.CallOpts, data, sig)
}

// MatchSender is a free data retrieval call binding the contract method 0xf96ee26d.
//
// Solidity: function MatchSender(bytes data, bytes sig) view returns(bool)
func (_Verify *VerifyCaller) MatchSender(opts *bind.CallOpts, data []byte, sig []byte) (bool, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "MatchSender", data, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MatchSender is a free data retrieval call binding the contract method 0xf96ee26d.
//
// Solidity: function MatchSender(bytes data, bytes sig) view returns(bool)
func (_Verify *VerifySession) MatchSender(data []byte, sig []byte) (bool, error) {
	return _Verify.Contract.MatchSender(&_Verify.CallOpts, data, sig)
}

// MatchSender is a free data retrieval call binding the contract method 0xf96ee26d.
//
// Solidity: function MatchSender(bytes data, bytes sig) view returns(bool)
func (_Verify *VerifyCallerSession) MatchSender(data []byte, sig []byte) (bool, error) {
	return _Verify.Contract.MatchSender(&_Verify.CallOpts, data, sig)
}
