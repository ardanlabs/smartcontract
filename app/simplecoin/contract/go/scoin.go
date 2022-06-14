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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"EventFrozenAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"EventTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizedAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowanceAmount\",\"type\":\"uint256\"}],\"name\":\"Authorize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CoinBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"freeze\",\"type\":\"bool\"}],\"name\":\"FreezeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FrozenAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200192338038062001923833981810160405281019062000037919062000232565b33600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000ac600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682620000b360201b60201c565b506200031e565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146200010e57600080fd5b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546200015e919062000293565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a4583604051620001e6919062000301565b60405180910390a35050565b600080fd5b6000819050919050565b6200020c81620001f7565b81146200021857600080fd5b50565b6000815190506200022c8162000201565b92915050565b6000602082840312156200024b576200024a620001f2565b5b60006200025b848285016200021b565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620002a082620001f7565b9150620002ad83620001f7565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620002e557620002e462000264565b5b828201905092915050565b620002fb81620001f7565b82525050565b6000602082019050620003186000830184620002f0565b92915050565b6115f5806200032e6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638bdf1534116100665780638bdf153414610130578063ab31471e14610160578063b4a99a4e14610190578063c0d84ce5146101ae578063d16a7a4b146101ca57610093565b80630f6798a5146100985780636386a433146100b457806369ca02dd146100e45780636dcdd00f14610100575b600080fd5b6100b260048036038101906100ad9190610f12565b6101e6565b005b6100ce60048036038101906100c99190610f12565b610320565b6040516100db9190610f6d565b60405180910390f35b6100fe60048036038101906100f99190610f12565b6103ad565b005b61011a60048036038101906101159190610f88565b610585565b6040516101279190610fd7565b60405180910390f35b61014a60048036038101906101459190610ff2565b6105aa565b6040516101579190610fd7565b60405180910390f35b61017a60048036038101906101759190610ff2565b6105c2565b6040516101879190610f6d565b60405180910390f35b6101986105e2565b6040516101a5919061102e565b60405180910390f35b6101c860048036038101906101c39190611049565b610608565b005b6101e460048036038101906101df91906110c8565b610874565b005b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461024057600080fd5b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461028e9190611137565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45836040516103149190610fd7565b60405180910390a35050565b600081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001905092915050565b60006103ba338484610962565b90508060000151156104075780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103fe9190611226565b60405180910390fd5b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a60405161043490611294565b60405180910390a1816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461048a91906112b4565b92505081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104df9190611137565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a60405161051390611334565b60405180910390a18273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45846040516105789190610fd7565b60405180910390a3505050565b6001602052816000526040600020602052806000526040600020600091509150505481565b60006020528060005260406000206000915090505481565b60026020528060005260406000206000915054906101000a900460ff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000610615848484610962565b90508060000151156106625780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106599190611226565b60405180910390fd5b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a60405161068f90611294565b60405180910390a1816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106e591906112b4565b92505081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461073a9190611137565b9250508190555081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107cd91906112b4565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a60405161080190611334565b60405180910390a18273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45846040516108669190610fd7565b60405180910390a350505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108ce57600080fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d8282604051610956929190611354565b60405180910390a15050565b61096a610e5d565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036109e3576109dc6040518060400160405280601f81526020017f63616e27742073656e64206d6f6e657920746f20616464726573732030783000815250610be2565b9050610bdb565b816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610af457610aed6040518060400160405280601781526020017f696e737566666963656e742066756e6473202062616c3a000000000000000000815250610aa96000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c09565b6040518060400160405280600a81526020017f2020616d6f756e743a2000000000000000000000000000000000000000000000815250610ae886610c09565b610d91565b9050610bdb565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610b7d9190611137565b11610bd057610bc96040518060400160405280601081526020017f696e76616c696420616d6f756e743a2000000000000000000000000000000000815250610bc484610c09565b610de0565b9050610bdb565b610bd8610e29565b90505b9392505050565b610bea610e5d565b6040518060400160405280600115158152602001838152509050919050565b606060008203610c50576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610d8c565b600082905060005b60008214610c82578080610c6b9061137d565b915050600a82610c7b91906113f4565b9150610c58565b60008167ffffffffffffffff811115610c9e57610c9d611425565b5b6040519080825280601f01601f191660200182016040528015610cd05781602001600182028036833780820191505090505b50905060008290505b60008614610d8457600181610cee91906112b4565b90506000600a8088610d0091906113f4565b610d0a9190611454565b87610d1591906112b4565b6030610d2191906114bb565b905060008160f81b905080848481518110610d3f57610d3e6114f2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610d7b91906113f4565b97505050610cd9565b819450505050505b919050565b610d99610e5d565b604051806040016040528060011515815260200186868686604051602001610dc4949392919061155d565b6040516020818303038152906040528152509050949350505050565b610de8610e5d565b60405180604001604052806001151581526020018484604051602001610e0f92919061159b565b604051602081830303815290604052815250905092915050565b610e31610e5d565b604051806040016040528060001515815260200160405180602001604052806000815250815250905090565b6040518060400160405280600015158152602001606081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ea982610e7e565b9050919050565b610eb981610e9e565b8114610ec457600080fd5b50565b600081359050610ed681610eb0565b92915050565b6000819050919050565b610eef81610edc565b8114610efa57600080fd5b50565b600081359050610f0c81610ee6565b92915050565b60008060408385031215610f2957610f28610e79565b5b6000610f3785828601610ec7565b9250506020610f4885828601610efd565b9150509250929050565b60008115159050919050565b610f6781610f52565b82525050565b6000602082019050610f826000830184610f5e565b92915050565b60008060408385031215610f9f57610f9e610e79565b5b6000610fad85828601610ec7565b9250506020610fbe85828601610ec7565b9150509250929050565b610fd181610edc565b82525050565b6000602082019050610fec6000830184610fc8565b92915050565b60006020828403121561100857611007610e79565b5b600061101684828501610ec7565b91505092915050565b61102881610e9e565b82525050565b6000602082019050611043600083018461101f565b92915050565b60008060006060848603121561106257611061610e79565b5b600061107086828701610ec7565b935050602061108186828701610ec7565b925050604061109286828701610efd565b9150509250925092565b6110a581610f52565b81146110b057600080fd5b50565b6000813590506110c28161109c565b92915050565b600080604083850312156110df576110de610e79565b5b60006110ed85828601610ec7565b92505060206110fe858286016110b3565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061114282610edc565b915061114d83610edc565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561118257611181611108565b5b828201905092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156111c75780820151818401526020810190506111ac565b838111156111d6576000848401525b50505050565b6000601f19601f8301169050919050565b60006111f88261118d565b6112028185611198565b93506112128185602086016111a9565b61121b816111dc565b840191505092915050565b6000602082019050818103600083015261124081846111ed565b905092915050565b7f7374617274696e67207472616e73666572000000000000000000000000000000600082015250565b600061127e601183611198565b915061128982611248565b602082019050919050565b600060208201905081810360008301526112ad81611271565b9050919050565b60006112bf82610edc565b91506112ca83610edc565b9250828210156112dd576112dc611108565b5b828203905092915050565b7f656e64696e67207472616e736665720000000000000000000000000000000000600082015250565b600061131e600f83611198565b9150611329826112e8565b602082019050919050565b6000602082019050818103600083015261134d81611311565b9050919050565b6000604082019050611369600083018561101f565b6113766020830184610f5e565b9392505050565b600061138882610edc565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036113ba576113b9611108565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006113ff82610edc565b915061140a83610edc565b92508261141a576114196113c5565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600061145f82610edc565b915061146a83610edc565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156114a3576114a2611108565b5b828202905092915050565b600060ff82169050919050565b60006114c6826114ae565b91506114d1836114ae565b92508260ff038211156114e7576114e6611108565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081905092915050565b60006115378261118d565b6115418185611521565b93506115518185602086016111a9565b80840191505092915050565b6000611569828761152c565b9150611575828661152c565b9150611581828561152c565b915061158d828461152c565b915081905095945050505050565b60006115a7828561152c565b91506115b3828461152c565b9150819050939250505056fea26469706673582212204b2df0c5f414f89f6c8dacc9bb6dcd8dea1c4d70c03661a7107f5e7e1cb5ec8164736f6c634300080e0033",
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

// Allowance is a free data retrieval call binding the contract method 0x6dcdd00f.
//
// Solidity: function Allowance(address , address ) view returns(uint256)
func (_Scoin *ScoinCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scoin.contract.Call(opts, &out, "Allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x6dcdd00f.
//
// Solidity: function Allowance(address , address ) view returns(uint256)
func (_Scoin *ScoinSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Scoin.Contract.Allowance(&_Scoin.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0x6dcdd00f.
//
// Solidity: function Allowance(address , address ) view returns(uint256)
func (_Scoin *ScoinCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Scoin.Contract.Allowance(&_Scoin.CallOpts, arg0, arg1)
}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_Scoin *ScoinCaller) CoinBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Scoin.contract.Call(opts, &out, "CoinBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_Scoin *ScoinSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _Scoin.Contract.CoinBalance(&_Scoin.CallOpts, arg0)
}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_Scoin *ScoinCallerSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _Scoin.Contract.CoinBalance(&_Scoin.CallOpts, arg0)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_Scoin *ScoinCaller) FrozenAccount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Scoin.contract.Call(opts, &out, "FrozenAccount", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_Scoin *ScoinSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _Scoin.Contract.FrozenAccount(&_Scoin.CallOpts, arg0)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_Scoin *ScoinCallerSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _Scoin.Contract.FrozenAccount(&_Scoin.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Scoin *ScoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Scoin.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Scoin *ScoinSession) Owner() (common.Address, error) {
	return _Scoin.Contract.Owner(&_Scoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Scoin *ScoinCallerSession) Owner() (common.Address, error) {
	return _Scoin.Contract.Owner(&_Scoin.CallOpts)
}

// Authorize is a paid mutator transaction binding the contract method 0x6386a433.
//
// Solidity: function Authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_Scoin *ScoinTransactor) Authorize(opts *bind.TransactOpts, authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _Scoin.contract.Transact(opts, "Authorize", authorizedAccount, allowanceAmount)
}

// Authorize is a paid mutator transaction binding the contract method 0x6386a433.
//
// Solidity: function Authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_Scoin *ScoinSession) Authorize(authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Authorize(&_Scoin.TransactOpts, authorizedAccount, allowanceAmount)
}

// Authorize is a paid mutator transaction binding the contract method 0x6386a433.
//
// Solidity: function Authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_Scoin *ScoinTransactorSession) Authorize(authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Authorize(&_Scoin.TransactOpts, authorizedAccount, allowanceAmount)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_Scoin *ScoinTransactor) FreezeAccount(opts *bind.TransactOpts, target common.Address, freeze bool) (*types.Transaction, error) {
	return _Scoin.contract.Transact(opts, "FreezeAccount", target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_Scoin *ScoinSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _Scoin.Contract.FreezeAccount(&_Scoin.TransactOpts, target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_Scoin *ScoinTransactorSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _Scoin.Contract.FreezeAccount(&_Scoin.TransactOpts, target, freeze)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_Scoin *ScoinTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Scoin.contract.Transact(opts, "Mint", recipient, mintedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_Scoin *ScoinSession) Mint(recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Mint(&_Scoin.TransactOpts, recipient, mintedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_Scoin *ScoinTransactorSession) Mint(recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Mint(&_Scoin.TransactOpts, recipient, mintedAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_Scoin *ScoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.contract.Transact(opts, "Transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_Scoin *ScoinSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Transfer(&_Scoin.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_Scoin *ScoinTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.Transfer(&_Scoin.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xc0d84ce5.
//
// Solidity: function TransferFrom(address from, address to, uint256 amount) returns()
func (_Scoin *ScoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.contract.Transact(opts, "TransferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xc0d84ce5.
//
// Solidity: function TransferFrom(address from, address to, uint256 amount) returns()
func (_Scoin *ScoinSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.TransferFrom(&_Scoin.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xc0d84ce5.
//
// Solidity: function TransferFrom(address from, address to, uint256 amount) returns()
func (_Scoin *ScoinTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Scoin.Contract.TransferFrom(&_Scoin.TransactOpts, from, to, amount)
}

// ScoinEventFrozenAccountIterator is returned from FilterEventFrozenAccount and is used to iterate over the raw logs and unpacked data for EventFrozenAccount events raised by the Scoin contract.
type ScoinEventFrozenAccountIterator struct {
	Event *ScoinEventFrozenAccount // Event containing the contract specifics and raw log

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
func (it *ScoinEventFrozenAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoinEventFrozenAccount)
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
		it.Event = new(ScoinEventFrozenAccount)
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
func (it *ScoinEventFrozenAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoinEventFrozenAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoinEventFrozenAccount represents a EventFrozenAccount event raised by the Scoin contract.
type ScoinEventFrozenAccount struct {
	Target common.Address
	Frozen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventFrozenAccount is a free log retrieval operation binding the contract event 0x3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d.
//
// Solidity: event EventFrozenAccount(address target, bool frozen)
func (_Scoin *ScoinFilterer) FilterEventFrozenAccount(opts *bind.FilterOpts) (*ScoinEventFrozenAccountIterator, error) {

	logs, sub, err := _Scoin.contract.FilterLogs(opts, "EventFrozenAccount")
	if err != nil {
		return nil, err
	}
	return &ScoinEventFrozenAccountIterator{contract: _Scoin.contract, event: "EventFrozenAccount", logs: logs, sub: sub}, nil
}

// WatchEventFrozenAccount is a free log subscription operation binding the contract event 0x3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d.
//
// Solidity: event EventFrozenAccount(address target, bool frozen)
func (_Scoin *ScoinFilterer) WatchEventFrozenAccount(opts *bind.WatchOpts, sink chan<- *ScoinEventFrozenAccount) (event.Subscription, error) {

	logs, sub, err := _Scoin.contract.WatchLogs(opts, "EventFrozenAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoinEventFrozenAccount)
				if err := _Scoin.contract.UnpackLog(event, "EventFrozenAccount", log); err != nil {
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
func (_Scoin *ScoinFilterer) ParseEventFrozenAccount(log types.Log) (*ScoinEventFrozenAccount, error) {
	event := new(ScoinEventFrozenAccount)
	if err := _Scoin.contract.UnpackLog(event, "EventFrozenAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScoinEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Scoin contract.
type ScoinEventLogIterator struct {
	Event *ScoinEventLog // Event containing the contract specifics and raw log

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
func (it *ScoinEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoinEventLog)
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
		it.Event = new(ScoinEventLog)
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
func (it *ScoinEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoinEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoinEventLog represents a EventLog event raised by the Scoin contract.
type ScoinEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Scoin *ScoinFilterer) FilterEventLog(opts *bind.FilterOpts) (*ScoinEventLogIterator, error) {

	logs, sub, err := _Scoin.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &ScoinEventLogIterator{contract: _Scoin.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Scoin *ScoinFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *ScoinEventLog) (event.Subscription, error) {

	logs, sub, err := _Scoin.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoinEventLog)
				if err := _Scoin.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_Scoin *ScoinFilterer) ParseEventLog(log types.Log) (*ScoinEventLog, error) {
	event := new(ScoinEventLog)
	if err := _Scoin.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScoinEventTransferIterator is returned from FilterEventTransfer and is used to iterate over the raw logs and unpacked data for EventTransfer events raised by the Scoin contract.
type ScoinEventTransferIterator struct {
	Event *ScoinEventTransfer // Event containing the contract specifics and raw log

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
func (it *ScoinEventTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoinEventTransfer)
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
		it.Event = new(ScoinEventTransfer)
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
func (it *ScoinEventTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoinEventTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoinEventTransfer represents a EventTransfer event raised by the Scoin contract.
type ScoinEventTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventTransfer is a free log retrieval operation binding the contract event 0x64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45.
//
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 value)
func (_Scoin *ScoinFilterer) FilterEventTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScoinEventTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Scoin.contract.FilterLogs(opts, "EventTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScoinEventTransferIterator{contract: _Scoin.contract, event: "EventTransfer", logs: logs, sub: sub}, nil
}

// WatchEventTransfer is a free log subscription operation binding the contract event 0x64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45.
//
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 value)
func (_Scoin *ScoinFilterer) WatchEventTransfer(opts *bind.WatchOpts, sink chan<- *ScoinEventTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Scoin.contract.WatchLogs(opts, "EventTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoinEventTransfer)
				if err := _Scoin.contract.UnpackLog(event, "EventTransfer", log); err != nil {
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
func (_Scoin *ScoinFilterer) ParseEventTransfer(log types.Log) (*ScoinEventTransfer, error) {
	event := new(ScoinEventTransfer)
	if err := _Scoin.contract.UnpackLog(event, "EventTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
