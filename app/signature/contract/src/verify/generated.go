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

// ErrorMetaData contains all meta data concerning the Error contract.
var ErrorMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a8b1a096616f436c4390165e68440e808cb6e0760d9f1f145610d31dfac1f4df64736f6c63430008110033",
}

// ErrorABI is the input ABI used to generate the binding from.
// Deprecated: Use ErrorMetaData.ABI instead.
var ErrorABI = ErrorMetaData.ABI

// ErrorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ErrorMetaData.Bin instead.
var ErrorBin = ErrorMetaData.Bin

// DeployError deploys a new Ethereum contract, binding an instance of Error to it.
func DeployError(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Error, error) {
	parsed, err := ErrorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ErrorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Error{ErrorCaller: ErrorCaller{contract: contract}, ErrorTransactor: ErrorTransactor{contract: contract}, ErrorFilterer: ErrorFilterer{contract: contract}}, nil
}

// Error is an auto generated Go binding around an Ethereum contract.
type Error struct {
	ErrorCaller     // Read-only binding to the contract
	ErrorTransactor // Write-only binding to the contract
	ErrorFilterer   // Log filterer for contract events
}

// ErrorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErrorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErrorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErrorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErrorSession struct {
	Contract     *Error            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErrorCallerSession struct {
	Contract *ErrorCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ErrorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErrorTransactorSession struct {
	Contract     *ErrorTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErrorRaw struct {
	Contract *Error // Generic contract binding to access the raw methods on
}

// ErrorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErrorCallerRaw struct {
	Contract *ErrorCaller // Generic read-only contract binding to access the raw methods on
}

// ErrorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErrorTransactorRaw struct {
	Contract *ErrorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewError creates a new instance of Error, bound to a specific deployed contract.
func NewError(address common.Address, backend bind.ContractBackend) (*Error, error) {
	contract, err := bindError(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Error{ErrorCaller: ErrorCaller{contract: contract}, ErrorTransactor: ErrorTransactor{contract: contract}, ErrorFilterer: ErrorFilterer{contract: contract}}, nil
}

// NewErrorCaller creates a new read-only instance of Error, bound to a specific deployed contract.
func NewErrorCaller(address common.Address, caller bind.ContractCaller) (*ErrorCaller, error) {
	contract, err := bindError(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorCaller{contract: contract}, nil
}

// NewErrorTransactor creates a new write-only instance of Error, bound to a specific deployed contract.
func NewErrorTransactor(address common.Address, transactor bind.ContractTransactor) (*ErrorTransactor, error) {
	contract, err := bindError(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorTransactor{contract: contract}, nil
}

// NewErrorFilterer creates a new log filterer instance of Error, bound to a specific deployed contract.
func NewErrorFilterer(address common.Address, filterer bind.ContractFilterer) (*ErrorFilterer, error) {
	contract, err := bindError(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErrorFilterer{contract: contract}, nil
}

// bindError binds a generic wrapper to an already deployed contract.
func bindError(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErrorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Error *ErrorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Error.Contract.ErrorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Error *ErrorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Error.Contract.ErrorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Error *ErrorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Error.Contract.ErrorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Error *ErrorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Error.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Error *ErrorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Error.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Error *ErrorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Error.Contract.contract.Transact(opts, method, params...)
}

// VerifyMetaData contains all meta data concerning the Verify contract.
var VerifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MatchSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"abbddf3b": "Address(string,address,uint256,bytes)",
		"406b5523": "MatchSender(string,address,uint256,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061064f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063406b55231461003b578063abbddf3b14610063575b600080fd5b61004e610049366004610409565b61008e565b60405190151581526020015b60405180910390f35b610076610071366004610409565b61012c565b6040516001600160a01b03909116815260200161005a565b6000808686866040516020016100a693929190610550565b6040516020818303038152906040528051906020012090506000806100cc8387876101a0565b80519193509150156100ff57806020015160405162461bcd60e51b81526004016100f6919061057e565b60405180910390fd5b336001600160a01b0383160361011b5760019350505050610123565b600093505050505b95945050505050565b60008086868660405160200161014493929190610550565b60405160208183030381529060405280519060200120905060008061016a8387876101a0565b805191935091501561019457806020015160405162461bcd60e51b81526004016100f6919061057e565b50979650505050505050565b60408051808201909152600080825260606020830152906041831461022d5760006102246040518060400160405280601881526020017f696e76616c6964207369676e6174757265206c656e677468000000000000000081525060408051808201825260008152606060209182015281518083019092526001825281019190915290565b91509150610386565b60006040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090506000818760405160200161027c929190610598565b60408051601f19818403018152919052805160209182012091506000906102a59082888a6105ba565b6102ae916105e4565b905060006102c060406020898b6105ba565b6102c9916105e4565b90506000888860408181106102e0576102e0610603565b604080516000815260208101808352899052939091013560f81c9083018190526060830186905260808301859052925060019160a00190506020604051602081039080840390855afa15801561033a573d6000803e3d6000fd5b5050506020604051035161037c604080518082018252600080825260606020928301528251808401845281815283518084019094529083529081019190915290565b9650965050505050505b935093915050565b634e487b7160e01b600052604160045260246000fd5b80356001600160a01b03811681146103bb57600080fd5b919050565b60008083601f8401126103d257600080fd5b50813567ffffffffffffffff8111156103ea57600080fd5b60208301915083602082850101111561040257600080fd5b9250929050565b60008060008060006080868803121561042157600080fd5b853567ffffffffffffffff8082111561043957600080fd5b818801915088601f83011261044d57600080fd5b81358181111561045f5761045f61038e565b604051601f8201601f19908116603f011681019083821181831017156104875761048761038e565b816040528281528b60208487010111156104a057600080fd5b826020860160208301376000602084830101528099505050506104c5602089016103a4565b95506040880135945060608801359150808211156104e257600080fd5b506104ef888289016103c0565b969995985093965092949392505050565b60005b8381101561051b578181015183820152602001610503565b50506000910152565b6000815180845261053c816020860160208601610500565b601f01601f19169290920160200192915050565b6060815260006105636060830186610524565b6001600160a01b039490941660208301525060400152919050565b6020815260006105916020830184610524565b9392505050565b600083516105aa818460208801610500565b9190910191825250602001919050565b600080858511156105ca57600080fd5b838611156105d757600080fd5b5050820193919092039150565b803560208310156105fd57600019602084900360031b1b165b92915050565b634e487b7160e01b600052603260045260246000fdfea2646970667358221220d4c3ac82b8592b6169662025d2443015c36c9e07401b0cd900fde048b369776564736f6c63430008110033",
}

// VerifyABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyMetaData.ABI instead.
var VerifyABI = VerifyMetaData.ABI

// Deprecated: Use VerifyMetaData.Sigs instead.
// VerifyFuncSigs maps the 4-byte function signature to its string representation.
var VerifyFuncSigs = VerifyMetaData.Sigs

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

