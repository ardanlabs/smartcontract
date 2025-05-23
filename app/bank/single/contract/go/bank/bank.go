// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060400160405280600581526020017f302e312e300000000000000000000000000000000000000000000000000000008152506001908161009391906102d6565b506103a5565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061011457607f821691505b602082108103610127576101266100d0565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101897fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261014e565b610193868361014e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6101d76101d26101cd846101ab565b6101b4565b6101ab565b9050919050565b5f819050919050565b6101f0836101bd565b6102046101fc826101de565b84845461015a565b825550505050565b5f5f905090565b61021b61020c565b6102268184846101e7565b505050565b5b818110156102495761023e5f82610213565b60018101905061022c565b5050565b601f82111561028e5761025f8161012d565b6102688461013f565b81016020851015610277578190505b61028b6102838561013f565b83018261022b565b50505b505050565b5f82821c905092915050565b5f6102ae5f1984600802610293565b1980831691505092915050565b5f6102c6838361029f565b9150826002028217905092915050565b6102df82610099565b67ffffffffffffffff8111156102f8576102f76100a3565b5b61030282546100fd565b61030d82828561024d565b5f60209050601f83116001811461033e575f841561032c578287015190505b61033685826102bb565b86555061039d565b601f19841661034c8661012d565b5f5b828110156103735784890151825560018201915060208501945060208101905061034e565b86831015610390578489015161038c601f89168261029f565b8355505b6001600288020188555050505b505050505050565b611b90806103b25f395ff3fe60806040526004361061006f575f3560e01c8063bb62860d1161004d578063bb62860d146100d1578063e63f341f146100fb578063ed21248c14610137578063fa84fd8e146101415761006f565b80630ef678871461007357806357ea89b61461009d578063b4a99a4e146100a7575b5f5ffd5b34801561007e575f5ffd5b50610087610169565b6040516100949190610f30565b60405180910390f35b6100a56101ad565b005b3480156100b2575f5ffd5b506100bb610363565b6040516100c89190610f88565b60405180910390f35b3480156100dc575f5ffd5b506100e5610387565b6040516100f29190611011565b60405180910390f35b348015610106575f5ffd5b50610121600480360381019061011c9190611063565b610413565b60405161012e9190610f30565b60405180910390f35b61013f6104b0565b005b34801561014c575f5ffd5b5061016760048036038101906101629190611119565b6105aa565b005b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905090565b5f3390505f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610230576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610227906111e7565b60405180910390fd5b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490508173ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156102b4573d5f5f3e3d5ffd5b505f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61032133610b9c565b61032a83610d55565b60405160200161033b9291906112b1565b6040516020818303038152906040526040516103579190611011565b60405180910390a15050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600180546103949061132e565b80601f01602080910402602001604051908101604052809291908181526020018280546103c09061132e565b801561040b5780601f106103e25761010080835404028352916020019161040b565b820191905f5260205f20905b8154815290600101906020018083116103ee57829003601f168201915b505050505081565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461046b575f5ffd5b60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b3460025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546104fc919061138b565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61052d33610b9c565b61057360025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610d55565b60405160200161058492919061140a565b6040516020818303038152906040526040516105a09190611011565b60405180910390a1565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610601575f5ffd5b5f8290505f5f90505b858590508110156108c3578360025f88888581811061062c5761062b61145a565b5b90506020020160208101906106419190611063565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101561082d577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61070f60025f8989868181106106bb576106ba61145a565b5b90506020020160208101906106d09190611063565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610d55565b61071886610d55565b6040516020016107299291906114d3565b6040516020818303038152906040526040516107459190611011565b60405180910390a160025f8787848181106107635761076261145a565b5b90506020020160208101906107789190611063565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054826107bd919061138b565b91505f60025f8888858181106107d6576107d561145a565b5b90506020020160208101906107eb9190611063565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055506108b6565b8382610839919061138b565b91508360025f8888858181106108525761085161145a565b5b90506020020160208101906108679190611063565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546108ae9190611514565b925050819055505b808060010191505061060a565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6108ee84610d55565b6108f784610d55565b61090084610d55565b604051602001610912939291906115b9565b60405160208183030381529060405260405161092e9190611011565b60405180910390a15f8103610978576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096f90611695565b60405180910390fd5b81811015610a57578060025f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109ec919061138b565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610a1d82610d55565b604051602001610a2d9190611723565b604051602081830303815290604052604051610a499190611011565b60405180910390a150610b95565b8181610a639190611514565b90508060025f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ab1919061138b565b925050819055508160025f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610b24919061138b565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610b5582610d55565b610b5e84610d55565b604051602001610b6f92919061179f565b604051602081830303815290604052604051610b8b9190611011565b60405180910390a1505b5050505050565b60605f602867ffffffffffffffff811115610bba57610bb96117ef565b5b6040519080825280601f01601f191660200182016040528015610bec5781602001600182028036833780820191505090505b5090505f5f90505b6014811015610d4b575f816013610c0b9190611514565b6008610c17919061181c565b6002610c23919061198c565b8573ffffffffffffffffffffffffffffffffffffffff16610c449190611a03565b60f81b90505f60108260f81c610c5a9190611a3f565b60f81b90505f8160f81c6010610c709190611a6f565b8360f81c610c7e9190611aab565b60f81b9050610c8c82610ed3565b85856002610c9a919061181c565b81518110610cab57610caa61145a565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350610ce281610ed3565b856001866002610cf2919061181c565b610cfc919061138b565b81518110610d0d57610d0c61145a565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505050508080600101915050610bf4565b5080915050919050565b60605f8203610d9b576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610ece565b5f8290505f5b5f8214610dca578080610db390611adf565b915050600a82610dc39190611a03565b9150610da1565b5f8167ffffffffffffffff811115610de557610de46117ef565b5b6040519080825280601f01601f191660200182016040528015610e175781602001600182028036833780820191505090505b5090505f8290505b5f8614610ec657600181610e339190611514565b90505f600a8088610e449190611a03565b610e4e919061181c565b87610e599190611514565b6030610e659190611b26565b90505f8160f81b905080848481518110610e8257610e8161145a565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a88610ebd9190611a03565b97505050610e1f565b819450505050505b919050565b5f600a8260f81c60ff161015610efd5760308260f81c610ef39190611b26565b60f81b9050610f13565b60578260f81c610f0d9190611b26565b60f81b90505b919050565b5f819050919050565b610f2a81610f18565b82525050565b5f602082019050610f435f830184610f21565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f7282610f49565b9050919050565b610f8281610f68565b82525050565b5f602082019050610f9b5f830184610f79565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610fe382610fa1565b610fed8185610fab565b9350610ffd818560208601610fbb565b61100681610fc9565b840191505092915050565b5f6020820190508181035f8301526110298184610fd9565b905092915050565b5f5ffd5b5f5ffd5b61104281610f68565b811461104c575f5ffd5b50565b5f8135905061105d81611039565b92915050565b5f6020828403121561107857611077611031565b5b5f6110858482850161104f565b91505092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126110af576110ae61108e565b5b8235905067ffffffffffffffff8111156110cc576110cb611092565b5b6020830191508360208202830111156110e8576110e7611096565b5b9250929050565b6110f881610f18565b8114611102575f5ffd5b50565b5f81359050611113816110ef565b92915050565b5f5f5f5f5f6080868803121561113257611131611031565b5b5f61113f8882890161104f565b955050602086013567ffffffffffffffff8111156111605761115f611035565b5b61116c8882890161109a565b9450945050604061117f88828901611105565b925050606061119088828901611105565b9150509295509295909350565b7f6e6f7420656e6f7567682062616c616e636500000000000000000000000000005f82015250565b5f6111d1601283610fab565b91506111dc8261119d565b602082019050919050565b5f6020820190508181035f8301526111fe816111c5565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b5f81905092915050565b5f61123f82610fa1565b611249818561122b565b9350611259818560208601610fbb565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b5f6112bb82611205565b6009820191506112cb8285611235565b91506112d682611265565b6009820191506112e68284611235565b91506112f18261128b565b6001820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061134557607f821691505b60208210810361135857611357611301565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61139582610f18565b91506113a083610f18565b92508282019050808211156113b8576113b761135e565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b5f611414826113be565b6008820191506114248285611235565b915061142f826113e4565b600a8201915061143f8284611235565b915061144a8261128b565b6001820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f6163636f756e742062616c616e63652000000000000000000000000000000000815250565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b5f6114dd82611487565b6010820191506114ed8285611235565b91506114f8826114ad565b6013820191506115088284611235565b91508190509392505050565b5f61151e82610f18565b915061152983610f18565b92508282039050818111156115415761154061135e565b5b92915050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b5f6115c382611547565b6005820191506115d38286611235565b91506115de8261156d565b600a820191506115ee8285611235565b91506115f982611593565b6006820191506116098284611235565b91506116148261128b565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e206163636f5f8201527f756e742062616c616e6365730000000000000000000000000000000000000000602082015250565b5f61167f602c83610fab565b915061168a82611625565b604082019050919050565b5f6020820190508181035f8301526116ac81611673565b9050919050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e5f8201527f65725b0000000000000000000000000000000000000000000000000000000000602082015250565b5f61170d60238361122b565b9150611718826116b3565b602382019050919050565b5f61172d82611701565b91506117398284611235565b91506117448261128b565b60018201915081905092915050565b7f77696e6e65725b00000000000000000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b5f6117a982611753565b6007820191506117b98285611235565b91506117c482611779565b6008820191506117d48284611235565b91506117df8261128b565b6001820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f61182682610f18565b915061183183610f18565b925082820261183f81610f18565b915082820484148315176118565761185561135e565b5b5092915050565b5f8160011c9050919050565b5f5f8291508390505b60018511156118b25780860481111561188e5761188d61135e565b5b600185161561189d5780820291505b80810290506118ab8561185d565b9450611872565b94509492505050565b5f826118ca5760019050611985565b816118d7575f9050611985565b81600181146118ed57600281146118f757611926565b6001915050611985565b60ff8411156119095761190861135e565b5b8360020a9150848211156119205761191f61135e565b5b50611985565b5060208310610133831016604e8410600b841016171561195b5782820a9050838111156119565761195561135e565b5b611985565b6119688484846001611869565b9250905081840481111561197f5761197e61135e565b5b81810290505b9392505050565b5f61199682610f18565b91506119a183610f18565b92506119ce7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846118bb565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611a0d82610f18565b9150611a1883610f18565b925082611a2857611a276119d6565b5b828204905092915050565b5f60ff82169050919050565b5f611a4982611a33565b9150611a5483611a33565b925082611a6457611a636119d6565b5b828204905092915050565b5f611a7982611a33565b9150611a8483611a33565b9250828202611a9281611a33565b9150808214611aa457611aa361135e565b5b5092915050565b5f611ab582611a33565b9150611ac083611a33565b9250828203905060ff811115611ad957611ad861135e565b5b92915050565b5f611ae982610f18565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b1b57611b1a61135e565b5b600182019050919050565b5f611b3082611a33565b9150611b3b83611a33565b9250828201905060ff811115611b5457611b5361135e565b5b9291505056fea26469706673582212208ba993cdf08decbe8bcf0625643266161eb1060b1c32815c7dd077fa2609bd5464736f6c634300081e0033",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankMetaData.Bin instead.
var BankBin = BankMetaData.Bin

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCaller) AccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCallerSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCallerSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCallerSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankSession) Version() (string, error) {
	return _Bank.Contract.Version(&_Bank.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankCallerSession) Version() (string, error) {
	return _Bank.Contract.Version(&_Bank.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bank *BankTransactor) Reconcile(opts *bind.TransactOpts, winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Reconcile", winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bank *BankSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Reconcile(&_Bank.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Reconcile is a paid mutator transaction binding the contract method 0xfa84fd8e.
//
// Solidity: function Reconcile(address winner, address[] losers, uint256 anteWei, uint256 gameFeeWei) returns()
func (_Bank *BankTransactorSession) Reconcile(winner common.Address, losers []common.Address, anteWei *big.Int, gameFeeWei *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Reconcile(&_Bank.TransactOpts, winner, losers, anteWei, gameFeeWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// BankEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bank contract.
type BankEventLogIterator struct {
	Event *BankEventLog // Event containing the contract specifics and raw log

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
func (it *BankEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventLog)
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
		it.Event = new(BankEventLog)
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
func (it *BankEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventLog represents a EventLog event raised by the Bank contract.
type BankEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankEventLogIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankEventLogIterator{contract: _Bank.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankEventLog) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventLog)
				if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_Bank *BankFilterer) ParseEventLog(log types.Log) (*BankEventLog, error) {
	event := new(BankEventLog)
	if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
