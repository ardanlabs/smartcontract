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
	_ = abi.ConvertType
)

// SimplecoinMetaData contains all meta data concerning the Simplecoin contract.
var SimplecoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"EventFrozenAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EventTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CoinBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"freeze\",\"type\":\"bool\"}],\"name\":\"FreezeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FrozenAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001a2138038062001a21833981810160405281019062000037919062000235565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000b260008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682620000b9640100000000026401000000009004565b50620002ff565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146200011257600080fd5b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825462000163919062000296565b925050819055508173ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a4583604051620001e99190620002e2565b60405180910390a35050565b600080fd5b6000819050919050565b6200020f81620001fa565b81146200021b57600080fd5b50565b6000815190506200022f8162000204565b92915050565b6000602082840312156200024e576200024d620001f5565b5b60006200025e848285016200021e565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620002a382620001fa565b9150620002b083620001fa565b9250828201905080821115620002cb57620002ca62000267565b5b92915050565b620002dc81620001fa565b82525050565b6000602082019050620002f96000830184620002d1565b92915050565b611712806200030f6000396000f3fe608060405234801561001057600080fd5b506004361061007f576000357c0100000000000000000000000000000000000000000000000000000000900480630f6798a51461008457806369ca02dd146100a05780638bdf1534146100bc578063ab31471e146100ec578063b4a99a4e1461011c578063d16a7a4b1461013a575b600080fd5b61009e60048036038101906100999190610eee565b610156565b005b6100ba60048036038101906100b59190610eee565b61028d565b005b6100d660048036038101906100d19190610f2e565b61046f565b6040516100e39190610f6a565b60405180910390f35b61010660048036038101906101019190610f2e565b610487565b6040516101139190610fa0565b60405180910390f35b6101246104a7565b6040516101319190610fca565b60405180910390f35b610154600480360381019061014f9190611011565b6104cb565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101ae57600080fd5b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546101fd9190611080565b925050819055508173ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45836040516102819190610f6a565b60405180910390a35050565b600061029a3384846105b7565b90508060000151156102e75780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102de9190611144565b60405180910390fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546103369190611166565b9250508190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461038c9190611080565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103bd8361083b565b6103c6856109e2565b6103cf336109e2565b6040516020016103e193929190611248565b6040516020818303038152906040526040516103fd9190611144565b60405180910390a18273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45846040516104629190610f6a565b60405180910390a3505050565b60016020528060005260406000206000915090505481565b60026020528060005260406000206000915054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461052357600080fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d82826040516105ab9291906112a6565b60405180910390a15050565b6105bf610e39565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610638576106316040518060400160405280601f81526020017f63616e27742073656e64206d6f6e657920746f20616464726573732030783000815250610c62565b9050610834565b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561074b576107446040518060400160405280601781526020017f696e737566666963656e742066756e6473202062616c3a000000000000000000815250610700600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461083b565b6040518060400160405280600a81526020017f2020616d6f756e743a200000000000000000000000000000000000000000000081525061073f8661083b565b610c89565b9050610834565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546107d69190611080565b11610829576108226040518060400160405280601081526020017f696e76616c696420616d6f756e743a200000000000000000000000000000000081525061081d8461083b565b610cd8565b9050610834565b610831610d21565b90505b9392505050565b606060008203610882576040518060400160405280600181526020017f300000000000000000000000000000000000000000000000000000000000000081525090506109dd565b600082905060005b600082146108b457808061089d906112cf565b915050600a826108ad9190611346565b915061088a565b60008167ffffffffffffffff8111156108d0576108cf611377565b5b6040519080825280601f01601f1916602001820160405280156109025781602001600182028036833780820191505090505b50905060008290505b600086146109d5576001816109209190611166565b90506000600a80886109329190611346565b61093c91906113a6565b876109479190611166565b603061095391906113f5565b90506000817f0100000000000000000000000000000000000000000000000000000000000000029050808484815181106109905761098f61142a565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a886109cc9190611346565b9750505061090b565b819450505050505b919050565b60606000602867ffffffffffffffff811115610a0157610a00611377565b5b6040519080825280601f01601f191660200182016040528015610a335781602001600182028036833780820191505090505b50905060005b6014811015610c58576000816013610a519190611166565b6008610a5d91906113a6565b6002610a69919061158c565b8573ffffffffffffffffffffffffffffffffffffffff16610a8a9190611346565b7f010000000000000000000000000000000000000000000000000000000000000002905060006010827f01000000000000000000000000000000000000000000000000000000000000009004610ae091906115d7565b7f01000000000000000000000000000000000000000000000000000000000000000290506000817f010000000000000000000000000000000000000000000000000000000000000090046010610b369190611608565b837f01000000000000000000000000000000000000000000000000000000000000009004610b649190611645565b7f0100000000000000000000000000000000000000000000000000000000000000029050610b9182610d55565b85856002610b9f91906113a6565b81518110610bb057610baf61142a565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610be881610d55565b856001866002610bf891906113a6565b610c029190611080565b81518110610c1357610c1261142a565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610c50906112cf565b915050610a39565b5080915050919050565b610c6a610e39565b6040518060400160405280600115158152602001838152509050919050565b610c91610e39565b604051806040016040528060011515815260200186868686604051602001610cbc949392919061167a565b6040516020818303038152906040528152509050949350505050565b610ce0610e39565b60405180604001604052806001151581526020018484604051602001610d079291906116b8565b604051602081830303815290604052815250905092915050565b610d29610e39565b604051806040016040528060001515815260200160405180602001604052806000815250815250905090565b6000600a827f0100000000000000000000000000000000000000000000000000000000000000900460ff161015610ddf576030827f01000000000000000000000000000000000000000000000000000000000000009004610db691906113f5565b7f0100000000000000000000000000000000000000000000000000000000000000029050610e34565b6057827f01000000000000000000000000000000000000000000000000000000000000009004610e0f91906113f5565b7f01000000000000000000000000000000000000000000000000000000000000000290505b919050565b6040518060400160405280600015158152602001606081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e8582610e5a565b9050919050565b610e9581610e7a565b8114610ea057600080fd5b50565b600081359050610eb281610e8c565b92915050565b6000819050919050565b610ecb81610eb8565b8114610ed657600080fd5b50565b600081359050610ee881610ec2565b92915050565b60008060408385031215610f0557610f04610e55565b5b6000610f1385828601610ea3565b9250506020610f2485828601610ed9565b9150509250929050565b600060208284031215610f4457610f43610e55565b5b6000610f5284828501610ea3565b91505092915050565b610f6481610eb8565b82525050565b6000602082019050610f7f6000830184610f5b565b92915050565b60008115159050919050565b610f9a81610f85565b82525050565b6000602082019050610fb56000830184610f91565b92915050565b610fc481610e7a565b82525050565b6000602082019050610fdf6000830184610fbb565b92915050565b610fee81610f85565b8114610ff957600080fd5b50565b60008135905061100b81610fe5565b92915050565b6000806040838503121561102857611027610e55565b5b600061103685828601610ea3565b925050602061104785828601610ffc565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061108b82610eb8565b915061109683610eb8565b92508282019050808211156110ae576110ad611051565b5b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156110ee5780820151818401526020810190506110d3565b60008484015250505050565b6000601f19601f8301169050919050565b6000611116826110b4565b61112081856110bf565b93506111308185602086016110d0565b611139816110fa565b840191505092915050565b6000602082019050818103600083015261115e818461110b565b905092915050565b600061117182610eb8565b915061117c83610eb8565b925082820390508181111561119457611193611051565b5b92915050565b7f7472616e73666572656420000000000000000000000000000000000000000000815250565b600081905092915050565b60006111d6826110b4565b6111e081856111c0565b93506111f08185602086016110d0565b80840191505092915050565b7f20746f2000000000000000000000000000000000000000000000000000000000815250565b7f2066726f6d200000000000000000000000000000000000000000000000000000815250565b60006112538261119a565b600b8201915061126382866111cb565b915061126e826111fc565b60048201915061127e82856111cb565b915061128982611222565b60068201915061129982846111cb565b9150819050949350505050565b60006040820190506112bb6000830185610fbb565b6112c86020830184610f91565b9392505050565b60006112da82610eb8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361130c5761130b611051565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061135182610eb8565b915061135c83610eb8565b92508261136c5761136b611317565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006113b182610eb8565b91506113bc83610eb8565b92508282026113ca81610eb8565b915082820484148315176113e1576113e0611051565b5b5092915050565b600060ff82169050919050565b6000611400826113e8565b915061140b836113e8565b9250828201905060ff81111561142457611423611051565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000600282049050919050565b6000808291508390505b60018511156114b05780860481111561148c5761148b611051565b5b600185161561149b5780820291505b80810290506114a985611459565b9450611470565b94509492505050565b6000826114c95760019050611585565b816114d75760009050611585565b81600181146114ed57600281146114f757611526565b6001915050611585565b60ff84111561150957611508611051565b5b8360020a9150848211156115205761151f611051565b5b50611585565b5060208310610133831016604e8410600b841016171561155b5782820a90508381111561155657611555611051565b5b611585565b6115688484846001611466565b9250905081840481111561157f5761157e611051565b5b81810290505b9392505050565b600061159782610eb8565b91506115a283610eb8565b92506115cf7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846114b9565b905092915050565b60006115e2826113e8565b91506115ed836113e8565b9250826115fd576115fc611317565b5b828204905092915050565b6000611613826113e8565b915061161e836113e8565b925082820261162c816113e8565b915080821461163e5761163d611051565b5b5092915050565b6000611650826113e8565b915061165b836113e8565b9250828203905060ff81111561167457611673611051565b5b92915050565b600061168682876111cb565b915061169282866111cb565b915061169e82856111cb565b91506116aa82846111cb565b915081905095945050505050565b60006116c482856111cb565b91506116d082846111cb565b9150819050939250505056fea2646970667358221220ca0008913a51b05907a745b2de5b51ddfa1eafb17823ceaec106143abf59b32564736f6c63430008140033",
}

// SimplecoinABI is the input ABI used to generate the binding from.
// Deprecated: Use SimplecoinMetaData.ABI instead.
var SimplecoinABI = SimplecoinMetaData.ABI

// SimplecoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimplecoinMetaData.Bin instead.
var SimplecoinBin = SimplecoinMetaData.Bin

// DeploySimplecoin deploys a new Ethereum contract, binding an instance of Simplecoin to it.
func DeploySimplecoin(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *Simplecoin, error) {
	parsed, err := SimplecoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimplecoinBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simplecoin{SimplecoinCaller: SimplecoinCaller{contract: contract}, SimplecoinTransactor: SimplecoinTransactor{contract: contract}, SimplecoinFilterer: SimplecoinFilterer{contract: contract}}, nil
}

// Simplecoin is an auto generated Go binding around an Ethereum contract.
type Simplecoin struct {
	SimplecoinCaller     // Read-only binding to the contract
	SimplecoinTransactor // Write-only binding to the contract
	SimplecoinFilterer   // Log filterer for contract events
}

// SimplecoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplecoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplecoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplecoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplecoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplecoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplecoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplecoinSession struct {
	Contract     *Simplecoin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimplecoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplecoinCallerSession struct {
	Contract *SimplecoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SimplecoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplecoinTransactorSession struct {
	Contract     *SimplecoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimplecoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplecoinRaw struct {
	Contract *Simplecoin // Generic contract binding to access the raw methods on
}

// SimplecoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplecoinCallerRaw struct {
	Contract *SimplecoinCaller // Generic read-only contract binding to access the raw methods on
}

// SimplecoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplecoinTransactorRaw struct {
	Contract *SimplecoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplecoin creates a new instance of Simplecoin, bound to a specific deployed contract.
func NewSimplecoin(address common.Address, backend bind.ContractBackend) (*Simplecoin, error) {
	contract, err := bindSimplecoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simplecoin{SimplecoinCaller: SimplecoinCaller{contract: contract}, SimplecoinTransactor: SimplecoinTransactor{contract: contract}, SimplecoinFilterer: SimplecoinFilterer{contract: contract}}, nil
}

// NewSimplecoinCaller creates a new read-only instance of Simplecoin, bound to a specific deployed contract.
func NewSimplecoinCaller(address common.Address, caller bind.ContractCaller) (*SimplecoinCaller, error) {
	contract, err := bindSimplecoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplecoinCaller{contract: contract}, nil
}

// NewSimplecoinTransactor creates a new write-only instance of Simplecoin, bound to a specific deployed contract.
func NewSimplecoinTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplecoinTransactor, error) {
	contract, err := bindSimplecoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplecoinTransactor{contract: contract}, nil
}

// NewSimplecoinFilterer creates a new log filterer instance of Simplecoin, bound to a specific deployed contract.
func NewSimplecoinFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplecoinFilterer, error) {
	contract, err := bindSimplecoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplecoinFilterer{contract: contract}, nil
}

// bindSimplecoin binds a generic wrapper to an already deployed contract.
func bindSimplecoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimplecoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplecoin *SimplecoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplecoin.Contract.SimplecoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplecoin *SimplecoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplecoin.Contract.SimplecoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplecoin *SimplecoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplecoin.Contract.SimplecoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplecoin *SimplecoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplecoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplecoin *SimplecoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplecoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplecoin *SimplecoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplecoin.Contract.contract.Transact(opts, method, params...)
}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_Simplecoin *SimplecoinCaller) CoinBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Simplecoin.contract.Call(opts, &out, "CoinBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_Simplecoin *SimplecoinSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _Simplecoin.Contract.CoinBalance(&_Simplecoin.CallOpts, arg0)
}

// CoinBalance is a free data retrieval call binding the contract method 0x8bdf1534.
//
// Solidity: function CoinBalance(address ) view returns(uint256)
func (_Simplecoin *SimplecoinCallerSession) CoinBalance(arg0 common.Address) (*big.Int, error) {
	return _Simplecoin.Contract.CoinBalance(&_Simplecoin.CallOpts, arg0)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_Simplecoin *SimplecoinCaller) FrozenAccount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Simplecoin.contract.Call(opts, &out, "FrozenAccount", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_Simplecoin *SimplecoinSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _Simplecoin.Contract.FrozenAccount(&_Simplecoin.CallOpts, arg0)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xab31471e.
//
// Solidity: function FrozenAccount(address ) view returns(bool)
func (_Simplecoin *SimplecoinCallerSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _Simplecoin.Contract.FrozenAccount(&_Simplecoin.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Simplecoin *SimplecoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simplecoin.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Simplecoin *SimplecoinSession) Owner() (common.Address, error) {
	return _Simplecoin.Contract.Owner(&_Simplecoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Simplecoin *SimplecoinCallerSession) Owner() (common.Address, error) {
	return _Simplecoin.Contract.Owner(&_Simplecoin.CallOpts)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_Simplecoin *SimplecoinTransactor) FreezeAccount(opts *bind.TransactOpts, target common.Address, freeze bool) (*types.Transaction, error) {
	return _Simplecoin.contract.Transact(opts, "FreezeAccount", target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_Simplecoin *SimplecoinSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _Simplecoin.Contract.FreezeAccount(&_Simplecoin.TransactOpts, target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xd16a7a4b.
//
// Solidity: function FreezeAccount(address target, bool freeze) returns()
func (_Simplecoin *SimplecoinTransactorSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _Simplecoin.Contract.FreezeAccount(&_Simplecoin.TransactOpts, target, freeze)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_Simplecoin *SimplecoinTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.contract.Transact(opts, "Mint", recipient, mintedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_Simplecoin *SimplecoinSession) Mint(recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.Mint(&_Simplecoin.TransactOpts, recipient, mintedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x0f6798a5.
//
// Solidity: function Mint(address recipient, uint256 mintedAmount) returns()
func (_Simplecoin *SimplecoinTransactorSession) Mint(recipient common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.Mint(&_Simplecoin.TransactOpts, recipient, mintedAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_Simplecoin *SimplecoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.contract.Transact(opts, "Transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_Simplecoin *SimplecoinSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.Transfer(&_Simplecoin.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address to, uint256 amount) returns()
func (_Simplecoin *SimplecoinTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.Transfer(&_Simplecoin.TransactOpts, to, amount)
}

// SimplecoinEventFrozenAccountIterator is returned from FilterEventFrozenAccount and is used to iterate over the raw logs and unpacked data for EventFrozenAccount events raised by the Simplecoin contract.
type SimplecoinEventFrozenAccountIterator struct {
	Event *SimplecoinEventFrozenAccount // Event containing the contract specifics and raw log

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
func (it *SimplecoinEventFrozenAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplecoinEventFrozenAccount)
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
		it.Event = new(SimplecoinEventFrozenAccount)
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
func (it *SimplecoinEventFrozenAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplecoinEventFrozenAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplecoinEventFrozenAccount represents a EventFrozenAccount event raised by the Simplecoin contract.
type SimplecoinEventFrozenAccount struct {
	Target common.Address
	Frozen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventFrozenAccount is a free log retrieval operation binding the contract event 0x3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d.
//
// Solidity: event EventFrozenAccount(address target, bool frozen)
func (_Simplecoin *SimplecoinFilterer) FilterEventFrozenAccount(opts *bind.FilterOpts) (*SimplecoinEventFrozenAccountIterator, error) {

	logs, sub, err := _Simplecoin.contract.FilterLogs(opts, "EventFrozenAccount")
	if err != nil {
		return nil, err
	}
	return &SimplecoinEventFrozenAccountIterator{contract: _Simplecoin.contract, event: "EventFrozenAccount", logs: logs, sub: sub}, nil
}

// WatchEventFrozenAccount is a free log subscription operation binding the contract event 0x3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d.
//
// Solidity: event EventFrozenAccount(address target, bool frozen)
func (_Simplecoin *SimplecoinFilterer) WatchEventFrozenAccount(opts *bind.WatchOpts, sink chan<- *SimplecoinEventFrozenAccount) (event.Subscription, error) {

	logs, sub, err := _Simplecoin.contract.WatchLogs(opts, "EventFrozenAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplecoinEventFrozenAccount)
				if err := _Simplecoin.contract.UnpackLog(event, "EventFrozenAccount", log); err != nil {
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
func (_Simplecoin *SimplecoinFilterer) ParseEventFrozenAccount(log types.Log) (*SimplecoinEventFrozenAccount, error) {
	event := new(SimplecoinEventFrozenAccount)
	if err := _Simplecoin.contract.UnpackLog(event, "EventFrozenAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimplecoinEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Simplecoin contract.
type SimplecoinEventLogIterator struct {
	Event *SimplecoinEventLog // Event containing the contract specifics and raw log

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
func (it *SimplecoinEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplecoinEventLog)
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
		it.Event = new(SimplecoinEventLog)
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
func (it *SimplecoinEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplecoinEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplecoinEventLog represents a EventLog event raised by the Simplecoin contract.
type SimplecoinEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Simplecoin *SimplecoinFilterer) FilterEventLog(opts *bind.FilterOpts) (*SimplecoinEventLogIterator, error) {

	logs, sub, err := _Simplecoin.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &SimplecoinEventLogIterator{contract: _Simplecoin.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Simplecoin *SimplecoinFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *SimplecoinEventLog) (event.Subscription, error) {

	logs, sub, err := _Simplecoin.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplecoinEventLog)
				if err := _Simplecoin.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_Simplecoin *SimplecoinFilterer) ParseEventLog(log types.Log) (*SimplecoinEventLog, error) {
	event := new(SimplecoinEventLog)
	if err := _Simplecoin.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimplecoinEventTransferIterator is returned from FilterEventTransfer and is used to iterate over the raw logs and unpacked data for EventTransfer events raised by the Simplecoin contract.
type SimplecoinEventTransferIterator struct {
	Event *SimplecoinEventTransfer // Event containing the contract specifics and raw log

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
func (it *SimplecoinEventTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplecoinEventTransfer)
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
		it.Event = new(SimplecoinEventTransfer)
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
func (it *SimplecoinEventTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplecoinEventTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplecoinEventTransfer represents a EventTransfer event raised by the Simplecoin contract.
type SimplecoinEventTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventTransfer is a free log retrieval operation binding the contract event 0x64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45.
//
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 amount)
func (_Simplecoin *SimplecoinFilterer) FilterEventTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimplecoinEventTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Simplecoin.contract.FilterLogs(opts, "EventTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SimplecoinEventTransferIterator{contract: _Simplecoin.contract, event: "EventTransfer", logs: logs, sub: sub}, nil
}

// WatchEventTransfer is a free log subscription operation binding the contract event 0x64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45.
//
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 amount)
func (_Simplecoin *SimplecoinFilterer) WatchEventTransfer(opts *bind.WatchOpts, sink chan<- *SimplecoinEventTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Simplecoin.contract.WatchLogs(opts, "EventTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplecoinEventTransfer)
				if err := _Simplecoin.contract.UnpackLog(event, "EventTransfer", log); err != nil {
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
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 amount)
func (_Simplecoin *SimplecoinFilterer) ParseEventTransfer(log types.Log) (*SimplecoinEventTransfer, error) {
	event := new(SimplecoinEventTransfer)
	if err := _Simplecoin.contract.UnpackLog(event, "EventTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
