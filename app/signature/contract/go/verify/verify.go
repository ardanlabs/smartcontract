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
	Bin: "0x608060405234801561001057600080fd5b50610b18806100206000396000f3fe608060405234801561001057600080fd5b5060043610610053576000357c010000000000000000000000000000000000000000000000000000000090048063406b552314610058578063abbddf3b14610088575b600080fd5b610072600480360381019061006d91906106f8565b6100b8565b60405161007f91906107b7565b60405180910390f35b6100a2600480360381019061009d91906106f8565b610195565b6040516100af91906107e1565b60405180910390f35b6000808686866040516020016100d09392919061088a565b6040516020818303038152906040528051906020012090506000806100f6838787610231565b915091508060000151156101455780602001516040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013c91906108c8565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610184576001935050505061018c565b600093505050505b95945050505050565b6000808686866040516020016101ad9392919061088a565b6040516020818303038152906040528051906020012090506000806101d3838787610231565b915091508060000151156102225780602001516040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021991906108c8565b60405180910390fd5b81935050505095945050505050565b600061023b61048e565b6041848490501461028f5760006102866040518060400160405280601881526020017f696e76616c6964207369676e6174757265206c656e6774680000000000000000815250610433565b9150915061042b565b60006040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081876040516020016102de92919061095c565b6040516020818303038152906040528051906020012090506000868660009060209261030c9392919061098e565b9061031791906109e4565b90506000878760209060409261032f9392919061098e565b9061033a91906109e4565b905060008888604081811061035257610351610a43565b5b905001357f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027f010000000000000000000000000000000000000000000000000000000000000090049050600184828585604051600081526020016040526000604051602001526040516103ec9493929190610a9d565b60206040516020810390808403906000866161da5a03f115801561040f57600080fd5b5050506020604051035161042161045a565b9650965050505050505b935093915050565b61043b61048e565b6040518060400160405280600115158152602001838152509050919050565b61046261048e565b604051806040016040528060001515815260200160405180602001604052806000815250815250905090565b6040518060400160405280600015158152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610511826104c8565b810181811067ffffffffffffffff821117156105305761052f6104d9565b5b80604052505050565b60006105436104aa565b905061054f8282610508565b919050565b600067ffffffffffffffff82111561056f5761056e6104d9565b5b610578826104c8565b9050602081019050919050565b82818337600083830152505050565b60006105a76105a284610554565b610539565b9050828152602081018484840111156105c3576105c26104c3565b5b6105ce848285610585565b509392505050565b600082601f8301126105eb576105ea6104be565b5b81356105fb848260208601610594565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061062f82610604565b9050919050565b61063f81610624565b811461064a57600080fd5b50565b60008135905061065c81610636565b92915050565b6000819050919050565b61067581610662565b811461068057600080fd5b50565b6000813590506106928161066c565b92915050565b600080fd5b600080fd5b60008083601f8401126106b8576106b76104be565b5b8235905067ffffffffffffffff8111156106d5576106d4610698565b5b6020830191508360018202830111156106f1576106f061069d565b5b9250929050565b600080600080600060808688031215610714576107136104b4565b5b600086013567ffffffffffffffff811115610732576107316104b9565b5b61073e888289016105d6565b955050602061074f8882890161064d565b945050604061076088828901610683565b935050606086013567ffffffffffffffff811115610781576107806104b9565b5b61078d888289016106a2565b92509250509295509295909350565b60008115159050919050565b6107b18161079c565b82525050565b60006020820190506107cc60008301846107a8565b92915050565b6107db81610624565b82525050565b60006020820190506107f660008301846107d2565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561083657808201518184015260208101905061081b565b60008484015250505050565b600061084d826107fc565b6108578185610807565b9350610867818560208601610818565b610870816104c8565b840191505092915050565b61088481610662565b82525050565b600060608201905081810360008301526108a48186610842565b90506108b360208301856107d2565b6108c0604083018461087b565b949350505050565b600060208201905081810360008301526108e28184610842565b905092915050565b600081519050919050565b600081905092915050565b600061090b826108ea565b61091581856108f5565b9350610925818560208601610818565b80840191505092915050565b6000819050919050565b6000819050919050565b61095661095182610931565b61093b565b82525050565b60006109688285610900565b91506109748284610945565b6020820191508190509392505050565b600080fd5b600080fd5b600080858511156109a2576109a1610984565b5b838611156109b3576109b2610989565b5b6001850283019150848603905094509492505050565b600082905092915050565b60008160020a8302905092915050565b60006109f083836109c9565b826109fb8135610931565b92506020821015610a3b57610a367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff836020036008026109d4565b831692505b505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b610a7b81610931565b82525050565b600060ff82169050919050565b610a9781610a81565b82525050565b6000608082019050610ab26000830187610a72565b610abf6020830186610a8e565b610acc6040830185610a72565b610ad96060830184610a72565b9594505050505056fea2646970667358221220e41e2d4b9586ba8bcc280a8d992fc69dbe9571d76e442c42d5cdcef63b69e84964736f6c63430008140033",
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
