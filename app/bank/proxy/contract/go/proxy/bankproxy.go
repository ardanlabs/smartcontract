// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bankproxy

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

// BankproxyMetaData contains all meta data concerning the Bankproxy contract.
var BankproxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"api\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"API\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"api\",\"type\":\"address\"}],\"name\":\"UpgradeAPI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001a0438038062001a04833981810160405281019062000037919062000129565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200015b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000f182620000c4565b9050919050565b6200010381620000e4565b81146200010f57600080fd5b50565b6000815190506200012381620000f8565b92915050565b600060208284031215620001425762000141620000bf565b5b6000620001528482850162000112565b91505092915050565b611899806200016b6000396000f3fe60806040526004361061007b5760003560e01c8063c32ab72e1161004e578063c32ab72e1461010b578063e63f341f14610134578063ed21248c14610171578063fa84fd8e1461017b5761007b565b80630ef678871461008057806357ea89b6146100ab5780637d7b0099146100b5578063b4a99a4e146100e0575b600080fd5b34801561008c57600080fd5b506100956101a4565b6040516100a29190610c9d565b60405180910390f35b6100b36101eb565b005b3480156100c157600080fd5b506100ca6103ae565b6040516100d79190610cf9565b60405180910390f35b3480156100ec57600080fd5b506100f56103d4565b6040516101029190610cf9565b60405180910390f35b34801561011757600080fd5b50610132600480360381019061012d9190610d54565b6103f8565b005b34801561014057600080fd5b5061015b60048036038101906101569190610d54565b6104f2565b6040516101689190610c9d565b60405180910390f35b610179610594565b005b34801561018757600080fd5b506101a2600480360381019061019d9190610f06565b610693565b005b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003390506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403610272576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026990610fe6565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156102fc573d6000803e3d6000fd5b506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61036c33610870565b61037583610a33565b6040516020016103869291906110e9565b6040516020818303038152906040526040516103a29190611173565b60405180910390a15050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461045057600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6104bb82610870565b6040516020016104cb91906111bb565b6040516020818303038152906040526040516104e79190611173565b60405180910390a150565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461054d57600080fd5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b34600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105e3919061121f565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61061433610870565b61065c600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610a33565b60405160200161066d92919061129f565b6040516020818303038152906040526040516106899190611173565b60405180910390a1565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106eb57600080fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168585858560405160240161073d94939291906113ae565b6040516020818303038152906040527ffa84fd8e000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516107c79190611441565b600060405180830381855af49150503d8060008114610802576040519150601f19603f3d011682016040523d82523d6000602084013e610807565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61083582610bbb565b604051602001610845919061147e565b6040516020818303038152906040526040516108619190611173565b60405180910390a15050505050565b60606000602867ffffffffffffffff81111561088f5761088e610d97565b5b6040519080825280601f01601f1916602001820160405280156108c15781602001600182028036833780820191505090505b50905060005b6014811015610a295760008160136108df91906114b3565b60086108eb91906114e7565b60026108f7919061165c565b8573ffffffffffffffffffffffffffffffffffffffff1661091891906116d6565b60f81b9050600060108260f81c61092f9190611714565b60f81b905060008160f81c60106109469190611745565b8360f81c6109549190611782565b60f81b905061096282610c3e565b8585600261097091906114e7565b81518110610981576109806117b7565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506109b981610c3e565b8560018660026109c991906114e7565b6109d3919061121f565b815181106109e4576109e36117b7565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610a21906117e6565b9150506108c7565b5080915050919050565b606060008203610a7a576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610bb6565b600082905060005b60008214610aac578080610a95906117e6565b915050600a82610aa591906116d6565b9150610a82565b60008167ffffffffffffffff811115610ac857610ac7610d97565b5b6040519080825280601f01601f191660200182016040528015610afa5781602001600182028036833780820191505090505b50905060008290505b60008614610bae57600181610b1891906114b3565b90506000600a8088610b2a91906116d6565b610b3491906114e7565b87610b3f91906114b3565b6030610b4b919061182e565b905060008160f81b905080848481518110610b6957610b686117b7565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610ba591906116d6565b97505050610b03565b819450505050505b919050565b60608115610c00576040518060400160405280600481526020017f74727565000000000000000000000000000000000000000000000000000000008152509050610c39565b6040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525090505b919050565b6000600a8260f81c60ff161015610c695760308260f81c610c5f919061182e565b60f81b9050610c7f565b60578260f81c610c79919061182e565b60f81b90505b919050565b6000819050919050565b610c9781610c84565b82525050565b6000602082019050610cb26000830184610c8e565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ce382610cb8565b9050919050565b610cf381610cd8565b82525050565b6000602082019050610d0e6000830184610cea565b92915050565b6000604051905090565b600080fd5b600080fd5b610d3181610cd8565b8114610d3c57600080fd5b50565b600081359050610d4e81610d28565b92915050565b600060208284031215610d6a57610d69610d1e565b5b6000610d7884828501610d3f565b91505092915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610dcf82610d86565b810181811067ffffffffffffffff82111715610dee57610ded610d97565b5b80604052505050565b6000610e01610d14565b9050610e0d8282610dc6565b919050565b600067ffffffffffffffff821115610e2d57610e2c610d97565b5b602082029050602081019050919050565b600080fd5b6000610e56610e5184610e12565b610df7565b90508083825260208201905060208402830185811115610e7957610e78610e3e565b5b835b81811015610ea25780610e8e8882610d3f565b845260208401935050602081019050610e7b565b5050509392505050565b600082601f830112610ec157610ec0610d81565b5b8135610ed1848260208601610e43565b91505092915050565b610ee381610c84565b8114610eee57600080fd5b50565b600081359050610f0081610eda565b92915050565b60008060008060808587031215610f2057610f1f610d1e565b5b6000610f2e87828801610d3f565b945050602085013567ffffffffffffffff811115610f4f57610f4e610d23565b5b610f5b87828801610eac565b9350506040610f6c87828801610ef1565b9250506060610f7d87828801610ef1565b91505092959194509250565b600082825260208201905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000610fd0601283610f89565b9150610fdb82610f9a565b602082019050919050565b60006020820190508181036000830152610fff81610fc3565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b83811015611060578082015181840152602081019050611045565b60008484015250505050565b60006110778261102c565b6110818185611037565b9350611091818560208601611042565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b60006110f482611006565b600982019150611104828561106c565b915061110f8261109d565b60098201915061111f828461106c565b915061112a826110c3565b6001820191508190509392505050565b60006111458261102c565b61114f8185610f89565b935061115f818560208601611042565b61116881610d86565b840191505092915050565b6000602082019050818103600083015261118d818461113a565b905092915050565b7f757067726164655b000000000000000000000000000000000000000000000000815250565b60006111c682611195565b6008820191506111d6828461106c565b91506111e1826110c3565b60018201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061122a82610c84565b915061123583610c84565b925082820190508082111561124d5761124c6111f0565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b60006112aa82611253565b6008820191506112ba828561106c565b91506112c582611279565b600a820191506112d5828461106c565b91506112e0826110c3565b6001820191508190509392505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61132581610cd8565b82525050565b6000611337838361131c565b60208301905092915050565b6000602082019050919050565b600061135b826112f0565b61136581856112fb565b93506113708361130c565b8060005b838110156113a1578151611388888261132b565b975061139383611343565b925050600181019050611374565b5085935050505092915050565b60006080820190506113c36000830187610cea565b81810360208301526113d58186611350565b90506113e46040830185610c8e565b6113f16060830184610c8e565b95945050505050565b600081519050919050565b600081905092915050565b600061141b826113fa565b6114258185611405565b9350611435818560208601611042565b80840191505092915050565b600061144d8284611410565b915081905092915050565b7f7265636f6e63696c655b00000000000000000000000000000000000000000000815250565b600061148982611458565b600a82019150611499828461106c565b91506114a4826110c3565b60018201915081905092915050565b60006114be82610c84565b91506114c983610c84565b92508282039050818111156114e1576114e06111f0565b5b92915050565b60006114f282610c84565b91506114fd83610c84565b925082820261150b81610c84565b91508282048414831517611522576115216111f0565b5b5092915050565b60008160011c9050919050565b6000808291508390505b60018511156115805780860481111561155c5761155b6111f0565b5b600185161561156b5780820291505b808102905061157985611529565b9450611540565b94509492505050565b6000826115995760019050611655565b816115a75760009050611655565b81600181146115bd57600281146115c7576115f6565b6001915050611655565b60ff8411156115d9576115d86111f0565b5b8360020a9150848211156115f0576115ef6111f0565b5b50611655565b5060208310610133831016604e8410600b841016171561162b5782820a905083811115611626576116256111f0565b5b611655565b6116388484846001611536565b9250905081840481111561164f5761164e6111f0565b5b81810290505b9392505050565b600061166782610c84565b915061167283610c84565b925061169f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611589565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006116e182610c84565b91506116ec83610c84565b9250826116fc576116fb6116a7565b5b828204905092915050565b600060ff82169050919050565b600061171f82611707565b915061172a83611707565b92508261173a576117396116a7565b5b828204905092915050565b600061175082611707565b915061175b83611707565b925082820261176981611707565b915080821461177b5761177a6111f0565b5b5092915050565b600061178d82611707565b915061179883611707565b9250828203905060ff8111156117b1576117b06111f0565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006117f182610c84565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611823576118226111f0565b5b600182019050919050565b600061183982611707565b915061184483611707565b9250828201905060ff81111561185d5761185c6111f0565b5b9291505056fea26469706673582212209b97a29c01934d692d08ea9dbec062488f6bf8ecfcff2c5715cdd6e0bad9bfca64736f6c63430008110033",
}

// BankproxyABI is the input ABI used to generate the binding from.
// Deprecated: Use BankproxyMetaData.ABI instead.
var BankproxyABI = BankproxyMetaData.ABI

// BankproxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankproxyMetaData.Bin instead.
var BankproxyBin = BankproxyMetaData.Bin

// DeployBankproxy deploys a new Ethereum contract, binding an instance of Bankproxy to it.
func DeployBankproxy(auth *bind.TransactOpts, backend bind.ContractBackend, api common.Address) (common.Address, *types.Transaction, *Bankproxy, error) {
	parsed, err := BankproxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankproxyBin), backend, api)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bankproxy{BankproxyCaller: BankproxyCaller{contract: contract}, BankproxyTransactor: BankproxyTransactor{contract: contract}, BankproxyFilterer: BankproxyFilterer{contract: contract}}, nil
}

// Bankproxy is an auto generated Go binding around an Ethereum contract.
type Bankproxy struct {
	BankproxyCaller     // Read-only binding to the contract
	BankproxyTransactor // Write-only binding to the contract
	BankproxyFilterer   // Log filterer for contract events
}

// BankproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankproxySession struct {
	Contract     *Bankproxy        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankproxyCallerSession struct {
	Contract *BankproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BankproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankproxyTransactorSession struct {
	Contract     *BankproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BankproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankproxyRaw struct {
	Contract *Bankproxy // Generic contract binding to access the raw methods on
}

// BankproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankproxyCallerRaw struct {
	Contract *BankproxyCaller // Generic read-only contract binding to access the raw methods on
}

// BankproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankproxyTransactorRaw struct {
	Contract *BankproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBankproxy creates a new instance of Bankproxy, bound to a specific deployed contract.
func NewBankproxy(address common.Address, backend bind.ContractBackend) (*Bankproxy, error) {
	contract, err := bindBankproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bankproxy{BankproxyCaller: BankproxyCaller{contract: contract}, BankproxyTransactor: BankproxyTransactor{contract: contract}, BankproxyFilterer: BankproxyFilterer{contract: contract}}, nil
}

// NewBankproxyCaller creates a new read-only instance of Bankproxy, bound to a specific deployed contract.
func NewBankproxyCaller(address common.Address, caller bind.ContractCaller) (*BankproxyCaller, error) {
	contract, err := bindBankproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankproxyCaller{contract: contract}, nil
}

// NewBankproxyTransactor creates a new write-only instance of Bankproxy, bound to a specific deployed contract.
func NewBankproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*BankproxyTransactor, error) {
	contract, err := bindBankproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankproxyTransactor{contract: contract}, nil
}

// NewBankproxyFilterer creates a new log filterer instance of Bankproxy, bound to a specific deployed contract.
func NewBankproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*BankproxyFilterer, error) {
	contract, err := bindBankproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankproxyFilterer{contract: contract}, nil
}

// bindBankproxy binds a generic wrapper to an already deployed contract.
func bindBankproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankproxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankproxy *BankproxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankproxy.Contract.BankproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankproxy *BankproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankproxy.Contract.BankproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankproxy *BankproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankproxy.Contract.BankproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankproxy *BankproxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankproxy *BankproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankproxy *BankproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankproxy.Contract.contract.Transact(opts, method, params...)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankproxy *BankproxyCaller) API(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankproxy.contract.Call(opts, &out, "API")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankproxy *BankproxySession) API() (common.Address, error) {
	return _Bankproxy.Contract.API(&_Bankproxy.CallOpts)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankproxy *BankproxyCallerSession) API() (common.Address, error) {
	return _Bankproxy.Contract.API(&_Bankproxy.CallOpts)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bankproxy *BankproxyCaller) AccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bankproxy.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bankproxy *BankproxySession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bankproxy.Contract.AccountBalance(&_Bankproxy.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bankproxy *BankproxyCallerSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bankproxy.Contract.AccountBalance(&_Bankproxy.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bankproxy *BankproxyCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bankproxy.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bankproxy *BankproxySession) Balance() (*big.Int, error) {
	return _Bankproxy.Contract.Balance(&_Bankproxy.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bankproxy *BankproxyCallerSession) Balance() (*big.Int, error) {
	return _Bankproxy.Contract.Balance(&_Bankproxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankproxy *BankproxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankproxy.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankproxy *BankproxySession) Owner() (common.Address, error) {
	return _Bankproxy.Contract.Owner(&_Bankproxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankproxy *BankproxyCallerSession) Owner() (common.Address, error) {
	return _Bankproxy.Contract.Owner(&_Bankproxy.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankproxy *BankproxyTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankproxy.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankproxy *BankproxySession) Deposit() (*types.Transaction, error) {
	return _Bankproxy.Contract.Deposit(&_Bankproxy.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankproxy *BankproxyTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bankproxy.Contract.Deposit(&_Bankproxy.TransactOpts)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address _winner, address[] _losers, uint256 _anteWei, uint256 _gameFeeWei) returns()
func (_Bankproxy *BankproxyTransactor) Reconcile(opts *bind.TransactOpts, _winner common.Address, _losers []common.Address, _anteWei *big.Int, _gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankproxy.contract.Transact(opts, "Reconcile", _winner, _losers, _anteWei, _gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address _winner, address[] _losers, uint256 _anteWei, uint256 _gameFeeWei) returns()
func (_Bankproxy *BankproxySession) Reconcile(_winner common.Address, _losers []common.Address, _anteWei *big.Int, _gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankproxy.Contract.Reconcile(&_Bankproxy.TransactOpts, _winner, _losers, _anteWei, _gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address _winner, address[] _losers, uint256 _anteWei, uint256 _gameFeeWei) returns()
func (_Bankproxy *BankproxyTransactorSession) Reconcile(_winner common.Address, _losers []common.Address, _anteWei *big.Int, _gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bankproxy.Contract.Reconcile(&_Bankproxy.TransactOpts, _winner, _losers, _anteWei, _gameFeeWei)
}

// UpgradeAPI is a paid mutator transaction binding the contract method 0xc32ab72e.
//
// Solidity: function UpgradeAPI(address api) returns()
func (_Bankproxy *BankproxyTransactor) UpgradeAPI(opts *bind.TransactOpts, api common.Address) (*types.Transaction, error) {
	return _Bankproxy.contract.Transact(opts, "UpgradeAPI", api)
}

// UpgradeAPI is a paid mutator transaction binding the contract method 0xc32ab72e.
//
// Solidity: function UpgradeAPI(address api) returns()
func (_Bankproxy *BankproxySession) UpgradeAPI(api common.Address) (*types.Transaction, error) {
	return _Bankproxy.Contract.UpgradeAPI(&_Bankproxy.TransactOpts, api)
}

// UpgradeAPI is a paid mutator transaction binding the contract method 0xc32ab72e.
//
// Solidity: function UpgradeAPI(address api) returns()
func (_Bankproxy *BankproxyTransactorSession) UpgradeAPI(api common.Address) (*types.Transaction, error) {
	return _Bankproxy.Contract.UpgradeAPI(&_Bankproxy.TransactOpts, api)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankproxy *BankproxyTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankproxy.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankproxy *BankproxySession) Withdraw() (*types.Transaction, error) {
	return _Bankproxy.Contract.Withdraw(&_Bankproxy.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankproxy *BankproxyTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bankproxy.Contract.Withdraw(&_Bankproxy.TransactOpts)
}

// BankproxyEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bankproxy contract.
type BankproxyEventLogIterator struct {
	Event *BankproxyEventLog // Event containing the contract specifics and raw log

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
func (it *BankproxyEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankproxyEventLog)
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
		it.Event = new(BankproxyEventLog)
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
func (it *BankproxyEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankproxyEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankproxyEventLog represents a EventLog event raised by the Bankproxy contract.
type BankproxyEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankproxy *BankproxyFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankproxyEventLogIterator, error) {

	logs, sub, err := _Bankproxy.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankproxyEventLogIterator{contract: _Bankproxy.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankproxy *BankproxyFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankproxyEventLog) (event.Subscription, error) {

	logs, sub, err := _Bankproxy.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankproxyEventLog)
				if err := _Bankproxy.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_Bankproxy *BankproxyFilterer) ParseEventLog(log types.Log) (*BankproxyEventLog, error) {
	event := new(BankproxyEventLog)
	if err := _Bankproxy.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
