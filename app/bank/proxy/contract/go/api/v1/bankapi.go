// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bankapi

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

// BankapiMetaData contains all meta data concerning the Bankapi contract.
var BankapiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"API\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506110cc806100606000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80637d7b009914610046578063b4a99a4e14610064578063fa84fd8e14610082575b600080fd5b61004e61009e565b60405161005b9190610869565b60405180910390f35b61006c6100c4565b6040516100799190610869565b60405180910390f35b61009c60048036038101906100979190610955565b6100e8565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600082905060005b858590508110156103bb578360026000888885818110610113576101126109dd565b5b90506020020160208101906101289190610a0c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561031c577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6101f9600260008989868181106101a4576101a36109dd565b5b90506020020160208101906101b99190610a0c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546106a0565b610202866106a0565b604051602001610213929190610af6565b60405160208183030381529060405260405161022f9190610b93565b60405180910390a16002600087878481811061024e5761024d6109dd565b5b90506020020160208101906102639190610a0c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826102a99190610be4565b91506000600260008888858181106102c4576102c36109dd565b5b90506020020160208101906102d99190610a0c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506103a8565b83826103289190610be4565b91508360026000888885818110610342576103416109dd565b5b90506020020160208101906103579190610a0c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546103a09190610c18565b925050819055505b80806103b390610c4c565b9150506100f0565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103e6846106a0565b6103ef846106a0565b6103f8846106a0565b60405160200161040a93929190610d2c565b6040516020818303038152906040526040516104269190610b93565b60405180910390a160008103610471576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046890610e0b565b60405180910390fd5b818110156105545780600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104e99190610be4565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61051a826106a0565b60405160200161052a9190610e9d565b6040516020818303038152906040526040516105469190610b93565b60405180910390a150610699565b81816105609190610c18565b905080600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105b19190610be4565b9250508190555081600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106289190610be4565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610659826106a0565b610662846106a0565b604051602001610673929190610f1a565b60405160208183030381529060405260405161068f9190610b93565b60405180910390a1505b5050505050565b6060600082036106e7576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610823565b600082905060005b6000821461071957808061070290610c4c565b915050600a826107129190610f9a565b91506106ef565b60008167ffffffffffffffff81111561073557610734610fcb565b5b6040519080825280601f01601f1916602001820160405280156107675781602001600182028036833780820191505090505b50905060008290505b6000861461081b576001816107859190610c18565b90506000600a80886107979190610f9a565b6107a19190610ffa565b876107ac9190610c18565b60306107b89190611061565b905060008160f81b9050808484815181106107d6576107d56109dd565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a886108129190610f9a565b97505050610770565b819450505050505b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061085382610828565b9050919050565b61086381610848565b82525050565b600060208201905061087e600083018461085a565b92915050565b600080fd5b600080fd5b61089781610848565b81146108a257600080fd5b50565b6000813590506108b48161088e565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126108df576108de6108ba565b5b8235905067ffffffffffffffff8111156108fc576108fb6108bf565b5b602083019150836020820283011115610918576109176108c4565b5b9250929050565b6000819050919050565b6109328161091f565b811461093d57600080fd5b50565b60008135905061094f81610929565b92915050565b60008060008060006080868803121561097157610970610884565b5b600061097f888289016108a5565b955050602086013567ffffffffffffffff8111156109a05761099f610889565b5b6109ac888289016108c9565b945094505060406109bf88828901610940565b92505060606109d088828901610940565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610a2257610a21610884565b5b6000610a30848285016108a5565b91505092915050565b7f6163636f756e742062616c616e63652000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b83811015610a93578082015181840152602081019050610a78565b60008484015250505050565b6000610aaa82610a5f565b610ab48185610a6a565b9350610ac4818560208601610a75565b80840191505092915050565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b6000610b0182610a39565b601082019150610b118285610a9f565b9150610b1c82610ad0565b601382019150610b2c8284610a9f565b91508190509392505050565b600082825260208201905092915050565b6000601f19601f8301169050919050565b6000610b6582610a5f565b610b6f8185610b38565b9350610b7f818560208601610a75565b610b8881610b49565b840191505092915050565b60006020820190508181036000830152610bad8184610b5a565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610bef8261091f565b9150610bfa8361091f565b9250828201905080821115610c1257610c11610bb5565b5b92915050565b6000610c238261091f565b9150610c2e8361091f565b9250828203905081811115610c4657610c45610bb5565b5b92915050565b6000610c578261091f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c8957610c88610bb5565b5b600182019050919050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b6000610d3782610c94565b600582019150610d478286610a9f565b9150610d5282610cba565b600a82019150610d628285610a9f565b9150610d6d82610ce0565b600682019150610d7d8284610a9f565b9150610d8882610d06565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e206163636f60008201527f756e742062616c616e6365730000000000000000000000000000000000000000602082015250565b6000610df5602c83610b38565b9150610e0082610d99565b604082019050919050565b60006020820190508181036000830152610e2481610de8565b9050919050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e60008201527f65725b0000000000000000000000000000000000000000000000000000000000602082015250565b6000610e87602383610a6a565b9150610e9282610e2b565b602382019050919050565b6000610ea882610e7a565b9150610eb48284610a9f565b9150610ebf82610d06565b60018201915081905092915050565b7f77696e6e65725b00000000000000000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b6000610f2582610ece565b600782019150610f358285610a9f565b9150610f4082610ef4565b600882019150610f508284610a9f565b9150610f5b82610d06565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610fa58261091f565b9150610fb08361091f565b925082610fc057610fbf610f6b565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006110058261091f565b91506110108361091f565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561104957611048610bb5565b5b828202905092915050565b600060ff82169050919050565b600061106c82611054565b915061107783611054565b9250828201905060ff8111156110905761108f610bb5565b5b9291505056fea2646970667358221220677da024edd1953eb6fae5754c87c75ba253e6583422cda980f92d5bd358e12b64736f6c63430008100033",
}

// BankapiABI is the input ABI used to generate the binding from.
// Deprecated: Use BankapiMetaData.ABI instead.
var BankapiABI = BankapiMetaData.ABI

// BankapiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankapiMetaData.Bin instead.
var BankapiBin = BankapiMetaData.Bin

// DeployBankapi deploys a new Ethereum contract, binding an instance of Bankapi to it.
func DeployBankapi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bankapi, error) {
	parsed, err := BankapiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankapiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bankapi{BankapiCaller: BankapiCaller{contract: contract}, BankapiTransactor: BankapiTransactor{contract: contract}, BankapiFilterer: BankapiFilterer{contract: contract}}, nil
}

// Bankapi is an auto generated Go binding around an Ethereum contract.
type Bankapi struct {
	BankapiCaller     // Read-only binding to the contract
	BankapiTransactor // Write-only binding to the contract
	BankapiFilterer   // Log filterer for contract events
}

// BankapiCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankapiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankapiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankapiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankapiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankapiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankapiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankapiSession struct {
	Contract     *Bankapi          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankapiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankapiCallerSession struct {
	Contract *BankapiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BankapiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankapiTransactorSession struct {
	Contract     *BankapiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BankapiRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankapiRaw struct {
	Contract *Bankapi // Generic contract binding to access the raw methods on
}

// BankapiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankapiCallerRaw struct {
	Contract *BankapiCaller // Generic read-only contract binding to access the raw methods on
}

// BankapiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankapiTransactorRaw struct {
	Contract *BankapiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBankapi creates a new instance of Bankapi, bound to a specific deployed contract.
func NewBankapi(address common.Address, backend bind.ContractBackend) (*Bankapi, error) {
	contract, err := bindBankapi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bankapi{BankapiCaller: BankapiCaller{contract: contract}, BankapiTransactor: BankapiTransactor{contract: contract}, BankapiFilterer: BankapiFilterer{contract: contract}}, nil
}

// NewBankapiCaller creates a new read-only instance of Bankapi, bound to a specific deployed contract.
func NewBankapiCaller(address common.Address, caller bind.ContractCaller) (*BankapiCaller, error) {
	contract, err := bindBankapi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankapiCaller{contract: contract}, nil
}

// NewBankapiTransactor creates a new write-only instance of Bankapi, bound to a specific deployed contract.
func NewBankapiTransactor(address common.Address, transactor bind.ContractTransactor) (*BankapiTransactor, error) {
	contract, err := bindBankapi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankapiTransactor{contract: contract}, nil
}

// NewBankapiFilterer creates a new log filterer instance of Bankapi, bound to a specific deployed contract.
func NewBankapiFilterer(address common.Address, filterer bind.ContractFilterer) (*BankapiFilterer, error) {
	contract, err := bindBankapi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankapiFilterer{contract: contract}, nil
}

// bindBankapi binds a generic wrapper to an already deployed contract.
func bindBankapi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankapiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankapi *BankapiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankapi.Contract.BankapiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankapi *BankapiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.Contract.BankapiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankapi *BankapiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankapi.Contract.BankapiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankapi *BankapiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankapi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankapi *BankapiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankapi *BankapiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankapi.Contract.contract.Transact(opts, method, params...)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankapi *BankapiCaller) API(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankapi.contract.Call(opts, &out, "API")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankapi *BankapiSession) API() (common.Address, error) {
	return _Bankapi.Contract.API(&_Bankapi.CallOpts)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankapi *BankapiCallerSession) API() (common.Address, error) {
	return _Bankapi.Contract.API(&_Bankapi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankapi *BankapiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankapi.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankapi *BankapiSession) Owner() (common.Address, error) {
	return _Bankapi.Contract.Owner(&_Bankapi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankapi *BankapiCallerSession) Owner() (common.Address, error) {
	return _Bankapi.Contract.Owner(&_Bankapi.CallOpts)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bankapi *BankapiTransactor) Reconcile(opts *bind.TransactOpts, winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankapi.contract.Transact(opts, "Reconcile", winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bankapi *BankapiSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankapi.Contract.Reconcile(&_Bankapi.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bankapi *BankapiTransactorSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankapi.Contract.Reconcile(&_Bankapi.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// BankapiEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bankapi contract.
type BankapiEventLogIterator struct {
	Event *BankapiEventLog // Event containing the contract specifics and raw log

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
func (it *BankapiEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankapiEventLog)
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
		it.Event = new(BankapiEventLog)
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
func (it *BankapiEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankapiEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankapiEventLog represents a EventLog event raised by the Bankapi contract.
type BankapiEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankapi *BankapiFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankapiEventLogIterator, error) {

	logs, sub, err := _Bankapi.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankapiEventLogIterator{contract: _Bankapi.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankapi *BankapiFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankapiEventLog) (event.Subscription, error) {

	logs, sub, err := _Bankapi.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankapiEventLog)
				if err := _Bankapi.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_Bankapi *BankapiFilterer) ParseEventLog(log types.Log) (*BankapiEventLog, error) {
	event := new(BankapiEventLog)
	if err := _Bankapi.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
