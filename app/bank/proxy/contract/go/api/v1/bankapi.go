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
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506110fc806100606000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80637d7b009914610046578063b4a99a4e14610064578063fa84fd8e14610082575b600080fd5b61004e61009e565b60405161005b9190610825565b60405180910390f35b61006c6100c4565b6040516100799190610825565b60405180910390f35b61009c60048036038101906100979190610a0f565b6100e8565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600082905060005b845181101561037857836002600087848151811061011157610110610a92565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156102e6577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6101dd6002600088858151811061019557610194610a92565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461065c565b6101e68661065c565b6040516020016101f7929190610b7e565b6040516020818303038152906040526040516102139190610c0a565b60405180910390a16002600086838151811061023257610231610a92565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826102809190610c5b565b915060006002600087848151811061029b5761029a610a92565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610365565b83826102f29190610c5b565b9150836002600087848151811061030c5761030b610a92565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461035d9190610c8f565b925050819055505b808061037090610cc3565b9150506100f0565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103a38461065c565b6103ac8461065c565b6103b58461065c565b6040516020016103c793929190610da3565b6040516020818303038152906040526040516103e39190610c0a565b60405180910390a16000810361042e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042590610e82565b60405180910390fd5b818110156105115780600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104a69190610c5b565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6104d78261065c565b6040516020016104e79190610f14565b6040516020818303038152906040526040516105039190610c0a565b60405180910390a150610656565b818161051d9190610c8f565b905080600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461056e9190610c5b565b9250508190555081600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105e59190610c5b565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6106168261065c565b61061f8461065c565b604051602001610630929190610f91565b60405160208183030381529060405260405161064c9190610c0a565b60405180910390a1505b50505050565b6060600082036106a3576040518060400160405280600181526020017f300000000000000000000000000000000000000000000000000000000000000081525090506107df565b600082905060005b600082146106d55780806106be90610cc3565b915050600a826106ce9190611011565b91506106ab565b60008167ffffffffffffffff8111156106f1576106f0610896565b5b6040519080825280601f01601f1916602001820160405280156107235781602001600182028036833780820191505090505b50905060008290505b600086146107d7576001816107419190610c8f565b90506000600a80886107539190611011565b61075d9190611042565b876107689190610c8f565b60306107749190611091565b905060008160f81b90508084848151811061079257610791610a92565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a886107ce9190611011565b9750505061072c565b819450505050505b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061080f826107e4565b9050919050565b61081f81610804565b82525050565b600060208201905061083a6000830184610816565b92915050565b6000604051905090565b600080fd5b600080fd5b61085d81610804565b811461086857600080fd5b50565b60008135905061087a81610854565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6108ce82610885565b810181811067ffffffffffffffff821117156108ed576108ec610896565b5b80604052505050565b6000610900610840565b905061090c82826108c5565b919050565b600067ffffffffffffffff82111561092c5761092b610896565b5b602082029050602081019050919050565b600080fd5b600061095561095084610911565b6108f6565b905080838252602082019050602084028301858111156109785761097761093d565b5b835b818110156109a1578061098d888261086b565b84526020840193505060208101905061097a565b5050509392505050565b600082601f8301126109c0576109bf610880565b5b81356109d0848260208601610942565b91505092915050565b6000819050919050565b6109ec816109d9565b81146109f757600080fd5b50565b600081359050610a09816109e3565b92915050565b60008060008060808587031215610a2957610a2861084a565b5b6000610a378782880161086b565b945050602085013567ffffffffffffffff811115610a5857610a5761084f565b5b610a64878288016109ab565b9350506040610a75878288016109fa565b9250506060610a86878288016109fa565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f6163636f756e742062616c616e63652000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b83811015610b1b578082015181840152602081019050610b00565b60008484015250505050565b6000610b3282610ae7565b610b3c8185610af2565b9350610b4c818560208601610afd565b80840191505092915050565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b6000610b8982610ac1565b601082019150610b998285610b27565b9150610ba482610b58565b601382019150610bb48284610b27565b91508190509392505050565b600082825260208201905092915050565b6000610bdc82610ae7565b610be68185610bc0565b9350610bf6818560208601610afd565b610bff81610885565b840191505092915050565b60006020820190508181036000830152610c248184610bd1565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c66826109d9565b9150610c71836109d9565b9250828201905080821115610c8957610c88610c2c565b5b92915050565b6000610c9a826109d9565b9150610ca5836109d9565b9250828203905081811115610cbd57610cbc610c2c565b5b92915050565b6000610cce826109d9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d0057610cff610c2c565b5b600182019050919050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b6000610dae82610d0b565b600582019150610dbe8286610b27565b9150610dc982610d31565b600a82019150610dd98285610b27565b9150610de482610d57565b600682019150610df48284610b27565b9150610dff82610d7d565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e206163636f60008201527f756e742062616c616e6365730000000000000000000000000000000000000000602082015250565b6000610e6c602c83610bc0565b9150610e7782610e10565b604082019050919050565b60006020820190508181036000830152610e9b81610e5f565b9050919050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e60008201527f65725b0000000000000000000000000000000000000000000000000000000000602082015250565b6000610efe602383610af2565b9150610f0982610ea2565b602382019050919050565b6000610f1f82610ef1565b9150610f2b8284610b27565b9150610f3682610d7d565b60018201915081905092915050565b7f77696e6e65725b00000000000000000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b6000610f9c82610f45565b600782019150610fac8285610b27565b9150610fb782610f6b565b600882019150610fc78284610b27565b9150610fd282610d7d565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061101c826109d9565b9150611027836109d9565b92508261103757611036610fe2565b5b828204905092915050565b600061104d826109d9565b9150611058836109d9565b9250828202611066816109d9565b9150828204841483151761107d5761107c610c2c565b5b5092915050565b600060ff82169050919050565b600061109c82611084565b91506110a783611084565b9250828201905060ff8111156110c0576110bf610c2c565b5b9291505056fea2646970667358221220e9bcc21e063d7bdff810643d35074fc664cd9c43e641e61aaf2eb7ddba39e80e64736f6c63430008110033",
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
