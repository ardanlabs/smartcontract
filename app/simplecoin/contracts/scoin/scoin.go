// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scoin

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

// ScoinMetaData contains all meta data concerning the Scoin contract.
var ScoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"coinBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610d56380380610d56833981810160405281019061003291906100b6565b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550506100e3565b600080fd5b6000819050919050565b61009381610080565b811461009e57600080fd5b50565b6000815190506100b08161008a565b92915050565b6000602082840312156100cc576100cb61007b565b5b60006100da848285016100a1565b91505092915050565b610c64806100f26000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063a9059cbb1461003b578063fabde80c14610057575b600080fd5b6100556004803603810190610050919061062b565b610087565b005b610071600480360381019061006c919061066b565b6103f2565b60405161007e91906106a7565b60405180910390f35b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561017e5760006101166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461040a565b61011f8361040a565b6040516020016101309291906107d4565b6040516020818303038152906040529050806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101759190610869565b60405180910390fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461020791906108ba565b106102755760006102178261040a565b604051602001610227919061095c565b6040516020818303038152906040529050806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026c9190610869565b60405180910390fd5b7fcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab6040516102a2906109ca565b60405180910390a1806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102f891906109ea565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461034d91906108ba565b925050819055507fcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab60405161038190610a6a565b60405180910390a18173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516103e691906106a7565b60405180910390a35050565b60006020528060005260406000206000915090505481565b606060008203610451576040518060400160405280600181526020017f3000000000000000000000000000000000000000000000000000000000000000815250905061058d565b600082905060005b6000821461048357808061046c90610a8a565b915050600a8261047c9190610b01565b9150610459565b60008167ffffffffffffffff81111561049f5761049e610b32565b5b6040519080825280601f01601f1916602001820160405280156104d15781602001600182028036833780820191505090505b50905060008290505b60008614610585576001816104ef91906109ea565b90506000600a80886105019190610b01565b61050b9190610b61565b8761051691906109ea565b60306105229190610bc8565b905060008160f81b9050808484815181106105405761053f610bff565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a8861057c9190610b01565b975050506104da565b819450505050505b919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105c282610597565b9050919050565b6105d2816105b7565b81146105dd57600080fd5b50565b6000813590506105ef816105c9565b92915050565b6000819050919050565b610608816105f5565b811461061357600080fd5b50565b600081359050610625816105ff565b92915050565b6000806040838503121561064257610641610592565b5b6000610650858286016105e0565b925050602061066185828601610616565b9150509250929050565b60006020828403121561068157610680610592565b5b600061068f848285016105e0565b91505092915050565b6106a1816105f5565b82525050565b60006020820190506106bc6000830184610698565b92915050565b600081905092915050565b7f696e737566666963656e742066756e6473202062616c3a000000000000000000600082015250565b60006107036017836106c2565b915061070e826106cd565b601782019050919050565b600081519050919050565b60005b83811015610742578082015181840152602081019050610727565b83811115610751576000848401525b50505050565b600061076282610719565b61076c81856106c2565b935061077c818560208601610724565b80840191505092915050565b7f2020616d6f756e743a2000000000000000000000000000000000000000000000600082015250565b60006107be600a836106c2565b91506107c982610788565b600a82019050919050565b60006107df826106f6565b91506107eb8285610757565b91506107f6826107b1565b91506108028284610757565b91508190509392505050565b600082825260208201905092915050565b6000601f19601f8301169050919050565b600061083b82610719565b610845818561080e565b9350610855818560208601610724565b61085e8161081f565b840191505092915050565b600060208201905081810360008301526108838184610830565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006108c5826105f5565b91506108d0836105f5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156109055761090461088b565b5b828201905092915050565b7f61726974686d65746963206f766572666c6f772020616d6f756e743a20000000600082015250565b6000610946601d836106c2565b915061095182610910565b601d82019050919050565b600061096782610939565b91506109738284610757565b915081905092915050565b7f7374617274696e67207472616e73666572000000000000000000000000000000600082015250565b60006109b460118361080e565b91506109bf8261097e565b602082019050919050565b600060208201905081810360008301526109e3816109a7565b9050919050565b60006109f5826105f5565b9150610a00836105f5565b925082821015610a1357610a1261088b565b5b828203905092915050565b7f656e64696e67207472616e736665720000000000000000000000000000000000600082015250565b6000610a54600f8361080e565b9150610a5f82610a1e565b602082019050919050565b60006020820190508181036000830152610a8381610a47565b9050919050565b6000610a95826105f5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ac757610ac661088b565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610b0c826105f5565b9150610b17836105f5565b925082610b2757610b26610ad2565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000610b6c826105f5565b9150610b77836105f5565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610bb057610baf61088b565b5b828202905092915050565b600060ff82169050919050565b6000610bd382610bbb565b9150610bde83610bbb565b92508260ff03821115610bf457610bf361088b565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea26469706673582212201f8898128a1a120b0b232f80ce62d9d9b30a2c2c6eb20d39adb50bb6bc6ed18064736f6c634300080e0033",
}

// ScoinABI is the input ABI used to generate the binding from.
// Deprecated: Use ScoinMetaData.ABI instead.
var ScoinABI = ScoinMetaData.ABI

// ScoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ScoinMetaData.Bin instead.
var ScoinBin = ScoinMetaData.Bin

// DeployScoin deploys a new Ethereum contract, binding an instance of Scoin to it.
func DeployScoin(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *Scoin, error) {
	parsed, err := ScoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScoinBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Scoin{ScoinCaller: ScoinCaller{contract: contract}, ScoinTransactor: ScoinTransactor{contract: contract}, ScoinFilterer: ScoinFilterer{contract: contract}}, nil
}

// Scoin is an auto generated Go binding around an Ethereum contract.
type Scoin struct {
	ScoinCaller     // Read-only binding to the contract
	ScoinTransactor // Write-only binding to the contract
	ScoinFilterer   // Log filterer for contract events
}

// ScoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScoinSession struct {
	Contract     *Scoin            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScoinCallerSession struct {
	Contract *ScoinCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ScoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScoinTransactorSession struct {
	Contract     *ScoinTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScoinRaw struct {
	Contract *Scoin // Generic contract binding to access the raw methods on
}

// ScoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScoinCallerRaw struct {
	Contract *ScoinCaller // Generic read-only contract binding to access the raw methods on
}

// ScoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScoinTransactorRaw struct {
	Contract *ScoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScoin creates a new instance of Scoin, bound to a specific deployed contract.
func NewScoin(address common.Address, backend bind.ContractBackend) (*Scoin, error) {
	contract, err := bindScoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Scoin{ScoinCaller: ScoinCaller{contract: contract}, ScoinTransactor: ScoinTransactor{contract: contract}, ScoinFilterer: ScoinFilterer{contract: contract}}, nil
}

// NewScoinCaller creates a new read-only instance of Scoin, bound to a specific deployed contract.
func NewScoinCaller(address common.Address, caller bind.ContractCaller) (*ScoinCaller, error) {
	contract, err := bindScoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScoinCaller{contract: contract}, nil
}

// NewScoinTransactor creates a new write-only instance of Scoin, bound to a specific deployed contract.
func NewScoinTransactor(address common.Address, transactor bind.ContractTransactor) (*ScoinTransactor, error) {
	contract, err := bindScoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScoinTransactor{contract: contract}, nil
}

// NewScoinFilterer creates a new log filterer instance of Scoin, bound to a specific deployed contract.
func NewScoinFilterer(address common.Address, filterer bind.ContractFilterer) (*ScoinFilterer, error) {
	contract, err := bindScoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScoinFilterer{contract: contract}, nil
}

// bindScoin binds a generic wrapper to an already deployed contract.
func bindScoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scoin *ScoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scoin.Contract.ScoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scoin *ScoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scoin.Contract.ScoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scoin *ScoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scoin.Contract.ScoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scoin *ScoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scoin *ScoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scoin *ScoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scoin.Contract.contract.Transact(opts, method, params...)
}

// CoinBalance is a free data retrieval call binding the contract method 0xfabde80c.
//
// Solidity: function coinBalance(address ) view returns(uint256)
func (_Scoin *ScoinCaller) CoinBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scoin.contract.Call(opts, &out, "coinBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinBalance is a free data retrieval call binding the contract method 0xfabde80c.
//
// Solidity: function coinBalance(address ) view returns(uint256)
func (_Scoin *ScoinSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _Scoin.Contract.CoinBalance(&_Scoin.CallOpts, arg0)
}

// CoinBalance is a free data retrieval call binding the contract method 0xfabde80c.
//
// Solidity: function coinBalance(address ) view returns(uint256)
func (_Scoin *ScoinCallerSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _Scoin.Contract.CoinBalance(&_Scoin.CallOpts, arg0)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Scoin *ScoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Scoin *ScoinSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Transfer(&_Scoin.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns()
func (_Scoin *ScoinTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Transfer(&_Scoin.TransactOpts, to, amount)
}

// ScoinLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the Scoin contract.
type ScoinLogIterator struct {
	Event *ScoinLog // Event containing the contract specifics and raw log

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
func (it *ScoinLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoinLog)
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
		it.Event = new(ScoinLog)
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
func (it *ScoinLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoinLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoinLog represents a Log event raised by the Scoin contract.
type ScoinLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string value)
func (_Scoin *ScoinFilterer) FilterLog(opts *bind.FilterOpts) (*ScoinLogIterator, error) {

	logs, sub, err := _Scoin.contract.FilterLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return &ScoinLogIterator{contract: _Scoin.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string value)
func (_Scoin *ScoinFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *ScoinLog) (event.Subscription, error) {

	logs, sub, err := _Scoin.contract.WatchLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoinLog)
				if err := _Scoin.contract.UnpackLog(event, "Log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0xcf34ef537ac33ee1ac626ca1587a0a7e8e51561e5514f8cb36afa1c5102b3bab.
//
// Solidity: event Log(string value)
func (_Scoin *ScoinFilterer) ParseLog(log types.Log) (*ScoinLog, error) {
	event := new(ScoinLog)
	if err := _Scoin.contract.UnpackLog(event, "Log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Scoin contract.
type ScoinTransferIterator struct {
	Event *ScoinTransfer // Event containing the contract specifics and raw log

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
func (it *ScoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoinTransfer)
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
		it.Event = new(ScoinTransfer)
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
func (it *ScoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoinTransfer represents a Transfer event raised by the Scoin contract.
type ScoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scoin *ScoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Scoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScoinTransferIterator{contract: _Scoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scoin *ScoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ScoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Scoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoinTransfer)
				if err := _Scoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Scoin *ScoinFilterer) ParseTransfer(log types.Log) (*ScoinTransfer, error) {
	event := new(ScoinTransfer)
	if err := _Scoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
