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
	Bin: "0x6080604052348015600e575f80fd5b503360025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611bd08061005c5f395ff3fe608060405260043610610085575f3560e01c8063bb62860d11610058578063bb62860d14610111578063d2aadb3c1461013b578063e63f341f14610163578063ed21248c1461019f578063fa84fd8e146101a957610085565b80630ef678871461008957806357ea89b6146100b35780637d7b0099146100bd578063b4a99a4e146100e7575b5f80fd5b348015610094575f80fd5b5061009d6101d1565b6040516100aa9190610d6d565b60405180910390f35b6100bb610215565b005b3480156100c8575f80fd5b506100d161037f565b6040516100de9190610dc5565b60405180910390f35b3480156100f2575f80fd5b506100fb6103a2565b6040516101089190610dc5565b60405180910390f35b34801561011c575f80fd5b506101256103c7565b6040516101329190610e4e565b60405180910390f35b348015610146575f80fd5b50610161600480360381019061015c9190610ea9565b610453565b005b34801561016e575f80fd5b5061018960048036038101906101849190610ea9565b6106f8565b6040516101969190610d6d565b60405180910390f35b6101a7610796565b005b3480156101b4575f80fd5b506101cf60048036038101906101ca919061103e565b610900565b005b5f60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f57ea89b6000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516102dd9190611102565b5f60405180830381855af49150503d805f8114610315576040519150601f19603f3d011682016040523d82523d5f602084013e61031a565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61034882610ad7565b604051602001610358919061119e565b6040516020818303038152906040526040516103749190610e4e565b60405180910390a150565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600180546103d4906111ff565b80601f0160208091040260200160405190810160405280929190818152602001828054610400906111ff565b801561044b5780601f106104225761010080835404028352916020019161044b565b820191905f5260205f20905b81548152906001019060200180831161042e57829003601f168201915b505050505081565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ab575f80fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fbb62860d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516105b29190611102565b5f604051808303815f865af19150503d805f81146105eb576040519150601f19603f3d011682016040523d82523d5f602084013e6105f0565b606091505b50915091508115610623578080602001905181019061060f91906112d1565b6001908161061d91906114b5565b50610669565b6040518060400160405280600781526020017f756e6b6e6f776e000000000000000000000000000000000000000000000000008152506001908161066791906114b5565b505b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6106b25f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610b5a565b6106bb84610ad7565b60016040516020016106cf93929190611676565b6040516020818303038152906040526040516106eb9190610e4e565b60405180910390a1505050565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610751575f80fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fed21248c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161085e9190611102565b5f60405180830381855af49150503d805f8114610896576040519150601f19603f3d011682016040523d82523d5f602084013e61089b565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6108c982610ad7565b6040516020016108d9919061119e565b6040516020818303038152906040526040516108f59190610e4e565b60405180910390a150565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610958575f80fd5b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16858585856040516024016109a79493929190611799565b6040516020818303038152906040527ffa84fd8e000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610a319190611102565b5f60405180830381855af49150503d805f8114610a69576040519150601f19603f3d011682016040523d82523d5f602084013e610a6e565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610a9c82610ad7565b604051602001610aac919061119e565b604051602081830303815290604052604051610ac89190610e4e565b60405180910390a15050505050565b60608115610b1c576040518060400160405280600481526020017f74727565000000000000000000000000000000000000000000000000000000008152509050610b55565b6040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525090505b919050565b60605f602867ffffffffffffffff811115610b7857610b77610ed8565b5b6040519080825280601f01601f191660200182016040528015610baa5781602001600182028036833780820191505090505b5090505f5b6014811015610d06575f816013610bc69190611810565b6008610bd29190611843565b6002610bde91906119b3565b8573ffffffffffffffffffffffffffffffffffffffff16610bff9190611a2a565b60f81b90505f60108260f81c610c159190611a66565b60f81b90505f8160f81c6010610c2b9190611a96565b8360f81c610c399190611ad2565b60f81b9050610c4782610d10565b85856002610c559190611843565b81518110610c6657610c65611b06565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350610c9d81610d10565b856001866002610cad9190611843565b610cb79190611b33565b81518110610cc857610cc7611b06565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505050508080600101915050610baf565b5080915050919050565b5f600a8260f81c60ff161015610d3a5760308260f81c610d309190611b66565b60f81b9050610d50565b60578260f81c610d4a9190611b66565b60f81b90505b919050565b5f819050919050565b610d6781610d55565b82525050565b5f602082019050610d805f830184610d5e565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610daf82610d86565b9050919050565b610dbf81610da5565b82525050565b5f602082019050610dd85f830184610db6565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e2082610dde565b610e2a8185610de8565b9350610e3a818560208601610df8565b610e4381610e06565b840191505092915050565b5f6020820190508181035f830152610e668184610e16565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b610e8881610da5565b8114610e92575f80fd5b50565b5f81359050610ea381610e7f565b92915050565b5f60208284031215610ebe57610ebd610e77565b5b5f610ecb84828501610e95565b91505092915050565b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610f0e82610e06565b810181811067ffffffffffffffff82111715610f2d57610f2c610ed8565b5b80604052505050565b5f610f3f610e6e565b9050610f4b8282610f05565b919050565b5f67ffffffffffffffff821115610f6a57610f69610ed8565b5b602082029050602081019050919050565b5f80fd5b5f610f91610f8c84610f50565b610f36565b90508083825260208201905060208402830185811115610fb457610fb3610f7b565b5b835b81811015610fdd5780610fc98882610e95565b845260208401935050602081019050610fb6565b5050509392505050565b5f82601f830112610ffb57610ffa610ed4565b5b813561100b848260208601610f7f565b91505092915050565b61101d81610d55565b8114611027575f80fd5b50565b5f8135905061103881611014565b92915050565b5f805f806080858703121561105657611055610e77565b5b5f61106387828801610e95565b945050602085013567ffffffffffffffff81111561108457611083610e7b565b5b61109087828801610fe7565b93505060406110a18782880161102a565b92505060606110b28782880161102a565b91505092959194509250565b5f81519050919050565b5f81905092915050565b5f6110dc826110be565b6110e681856110c8565b93506110f6818560208601610df8565b80840191505092915050565b5f61110d82846110d2565b915081905092915050565b7f737563636573735b000000000000000000000000000000000000000000000000815250565b5f81905092915050565b5f61115282610dde565b61115c818561113e565b935061116c818560208601610df8565b80840191505092915050565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b5f6111a882611118565b6008820191506111b88284611148565b91506111c382611178565b60018201915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061121657607f821691505b602082108103611229576112286111d2565b5b50919050565b5f80fd5b5f67ffffffffffffffff82111561124d5761124c610ed8565b5b61125682610e06565b9050602081019050919050565b5f61127561127084611233565b610f36565b9050828152602081018484840111156112915761129061122f565b5b61129c848285610df8565b509392505050565b5f82601f8301126112b8576112b7610ed4565b5b81516112c8848260208601611263565b91505092915050565b5f602082840312156112e6576112e5610e77565b5b5f82015167ffffffffffffffff81111561130357611302610e7b565b5b61130f848285016112a4565b91505092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026113747fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611339565b61137e8683611339565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6113b96113b46113af84610d55565b611396565b610d55565b9050919050565b5f819050919050565b6113d28361139f565b6113e66113de826113c0565b848454611345565b825550505050565b5f90565b6113fa6113ee565b6114058184846113c9565b505050565b5b818110156114285761141d5f826113f2565b60018101905061140b565b5050565b601f82111561146d5761143e81611318565b6114478461132a565b81016020851015611456578190505b61146a6114628561132a565b83018261140a565b50505b505050565b5f82821c905092915050565b5f61148d5f1984600802611472565b1980831691505092915050565b5f6114a5838361147e565b9150826002028217905092915050565b6114be82610dde565b67ffffffffffffffff8111156114d7576114d6610ed8565b5b6114e182546111ff565b6114ec82828561142c565b5f60209050601f83116001811461151d575f841561150b578287015190505b611515858261149a565b86555061157c565b601f19841661152b86611318565b5f5b828110156115525784890151825560018201915060208501945060208101905061152d565b8683101561156f578489015161156b601f89168261147e565b8355505b6001600288020188555050505b505050505050565b7f636f6e74726163745b0000000000000000000000000000000000000000000000815250565b7f5d20737563636573735b00000000000000000000000000000000000000000000815250565b7f5d2076657273696f6e5b00000000000000000000000000000000000000000000815250565b5f8154611602816111ff565b61160c818661113e565b9450600182165f8114611626576001811461163b5761166d565b60ff198316865281151582028601935061166d565b61164485611318565b5f5b8381101561166557815481890152600182019150602081019050611646565b838801955050505b50505092915050565b5f61168082611584565b6009820191506116908286611148565b915061169b826115aa565b600a820191506116ab8285611148565b91506116b6826115d0565b600a820191506116c682846115f6565b91506116d182611178565b600182019150819050949350505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61171481610da5565b82525050565b5f611725838361170b565b60208301905092915050565b5f602082019050919050565b5f611747826116e2565b61175181856116ec565b935061175c836116fc565b805f5b8381101561178c578151611773888261171a565b975061177e83611731565b92505060018101905061175f565b5085935050505092915050565b5f6080820190506117ac5f830187610db6565b81810360208301526117be818661173d565b90506117cd6040830185610d5e565b6117da6060830184610d5e565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61181a82610d55565b915061182583610d55565b925082820390508181111561183d5761183c6117e3565b5b92915050565b5f61184d82610d55565b915061185883610d55565b925082820261186681610d55565b9150828204841483151761187d5761187c6117e3565b5b5092915050565b5f8160011c9050919050565b5f808291508390505b60018511156118d9578086048111156118b5576118b46117e3565b5b60018516156118c45780820291505b80810290506118d285611884565b9450611899565b94509492505050565b5f826118f157600190506119ac565b816118fe575f90506119ac565b8160018114611914576002811461191e5761194d565b60019150506119ac565b60ff8411156119305761192f6117e3565b5b8360020a915084821115611947576119466117e3565b5b506119ac565b5060208310610133831016604e8410600b84101617156119825782820a90508381111561197d5761197c6117e3565b5b6119ac565b61198f8484846001611890565b925090508184048111156119a6576119a56117e3565b5b81810290505b9392505050565b5f6119bd82610d55565b91506119c883610d55565b92506119f57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846118e2565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611a3482610d55565b9150611a3f83610d55565b925082611a4f57611a4e6119fd565b5b828204905092915050565b5f60ff82169050919050565b5f611a7082611a5a565b9150611a7b83611a5a565b925082611a8b57611a8a6119fd565b5b828204905092915050565b5f611aa082611a5a565b9150611aab83611a5a565b9250828202611ab981611a5a565b9150808214611acb57611aca6117e3565b5b5092915050565b5f611adc82611a5a565b9150611ae783611a5a565b9250828203905060ff811115611b0057611aff6117e3565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f611b3d82610d55565b9150611b4883610d55565b9250828201905080821115611b6057611b5f6117e3565b5b92915050565b5f611b7082611a5a565b9150611b7b83611a5a565b9250828201905060ff811115611b9457611b936117e3565b5b9291505056fea2646970667358221220860607ee2801f8235c41515a8e466dba4e10c8c3504f6a2fd869aee0618d127f64736f6c63430008190033",
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
