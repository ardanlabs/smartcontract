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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"API\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"SetContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611cd6806100616000396000f3fe6080604052600436106100865760003560e01c8063bb62860d11610059578063bb62860d14610116578063d2aadb3c14610141578063e63f341f1461016a578063ed21248c146101a7578063fa84fd8e146101b157610086565b80630ef678871461008b57806357ea89b6146100b65780637d7b0099146100c0578063b4a99a4e146100eb575b600080fd5b34801561009757600080fd5b506100a06101da565b6040516100ad9190610dac565b60405180910390f35b6100be610221565b005b3480156100cc57600080fd5b506100d5610390565b6040516100e29190610e08565b60405180910390f35b3480156100f757600080fd5b506101006103b4565b60405161010d9190610e08565b60405180910390f35b34801561012257600080fd5b5061012b6103da565b6040516101389190610eb3565b60405180910390f35b34801561014d57600080fd5b5061016860048036038101906101639190610f15565b610468565b005b34801561017657600080fd5b50610191600480360381019061018c9190610f15565b610717565b60405161019e9190610dac565b60405180910390f35b6101af6107ba565b005b3480156101bd57600080fd5b506101d860048036038101906101d391906110b6565b610929565b005b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f57ea89b6000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516102eb9190611180565b600060405180830381855af49150503d8060008114610326576040519150601f19603f3d011682016040523d82523d6000602084013e61032b565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61035982610b07565b604051602001610369919061121f565b6040516020818303038152906040526040516103859190610eb3565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600180546103e790611283565b80601f016020809104026020016040519081016040528092919081815260200182805461041390611283565b80156104605780601f1061043557610100808354040283529160200191610460565b820191906000526020600020905b81548152906001019060200180831161044357829003601f168201915b505050505081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104c257600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fbb62860d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516105cc9190611180565b6000604051808303816000865af19150503d8060008114610609576040519150601f19603f3d011682016040523d82523d6000602084013e61060e565b606091505b50915091508115610641578080602001905181019061062d919061135a565b6001908161063b919061154f565b50610687565b6040518060400160405280600781526020017f756e6b6e6f776e0000000000000000000000000000000000000000000000000081525060019081610685919061154f565b505b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6106d160008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610b8a565b6106da84610b07565b60016040516020016106ee93929190611716565b60405160208183030381529060405260405161070a9190610eb3565b60405180910390a1505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461077357600080fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fed21248c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516108849190611180565b600060405180830381855af49150503d80600081146108bf576040519150601f19603f3d011682016040523d82523d6000602084013e6108c4565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6108f282610b07565b604051602001610902919061121f565b60405160208183030381529060405260405161091e9190610eb3565b60405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461098357600080fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16858585856040516024016109d49493929190611841565b6040516020818303038152906040527ffa84fd8e000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610a5e9190611180565b600060405180830381855af49150503d8060008114610a99576040519150601f19603f3d011682016040523d82523d6000602084013e610a9e565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610acc82610b07565b604051602001610adc919061121f565b604051602081830303815290604052604051610af89190610eb3565b60405180910390a15050505050565b60608115610b4c576040518060400160405280600481526020017f74727565000000000000000000000000000000000000000000000000000000008152509050610b85565b6040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525090505b919050565b60606000602867ffffffffffffffff811115610ba957610ba8610f47565b5b6040519080825280601f01601f191660200182016040528015610bdb5781602001600182028036833780820191505090505b50905060005b6014811015610d43576000816013610bf991906118bc565b6008610c0591906118f0565b6002610c119190611a65565b8573ffffffffffffffffffffffffffffffffffffffff16610c329190611adf565b60f81b9050600060108260f81c610c499190611b1d565b60f81b905060008160f81c6010610c609190611b4e565b8360f81c610c6e9190611b8b565b60f81b9050610c7c82610d4d565b85856002610c8a91906118f0565b81518110610c9b57610c9a611bc0565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610cd381610d4d565b856001866002610ce391906118f0565b610ced9190611bef565b81518110610cfe57610cfd611bc0565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610d3b90611c23565b915050610be1565b5080915050919050565b6000600a8260f81c60ff161015610d785760308260f81c610d6e9190611c6b565b60f81b9050610d8e565b60578260f81c610d889190611c6b565b60f81b90505b919050565b6000819050919050565b610da681610d93565b82525050565b6000602082019050610dc16000830184610d9d565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610df282610dc7565b9050919050565b610e0281610de7565b82525050565b6000602082019050610e1d6000830184610df9565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610e5d578082015181840152602081019050610e42565b60008484015250505050565b6000601f19601f8301169050919050565b6000610e8582610e23565b610e8f8185610e2e565b9350610e9f818560208601610e3f565b610ea881610e69565b840191505092915050565b60006020820190508181036000830152610ecd8184610e7a565b905092915050565b6000604051905090565b600080fd5b600080fd5b610ef281610de7565b8114610efd57600080fd5b50565b600081359050610f0f81610ee9565b92915050565b600060208284031215610f2b57610f2a610edf565b5b6000610f3984828501610f00565b91505092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f7f82610e69565b810181811067ffffffffffffffff82111715610f9e57610f9d610f47565b5b80604052505050565b6000610fb1610ed5565b9050610fbd8282610f76565b919050565b600067ffffffffffffffff821115610fdd57610fdc610f47565b5b602082029050602081019050919050565b600080fd5b600061100661100184610fc2565b610fa7565b9050808382526020820190506020840283018581111561102957611028610fee565b5b835b81811015611052578061103e8882610f00565b84526020840193505060208101905061102b565b5050509392505050565b600082601f83011261107157611070610f42565b5b8135611081848260208601610ff3565b91505092915050565b61109381610d93565b811461109e57600080fd5b50565b6000813590506110b08161108a565b92915050565b600080600080608085870312156110d0576110cf610edf565b5b60006110de87828801610f00565b945050602085013567ffffffffffffffff8111156110ff576110fe610ee4565b5b61110b8782880161105c565b935050604061111c878288016110a1565b925050606061112d878288016110a1565b91505092959194509250565b600081519050919050565b600081905092915050565b600061115a82611139565b6111648185611144565b9350611174818560208601610e3f565b80840191505092915050565b600061118c828461114f565b915081905092915050565b7f737563636573735b000000000000000000000000000000000000000000000000815250565b600081905092915050565b60006111d382610e23565b6111dd81856111bd565b93506111ed818560208601610e3f565b80840191505092915050565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b600061122a82611197565b60088201915061123a82846111c8565b9150611245826111f9565b60018201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061129b57607f821691505b6020821081036112ae576112ad611254565b5b50919050565b600080fd5b600067ffffffffffffffff8211156112d4576112d3610f47565b5b6112dd82610e69565b9050602081019050919050565b60006112fd6112f8846112b9565b610fa7565b905082815260208101848484011115611319576113186112b4565b5b611324848285610e3f565b509392505050565b600082601f83011261134157611340610f42565b5b81516113518482602086016112ea565b91505092915050565b6000602082840312156113705761136f610edf565b5b600082015167ffffffffffffffff81111561138e5761138d610ee4565b5b61139a8482850161132c565b91505092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026114057fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826113c8565b61140f86836113c8565b95508019841693508086168417925050509392505050565b6000819050919050565b600061144c61144761144284610d93565b611427565b610d93565b9050919050565b6000819050919050565b61146683611431565b61147a61147282611453565b8484546113d5565b825550505050565b600090565b61148f611482565b61149a81848461145d565b505050565b5b818110156114be576114b3600082611487565b6001810190506114a0565b5050565b601f821115611503576114d4816113a3565b6114dd846113b8565b810160208510156114ec578190505b6115006114f8856113b8565b83018261149f565b50505b505050565b600082821c905092915050565b600061152660001984600802611508565b1980831691505092915050565b600061153f8383611515565b9150826002028217905092915050565b61155882610e23565b67ffffffffffffffff81111561157157611570610f47565b5b61157b8254611283565b6115868282856114c2565b600060209050601f8311600181146115b957600084156115a7578287015190505b6115b18582611533565b865550611619565b601f1984166115c7866113a3565b60005b828110156115ef578489015182556001820191506020850194506020810190506115ca565b8683101561160c5784890151611608601f891682611515565b8355505b6001600288020188555050505b505050505050565b7f636f6e74726163745b0000000000000000000000000000000000000000000000815250565b7f5d20737563636573735b00000000000000000000000000000000000000000000815250565b7f5d2076657273696f6e5b00000000000000000000000000000000000000000000815250565b600081546116a081611283565b6116aa81866111bd565b945060018216600081146116c557600181146116da5761170d565b60ff198316865281151582028601935061170d565b6116e3856113a3565b60005b83811015611705578154818901526001820191506020810190506116e6565b838801955050505b50505092915050565b600061172182611621565b60098201915061173182866111c8565b915061173c82611647565b600a8201915061174c82856111c8565b91506117578261166d565b600a820191506117678284611693565b9150611772826111f9565b600182019150819050949350505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6117b881610de7565b82525050565b60006117ca83836117af565b60208301905092915050565b6000602082019050919050565b60006117ee82611783565b6117f8818561178e565b93506118038361179f565b8060005b8381101561183457815161181b88826117be565b9750611826836117d6565b925050600181019050611807565b5085935050505092915050565b60006080820190506118566000830187610df9565b818103602083015261186881866117e3565b90506118776040830185610d9d565b6118846060830184610d9d565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118c782610d93565b91506118d283610d93565b92508282039050818111156118ea576118e961188d565b5b92915050565b60006118fb82610d93565b915061190683610d93565b925082820261191481610d93565b9150828204841483151761192b5761192a61188d565b5b5092915050565b60008160011c9050919050565b6000808291508390505b6001851115611989578086048111156119655761196461188d565b5b60018516156119745780820291505b808102905061198285611932565b9450611949565b94509492505050565b6000826119a25760019050611a5e565b816119b05760009050611a5e565b81600181146119c657600281146119d0576119ff565b6001915050611a5e565b60ff8411156119e2576119e161188d565b5b8360020a9150848211156119f9576119f861188d565b5b50611a5e565b5060208310610133831016604e8410600b8410161715611a345782820a905083811115611a2f57611a2e61188d565b5b611a5e565b611a41848484600161193f565b92509050818404811115611a5857611a5761188d565b5b81810290505b9392505050565b6000611a7082610d93565b9150611a7b83610d93565b9250611aa87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611992565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611aea82610d93565b9150611af583610d93565b925082611b0557611b04611ab0565b5b828204905092915050565b600060ff82169050919050565b6000611b2882611b10565b9150611b3383611b10565b925082611b4357611b42611ab0565b5b828204905092915050565b6000611b5982611b10565b9150611b6483611b10565b9250828202611b7281611b10565b9150808214611b8457611b8361188d565b5b5092915050565b6000611b9682611b10565b9150611ba183611b10565b9250828203905060ff811115611bba57611bb961188d565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611bfa82610d93565b9150611c0583610d93565b9250828201905080821115611c1d57611c1c61188d565b5b92915050565b6000611c2e82610d93565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c6057611c5f61188d565b5b600182019050919050565b6000611c7682611b10565b9150611c8183611b10565b9250828201905060ff811115611c9a57611c9961188d565b5b9291505056fea26469706673582212206780ab522eb833a65d34824e1e6401b24ae8a88232ebbd88b888f568cce401b664736f6c63430008140033",
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

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bank *BankCaller) API(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "API")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bank *BankSession) API() (common.Address, error) {
	return _Bank.Contract.API(&_Bank.CallOpts)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bank *BankCallerSession) API() (common.Address, error) {
	return _Bank.Contract.API(&_Bank.CallOpts)
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

// SetContract is a paid mutator transaction binding the contract method 0xd2aadb3c.
//
// Solidity: function SetContract(address contractAddr) returns()
func (_Bank *BankTransactor) SetContract(opts *bind.TransactOpts, contractAddr common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "SetContract", contractAddr)
}

// SetContract is a paid mutator transaction binding the contract method 0xd2aadb3c.
//
// Solidity: function SetContract(address contractAddr) returns()
func (_Bank *BankSession) SetContract(contractAddr common.Address) (*types.Transaction, error) {
	return _Bank.Contract.SetContract(&_Bank.TransactOpts, contractAddr)
}

// SetContract is a paid mutator transaction binding the contract method 0xd2aadb3c.
//
// Solidity: function SetContract(address contractAddr) returns()
func (_Bank *BankTransactorSession) SetContract(contractAddr common.Address) (*types.Transaction, error) {
	return _Bank.Contract.SetContract(&_Bank.TransactOpts, contractAddr)
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
