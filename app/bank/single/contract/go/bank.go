// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package banksingle

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

// BanksingleMetaData contains all meta data concerning the Banksingle contract.
var BanksingleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611b2d806100606000396000f3fe6080604052600436106100555760003560e01c80630ef678871461005a57806357ea89b614610085578063b4a99a4e1461008f578063e63f341f146100ba578063ed21248c146100f7578063fa84fd8e14610101575b600080fd5b34801561006657600080fd5b5061006f61012a565b60405161007c9190610eb3565b60405180910390f35b61008d610171565b005b34801561009b57600080fd5b506100a4610334565b6040516100b19190610f0f565b60405180910390f35b3480156100c657600080fd5b506100e160048036038101906100dc9190610f60565b610358565b6040516100ee9190610eb3565b60405180910390f35b6100ff6103fa565b005b34801561010d57600080fd5b506101286004803603810190610123919061101e565b6104f9565b005b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003390506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054036101f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ef90611103565b60405180910390fd5b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610282573d6000803e3d6000fd5b506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6102f233610b09565b6102fb83610ccc565b60405160200161030c929190611206565b60405160208183030381529060405260405161032891906112a1565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103b357600080fd5b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461044991906112f2565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61047a33610b09565b6104c2600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ccc565b6040516020016104d3929190611372565b6040516020818303038152906040526040516104ef91906112a1565b60405180910390a1565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461055157600080fd5b600082905060005b8585905081101561082457836001600088888581811061057c5761057b6113c3565b5b90506020020160208101906105919190610f60565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610785577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6106626001600089898681811061060d5761060c6113c3565b5b90506020020160208101906106229190610f60565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ccc565b61066b86610ccc565b60405160200161067c92919061143e565b60405160208183030381529060405260405161069891906112a1565b60405180910390a1600160008787848181106106b7576106b66113c3565b5b90506020020160208101906106cc9190610f60565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548261071291906112f2565b915060006001600088888581811061072d5761072c6113c3565b5b90506020020160208101906107429190610f60565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610811565b838261079191906112f2565b915083600160008888858181106107ab576107aa6113c3565b5b90506020020160208101906107c09190610f60565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108099190611480565b925050819055505b808061081c906114b4565b915050610559565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61084f84610ccc565b61085884610ccc565b61086184610ccc565b6040516020016108739392919061156e565b60405160208183030381529060405260405161088f91906112a1565b60405180910390a1600081036108da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d19061164d565b60405180910390fd5b818110156109bd5780600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461095291906112f2565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61098382610ccc565b60405160200161099391906116df565b6040516020818303038152906040526040516109af91906112a1565b60405180910390a150610b02565b81816109c99190611480565b905080600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a1a91906112f2565b9250508190555081600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a9191906112f2565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610ac282610ccc565b610acb84610ccc565b604051602001610adc92919061175c565b604051602081830303815290604052604051610af891906112a1565b60405180910390a1505b5050505050565b60606000602867ffffffffffffffff811115610b2857610b276117ad565b5b6040519080825280601f01601f191660200182016040528015610b5a5781602001600182028036833780820191505090505b50905060005b6014811015610cc2576000816013610b789190611480565b6008610b8491906117dc565b6002610b909190611969565b8573ffffffffffffffffffffffffffffffffffffffff16610bb191906119e3565b60f81b9050600060108260f81c610bc89190611a21565b60f81b905060008160f81c6010610bdf9190611a52565b8360f81c610bed9190611a8d565b60f81b9050610bfb82610e54565b85856002610c0991906117dc565b81518110610c1a57610c196113c3565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610c5281610e54565b856001866002610c6291906117dc565b610c6c91906112f2565b81518110610c7d57610c7c6113c3565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610cba906114b4565b915050610b60565b5080915050919050565b606060008203610d13576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610e4f565b600082905060005b60008214610d45578080610d2e906114b4565b915050600a82610d3e91906119e3565b9150610d1b565b60008167ffffffffffffffff811115610d6157610d606117ad565b5b6040519080825280601f01601f191660200182016040528015610d935781602001600182028036833780820191505090505b50905060008290505b60008614610e4757600181610db19190611480565b90506000600a8088610dc391906119e3565b610dcd91906117dc565b87610dd89190611480565b6030610de49190611ac2565b905060008160f81b905080848481518110610e0257610e016113c3565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610e3e91906119e3565b97505050610d9c565b819450505050505b919050565b6000600a8260f81c60ff161015610e7f5760308260f81c610e759190611ac2565b60f81b9050610e95565b60578260f81c610e8f9190611ac2565b60f81b90505b919050565b6000819050919050565b610ead81610e9a565b82525050565b6000602082019050610ec86000830184610ea4565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ef982610ece565b9050919050565b610f0981610eee565b82525050565b6000602082019050610f246000830184610f00565b92915050565b600080fd5b600080fd5b610f3d81610eee565b8114610f4857600080fd5b50565b600081359050610f5a81610f34565b92915050565b600060208284031215610f7657610f75610f2a565b5b6000610f8484828501610f4b565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610fb257610fb1610f8d565b5b8235905067ffffffffffffffff811115610fcf57610fce610f92565b5b602083019150836020820283011115610feb57610fea610f97565b5b9250929050565b610ffb81610e9a565b811461100657600080fd5b50565b60008135905061101881610ff2565b92915050565b60008060008060006080868803121561103a57611039610f2a565b5b600061104888828901610f4b565b955050602086013567ffffffffffffffff81111561106957611068610f2f565b5b61107588828901610f9c565b9450945050604061108888828901611009565b925050606061109988828901611009565b9150509295509295909350565b600082825260208201905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b60006110ed6012836110a6565b91506110f8826110b7565b602082019050919050565b6000602082019050818103600083015261111c816110e0565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b8381101561117d578082015181840152602081019050611162565b60008484015250505050565b600061119482611149565b61119e8185611154565b93506111ae81856020860161115f565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b600061121182611123565b6009820191506112218285611189565b915061122c826111ba565b60098201915061123c8284611189565b9150611247826111e0565b6001820191508190509392505050565b6000601f19601f8301169050919050565b600061127382611149565b61127d81856110a6565b935061128d81856020860161115f565b61129681611257565b840191505092915050565b600060208201905081810360008301526112bb8184611268565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006112fd82610e9a565b915061130883610e9a565b92508282019050808211156113205761131f6112c3565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b600061137d82611326565b60088201915061138d8285611189565b91506113988261134c565b600a820191506113a88284611189565b91506113b3826111e0565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f6163636f756e742062616c616e63652000000000000000000000000000000000815250565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b6000611449826113f2565b6010820191506114598285611189565b915061146482611418565b6013820191506114748284611189565b91508190509392505050565b600061148b82610e9a565b915061149683610e9a565b92508282039050818111156114ae576114ad6112c3565b5b92915050565b60006114bf82610e9a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036114f1576114f06112c3565b5b600182019050919050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b6000611579826114fc565b6005820191506115898286611189565b915061159482611522565b600a820191506115a48285611189565b91506115af82611548565b6006820191506115bf8284611189565b91506115ca826111e0565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e206163636f60008201527f756e742062616c616e6365730000000000000000000000000000000000000000602082015250565b6000611637602c836110a6565b9150611642826115db565b604082019050919050565b600060208201905081810360008301526116668161162a565b9050919050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e60008201527f65725b0000000000000000000000000000000000000000000000000000000000602082015250565b60006116c9602383611154565b91506116d48261166d565b602382019050919050565b60006116ea826116bc565b91506116f68284611189565b9150611701826111e0565b60018201915081905092915050565b7f77696e6e65725b00000000000000000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b600061176782611710565b6007820191506117778285611189565b915061178282611736565b6008820191506117928284611189565b915061179d826111e0565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006117e782610e9a565b91506117f283610e9a565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561182b5761182a6112c3565b5b828202905092915050565b60008160011c9050919050565b6000808291508390505b600185111561188d57808604811115611869576118686112c3565b5b60018516156118785780820291505b808102905061188685611836565b945061184d565b94509492505050565b6000826118a65760019050611962565b816118b45760009050611962565b81600181146118ca57600281146118d457611903565b6001915050611962565b60ff8411156118e6576118e56112c3565b5b8360020a9150848211156118fd576118fc6112c3565b5b50611962565b5060208310610133831016604e8410600b84101617156119385782820a905083811115611933576119326112c3565b5b611962565b6119458484846001611843565b9250905081840481111561195c5761195b6112c3565b5b81810290505b9392505050565b600061197482610e9a565b915061197f83610e9a565b92506119ac7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611896565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006119ee82610e9a565b91506119f983610e9a565b925082611a0957611a086119b4565b5b828204905092915050565b600060ff82169050919050565b6000611a2c82611a14565b9150611a3783611a14565b925082611a4757611a466119b4565b5b828204905092915050565b6000611a5d82611a14565b9150611a6883611a14565b92508160ff0483118215151615611a8257611a816112c3565b5b828202905092915050565b6000611a9882611a14565b9150611aa383611a14565b9250828203905060ff811115611abc57611abb6112c3565b5b92915050565b6000611acd82611a14565b9150611ad883611a14565b9250828201905060ff811115611af157611af06112c3565b5b9291505056fea2646970667358221220b2c94f8819a39f68eae46dd7b7a681cbdc5b8325268e760021d3ec668a95bf9064736f6c63430008100033",
}

// BanksingleABI is the input ABI used to generate the binding from.
// Deprecated: Use BanksingleMetaData.ABI instead.
var BanksingleABI = BanksingleMetaData.ABI

// BanksingleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BanksingleMetaData.Bin instead.
var BanksingleBin = BanksingleMetaData.Bin

// DeployBanksingle deploys a new Ethereum contract, binding an instance of Banksingle to it.
func DeployBanksingle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Banksingle, error) {
	parsed, err := BanksingleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BanksingleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Banksingle{BanksingleCaller: BanksingleCaller{contract: contract}, BanksingleTransactor: BanksingleTransactor{contract: contract}, BanksingleFilterer: BanksingleFilterer{contract: contract}}, nil
}

// Banksingle is an auto generated Go binding around an Ethereum contract.
type Banksingle struct {
	BanksingleCaller     // Read-only binding to the contract
	BanksingleTransactor // Write-only binding to the contract
	BanksingleFilterer   // Log filterer for contract events
}

// BanksingleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BanksingleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BanksingleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BanksingleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BanksingleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BanksingleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BanksingleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BanksingleSession struct {
	Contract     *Banksingle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BanksingleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BanksingleCallerSession struct {
	Contract *BanksingleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BanksingleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BanksingleTransactorSession struct {
	Contract     *BanksingleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BanksingleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BanksingleRaw struct {
	Contract *Banksingle // Generic contract binding to access the raw methods on
}

// BanksingleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BanksingleCallerRaw struct {
	Contract *BanksingleCaller // Generic read-only contract binding to access the raw methods on
}

// BanksingleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BanksingleTransactorRaw struct {
	Contract *BanksingleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBanksingle creates a new instance of Banksingle, bound to a specific deployed contract.
func NewBanksingle(address common.Address, backend bind.ContractBackend) (*Banksingle, error) {
	contract, err := bindBanksingle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Banksingle{BanksingleCaller: BanksingleCaller{contract: contract}, BanksingleTransactor: BanksingleTransactor{contract: contract}, BanksingleFilterer: BanksingleFilterer{contract: contract}}, nil
}

// NewBanksingleCaller creates a new read-only instance of Banksingle, bound to a specific deployed contract.
func NewBanksingleCaller(address common.Address, caller bind.ContractCaller) (*BanksingleCaller, error) {
	contract, err := bindBanksingle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BanksingleCaller{contract: contract}, nil
}

// NewBanksingleTransactor creates a new write-only instance of Banksingle, bound to a specific deployed contract.
func NewBanksingleTransactor(address common.Address, transactor bind.ContractTransactor) (*BanksingleTransactor, error) {
	contract, err := bindBanksingle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BanksingleTransactor{contract: contract}, nil
}

// NewBanksingleFilterer creates a new log filterer instance of Banksingle, bound to a specific deployed contract.
func NewBanksingleFilterer(address common.Address, filterer bind.ContractFilterer) (*BanksingleFilterer, error) {
	contract, err := bindBanksingle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BanksingleFilterer{contract: contract}, nil
}

// bindBanksingle binds a generic wrapper to an already deployed contract.
func bindBanksingle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BanksingleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Banksingle *BanksingleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Banksingle.Contract.BanksingleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Banksingle *BanksingleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Banksingle.Contract.BanksingleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Banksingle *BanksingleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Banksingle.Contract.BanksingleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Banksingle *BanksingleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Banksingle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Banksingle *BanksingleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Banksingle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Banksingle *BanksingleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Banksingle.Contract.contract.Transact(opts, method, params...)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Banksingle *BanksingleCaller) AccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Banksingle.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Banksingle *BanksingleSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Banksingle.Contract.AccountBalance(&_Banksingle.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Banksingle *BanksingleCallerSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Banksingle.Contract.AccountBalance(&_Banksingle.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Banksingle *BanksingleCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Banksingle.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Banksingle *BanksingleSession) Balance() (*big.Int, error) {
	return _Banksingle.Contract.Balance(&_Banksingle.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Banksingle *BanksingleCallerSession) Balance() (*big.Int, error) {
	return _Banksingle.Contract.Balance(&_Banksingle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Banksingle *BanksingleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Banksingle.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Banksingle *BanksingleSession) Owner() (common.Address, error) {
	return _Banksingle.Contract.Owner(&_Banksingle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Banksingle *BanksingleCallerSession) Owner() (common.Address, error) {
	return _Banksingle.Contract.Owner(&_Banksingle.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Banksingle *BanksingleTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Banksingle.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Banksingle *BanksingleSession) Deposit() (*types.Transaction, error) {
	return _Banksingle.Contract.Deposit(&_Banksingle.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Banksingle *BanksingleTransactorSession) Deposit() (*types.Transaction, error) {
	return _Banksingle.Contract.Deposit(&_Banksingle.TransactOpts)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Banksingle *BanksingleTransactor) Reconcile(opts *bind.TransactOpts, winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Banksingle.contract.Transact(opts, "Reconcile", winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Banksingle *BanksingleSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Banksingle.Contract.Reconcile(&_Banksingle.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Banksingle *BanksingleTransactorSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Banksingle.Contract.Reconcile(&_Banksingle.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Banksingle *BanksingleTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Banksingle.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Banksingle *BanksingleSession) Withdraw() (*types.Transaction, error) {
	return _Banksingle.Contract.Withdraw(&_Banksingle.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Banksingle *BanksingleTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Banksingle.Contract.Withdraw(&_Banksingle.TransactOpts)
}

// BanksingleEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Banksingle contract.
type BanksingleEventLogIterator struct {
	Event *BanksingleEventLog // Event containing the contract specifics and raw log

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
func (it *BanksingleEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BanksingleEventLog)
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
		it.Event = new(BanksingleEventLog)
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
func (it *BanksingleEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BanksingleEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BanksingleEventLog represents a EventLog event raised by the Banksingle contract.
type BanksingleEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Banksingle *BanksingleFilterer) FilterEventLog(opts *bind.FilterOpts) (*BanksingleEventLogIterator, error) {

	logs, sub, err := _Banksingle.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BanksingleEventLogIterator{contract: _Banksingle.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Banksingle *BanksingleFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BanksingleEventLog) (event.Subscription, error) {

	logs, sub, err := _Banksingle.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BanksingleEventLog)
				if err := _Banksingle.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_Banksingle *BanksingleFilterer) ParseEventLog(log types.Log) (*BanksingleEventLog, error) {
	event := new(BanksingleEventLog)
	if err := _Banksingle.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
