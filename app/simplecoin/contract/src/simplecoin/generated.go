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
)

// ErrorMetaData contains all meta data concerning the Error contract.
var ErrorMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220178b0bd27bb9e3967987260b1ae822bff57657fc388b257f8af76f67ef8d08bc64736f6c63430008110033",
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

// SimpleCoinMetaData contains all meta data concerning the SimpleCoin contract.
var SimpleCoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"EventFrozenAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"EventTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizedAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowanceAmount\",\"type\":\"uint256\"}],\"name\":\"Authorize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CoinBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"freeze\",\"type\":\"bool\"}],\"name\":\"FreezeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FrozenAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6dcdd00f": "Allowance(address,address)",
		"6386a433": "Authorize(address,uint256)",
		"8bdf1534": "CoinBalance(address)",
		"d16a7a4b": "FreezeAccount(address,bool)",
		"ab31471e": "FrozenAccount(address)",
		"0f6798a5": "Mint(address,uint256)",
		"b4a99a4e": "Owner()",
		"69ca02dd": "Transfer(address,uint256)",
		"c0d84ce5": "TransferFrom(address,address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161115b38038061115b83398101604081905261002f916100e0565b600080546001600160a01b0319163390811790915561004e9082610054565b50610120565b6000546001600160a01b0316331461006b57600080fd5b6001600160a01b038216600090815260016020526040812080548392906100939084906100f9565b90915550506000546040518281526001600160a01b038481169216907f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a459060200160405180910390a35050565b6000602082840312156100f257600080fd5b5051919050565b8082018082111561011a57634e487b7160e01b600052601160045260246000fd5b92915050565b61102c8061012f6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638bdf1534116100665780638bdf153414610121578063ab31471e14610141578063b4a99a4e14610164578063c0d84ce51461018f578063d16a7a4b146101a257600080fd5b80630f6798a5146100985780636386a433146100ad57806369ca02dd146100d55780636dcdd00f146100e8575b600080fd5b6100ab6100a6366004610b06565b6101b5565b005b6100c06100bb366004610b06565b610241565b60405190151581526020015b60405180910390f35b6100ab6100e3366004610b06565b610270565b6101136100f6366004610b30565b600260209081526000928352604080842090915290825290205481565b6040519081526020016100cc565b61011361012f366004610b63565b60016020526000908152604090205481565b6100c061014f366004610b63565b60036020526000908152604090205460ff1681565b600054610177906001600160a01b031681565b6040516001600160a01b0390911681526020016100cc565b6100ab61019d366004610b7e565b6103b5565b6100ab6101b0366004610bba565b610540565b6000546001600160a01b031633146101cc57600080fd5b6001600160a01b038216600090815260016020526040812080548392906101f4908490610c0c565b90915550506000546040518281526001600160a01b038481169216907f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a459060200160405180910390a35050565b3360009081526002602090815260408083206001600160a01b0386168452909152902081905560015b92915050565b600061027d3384846105ba565b8051909150156102ae57806020015160405162461bcd60e51b81526004016102a59190610c43565b60405180910390fd5b33600090815260016020526040812080548492906102cd908490610c76565b90915550506001600160a01b038316600090815260016020526040812080548492906102fa908490610c0c565b909155507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a905061032a83610791565b610333856108bd565b61033c336108bd565b60405160200161034e93929190610c89565b60408051601f198184030181529082905261036891610c43565b60405180910390a16040518281526001600160a01b0384169033907f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a459060200160405180910390a3505050565b60006103c28484846105ba565b8051909150156103ea57806020015160405162461bcd60e51b81526004016102a59190610c43565b6001600160a01b03841660009081526001602052604081208054849290610412908490610c76565b90915550506001600160a01b0383166000908152600160205260408120805484929061043f908490610c0c565b90915550506001600160a01b038416600090815260026020908152604080832033845290915281208054849290610477908490610c76565b909155507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a90506104a783610791565b6104b0856108bd565b6104b9876108bd565b6040516020016104cb93929190610d07565b60408051601f19818403018152908290526104e591610c43565b60405180910390a1826001600160a01b0316846001600160a01b03167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a458460405161053291815260200190565b60405180910390a350505050565b6000546001600160a01b0316331461055757600080fd5b6001600160a01b038216600081815260036020908152604091829020805460ff19168515159081179091558251938452908301527f3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d910160405180910390a15050565b6040805180820190915260008152606060208201526001600160a01b03831661063c57604080518082018252601f81527f63616e27742073656e64206d6f6e657920746f2061646472657373203078300060208083019190915282518084018452600081526060908201528251808401909352600183528201525b905061078a565b6001600160a01b0384166000908152600160205260409020548211156106f4576106356040518060400160405280601781526020017f696e737566666963656e742066756e6473202062616c3a0000000000000000008152506106c360016000336001600160a01b03166001600160a01b0316815260200190815260200160002054610791565b6040518060400160405280600a8152602001690101030b6b7bab73a1d160b51b8152506106ef86610791565b610a04565b6001600160a01b0383166000908152600160205260409020546107178382610c0c565b11610756576106356040518060400160405280601081526020016f034b73b30b634b21030b6b7bab73a1d160851b81525061075184610791565b610a5f565b5060408051808201825260008082526060602092830152825180840184528181528351808401909452908352908101919091525b9392505050565b6060816000036107b85750506040805180820190915260018152600360fc1b602082015290565b8160005b81156107e257806107cc81610d8a565b91506107db9050600a83610db9565b91506107bc565b60008167ffffffffffffffff8111156107fd576107fd610dcd565b6040519080825280601f01601f191660200182016040528015610827576020820181803683370190505b509050815b85156108b45761083d600182610c76565b9050600061084c600a88610db9565b61085790600a610de3565b6108619088610c76565b61086c906030610dfa565b905060008160f81b90508084848151811061088957610889610e13565b60200101906001600160f81b031916908160001a9053506108ab600a89610db9565b9750505061082c565b50949350505050565b60408051602880825260608281019093526000919060208201818036833701905050905060005b60148110156109fd5760006108fa826013610c76565b610905906008610de3565b610910906002610f0d565b610923906001600160a01b038716610db9565b60f81b9050600060108260f81c61093a9190610f19565b60f81b905060008160f81c60106109519190610f3b565b8360f81c61095f9190610f57565b60f81b905061096d82610ab4565b85610979866002610de3565b8151811061098957610989610e13565b60200101906001600160f81b031916908160001a9053506109a981610ab4565b856109b5866002610de3565b6109c0906001610c0c565b815181106109d0576109d0610e13565b60200101906001600160f81b031916908160001a90535050505080806109f590610d8a565b9150506108e4565b5092915050565b604080518082019091526000815260606020820152604051806040016040528060011515815260200186868686604051602001610a449493929190610f70565b60408051601f19818403018152919052905295945050505050565b60408051808201909152600081526060602082015260405180604001604052806001151581526020018484604051602001610a9b929190610fc7565b60408051601f1981840301815291905290529392505050565b6000600a60f883901c1015610adb57610ad260f883901c6030610dfa565b60f81b92915050565b610ad260f883901c6057610dfa565b80356001600160a01b0381168114610b0157600080fd5b919050565b60008060408385031215610b1957600080fd5b610b2283610aea565b946020939093013593505050565b60008060408385031215610b4357600080fd5b610b4c83610aea565b9150610b5a60208401610aea565b90509250929050565b600060208284031215610b7557600080fd5b61078a82610aea565b600080600060608486031215610b9357600080fd5b610b9c84610aea565b9250610baa60208501610aea565b9150604084013590509250925092565b60008060408385031215610bcd57600080fd5b610bd683610aea565b915060208301358015158114610beb57600080fd5b809150509250929050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561026a5761026a610bf6565b60005b83811015610c3a578181015183820152602001610c22565b50506000910152565b6020815260008251806020840152610c62816040850160208701610c1f565b601f01601f19169190910160400192915050565b8181038181111561026a5761026a610bf6565b6a03a3930b739b332b932b2160ad1b815260008451610caf81600b850160208901610c1f565b630103a37960e51b600b918401918201528451610cd381600f840160208901610c1f565b65010333937b6960d51b600f92909101918201528351610cfa816015840160208801610c1f565b0160150195945050505050565b6f03a3930b739b332b932b210333937b6960851b815260008451610d32816010850160208901610c1f565b630103a37960e51b6010918401918201528451610d56816014840160208901610c1f565b65010333937b6960d51b601492909101918201528351610d7d81601a840160208801610c1f565b01601a0195945050505050565b600060018201610d9c57610d9c610bf6565b5060010190565b634e487b7160e01b600052601260045260246000fd5b600082610dc857610dc8610da3565b500490565b634e487b7160e01b600052604160045260246000fd5b808202811582820484141761026a5761026a610bf6565b60ff818116838216019081111561026a5761026a610bf6565b634e487b7160e01b600052603260045260246000fd5b600181815b80851115610e64578160001904821115610e4a57610e4a610bf6565b80851615610e5757918102915b93841c9390800290610e2e565b509250929050565b600082610e7b5750600161026a565b81610e885750600061026a565b8160018114610e9e5760028114610ea857610ec4565b600191505061026a565b60ff841115610eb957610eb9610bf6565b50506001821b61026a565b5060208310610133831016604e8410600b8410161715610ee7575081810a61026a565b610ef18383610e29565b8060001904821115610f0557610f05610bf6565b029392505050565b600061078a8383610e6c565b600060ff831680610f2c57610f2c610da3565b8060ff84160491505092915050565b60ff81811683821602908116908181146109fd576109fd610bf6565b60ff828116828216039081111561026a5761026a610bf6565b60008551610f82818460208a01610c1f565b855190830190610f96818360208a01610c1f565b8551910190610fa9818360208901610c1f565b8451910190610fbc818360208801610c1f565b019695505050505050565b60008351610fd9818460208801610c1f565b835190830190610fed818360208801610c1f565b0194935050505056fea2646970667358221220e573ea8393060072625646a55ced6a23718e97a4606a2a414491de4098cab90c64736f6c63430008110033",
}

// SimpleCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleCoinMetaData.ABI instead.
var SimpleCoinABI = SimpleCoinMetaData.ABI

// Deprecated: Use SimpleCoinMetaData.Sigs instead.
// SimpleCoinFuncSigs maps the 4-byte function signature to its string representation.
var SimpleCoinFuncSigs = SimpleCoinMetaData.Sigs

// SimpleCoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleCoinMetaData.Bin instead.
var SimpleCoinBin = SimpleCoinMetaData.Bin

// DeploySimpleCoin deploys a new Ethereum contract, binding an instance of SimpleCoin to it.
func DeploySimpleCoin(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *SimpleCoin, error) {
	parsed, err := SimpleCoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleCoinBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleCoin{SimpleCoinCaller: SimpleCoinCaller{contract: contract}, SimpleCoinTransactor: SimpleCoinTransactor{contract: contract}, SimpleCoinFilterer: SimpleCoinFilterer{contract: contract}}, nil
}

// SimpleCoin is an auto generated Go binding around an Ethereum contract.
type SimpleCoin struct {
	SimpleCoinCaller     // Read-only binding to the contract
	SimpleCoinTransactor // Write-only binding to the contract
	SimpleCoinFilterer   // Log filterer for contract events
}

// SimpleCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleCoinSession struct {
	Contract     *SimpleCoin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleCoinCallerSession struct {
	Contract *SimpleCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SimpleCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleCoinTransactorSession struct {
	Contract     *SimpleCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimpleCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleCoinRaw struct {
	Contract *SimpleCoin // Generic contract binding to access the raw methods on
}

// SimpleCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleCoinCallerRaw struct {
	Contract *SimpleCoinCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleCoinTransactorRaw struct {
	Contract *SimpleCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleCoin creates a new instance of SimpleCoin, bound to a specific deployed contract.
func NewSimpleCoin(address common.Address, backend bind.ContractBackend) (*SimpleCoin, error) {
	contract, err := bindSimpleCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleCoin{SimpleCoinCaller: SimpleCoinCaller{contract: contract}, SimpleCoinTransactor: SimpleCoinTransactor{contract: contract}, SimpleCoinFilterer: SimpleCoinFilterer{contract: contract}}, nil
}

// NewSimpleCoinCaller creates a new read-only instance of SimpleCoin, bound to a specific deployed contract.
func NewSimpleCoinCaller(address common.Address, caller bind.ContractCaller) (*SimpleCoinCaller, error) {
	contract, err := bindSimpleCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCoinCaller{contract: contract}, nil
}

// NewSimpleCoinTransactor creates a new write-only instance of SimpleCoin, bound to a specific deployed contract.
func NewSimpleCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleCoinTransactor, error) {
	contract, err := bindSimpleCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCoinTransactor{contract: contract}, nil
}

// NewSimpleCoinFilterer creates a new log filterer instance of SimpleCoin, bound to a specific deployed contract.
func NewSimpleCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleCoinFilterer, error) {
	contract, err := bindSimpleCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleCoinFilterer{contract: contract}, nil
}

// bindSimpleCoin binds a generic wrapper to an already deployed contract.
func bindSimpleCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleCoin *SimpleCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleCoin.Contract.SimpleCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleCoin *SimpleCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleCoin.Contract.SimpleCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleCoin *SimpleCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleCoin.Contract.SimpleCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleCoin *SimpleCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleCoin *SimpleCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleCoin *SimpleCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleCoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x6dcdd00f.
//
// Solidity: function Allowance(address , address ) view returns(uint256)
func (_SimpleCoin *SimpleCoinCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SimpleCoin.contract.Call(opts, &out, "Allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x6dcdd00f.
//
// Solidity: function Allowance(address , address ) view returns(uint256)
func (_SimpleCoin *SimpleCoinSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SimpleCoin.Contract.Allowance(&_SimpleCoin.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0x6dcdd00f.
//
// Solidity: function Allowance(address , address ) view returns(uint256)
func (_SimpleCoin *SimpleCoinCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SimpleCoin.Contract.Allowance(&_SimpleCoin.CallOpts, arg0, arg1)
}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_SimpleCoin *SimpleCoinCaller) CoinBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SimpleCoin.contract.Call(opts, &out, "CoinBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_SimpleCoin *SimpleCoinSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _SimpleCoin.Contract.CoinBalance(&_SimpleCoin.CallOpts, arg0)
}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_SimpleCoin *SimpleCoinCallerSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _SimpleCoin.Contract.CoinBalance(&_SimpleCoin.CallOpts, arg0)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_SimpleCoin *SimpleCoinCaller) FrozenAccount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SimpleCoin.contract.Call(opts, &out, "FrozenAccount", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_SimpleCoin *SimpleCoinSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _SimpleCoin.Contract.FrozenAccount(&_SimpleCoin.CallOpts, arg0)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_SimpleCoin *SimpleCoinCallerSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _SimpleCoin.Contract.FrozenAccount(&_SimpleCoin.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_SimpleCoin *SimpleCoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleCoin.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_SimpleCoin *SimpleCoinSession) Owner() (common.Address, error) {
	return _SimpleCoin.Contract.Owner(&_SimpleCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_SimpleCoin *SimpleCoinCallerSession) Owner() (common.Address, error) {
	return _SimpleCoin.Contract.Owner(&_SimpleCoin.CallOpts)
}

// Authorize is a paid mutator transaction binding the contract method 0x6386a433.
//
// Solidity: function Authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_SimpleCoin *SimpleCoinTransactor) Authorize(opts *bind.TransactOpts, authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.contract.Transact(opts, "Authorize", authorizedAccount, allowanceAmount)
}

// Authorize is a paid mutator transaction binding the contract method 0x6386a433.
//
// Solidity: function Authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_SimpleCoin *SimpleCoinSession) Authorize(authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.Contract.Authorize(&_SimpleCoin.TransactOpts, authorizedAccount, allowanceAmount)
}

// Authorize is a paid mutator transaction binding the contract method 0x6386a433.
//
// Solidity: function Authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_SimpleCoin *SimpleCoinTransactorSession) Authorize(authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.Contract.Authorize(&_SimpleCoin.TransactOpts, authorizedAccount, allowanceAmount)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_SimpleCoin *SimpleCoinTransactor) FreezeAccount(opts *bind.TransactOpts, target common.Address, freeze bool) (*types.Transaction, error) {
	return _SimpleCoin.contract.Transact(opts, "FreezeAccount", target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_SimpleCoin *SimpleCoinSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _SimpleCoin.Contract.FreezeAccount(&_SimpleCoin.TransactOpts, target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_SimpleCoin *SimpleCoinTransactorSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _SimpleCoin.Contract.FreezeAccount(&_SimpleCoin.TransactOpts, target, freeze)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_SimpleCoin *SimpleCoinTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.contract.Transact(opts, "Mint", recipient, mintedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_SimpleCoin *SimpleCoinSession) Mint(recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.Contract.Mint(&_SimpleCoin.TransactOpts, recipient, mintedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_SimpleCoin *SimpleCoinTransactorSession) Mint(recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.Contract.Mint(&_SimpleCoin.TransactOpts, recipient, mintedAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_SimpleCoin *SimpleCoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.contract.Transact(opts, "Transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_SimpleCoin *SimpleCoinSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.Contract.Transfer(&_SimpleCoin.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_SimpleCoin *SimpleCoinTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.Contract.Transfer(&_SimpleCoin.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xc0d84ce5.
//
// Solidity: function TransferFrom(address from, address to, uint256 amount) returns()
func (_SimpleCoin *SimpleCoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.contract.Transact(opts, "TransferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xc0d84ce5.
//
// Solidity: function TransferFrom(address from, address to, uint256 amount) returns()
func (_SimpleCoin *SimpleCoinSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.Contract.TransferFrom(&_SimpleCoin.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xc0d84ce5.
//
// Solidity: function TransferFrom(address from, address to, uint256 amount) returns()
func (_SimpleCoin *SimpleCoinTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.Contract.TransferFrom(&_SimpleCoin.TransactOpts, from, to, amount)
}

// SimpleCoinEventFrozenAccountIterator is returned from FilterEventFrozenAccount and is used to iterate over the raw logs and unpacked data for EventFrozenAccount events raised by the SimpleCoin contract.
type SimpleCoinEventFrozenAccountIterator struct {
	Event *SimpleCoinEventFrozenAccount // Event containing the contract specifics and raw log

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
func (it *SimpleCoinEventFrozenAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCoinEventFrozenAccount)
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
		it.Event = new(SimpleCoinEventFrozenAccount)
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
func (it *SimpleCoinEventFrozenAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCoinEventFrozenAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCoinEventFrozenAccount represents a EventFrozenAccount event raised by the SimpleCoin contract.
type SimpleCoinEventFrozenAccount struct {
	Target common.Address
	Frozen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventFrozenAccount is a free log retrieval operation binding the contract event 0x3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d.
//
// Solidity: event EventFrozenAccount(address target, bool frozen)
func (_SimpleCoin *SimpleCoinFilterer) FilterEventFrozenAccount(opts *bind.FilterOpts) (*SimpleCoinEventFrozenAccountIterator, error) {

	logs, sub, err := _SimpleCoin.contract.FilterLogs(opts, "EventFrozenAccount")
	if err != nil {
		return nil, err
	}
	return &SimpleCoinEventFrozenAccountIterator{contract: _SimpleCoin.contract, event: "EventFrozenAccount", logs: logs, sub: sub}, nil
}

// WatchEventFrozenAccount is a free log subscription operation binding the contract event 0x3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d.
//
// Solidity: event EventFrozenAccount(address target, bool frozen)
func (_SimpleCoin *SimpleCoinFilterer) WatchEventFrozenAccount(opts *bind.WatchOpts, sink chan<- *SimpleCoinEventFrozenAccount) (event.Subscription, error) {

	logs, sub, err := _SimpleCoin.contract.WatchLogs(opts, "EventFrozenAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCoinEventFrozenAccount)
				if err := _SimpleCoin.contract.UnpackLog(event, "EventFrozenAccount", log); err != nil {
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
func (_SimpleCoin *SimpleCoinFilterer) ParseEventFrozenAccount(log types.Log) (*SimpleCoinEventFrozenAccount, error) {
	event := new(SimpleCoinEventFrozenAccount)
	if err := _SimpleCoin.contract.UnpackLog(event, "EventFrozenAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleCoinEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the SimpleCoin contract.
type SimpleCoinEventLogIterator struct {
	Event *SimpleCoinEventLog // Event containing the contract specifics and raw log

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
func (it *SimpleCoinEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCoinEventLog)
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
		it.Event = new(SimpleCoinEventLog)
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
func (it *SimpleCoinEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCoinEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCoinEventLog represents a EventLog event raised by the SimpleCoin contract.
type SimpleCoinEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_SimpleCoin *SimpleCoinFilterer) FilterEventLog(opts *bind.FilterOpts) (*SimpleCoinEventLogIterator, error) {

	logs, sub, err := _SimpleCoin.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &SimpleCoinEventLogIterator{contract: _SimpleCoin.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_SimpleCoin *SimpleCoinFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *SimpleCoinEventLog) (event.Subscription, error) {

	logs, sub, err := _SimpleCoin.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCoinEventLog)
				if err := _SimpleCoin.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_SimpleCoin *SimpleCoinFilterer) ParseEventLog(log types.Log) (*SimpleCoinEventLog, error) {
	event := new(SimpleCoinEventLog)
	if err := _SimpleCoin.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleCoinEventTransferIterator is returned from FilterEventTransfer and is used to iterate over the raw logs and unpacked data for EventTransfer events raised by the SimpleCoin contract.
type SimpleCoinEventTransferIterator struct {
	Event *SimpleCoinEventTransfer // Event containing the contract specifics and raw log

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
func (it *SimpleCoinEventTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCoinEventTransfer)
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
		it.Event = new(SimpleCoinEventTransfer)
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
func (it *SimpleCoinEventTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCoinEventTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCoinEventTransfer represents a EventTransfer event raised by the SimpleCoin contract.
type SimpleCoinEventTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventTransfer is a free log retrieval operation binding the contract event 0x64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45.
//
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 value)
func (_SimpleCoin *SimpleCoinFilterer) FilterEventTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleCoinEventTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleCoin.contract.FilterLogs(opts, "EventTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SimpleCoinEventTransferIterator{contract: _SimpleCoin.contract, event: "EventTransfer", logs: logs, sub: sub}, nil
}

// WatchEventTransfer is a free log subscription operation binding the contract event 0x64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45.
//
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 value)
func (_SimpleCoin *SimpleCoinFilterer) WatchEventTransfer(opts *bind.WatchOpts, sink chan<- *SimpleCoinEventTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleCoin.contract.WatchLogs(opts, "EventTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCoinEventTransfer)
				if err := _SimpleCoin.contract.UnpackLog(event, "EventTransfer", log); err != nil {
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
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 value)
func (_SimpleCoin *SimpleCoinFilterer) ParseEventTransfer(log types.Log) (*SimpleCoinEventTransfer, error) {
	event := new(SimpleCoinEventTransfer)
	if err := _SimpleCoin.contract.UnpackLog(event, "EventTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

