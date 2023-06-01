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
	Bin: "0x60806040523480156200001157600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060400160405280600581526020017f302e312e30000000000000000000000000000000000000000000000000000000815250600190816200009891906200031f565b5062000406565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200012157607f821691505b602082108103620001375762000136620000d9565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b60008160020a8302905092915050565b600060088302620001a47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000162565b620001b0868362000162565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620001fd620001f7620001f184620001c8565b620001d2565b620001c8565b9050919050565b6000819050919050565b6200021983620001dc565b62000231620002288262000204565b84845462000172565b825550505050565b600090565b6200024862000239565b620002558184846200020e565b505050565b5b818110156200027d57620002716000826200023e565b6001810190506200025b565b5050565b601f821115620002cc5762000296816200013d565b620002a18462000152565b81016020851015620002b1578190505b620002c9620002c08562000152565b8301826200025a565b50505b505050565b60008160020a8304905092915050565b6000620002f460001984600802620002d1565b1980831691505092915050565b60006200030f8383620002e1565b9150826002028217905092915050565b6200032a826200009f565b67ffffffffffffffff811115620003465762000345620000aa565b5b62000352825462000108565b6200035f82828562000281565b600060209050601f83116001811462000397576000841562000382578287015190505b6200038e858262000301565b865550620003fe565b601f198416620003a7866200013d565b60005b82811015620003d157848901518255600182019150602085019450602081019050620003aa565b86831015620003f15784890151620003ed601f891682620002e1565b8355505b6001600288020188555050505b505050505050565b611ddd80620004166000396000f3fe60806040526004361061008d576000357c010000000000000000000000000000000000000000000000000000000090048063bb62860d1161006b578063bb62860d146100f2578063e63f341f1461011d578063ed21248c1461015a578063fa84fd8e146101645761008d565b80630ef678871461009257806357ea89b6146100bd578063b4a99a4e146100c7575b600080fd5b34801561009e57600080fd5b506100a761018d565b6040516100b49190611119565b60405180910390f35b6100c56101d4565b005b3480156100d357600080fd5b506100dc610392565b6040516100e99190611175565b60405180910390f35b3480156100fe57600080fd5b506101076103b6565b6040516101149190611220565b60405180910390f35b34801561012957600080fd5b50610144600480360381019061013f9190611278565b610444565b6040516101519190611119565b60405180910390f35b6101626104e6565b005b34801561017057600080fd5b5061018b60048036038101906101869190611336565b6105e5565b005b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003390506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540361025b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102529061140a565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156102e057600080fd5b506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61035033610bf5565b61035983610e75565b60405160200161036a9291906114d8565b6040516020818303038152906040526040516103869190611220565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600180546103c390611558565b80601f01602080910402602001604051908101604052809291908181526020018280546103ef90611558565b801561043c5780601f106104115761010080835404028352916020019161043c565b820191906000526020600020905b81548152906001019060200180831161041f57829003601f168201915b505050505081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461049f57600080fd5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b34600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461053591906115b8565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61056633610bf5565b6105ae600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e75565b6040516020016105bf929190611638565b6040516020818303038152906040526040516105db9190611220565b60405180910390a1565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461063d57600080fd5b600082905060005b8585905081101561091057836002600088888581811061066857610667611689565b5b905060200201602081019061067d9190611278565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610871577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61074e600260008989868181106106f9576106f8611689565b5b905060200201602081019061070e9190611278565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e75565b61075786610e75565b604051602001610768929190611704565b6040516020818303038152906040526040516107849190611220565b60405180910390a1600260008787848181106107a3576107a2611689565b5b90506020020160208101906107b89190611278565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826107fe91906115b8565b915060006002600088888581811061081957610818611689565b5b905060200201602081019061082e9190611278565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506108fd565b838261087d91906115b8565b9150836002600088888581811061089757610896611689565b5b90506020020160208101906108ac9190611278565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108f59190611746565b925050819055505b80806109089061177a565b915050610645565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61093b84610e75565b61094484610e75565b61094d84610e75565b60405160200161095f93929190611834565b60405160208183030381529060405260405161097b9190611220565b60405180910390a1600081036109c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109bd90611913565b60405180910390fd5b81811015610aa95780600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a3e91906115b8565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610a6f82610e75565b604051602001610a7f91906119a5565b604051602081830303815290604052604051610a9b9190611220565b60405180910390a150610bee565b8181610ab59190611746565b905080600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b0691906115b8565b9250508190555081600260008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b7d91906115b8565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610bae82610e75565b610bb784610e75565b604051602001610bc8929190611a22565b604051602081830303815290604052604051610be49190611220565b60405180910390a1505b5050505050565b60606000602867ffffffffffffffff811115610c1457610c13611a73565b5b6040519080825280601f01601f191660200182016040528015610c465781602001600182028036833780820191505090505b50905060005b6014811015610e6b576000816013610c649190611746565b6008610c709190611aa2565b6002610c7c9190611c17565b8573ffffffffffffffffffffffffffffffffffffffff16610c9d9190611c91565b7f010000000000000000000000000000000000000000000000000000000000000002905060006010827f01000000000000000000000000000000000000000000000000000000000000009004610cf39190611ccf565b7f01000000000000000000000000000000000000000000000000000000000000000290506000817f010000000000000000000000000000000000000000000000000000000000000090046010610d499190611d00565b837f01000000000000000000000000000000000000000000000000000000000000009004610d779190611d3d565b7f0100000000000000000000000000000000000000000000000000000000000000029050610da48261101c565b85856002610db29190611aa2565b81518110610dc357610dc2611689565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610dfb8161101c565b856001866002610e0b9190611aa2565b610e1591906115b8565b81518110610e2657610e25611689565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610e639061177a565b915050610c4c565b5080915050919050565b606060008203610ebc576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611017565b600082905060005b60008214610eee578080610ed79061177a565b915050600a82610ee79190611c91565b9150610ec4565b60008167ffffffffffffffff811115610f0a57610f09611a73565b5b6040519080825280601f01601f191660200182016040528015610f3c5781602001600182028036833780820191505090505b50905060008290505b6000861461100f57600181610f5a9190611746565b90506000600a8088610f6c9190611c91565b610f769190611aa2565b87610f819190611746565b6030610f8d9190611d72565b90506000817f010000000000000000000000000000000000000000000000000000000000000002905080848481518110610fca57610fc9611689565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a886110069190611c91565b97505050610f45565b819450505050505b919050565b6000600a827f0100000000000000000000000000000000000000000000000000000000000000900460ff1610156110a6576030827f0100000000000000000000000000000000000000000000000000000000000000900461107d9190611d72565b7f01000000000000000000000000000000000000000000000000000000000000000290506110fb565b6057827f010000000000000000000000000000000000000000000000000000000000000090046110d69190611d72565b7f01000000000000000000000000000000000000000000000000000000000000000290505b919050565b6000819050919050565b61111381611100565b82525050565b600060208201905061112e600083018461110a565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061115f82611134565b9050919050565b61116f81611154565b82525050565b600060208201905061118a6000830184611166565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156111ca5780820151818401526020810190506111af565b60008484015250505050565b6000601f19601f8301169050919050565b60006111f282611190565b6111fc818561119b565b935061120c8185602086016111ac565b611215816111d6565b840191505092915050565b6000602082019050818103600083015261123a81846111e7565b905092915050565b600080fd5b600080fd5b61125581611154565b811461126057600080fd5b50565b6000813590506112728161124c565b92915050565b60006020828403121561128e5761128d611242565b5b600061129c84828501611263565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126112ca576112c96112a5565b5b8235905067ffffffffffffffff8111156112e7576112e66112aa565b5b602083019150836020820283011115611303576113026112af565b5b9250929050565b61131381611100565b811461131e57600080fd5b50565b6000813590506113308161130a565b92915050565b60008060008060006080868803121561135257611351611242565b5b600061136088828901611263565b955050602086013567ffffffffffffffff81111561138157611380611247565b5b61138d888289016112b4565b945094505060406113a088828901611321565b92505060606113b188828901611321565b9150509295509295909350565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b60006113f460128361119b565b91506113ff826113be565b602082019050919050565b60006020820190508181036000830152611423816113e7565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081905092915050565b600061146682611190565b6114708185611450565b93506114808185602086016111ac565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b60006114e38261142a565b6009820191506114f3828561145b565b91506114fe8261148c565b60098201915061150e828461145b565b9150611519826114b2565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061157057607f821691505b60208210810361158357611582611529565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006115c382611100565b91506115ce83611100565b92508282019050808211156115e6576115e5611589565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b6000611643826115ec565b600882019150611653828561145b565b915061165e82611612565b600a8201915061166e828461145b565b9150611679826114b2565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f6163636f756e742062616c616e63652000000000000000000000000000000000815250565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b600061170f826116b8565b60108201915061171f828561145b565b915061172a826116de565b60138201915061173a828461145b565b91508190509392505050565b600061175182611100565b915061175c83611100565b925082820390508181111561177457611773611589565b5b92915050565b600061178582611100565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117b7576117b6611589565b5b600182019050919050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b600061183f826117c2565b60058201915061184f828661145b565b915061185a826117e8565b600a8201915061186a828561145b565b91506118758261180e565b600682019150611885828461145b565b9150611890826114b2565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e206163636f60008201527f756e742062616c616e6365730000000000000000000000000000000000000000602082015250565b60006118fd602c8361119b565b9150611908826118a1565b604082019050919050565b6000602082019050818103600083015261192c816118f0565b9050919050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e60008201527f65725b0000000000000000000000000000000000000000000000000000000000602082015250565b600061198f602383611450565b915061199a82611933565b602382019050919050565b60006119b082611982565b91506119bc828461145b565b91506119c7826114b2565b60018201915081905092915050565b7f77696e6e65725b00000000000000000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b6000611a2d826119d6565b600782019150611a3d828561145b565b9150611a48826119fc565b600882019150611a58828461145b565b9150611a63826114b2565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000611aad82611100565b9150611ab883611100565b9250828202611ac681611100565b91508282048414831517611add57611adc611589565b5b5092915050565b6000600282049050919050565b6000808291508390505b6001851115611b3b57808604811115611b1757611b16611589565b5b6001851615611b265780820291505b8081029050611b3485611ae4565b9450611afb565b94509492505050565b600082611b545760019050611c10565b81611b625760009050611c10565b8160018114611b785760028114611b8257611bb1565b6001915050611c10565b60ff841115611b9457611b93611589565b5b8360020a915084821115611bab57611baa611589565b5b50611c10565b5060208310610133831016604e8410600b8410161715611be65782820a905083811115611be157611be0611589565b5b611c10565b611bf38484846001611af1565b92509050818404811115611c0a57611c09611589565b5b81810290505b9392505050565b6000611c2282611100565b9150611c2d83611100565b9250611c5a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611b44565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611c9c82611100565b9150611ca783611100565b925082611cb757611cb6611c62565b5b828204905092915050565b600060ff82169050919050565b6000611cda82611cc2565b9150611ce583611cc2565b925082611cf557611cf4611c62565b5b828204905092915050565b6000611d0b82611cc2565b9150611d1683611cc2565b9250828202611d2481611cc2565b9150808214611d3657611d35611589565b5b5092915050565b6000611d4882611cc2565b9150611d5383611cc2565b9250828203905060ff811115611d6c57611d6b611589565b5b92915050565b6000611d7d82611cc2565b9150611d8883611cc2565b9250828201905060ff811115611da157611da0611589565b5b9291505056fea2646970667358221220bc5d527fea8daec1cc3f2c4e78762a909953876ab575288e65ecffe3eb3e497964736f6c63430008140033",
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
