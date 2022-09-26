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
	Bin: "0x608060405234801561001057600080fd5b50610962806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063311b23ed1461003b578063f96ee26d1461006b575b600080fd5b61005560048036038101906100509190610588565b61009b565b6040516100629190610645565b60405180910390f35b61008560048036038101906100809190610588565b610106565b604051610092919061067b565b60405180910390f35b60008060006100ab8686866101b1565b915091508060000151156100fa5780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100f19190610715565b60405180910390fd5b81925050509392505050565b60008060006101168686866101b1565b915091508060000151156101655780602001516040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015c9190610715565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036101a3576001925050506101aa565b6000925050505b9392505050565b60006101bb6103b2565b6041848490501461020f5760006102066040518060400160405280601881526020017f696e76616c6964207369676e6174757265206c656e6774680000000000000000815250610357565b9150915061034f565b60008580519060200120905060006040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090506000818360405160200161026a9291906107a9565b60405160208183030381529060405280519060200120905060008787600090602092610298939291906107db565b906102a3919061082e565b9050600088886020906040926102bb939291906107db565b906102c6919061082e565b90506000898960408181106102de576102dd61088d565b5b9050013560f81c60f81b60f81c90506001848285856040516000815260200160405260405161031094939291906108e7565b6020604051602081039080840390855afa158015610332573d6000803e3d6000fd5b5050506020604051035161034461037e565b975097505050505050505b935093915050565b61035f6103b2565b6040518060400160405280600115158152602001838152509050919050565b6103866103b2565b604051806040016040528060001515815260200160405180602001604052806000815250815250905090565b6040518060400160405280600015158152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610435826103ec565b810181811067ffffffffffffffff82111715610454576104536103fd565b5b80604052505050565b60006104676103ce565b9050610473828261042c565b919050565b600067ffffffffffffffff821115610493576104926103fd565b5b61049c826103ec565b9050602081019050919050565b82818337600083830152505050565b60006104cb6104c684610478565b61045d565b9050828152602081018484840111156104e7576104e66103e7565b5b6104f28482856104a9565b509392505050565b600082601f83011261050f5761050e6103e2565b5b813561051f8482602086016104b8565b91505092915050565b600080fd5b600080fd5b60008083601f840112610548576105476103e2565b5b8235905067ffffffffffffffff81111561056557610564610528565b5b6020830191508360018202830111156105815761058061052d565b5b9250929050565b6000806000604084860312156105a1576105a06103d8565b5b600084013567ffffffffffffffff8111156105bf576105be6103dd565b5b6105cb868287016104fa565b935050602084013567ffffffffffffffff8111156105ec576105eb6103dd565b5b6105f886828701610532565b92509250509250925092565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061062f82610604565b9050919050565b61063f81610624565b82525050565b600060208201905061065a6000830184610636565b92915050565b60008115159050919050565b61067581610660565b82525050565b6000602082019050610690600083018461066c565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156106d05780820151818401526020810190506106b5565b60008484015250505050565b60006106e782610696565b6106f181856106a1565b93506107018185602086016106b2565b61070a816103ec565b840191505092915050565b6000602082019050818103600083015261072f81846106dc565b905092915050565b600081519050919050565b600081905092915050565b600061075882610737565b6107628185610742565b93506107728185602086016106b2565b80840191505092915050565b6000819050919050565b6000819050919050565b6107a361079e8261077e565b610788565b82525050565b60006107b5828561074d565b91506107c18284610792565b6020820191508190509392505050565b600080fd5b600080fd5b600080858511156107ef576107ee6107d1565b5b83861115610800576107ff6107d6565b5b6001850283019150848603905094509492505050565b600082905092915050565b600082821b905092915050565b600061083a8383610816565b82610845813561077e565b92506020821015610885576108807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83602003600802610821565b831692505b505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6108c58161077e565b82525050565b600060ff82169050919050565b6108e1816108cb565b82525050565b60006080820190506108fc60008301876108bc565b61090960208301866108d8565b61091660408301856108bc565b61092360608301846108bc565b9594505050505056fea26469706673582212206a6dc36cb788820db33479d60c46000562c6f20fb23fca40bee4cf74e2810a1864736f6c63430008100033",
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
