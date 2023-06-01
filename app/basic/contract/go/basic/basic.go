// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basic

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

// BasicMetaData contains all meta data concerning the Basic contract.
var BasicMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Items\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SetItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600381526020017f312e31000000000000000000000000000000000000000000000000000000000081525060009081620000589190620002df565b50620003c6565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620000e157607f821691505b602082108103620000f757620000f662000099565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b60008160020a8302905092915050565b600060088302620001647fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000122565b62000170868362000122565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620001bd620001b7620001b18462000188565b62000192565b62000188565b9050919050565b6000819050919050565b620001d9836200019c565b620001f1620001e882620001c4565b84845462000132565b825550505050565b600090565b62000208620001f9565b62000215818484620001ce565b505050565b5b818110156200023d5762000231600082620001fe565b6001810190506200021b565b5050565b601f8211156200028c576200025681620000fd565b620002618462000112565b8101602085101562000271578190505b62000289620002808562000112565b8301826200021a565b50505b505050565b60008160020a8304905092915050565b6000620002b46000198460080262000291565b1980831691505092915050565b6000620002cf8383620002a1565b9150826002028217905092915050565b620002ea826200005f565b67ffffffffffffffff8111156200030657620003056200006a565b5b620003128254620000c8565b6200031f82828562000241565b600060209050601f83116001811462000357576000841562000342578287015190505b6200034e8582620002c1565b865550620003be565b601f1984166200036786620000fd565b60005b8281101562000391578489015182556001820191506020850194506020810190506200036a565b86831015620003b15784890151620003ad601f891682620002a1565b8355505b6001600288020188555050505b505050505050565b61060280620003d66000396000f3fe608060405234801561001057600080fd5b506004361061005e576000357c0100000000000000000000000000000000000000000000000000000000900480634547a6b3146100635780638c5cf3ed14610093578063bb62860d146100af575b600080fd5b61007d60048036038101906100789190610343565b6100cd565b60405161008a91906103a5565b60405180910390f35b6100ad60048036038101906100a891906103ec565b6100fb565b005b6100b761015b565b6040516100c491906104c7565b60405180910390f35b6001818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b8060018360405161010c9190610525565b9081526020016040518091039020819055507f814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f828260405161014f92919061053c565b60405180910390a15050565b600080546101689061059b565b80601f01602080910402602001604051908101604052809291908181526020018280546101949061059b565b80156101e15780601f106101b6576101008083540402835291602001916101e1565b820191906000526020600020905b8154815290600101906020018083116101c457829003601f168201915b505050505081565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61025082610207565b810181811067ffffffffffffffff8211171561026f5761026e610218565b5b80604052505050565b60006102826101e9565b905061028e8282610247565b919050565b600067ffffffffffffffff8211156102ae576102ad610218565b5b6102b782610207565b9050602081019050919050565b82818337600083830152505050565b60006102e66102e184610293565b610278565b90508281526020810184848401111561030257610301610202565b5b61030d8482856102c4565b509392505050565b600082601f83011261032a576103296101fd565b5b813561033a8482602086016102d3565b91505092915050565b600060208284031215610359576103586101f3565b5b600082013567ffffffffffffffff811115610377576103766101f8565b5b61038384828501610315565b91505092915050565b6000819050919050565b61039f8161038c565b82525050565b60006020820190506103ba6000830184610396565b92915050565b6103c98161038c565b81146103d457600080fd5b50565b6000813590506103e6816103c0565b92915050565b60008060408385031215610403576104026101f3565b5b600083013567ffffffffffffffff811115610421576104206101f8565b5b61042d85828601610315565b925050602061043e858286016103d7565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610482578082015181840152602081019050610467565b60008484015250505050565b600061049982610448565b6104a38185610453565b93506104b3818560208601610464565b6104bc81610207565b840191505092915050565b600060208201905081810360008301526104e1818461048e565b905092915050565b600081905092915050565b60006104ff82610448565b61050981856104e9565b9350610519818560208601610464565b80840191505092915050565b600061053182846104f4565b915081905092915050565b60006040820190508181036000830152610556818561048e565b90506105656020830184610396565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806105b357607f821691505b6020821081036105c6576105c561056c565b5b5091905056fea264697066735822122079a0ac0885cef0ca3e8ea5b9c32b2769af46cd7c9864cc206b978ce0befe8bbf64736f6c63430008140033",
}

// BasicABI is the input ABI used to generate the binding from.
// Deprecated: Use BasicMetaData.ABI instead.
var BasicABI = BasicMetaData.ABI

// BasicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BasicMetaData.Bin instead.
var BasicBin = BasicMetaData.Bin

// DeployBasic deploys a new Ethereum contract, binding an instance of Basic to it.
func DeployBasic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Basic, error) {
	parsed, err := BasicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Basic{BasicCaller: BasicCaller{contract: contract}, BasicTransactor: BasicTransactor{contract: contract}, BasicFilterer: BasicFilterer{contract: contract}}, nil
}

// Basic is an auto generated Go binding around an Ethereum contract.
type Basic struct {
	BasicCaller     // Read-only binding to the contract
	BasicTransactor // Write-only binding to the contract
	BasicFilterer   // Log filterer for contract events
}

// BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicSession struct {
	Contract     *Basic            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicCallerSession struct {
	Contract *BasicCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTransactorSession struct {
	Contract     *BasicTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicRaw struct {
	Contract *Basic // Generic contract binding to access the raw methods on
}

// BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicCallerRaw struct {
	Contract *BasicCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTransactorRaw struct {
	Contract *BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasic creates a new instance of Basic, bound to a specific deployed contract.
func NewBasic(address common.Address, backend bind.ContractBackend) (*Basic, error) {
	contract, err := bindBasic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Basic{BasicCaller: BasicCaller{contract: contract}, BasicTransactor: BasicTransactor{contract: contract}, BasicFilterer: BasicFilterer{contract: contract}}, nil
}

// NewBasicCaller creates a new read-only instance of Basic, bound to a specific deployed contract.
func NewBasicCaller(address common.Address, caller bind.ContractCaller) (*BasicCaller, error) {
	contract, err := bindBasic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicCaller{contract: contract}, nil
}

// NewBasicTransactor creates a new write-only instance of Basic, bound to a specific deployed contract.
func NewBasicTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTransactor, error) {
	contract, err := bindBasic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTransactor{contract: contract}, nil
}

// NewBasicFilterer creates a new log filterer instance of Basic, bound to a specific deployed contract.
func NewBasicFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicFilterer, error) {
	contract, err := bindBasic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicFilterer{contract: contract}, nil
}

// bindBasic binds a generic wrapper to an already deployed contract.
func bindBasic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BasicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Basic *BasicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Basic.Contract.BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Basic *BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basic.Contract.BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Basic *BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Basic.Contract.BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Basic *BasicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Basic *BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Basic *BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Basic.Contract.contract.Transact(opts, method, params...)
}

// Items is a free data retrieval call binding the contract method 0x4547a6b3.
//
// Solidity: function Items(string ) view returns(uint256)
func (_Basic *BasicCaller) Items(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "Items", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x4547a6b3.
//
// Solidity: function Items(string ) view returns(uint256)
func (_Basic *BasicSession) Items(arg0 string) (*big.Int, error) {
	return _Basic.Contract.Items(&_Basic.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x4547a6b3.
//
// Solidity: function Items(string ) view returns(uint256)
func (_Basic *BasicCallerSession) Items(arg0 string) (*big.Int, error) {
	return _Basic.Contract.Items(&_Basic.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Basic *BasicCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Basic *BasicSession) Version() (string, error) {
	return _Basic.Contract.Version(&_Basic.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Basic *BasicCallerSession) Version() (string, error) {
	return _Basic.Contract.Version(&_Basic.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0x8c5cf3ed.
//
// Solidity: function SetItem(string key, uint256 value) returns()
func (_Basic *BasicTransactor) SetItem(opts *bind.TransactOpts, key string, value *big.Int) (*types.Transaction, error) {
	return _Basic.contract.Transact(opts, "SetItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0x8c5cf3ed.
//
// Solidity: function SetItem(string key, uint256 value) returns()
func (_Basic *BasicSession) SetItem(key string, value *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.SetItem(&_Basic.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0x8c5cf3ed.
//
// Solidity: function SetItem(string key, uint256 value) returns()
func (_Basic *BasicTransactorSession) SetItem(key string, value *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.SetItem(&_Basic.TransactOpts, key, value)
}

// BasicItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Basic contract.
type BasicItemSetIterator struct {
	Event *BasicItemSet // Event containing the contract specifics and raw log

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
func (it *BasicItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicItemSet)
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
		it.Event = new(BasicItemSet)
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
func (it *BasicItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicItemSet represents a ItemSet event raised by the Basic contract.
type BasicItemSet struct {
	Key   string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0x814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f.
//
// Solidity: event ItemSet(string key, uint256 value)
func (_Basic *BasicFilterer) FilterItemSet(opts *bind.FilterOpts) (*BasicItemSetIterator, error) {

	logs, sub, err := _Basic.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &BasicItemSetIterator{contract: _Basic.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0x814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f.
//
// Solidity: event ItemSet(string key, uint256 value)
func (_Basic *BasicFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *BasicItemSet) (event.Subscription, error) {

	logs, sub, err := _Basic.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicItemSet)
				if err := _Basic.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

// ParseItemSet is a log parse operation binding the contract event 0x814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f.
//
// Solidity: event ItemSet(string key, uint256 value)
func (_Basic *BasicFilterer) ParseItemSet(log types.Log) (*BasicItemSet, error) {
	event := new(BasicItemSet)
	if err := _Basic.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
