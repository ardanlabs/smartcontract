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
	Bin: "0x608060405234801561000f575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060400160405280600581526020017f302e312e300000000000000000000000000000000000000000000000000000008152506001908161009391906102d3565b506103a2565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061011457607f821691505b602082108103610127576101266100d0565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101897fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261014e565b610193868361014e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6101d76101d26101cd846101ab565b6101b4565b6101ab565b9050919050565b5f819050919050565b6101f0836101bd565b6102046101fc826101de565b84845461015a565b825550505050565b5f90565b61021861020c565b6102238184846101e7565b505050565b5b818110156102465761023b5f82610210565b600181019050610229565b5050565b601f82111561028b5761025c8161012d565b6102658461013f565b81016020851015610274578190505b6102886102808561013f565b830182610228565b50505b505050565b5f82821c905092915050565b5f6102ab5f1984600802610290565b1980831691505092915050565b5f6102c3838361029c565b9150826002028217905092915050565b6102dc82610099565b67ffffffffffffffff8111156102f5576102f46100a3565b5b6102ff82546100fd565b61030a82828561024a565b5f60209050601f83116001811461033b575f8415610329578287015190505b61033385826102b8565b86555061039a565b601f1984166103498661012d565b5f5b828110156103705784890151825560018201915060208501945060208101905061034b565b8683101561038d5784890151610389601f89168261029c565b8355505b6001600288020188555050505b505050505050565b611ba2806103af5f395ff3fe60806040526004361061006f575f3560e01c8063bb62860d1161004d578063bb62860d146100d1578063e63f341f146100fb578063ed21248c14610137578063fa84fd8e146101415761006f565b80630ef678871461007357806357ea89b61461009d578063b4a99a4e146100a7575b5f80fd5b34801561007e575f80fd5b50610087610169565b6040516100949190610f28565b60405180910390f35b6100a56101ad565b005b3480156100b2575f80fd5b506100bb610363565b6040516100c89190610f80565b60405180910390f35b3480156100dc575f80fd5b506100e5610386565b6040516100f29190611023565b60405180910390f35b348015610106575f80fd5b50610121600480360381019061011c9190611075565b610412565b60405161012e9190610f28565b60405180910390f35b61013f6104af565b005b34801561014c575f80fd5b506101676004803603810190610162919061112b565b6105a9565b005b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905090565b5f3390505f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610230576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610227906111f9565b60405180910390fd5b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490508173ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156102b4573d5f803e3d5ffd5b505f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61032133610b97565b61032a83610d4d565b60405160200161033b9291906112c3565b6040516020818303038152906040526040516103579190611023565b60405180910390a15050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001805461039390611340565b80601f01602080910402602001604051908101604052809291908181526020018280546103bf90611340565b801561040a5780601f106103e15761010080835404028352916020019161040a565b820191905f5260205f20905b8154815290600101906020018083116103ed57829003601f168201915b505050505081565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461046a575f80fd5b60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b3460025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546104fb919061139d565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61052c33610b97565b61057260025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610d4d565b60405160200161058392919061141c565b60405160208183030381529060405260405161059f9190611023565b60405180910390a1565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105ff575f80fd5b5f8290505f5b858590508110156108be578360025f8888858181106106275761062661146c565b5b905060200201602081019061063c9190611075565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610828577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61070a60025f8989868181106106b6576106b561146c565b5b90506020020160208101906106cb9190611075565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610d4d565b61071386610d4d565b6040516020016107249291906114e5565b6040516020818303038152906040526040516107409190611023565b60405180910390a160025f87878481811061075e5761075d61146c565b5b90506020020160208101906107739190611075565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054826107b8919061139d565b91505f60025f8888858181106107d1576107d061146c565b5b90506020020160208101906107e69190611075565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055506108b1565b8382610834919061139d565b91508360025f88888581811061084d5761084c61146c565b5b90506020020160208101906108629190611075565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546108a99190611526565b925050819055505b8080600101915050610605565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6108e984610d4d565b6108f284610d4d565b6108fb84610d4d565b60405160200161090d939291906115cb565b6040516020818303038152906040526040516109299190611023565b60405180910390a15f8103610973576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096a906116a7565b60405180910390fd5b81811015610a52578060025f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109e7919061139d565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610a1882610d4d565b604051602001610a289190611735565b604051602081830303815290604052604051610a449190611023565b60405180910390a150610b90565b8181610a5e9190611526565b90508060025f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610aac919061139d565b925050819055508160025f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610b1f919061139d565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610b5082610d4d565b610b5984610d4d565b604051602001610b6a9291906117b1565b604051602081830303815290604052604051610b869190611023565b60405180910390a1505b5050505050565b60605f602867ffffffffffffffff811115610bb557610bb4611801565b5b6040519080825280601f01601f191660200182016040528015610be75781602001600182028036833780820191505090505b5090505f5b6014811015610d43575f816013610c039190611526565b6008610c0f919061182e565b6002610c1b919061199e565b8573ffffffffffffffffffffffffffffffffffffffff16610c3c9190611a15565b60f81b90505f60108260f81c610c529190611a51565b60f81b90505f8160f81c6010610c689190611a81565b8360f81c610c769190611abd565b60f81b9050610c8482610ecb565b85856002610c92919061182e565b81518110610ca357610ca261146c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350610cda81610ecb565b856001866002610cea919061182e565b610cf4919061139d565b81518110610d0557610d0461146c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505050508080600101915050610bec565b5080915050919050565b60605f8203610d93576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610ec6565b5f8290505f5b5f8214610dc2578080610dab90611af1565b915050600a82610dbb9190611a15565b9150610d99565b5f8167ffffffffffffffff811115610ddd57610ddc611801565b5b6040519080825280601f01601f191660200182016040528015610e0f5781602001600182028036833780820191505090505b5090505f8290505b5f8614610ebe57600181610e2b9190611526565b90505f600a8088610e3c9190611a15565b610e46919061182e565b87610e519190611526565b6030610e5d9190611b38565b90505f8160f81b905080848481518110610e7a57610e7961146c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a88610eb59190611a15565b97505050610e17565b819450505050505b919050565b5f600a8260f81c60ff161015610ef55760308260f81c610eeb9190611b38565b60f81b9050610f0b565b60578260f81c610f059190611b38565b60f81b90505b919050565b5f819050919050565b610f2281610f10565b82525050565b5f602082019050610f3b5f830184610f19565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f6a82610f41565b9050919050565b610f7a81610f60565b82525050565b5f602082019050610f935f830184610f71565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610fd0578082015181840152602081019050610fb5565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610ff582610f99565b610fff8185610fa3565b935061100f818560208601610fb3565b61101881610fdb565b840191505092915050565b5f6020820190508181035f83015261103b8184610feb565b905092915050565b5f80fd5b5f80fd5b61105481610f60565b811461105e575f80fd5b50565b5f8135905061106f8161104b565b92915050565b5f6020828403121561108a57611089611043565b5b5f61109784828501611061565b91505092915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126110c1576110c06110a0565b5b8235905067ffffffffffffffff8111156110de576110dd6110a4565b5b6020830191508360208202830111156110fa576110f96110a8565b5b9250929050565b61110a81610f10565b8114611114575f80fd5b50565b5f8135905061112581611101565b92915050565b5f805f805f6080868803121561114457611143611043565b5b5f61115188828901611061565b955050602086013567ffffffffffffffff81111561117257611171611047565b5b61117e888289016110ac565b9450945050604061119188828901611117565b92505060606111a288828901611117565b9150509295509295909350565b7f6e6f7420656e6f7567682062616c616e636500000000000000000000000000005f82015250565b5f6111e3601283610fa3565b91506111ee826111af565b602082019050919050565b5f6020820190508181035f830152611210816111d7565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b5f81905092915050565b5f61125182610f99565b61125b818561123d565b935061126b818560208601610fb3565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b5f6112cd82611217565b6009820191506112dd8285611247565b91506112e882611277565b6009820191506112f88284611247565b91506113038261129d565b6001820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061135757607f821691505b60208210810361136a57611369611313565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6113a782610f10565b91506113b283610f10565b92508282019050808211156113ca576113c9611370565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b5f611426826113d0565b6008820191506114368285611247565b9150611441826113f6565b600a820191506114518284611247565b915061145c8261129d565b6001820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f6163636f756e742062616c616e63652000000000000000000000000000000000815250565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b5f6114ef82611499565b6010820191506114ff8285611247565b915061150a826114bf565b60138201915061151a8284611247565b91508190509392505050565b5f61153082610f10565b915061153b83610f10565b925082820390508181111561155357611552611370565b5b92915050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b5f6115d582611559565b6005820191506115e58286611247565b91506115f08261157f565b600a820191506116008285611247565b915061160b826115a5565b60068201915061161b8284611247565b91506116268261129d565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e206163636f5f8201527f756e742062616c616e6365730000000000000000000000000000000000000000602082015250565b5f611691602c83610fa3565b915061169c82611637565b604082019050919050565b5f6020820190508181035f8301526116be81611685565b9050919050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e5f8201527f65725b0000000000000000000000000000000000000000000000000000000000602082015250565b5f61171f60238361123d565b915061172a826116c5565b602382019050919050565b5f61173f82611713565b915061174b8284611247565b91506117568261129d565b60018201915081905092915050565b7f77696e6e65725b00000000000000000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b5f6117bb82611765565b6007820191506117cb8285611247565b91506117d68261178b565b6008820191506117e68284611247565b91506117f18261129d565b6001820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f61183882610f10565b915061184383610f10565b925082820261185181610f10565b9150828204841483151761186857611867611370565b5b5092915050565b5f8160011c9050919050565b5f808291508390505b60018511156118c4578086048111156118a05761189f611370565b5b60018516156118af5780820291505b80810290506118bd8561186f565b9450611884565b94509492505050565b5f826118dc5760019050611997565b816118e9575f9050611997565b81600181146118ff576002811461190957611938565b6001915050611997565b60ff84111561191b5761191a611370565b5b8360020a91508482111561193257611931611370565b5b50611997565b5060208310610133831016604e8410600b841016171561196d5782820a90508381111561196857611967611370565b5b611997565b61197a848484600161187b565b9250905081840481111561199157611990611370565b5b81810290505b9392505050565b5f6119a882610f10565b91506119b383610f10565b92506119e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846118cd565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611a1f82610f10565b9150611a2a83610f10565b925082611a3a57611a396119e8565b5b828204905092915050565b5f60ff82169050919050565b5f611a5b82611a45565b9150611a6683611a45565b925082611a7657611a756119e8565b5b828204905092915050565b5f611a8b82611a45565b9150611a9683611a45565b9250828202611aa481611a45565b9150808214611ab657611ab5611370565b5b5092915050565b5f611ac782611a45565b9150611ad283611a45565b9250828203905060ff811115611aeb57611aea611370565b5b92915050565b5f611afb82610f10565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b2d57611b2c611370565b5b600182019050919050565b5f611b4282611a45565b9150611b4d83611a45565b9250828201905060ff811115611b6657611b65611370565b5b9291505056fea26469706673582212204fe35accda6765044da1fa847c7b2c45f52f4af3ecd28d130d0d4c72fb00699964736f6c63430008190033",
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
