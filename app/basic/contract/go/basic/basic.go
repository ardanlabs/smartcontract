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
	Bin: "0x608060405234801561000f575f80fd5b506040518060400160405280600381526020017f312e3100000000000000000000000000000000000000000000000000000000008152505f90816100539190610293565b50610362565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100d457607f821691505b6020821081036100e7576100e6610090565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101497fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261010e565b610153868361010e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61019761019261018d8461016b565b610174565b61016b565b9050919050565b5f819050919050565b6101b08361017d565b6101c46101bc8261019e565b84845461011a565b825550505050565b5f90565b6101d86101cc565b6101e38184846101a7565b505050565b5b81811015610206576101fb5f826101d0565b6001810190506101e9565b5050565b601f82111561024b5761021c816100ed565b610225846100ff565b81016020851015610234578190505b610248610240856100ff565b8301826101e8565b50505b505050565b5f82821c905092915050565b5f61026b5f1984600802610250565b1980831691505092915050565b5f610283838361025c565b9150826002028217905092915050565b61029c82610059565b67ffffffffffffffff8111156102b5576102b4610063565b5b6102bf82546100bd565b6102ca82828561020a565b5f60209050601f8311600181146102fb575f84156102e9578287015190505b6102f38582610278565b86555061035a565b601f198416610309866100ed565b5f5b828110156103305784890151825560018201915060208501945060208101905061030b565b8683101561034d5784890151610349601f89168261025c565b8355505b6001600288020188555050505b505050505050565b6105b98061036f5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80634547a6b3146100435780638c5cf3ed14610073578063bb62860d1461008f575b5f80fd5b61005d60048036038101906100589190610312565b6100ad565b60405161006a9190610371565b60405180910390f35b61008d600480360381019061008891906103b4565b6100da565b005b61009761013a565b6040516100a49190610488565b60405180910390f35b6001818051602081018201805184825260208301602085012081835280955050505050505f915090505481565b806001836040516100eb91906104e2565b9081526020016040518091039020819055507f814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f828260405161012e9291906104f8565b60405180910390a15050565b5f805461014690610553565b80601f016020809104026020016040519081016040528092919081815260200182805461017290610553565b80156101bd5780601f10610194576101008083540402835291602001916101bd565b820191905f5260205f20905b8154815290600101906020018083116101a057829003601f168201915b505050505081565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610224826101de565b810181811067ffffffffffffffff82111715610243576102426101ee565b5b80604052505050565b5f6102556101c5565b9050610261828261021b565b919050565b5f67ffffffffffffffff8211156102805761027f6101ee565b5b610289826101de565b9050602081019050919050565b828183375f83830152505050565b5f6102b66102b184610266565b61024c565b9050828152602081018484840111156102d2576102d16101da565b5b6102dd848285610296565b509392505050565b5f82601f8301126102f9576102f86101d6565b5b81356103098482602086016102a4565b91505092915050565b5f60208284031215610327576103266101ce565b5b5f82013567ffffffffffffffff811115610344576103436101d2565b5b610350848285016102e5565b91505092915050565b5f819050919050565b61036b81610359565b82525050565b5f6020820190506103845f830184610362565b92915050565b61039381610359565b811461039d575f80fd5b50565b5f813590506103ae8161038a565b92915050565b5f80604083850312156103ca576103c96101ce565b5b5f83013567ffffffffffffffff8111156103e7576103e66101d2565b5b6103f3858286016102e5565b9250506020610404858286016103a0565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561044557808201518184015260208101905061042a565b5f8484015250505050565b5f61045a8261040e565b6104648185610418565b9350610474818560208601610428565b61047d816101de565b840191505092915050565b5f6020820190508181035f8301526104a08184610450565b905092915050565b5f81905092915050565b5f6104bc8261040e565b6104c681856104a8565b93506104d6818560208601610428565b80840191505092915050565b5f6104ed82846104b2565b915081905092915050565b5f6040820190508181035f8301526105108185610450565b905061051f6020830184610362565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061056a57607f821691505b60208210810361057d5761057c610526565b5b5091905056fea2646970667358221220dbd32c39f58d9a89ee20dac3738aed1b163c90921e6e60ee3881dd447950f1e064736f6c63430008190033",
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
