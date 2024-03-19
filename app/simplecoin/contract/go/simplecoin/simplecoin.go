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
	Bin: "0x608060405234801561000f575f80fd5b506040516117ca3803806117ca8339818101604052810190610031919061020d565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061009f5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff16826100a560201b60201c565b506102c0565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100fb575f80fd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546101479190610265565b925050819055508173ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45836040516101ca91906102a7565b60405180910390a35050565b5f80fd5b5f819050919050565b6101ec816101da565b81146101f6575f80fd5b50565b5f81519050610207816101e3565b92915050565b5f60208284031215610222576102216101d6565b5b5f61022f848285016101f9565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61026f826101da565b915061027a836101da565b925082820190508082111561029257610291610238565b5b92915050565b6102a1816101da565b82525050565b5f6020820190506102ba5f830184610298565b92915050565b6114fd806102cd5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630f6798a51461006457806369ca02dd146100805780638bdf15341461009c578063ab31471e146100cc578063b4a99a4e146100fc578063d16a7a4b1461011a575b5f80fd5b61007e60048036038101906100799190610d0e565b610136565b005b61009a60048036038101906100959190610d0e565b610267565b005b6100b660048036038101906100b19190610d4c565b610441565b6040516100c39190610d86565b60405180910390f35b6100e660048036038101906100e19190610d4c565b610456565b6040516100f39190610db9565b60405180910390f35b610104610473565b6040516101119190610de1565b60405180910390f35b610134600480360381019061012f9190610e24565b610496565b005b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461018c575f80fd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546101d89190610e8f565b925050819055508173ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a458360405161025b9190610d86565b60405180910390a35050565b5f61027333848461057d565b9050805f0151156102bf5780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b69190610f4c565b60405180910390fd5b8160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461030b9190610f6c565b925050819055508160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461035e9190610e8f565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61038f836107f8565b61039885610976565b6103a133610976565b6040516020016103b39392919061104b565b6040516020818303038152906040526040516103cf9190610f4c565b60405180910390a18273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45846040516104349190610d86565b60405180910390a3505050565b6001602052805f5260405f205f915090505481565b6002602052805f5260405f205f915054906101000a900460ff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ec575f80fd5b8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507f3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d82826040516105719291906110a8565b60405180910390a15050565b610585610c62565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036105fd576105f66040518060400160405280601f81526020017f63616e27742073656e64206d6f6e657920746f20616464726573732030783000815250610b2c565b90506107f1565b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101561070c576107056040518060400160405280601781526020017f696e737566666963656e742066756e6473202062616c3a0000000000000000008152506106c160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546107f8565b6040518060400160405280600a81526020017f2020616d6f756e743a2000000000000000000000000000000000000000000000815250610700866107f8565b610b53565b90506107f1565b60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548260015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546107939190610e8f565b116107e6576107df6040518060400160405280601081526020017f696e76616c696420616d6f756e743a20000000000000000000000000000000008152506107da846107f8565b610ba2565b90506107f1565b6107ee610beb565b90505b9392505050565b60605f820361083e576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610971565b5f8290505f5b5f821461086d578080610856906110cf565b915050600a826108669190611143565b9150610844565b5f8167ffffffffffffffff81111561088857610887611173565b5b6040519080825280601f01601f1916602001820160405280156108ba5781602001600182028036833780820191505090505b5090505f8290505b5f8614610969576001816108d69190610f6c565b90505f600a80886108e79190611143565b6108f191906111a0565b876108fc9190610f6c565b603061090891906111ed565b90505f8160f81b90508084848151811061092557610924611221565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a886109609190611143565b975050506108c2565b819450505050505b919050565b60605f602867ffffffffffffffff81111561099457610993611173565b5b6040519080825280601f01601f1916602001820160405280156109c65781602001600182028036833780820191505090505b5090505f5b6014811015610b22575f8160136109e29190610f6c565b60086109ee91906111a0565b60026109fa919061137d565b8573ffffffffffffffffffffffffffffffffffffffff16610a1b9190611143565b60f81b90505f60108260f81c610a3191906113c7565b60f81b90505f8160f81c6010610a4791906113f7565b8360f81c610a559190611433565b60f81b9050610a6382610c1d565b85856002610a7191906111a0565b81518110610a8257610a81611221565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350610ab981610c1d565b856001866002610ac991906111a0565b610ad39190610e8f565b81518110610ae457610ae3611221565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a90535050505080806001019150506109cb565b5080915050919050565b610b34610c62565b6040518060400160405280600115158152602001838152509050919050565b610b5b610c62565b604051806040016040528060011515815260200186868686604051602001610b869493929190611467565b6040516020818303038152906040528152509050949350505050565b610baa610c62565b60405180604001604052806001151581526020018484604051602001610bd19291906114a4565b604051602081830303815290604052815250905092915050565b610bf3610c62565b60405180604001604052805f1515815260200160405180602001604052805f815250815250905090565b5f600a8260f81c60ff161015610c475760308260f81c610c3d91906111ed565b60f81b9050610c5d565b60578260f81c610c5791906111ed565b60f81b90505b919050565b60405180604001604052805f15158152602001606081525090565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610caa82610c81565b9050919050565b610cba81610ca0565b8114610cc4575f80fd5b50565b5f81359050610cd581610cb1565b92915050565b5f819050919050565b610ced81610cdb565b8114610cf7575f80fd5b50565b5f81359050610d0881610ce4565b92915050565b5f8060408385031215610d2457610d23610c7d565b5b5f610d3185828601610cc7565b9250506020610d4285828601610cfa565b9150509250929050565b5f60208284031215610d6157610d60610c7d565b5b5f610d6e84828501610cc7565b91505092915050565b610d8081610cdb565b82525050565b5f602082019050610d995f830184610d77565b92915050565b5f8115159050919050565b610db381610d9f565b82525050565b5f602082019050610dcc5f830184610daa565b92915050565b610ddb81610ca0565b82525050565b5f602082019050610df45f830184610dd2565b92915050565b610e0381610d9f565b8114610e0d575f80fd5b50565b5f81359050610e1e81610dfa565b92915050565b5f8060408385031215610e3a57610e39610c7d565b5b5f610e4785828601610cc7565b9250506020610e5885828601610e10565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610e9982610cdb565b9150610ea483610cdb565b9250828201905080821115610ebc57610ebb610e62565b5b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610ef9578082015181840152602081019050610ede565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610f1e82610ec2565b610f288185610ecc565b9350610f38818560208601610edc565b610f4181610f04565b840191505092915050565b5f6020820190508181035f830152610f648184610f14565b905092915050565b5f610f7682610cdb565b9150610f8183610cdb565b9250828203905081811115610f9957610f98610e62565b5b92915050565b7f7472616e73666572656420000000000000000000000000000000000000000000815250565b5f81905092915050565b5f610fd982610ec2565b610fe38185610fc5565b9350610ff3818560208601610edc565b80840191505092915050565b7f20746f2000000000000000000000000000000000000000000000000000000000815250565b7f2066726f6d200000000000000000000000000000000000000000000000000000815250565b5f61105582610f9f565b600b820191506110658286610fcf565b915061107082610fff565b6004820191506110808285610fcf565b915061108b82611025565b60068201915061109b8284610fcf565b9150819050949350505050565b5f6040820190506110bb5f830185610dd2565b6110c86020830184610daa565b9392505050565b5f6110d982610cdb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361110b5761110a610e62565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61114d82610cdb565b915061115883610cdb565b92508261116857611167611116565b5b828204905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f6111aa82610cdb565b91506111b583610cdb565b92508282026111c381610cdb565b915082820484148315176111da576111d9610e62565b5b5092915050565b5f60ff82169050919050565b5f6111f7826111e1565b9150611202836111e1565b9250828201905060ff81111561121b5761121a610e62565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8160011c9050919050565b5f808291508390505b60018511156112a35780860481111561127f5761127e610e62565b5b600185161561128e5780820291505b808102905061129c8561124e565b9450611263565b94509492505050565b5f826112bb5760019050611376565b816112c8575f9050611376565b81600181146112de57600281146112e857611317565b6001915050611376565b60ff8411156112fa576112f9610e62565b5b8360020a91508482111561131157611310610e62565b5b50611376565b5060208310610133831016604e8410600b841016171561134c5782820a90508381111561134757611346610e62565b5b611376565b611359848484600161125a565b925090508184048111156113705761136f610e62565b5b81810290505b9392505050565b5f61138782610cdb565b915061139283610cdb565b92506113bf7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846112ac565b905092915050565b5f6113d1826111e1565b91506113dc836111e1565b9250826113ec576113eb611116565b5b828204905092915050565b5f611401826111e1565b915061140c836111e1565b925082820261141a816111e1565b915080821461142c5761142b610e62565b5b5092915050565b5f61143d826111e1565b9150611448836111e1565b9250828203905060ff81111561146157611460610e62565b5b92915050565b5f6114728287610fcf565b915061147e8286610fcf565b915061148a8285610fcf565b91506114968284610fcf565b915081905095945050505050565b5f6114af8285610fcf565b91506114bb8284610fcf565b9150819050939250505056fea26469706673582212205f21c9f776ac54bcfe4ab4f31601c5f1ea6d401f2be45e2a1765fc79b3de1ecd64736f6c63430008190033",
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
