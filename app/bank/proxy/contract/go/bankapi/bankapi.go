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
	_ = abi.ConvertType
)

// BankapiMetaData contains all meta data concerning the Bankapi contract.
var BankapiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"API\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600581526020017f302e332e3000000000000000000000000000000000000000000000000000000081525060019081620000589190620002d9565b50620003c0565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620000e157607f821691505b602082108103620000f757620000f662000099565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620001617fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000122565b6200016d868362000122565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620001ba620001b4620001ae8462000185565b6200018f565b62000185565b9050919050565b6000819050919050565b620001d68362000199565b620001ee620001e582620001c1565b8484546200012f565b825550505050565b600090565b62000205620001f6565b62000212818484620001cb565b505050565b5b818110156200023a576200022e600082620001fb565b60018101905062000218565b5050565b601f82111562000289576200025381620000fd565b6200025e8462000112565b810160208510156200026e578190505b620002866200027d8562000112565b83018262000217565b50505b505050565b600082821c905092915050565b6000620002ae600019846008026200028e565b1980831691505092915050565b6000620002c983836200029b565b9150826002028217905092915050565b620002e4826200005f565b67ffffffffffffffff8111156200030057620002ff6200006a565b5b6200030c8254620000c8565b620003198282856200023e565b600060209050601f8311600181146200035157600084156200033c578287015190505b620003488582620002bb565b865550620003b8565b601f1984166200036186620000fd565b60005b828110156200038b5784890151825560018201915060208501945060208101905062000364565b86831015620003ab5784890151620003a7601f8916826200029b565b8355505b6001600288020188555050505b505050505050565b611af880620003d06000396000f3fe6080604052600436106100555760003560e01c806357ea89b61461005a5780637d7b009914610064578063b4a99a4e1461008f578063bb62860d146100ba578063ed21248c146100e5578063fa84fd8e146100ef575b600080fd5b610062610118565b005b34801561007057600080fd5b506100796102db565b6040516100869190610dfa565b60405180910390f35b34801561009b57600080fd5b506100a46102ff565b6040516100b19190610dfa565b60405180910390f35b3480156100c657600080fd5b506100cf610325565b6040516100dc9190610ea5565b60405180910390f35b6100ed6103b3565b005b3480156100fb57600080fd5b5061011660048036038101906101119190611085565b6104b2565b005b60003390506000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540361019f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019690611154565b60405180910390fd5b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610229573d6000803e3d6000fd5b506000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61029933610a28565b6102a283610beb565b6040516020016102b3929190611222565b6040516020818303038152906040526040516102cf9190610ea5565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60018054610332906112a2565b80601f016020809104026020016040519081016040528092919081815260200182805461035e906112a2565b80156103ab5780601f10610380576101008083540402835291602001916103ab565b820191906000526020600020905b81548152906001019060200180831161038e57829003601f168201915b505050505081565b34600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104029190611302565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61043333610a28565b61047b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610beb565b60405160200161048c929190611382565b6040516020818303038152906040526040516104a89190610ea5565b60405180910390a1565b600082905060005b84518110156107425783600360008784815181106104db576104da6113d3565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156106b0577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6105a76003600088858151811061055f5761055e6113d3565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610beb565b6105b086610beb565b6040516020016105c192919061144e565b6040516020818303038152906040526040516105dd9190610ea5565b60405180910390a1600360008683815181106105fc576105fb6113d3565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548261064a9190611302565b9150600060036000878481518110610665576106646113d3565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061072f565b83826106bc9190611302565b915083600360008784815181106106d6576106d56113d3565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107279190611490565b925050819055505b808061073a906114c4565b9150506104ba565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61076d84610beb565b61077684610beb565b61077f84610beb565b6040516020016107919392919061157e565b6040516020818303038152906040526040516107ad9190610ea5565b60405180910390a1600081036107f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ef9061165d565b60405180910390fd5b818110156108dc578060036000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108719190611302565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6108a282610beb565b6040516020016108b291906116ef565b6040516020818303038152906040526040516108ce9190610ea5565b60405180910390a150610a22565b81816108e89190611490565b905080600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109399190611302565b925050819055508160036000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109b19190611302565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6109e282610beb565b6109eb84610beb565b6040516020016109fc92919061176c565b604051602081830303815290604052604051610a189190610ea5565b60405180910390a1505b50505050565b60606000602867ffffffffffffffff811115610a4757610a46610f0c565b5b6040519080825280601f01601f191660200182016040528015610a795781602001600182028036833780820191505090505b50905060005b6014811015610be1576000816013610a979190611490565b6008610aa391906117bd565b6002610aaf9190611932565b8573ffffffffffffffffffffffffffffffffffffffff16610ad091906119ac565b60f81b9050600060108260f81c610ae791906119ea565b60f81b905060008160f81c6010610afe9190611a1b565b8360f81c610b0c9190611a58565b60f81b9050610b1a82610d73565b85856002610b2891906117bd565b81518110610b3957610b386113d3565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610b7181610d73565b856001866002610b8191906117bd565b610b8b9190611302565b81518110610b9c57610b9b6113d3565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610bd9906114c4565b915050610a7f565b5080915050919050565b606060008203610c32576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610d6e565b600082905060005b60008214610c64578080610c4d906114c4565b915050600a82610c5d91906119ac565b9150610c3a565b60008167ffffffffffffffff811115610c8057610c7f610f0c565b5b6040519080825280601f01601f191660200182016040528015610cb25781602001600182028036833780820191505090505b50905060008290505b60008614610d6657600181610cd09190611490565b90506000600a8088610ce291906119ac565b610cec91906117bd565b87610cf79190611490565b6030610d039190611a8d565b905060008160f81b905080848481518110610d2157610d206113d3565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610d5d91906119ac565b97505050610cbb565b819450505050505b919050565b6000600a8260f81c60ff161015610d9e5760308260f81c610d949190611a8d565b60f81b9050610db4565b60578260f81c610dae9190611a8d565b60f81b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610de482610db9565b9050919050565b610df481610dd9565b82525050565b6000602082019050610e0f6000830184610deb565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610e4f578082015181840152602081019050610e34565b60008484015250505050565b6000601f19601f8301169050919050565b6000610e7782610e15565b610e818185610e20565b9350610e91818560208601610e31565b610e9a81610e5b565b840191505092915050565b60006020820190508181036000830152610ebf8184610e6c565b905092915050565b6000604051905090565b600080fd5b600080fd5b610ee481610dd9565b8114610eef57600080fd5b50565b600081359050610f0181610edb565b92915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f4482610e5b565b810181811067ffffffffffffffff82111715610f6357610f62610f0c565b5b80604052505050565b6000610f76610ec7565b9050610f828282610f3b565b919050565b600067ffffffffffffffff821115610fa257610fa1610f0c565b5b602082029050602081019050919050565b600080fd5b6000610fcb610fc684610f87565b610f6c565b90508083825260208201905060208402830185811115610fee57610fed610fb3565b5b835b8181101561101757806110038882610ef2565b845260208401935050602081019050610ff0565b5050509392505050565b600082601f83011261103657611035610f07565b5b8135611046848260208601610fb8565b91505092915050565b6000819050919050565b6110628161104f565b811461106d57600080fd5b50565b60008135905061107f81611059565b92915050565b6000806000806080858703121561109f5761109e610ed1565b5b60006110ad87828801610ef2565b945050602085013567ffffffffffffffff8111156110ce576110cd610ed6565b5b6110da87828801611021565b93505060406110eb87828801611070565b92505060606110fc87828801611070565b91505092959194509250565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b600061113e601283610e20565b915061114982611108565b602082019050919050565b6000602082019050818103600083015261116d81611131565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081905092915050565b60006111b082610e15565b6111ba818561119a565b93506111ca818560208601610e31565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b600061122d82611174565b60098201915061123d82856111a5565b9150611248826111d6565b60098201915061125882846111a5565b9150611263826111fc565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806112ba57607f821691505b6020821081036112cd576112cc611273565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061130d8261104f565b91506113188361104f565b92508282019050808211156113305761132f6112d3565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b600061138d82611336565b60088201915061139d82856111a5565b91506113a88261135c565b600a820191506113b882846111a5565b91506113c3826111fc565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f6163636f756e742062616c616e63652000000000000000000000000000000000815250565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b600061145982611402565b60108201915061146982856111a5565b915061147482611428565b60138201915061148482846111a5565b91508190509392505050565b600061149b8261104f565b91506114a68361104f565b92508282039050818111156114be576114bd6112d3565b5b92915050565b60006114cf8261104f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611501576115006112d3565b5b600182019050919050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b60006115898261150c565b60058201915061159982866111a5565b91506115a482611532565b600a820191506115b482856111a5565b91506115bf82611558565b6006820191506115cf82846111a5565b91506115da826111fc565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e206163636f60008201527f756e742062616c616e6365730000000000000000000000000000000000000000602082015250565b6000611647602c83610e20565b9150611652826115eb565b604082019050919050565b600060208201905081810360008301526116768161163a565b9050919050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e60008201527f65725b0000000000000000000000000000000000000000000000000000000000602082015250565b60006116d960238361119a565b91506116e48261167d565b602382019050919050565b60006116fa826116cc565b915061170682846111a5565b9150611711826111fc565b60018201915081905092915050565b7f77696e6e65725b00000000000000000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b600061177782611720565b60078201915061178782856111a5565b915061179282611746565b6008820191506117a282846111a5565b91506117ad826111fc565b6001820191508190509392505050565b60006117c88261104f565b91506117d38361104f565b92508282026117e18161104f565b915082820484148315176117f8576117f76112d3565b5b5092915050565b60008160011c9050919050565b6000808291508390505b600185111561185657808604811115611832576118316112d3565b5b60018516156118415780820291505b808102905061184f856117ff565b9450611816565b94509492505050565b60008261186f576001905061192b565b8161187d576000905061192b565b8160018114611893576002811461189d576118cc565b600191505061192b565b60ff8411156118af576118ae6112d3565b5b8360020a9150848211156118c6576118c56112d3565b5b5061192b565b5060208310610133831016604e8410600b84101617156119015782820a9050838111156118fc576118fb6112d3565b5b61192b565b61190e848484600161180c565b92509050818404811115611925576119246112d3565b5b81810290505b9392505050565b600061193d8261104f565b91506119488361104f565b92506119757fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461185f565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006119b78261104f565b91506119c28361104f565b9250826119d2576119d161197d565b5b828204905092915050565b600060ff82169050919050565b60006119f5826119dd565b9150611a00836119dd565b925082611a1057611a0f61197d565b5b828204905092915050565b6000611a26826119dd565b9150611a31836119dd565b9250828202611a3f816119dd565b9150808214611a5157611a506112d3565b5b5092915050565b6000611a63826119dd565b9150611a6e836119dd565b9250828203905060ff811115611a8757611a866112d3565b5b92915050565b6000611a98826119dd565b9150611aa3836119dd565b9250828201905060ff811115611abc57611abb6112d3565b5b9291505056fea26469706673582212204f53870e0e1b9bcb72893d6872eda80f8da5399215aee46a6ea97aba8240780464736f6c63430008140033",
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
	parsed, err := BankapiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bankapi *BankapiCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bankapi.contract.Call(opts, &out, "Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bankapi *BankapiSession) Version() (string, error) {
	return _Bankapi.Contract.Version(&_Bankapi.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bankapi *BankapiCallerSession) Version() (string, error) {
	return _Bankapi.Contract.Version(&_Bankapi.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankapi *BankapiTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankapi *BankapiSession) Deposit() (*types.Transaction, error) {
	return _Bankapi.Contract.Deposit(&_Bankapi.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankapi *BankapiTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bankapi.Contract.Deposit(&_Bankapi.TransactOpts)
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

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankapi *BankapiTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankapi *BankapiSession) Withdraw() (*types.Transaction, error) {
	return _Bankapi.Contract.Withdraw(&_Bankapi.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankapi *BankapiTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bankapi.Contract.Withdraw(&_Bankapi.TransactOpts)
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
