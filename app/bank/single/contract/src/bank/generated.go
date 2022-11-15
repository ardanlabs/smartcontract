// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e63f341f": "AccountBalance(address)",
		"0ef67887": "Balance()",
		"ed21248c": "Deposit()",
		"b4a99a4e": "Owner()",
		"fa84fd8e": "Reconcile(address,address[],uint256,uint256)",
		"bb62860d": "Version()",
		"57ea89b6": "Withdraw()",
	},
	Bin: "0x60806040523480156200001157600080fd5b50600080546001600160a01b031916331790556040805180820190915260058152640302e312e360dc1b60208201526001906200004f9082620000fb565b50620001c7565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200008157607f821691505b602082108103620000a257634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620000f657600081815260208120601f850160051c81016020861015620000d15750805b601f850160051c820191505b81811015620000f257828155600101620000dd565b5050505b505050565b81516001600160401b0381111562000117576200011762000056565b6200012f816200012884546200006c565b84620000a8565b602080601f8311600181146200016757600084156200014e5750858301515b600019600386901b1c1916600185901b178555620000f2565b600085815260208120601f198616915b82811015620001985788860151825594840194600190910190840162000177565b5085821015620001b75787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6110de80620001d76000396000f3fe6080604052600436106100705760003560e01c8063bb62860d1161004e578063bb62860d146100e8578063e63f341f1461010a578063ed21248c1461012a578063fa84fd8e1461013257600080fd5b80630ef678871461007557806357ea89b6146100a6578063b4a99a4e146100b0575b600080fd5b34801561008157600080fd5b50336000908152600260205260409020545b6040519081526020015b60405180910390f35b6100ae610152565b005b3480156100bc57600080fd5b506000546100d0906001600160a01b031681565b6040516001600160a01b03909116815260200161009d565b3480156100f457600080fd5b506100fd61025b565b60405161009d9190610abb565b34801561011657600080fd5b50610093610125366004610b05565b6102e9565b6100ae610321565b34801561013e57600080fd5b506100ae61014d366004610b27565b6103ac565b3360008181526002602052604081205490036101aa5760405162461bcd60e51b81526020600482015260126024820152716e6f7420656e6f7567682062616c616e636560701b60448201526064015b60405180910390fd5b3360009081526002602052604080822054905190916001600160a01b0384169183156108fc0291849190818181858888f193505050501580156101f1573d6000803e3d6000fd5b50336000818152600260205260408120556000805160206110898339815191529061021b906107ee565b61022483610935565b604051602001610235929190610bbc565b60408051601f198184030181529082905261024f91610abb565b60405180910390a15050565b6001805461026890610c23565b80601f016020809104026020016040519081016040528092919081815260200182805461029490610c23565b80156102e15780601f106102b6576101008083540402835291602001916102e1565b820191906000526020600020905b8154815290600101906020018083116102c457829003601f168201915b505050505081565b600080546001600160a01b0316331461030157600080fd5b506001600160a01b0381166000908152600260205260409020545b919050565b3360009081526002602052604081208054349290610340908490610c73565b90915550600080516020611089833981519152905061035e336107ee565b3360009081526002602052604090205461037790610935565b604051602001610388929190610c8c565b60408051601f19818403018152908290526103a291610abb565b60405180910390a1565b6000546001600160a01b031633146103c357600080fd5b8160005b848110156105dd5783600260008888858181106103e6576103e6610cd9565b90506020020160208101906103fb9190610b05565b6001600160a01b03166001600160a01b0316815260200190815260200160002054101561055e576000805160206110898339815191526104856002600089898681811061044a5761044a610cd9565b905060200201602081019061045f9190610b05565b6001600160a01b03166001600160a01b0316815260200190815260200160002054610935565b61048e86610935565b60405160200161049f929190610cef565b60408051601f19818403018152908290526104b991610abb565b60405180910390a1600260008787848181106104d7576104d7610cd9565b90506020020160208101906104ec9190610b05565b6001600160a01b031681526020810191909152604001600020546105109083610c73565b915060006002600088888581811061052a5761052a610cd9565b905060200201602081019061053f9190610b05565b6001600160a01b031681526020810191909152604001600020556105cb565b6105688483610c73565b9150836002600088888581811061058157610581610cd9565b90506020020160208101906105969190610b05565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008282546105c59190610d59565b90915550505b806105d581610d6c565b9150506103c7565b506000805160206110898339815191526105f684610935565b6105ff84610935565b61060884610935565b60405160200161061a93929190610d85565b60408051601f198184030181529082905261063491610abb565b60405180910390a1806000036106a15760405162461bcd60e51b815260206004820152602c60248201527f6e6f20706f74207761732063726561746564206261736564206f6e206163636f60448201526b756e742062616c616e63657360a01b60648201526084016101a1565b8181101561072757600080546001600160a01b0316815260026020526040812080548392906106d1908490610c73565b9091555060008051602061108983398151915290506106ef82610935565b6040516020016106ff9190610e11565b60408051601f198184030181529082905261071991610abb565b60405180910390a1506107e7565b6107318282610d59565b6001600160a01b03871660009081526002602052604081208054929350839290919061075e908490610c73565b9091555050600080546001600160a01b03168152600260205260408120805484929061078b908490610c73565b9091555060008051602061108983398151915290506107a982610935565b6107b284610935565b6040516020016107c3929190610e6d565b60408051601f19818403018152908290526107dd91610abb565b60405180910390a1505b5050505050565b60408051602880825260608281019093526000919060208201818036833701905050905060005b601481101561092e57600061082b826013610d59565b610836906008610ee7565b610841906002610fe2565b610854906001600160a01b038716611004565b60f81b9050600060108260f81c61086b9190611018565b60f81b905060008160f81c6010610882919061103a565b8360f81c6108909190611056565b60f81b905061089e82610a61565b856108aa866002610ee7565b815181106108ba576108ba610cd9565b60200101906001600160f81b031916908160001a9053506108da81610a61565b856108e6866002610ee7565b6108f1906001610c73565b8151811061090157610901610cd9565b60200101906001600160f81b031916908160001a905350505050808061092690610d6c565b915050610815565b5092915050565b60608160000361095c5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115610986578061097081610d6c565b915061097f9050600a83611004565b9150610960565b60008167ffffffffffffffff8111156109a1576109a1610ed1565b6040519080825280601f01601f1916602001820160405280156109cb576020820181803683370190505b509050815b8515610a58576109e1600182610d59565b905060006109f0600a88611004565b6109fb90600a610ee7565b610a059088610d59565b610a1090603061106f565b905060008160f81b905080848481518110610a2d57610a2d610cd9565b60200101906001600160f81b031916908160001a905350610a4f600a89611004565b975050506109d0565b50949350505050565b6000600a60f883901c1015610a8857610a7f60f883901c603061106f565b60f81b92915050565b610a7f60f883901c605761106f565b60005b83811015610ab2578181015183820152602001610a9a565b50506000910152565b6020815260008251806020840152610ada816040850160208701610a97565b601f01601f19169190910160400192915050565b80356001600160a01b038116811461031c57600080fd5b600060208284031215610b1757600080fd5b610b2082610aee565b9392505050565b600080600080600060808688031215610b3f57600080fd5b610b4886610aee565b9450602086013567ffffffffffffffff80821115610b6557600080fd5b818801915088601f830112610b7957600080fd5b813581811115610b8857600080fd5b8960208260051b8501011115610b9d57600080fd5b9699602092909201985095966040810135965060600135945092505050565b6877697468647261775b60b81b815260008351610be0816009850160208801610a97565b685d20616d6f756e745b60b81b6009918401918201528351610c09816012840160208801610a97565b605d60f81b60129290910191820152601301949350505050565b600181811c90821680610c3757607f821691505b602082108103610c5757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b80820180821115610c8657610c86610c5d565b92915050565b676465706f7369745b60c01b815260008351610caf816008850160208801610a97565b695d2062616c616e63655b60b01b6008918401918201528351610c09816012840160208801610a97565b634e487b7160e01b600052603260045260246000fd5b6f030b1b1b7bab73a103130b630b731b2960851b815260008351610d1a816010850160208801610a97565b7201034b9903632b9b9903a3430b71030b73a329606d1b6010918401918201528351610d4d816023840160208801610a97565b01602301949350505050565b81810381811115610c8657610c86610c5d565b600060018201610d7e57610d7e610c5d565b5060010190565b64616e74655b60d81b815260008451610da5816005850160208901610a97565b695d2067616d654665655b60b01b6005918401918201528451610dcf81600f840160208901610a97565b655d20706f745b60d01b600f92909101918201528351610df6816015840160208801610a97565b605d60f81b6015929091019182015260160195945050505050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e81526265725b60e81b602082015260008251610e55816023850160208701610a97565b605d60f81b6023939091019283015250602401919050565b6677696e6e65725b60c81b815260008351610e8f816007850160208801610a97565b675d206f776e65725b60c01b6007918401918201528351610eb781600f840160208801610a97565b605d60f81b600f9290910191820152601001949350505050565b634e487b7160e01b600052604160045260246000fd5b8082028115828204841417610c8657610c86610c5d565b600181815b80851115610f39578160001904821115610f1f57610f1f610c5d565b80851615610f2c57918102915b93841c9390800290610f03565b509250929050565b600082610f5057506001610c86565b81610f5d57506000610c86565b8160018114610f735760028114610f7d57610f99565b6001915050610c86565b60ff841115610f8e57610f8e610c5d565b50506001821b610c86565b5060208310610133831016604e8410600b8410161715610fbc575081810a610c86565b610fc68383610efe565b8060001904821115610fda57610fda610c5d565b029392505050565b6000610b208383610f41565b634e487b7160e01b600052601260045260246000fd5b60008261101357611013610fee565b500490565b600060ff83168061102b5761102b610fee565b8060ff84160491505092915050565b60ff818116838216029081169081811461092e5761092e610c5d565b60ff8281168282160390811115610c8657610c86610c5d565b60ff8181168382160190811115610c8657610c86610c5d56fed3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18aa264697066735822122063ebf996f02daae40fdeefafb06b96f95f9ecb433461f821d33fce4f0952bf6964736f6c63430008110033",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// Deprecated: Use BankMetaData.Sigs instead.
// BankFuncSigs maps the 4-byte function signature to its string representation.
var BankFuncSigs = BankMetaData.Sigs

// BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankMetaData.Bin instead.
var BankBin = BankMetaData.Bin

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCaller) AccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCallerSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCallerSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCallerSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankSession) Version() (string, error) {
	return _Bank.Contract.Version(&_Bank.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankCallerSession) Version() (string, error) {
	return _Bank.Contract.Version(&_Bank.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bank *BankTransactor) Reconcile(opts *bind.TransactOpts, winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Reconcile", winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bank *BankSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Reconcile(&_Bank.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bank *BankTransactorSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Reconcile(&_Bank.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// BankEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bank contract.
type BankEventLogIterator struct {
	Event *BankEventLog // Event containing the contract specifics and raw log

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
func (it *BankEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventLog)
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
		it.Event = new(BankEventLog)
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
func (it *BankEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventLog represents a EventLog event raised by the Bank contract.
type BankEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankEventLogIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankEventLogIterator{contract: _Bank.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankEventLog) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventLog)
				if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_Bank *BankFilterer) ParseEventLog(log types.Log) (*BankEventLog, error) {
	event := new(BankEventLog)
	if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErrorMetaData contains all meta data concerning the Error contract.
var ErrorMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122022b442d871437b3e7f209cc77bac269ea41432fe502ff2d0132764bb9909443664736f6c63430008110033",
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

