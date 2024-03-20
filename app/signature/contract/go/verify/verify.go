// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verify

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

// VerifyMetaData contains all meta data concerning the Verify contract.
var VerifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MatchSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50610a238061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063406b552314610038578063abbddf3b14610068575b5f80fd5b610052600480360381019061004d9190610643565b610098565b60405161005f91906106fd565b60405180910390f35b610082600480360381019061007d9190610643565b610171565b60405161008f9190610725565b60405180910390f35b5f808686866040516020016100af939291906107ad565b6040516020818303038152906040528051906020012090505f806100d483878761020a565b91509150805f0151156101225780602001516040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161011991906107e9565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036101615760019350505050610168565b5f93505050505b95945050505050565b5f80868686604051602001610188939291906107ad565b6040516020818303038152906040528051906020012090505f806101ad83878761020a565b91509150805f0151156101fb5780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f291906107e9565b60405180910390fd5b81935050505095945050505050565b5f6102136103f1565b60418484905014610266575f61025d6040518060400160405280601881526020017f696e76616c6964207369676e6174757265206c656e6774680000000000000000815250610398565b91509150610390565b5f6040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090505f81876040516020016102b3929190610876565b6040516020818303038152906040528051906020012090505f86865f906020926102df939291906108a5565b906102ea91906108f5565b90505f8787602090604092610301939291906108a5565b9061030c91906108f5565b90505f8888604081811061032357610322610953565b5b9050013560f81c60f81b60f81c90506001848285856040515f815260200160405260405161035494939291906109aa565b6020604051602081039080840390855afa158015610374573d5f803e3d5ffd5b505050602060405103516103866103bf565b9650965050505050505b935093915050565b6103a06103f1565b6040518060400160405280600115158152602001838152509050919050565b6103c76103f1565b60405180604001604052805f1515815260200160405180602001604052805f815250815250905090565b60405180604001604052805f15158152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61046b82610425565b810181811067ffffffffffffffff8211171561048a57610489610435565b5b80604052505050565b5f61049c61040c565b90506104a88282610462565b919050565b5f67ffffffffffffffff8211156104c7576104c6610435565b5b6104d082610425565b9050602081019050919050565b828183375f83830152505050565b5f6104fd6104f8846104ad565b610493565b90508281526020810184848401111561051957610518610421565b5b6105248482856104dd565b509392505050565b5f82601f8301126105405761053f61041d565b5b81356105508482602086016104eb565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61058282610559565b9050919050565b61059281610578565b811461059c575f80fd5b50565b5f813590506105ad81610589565b92915050565b5f819050919050565b6105c5816105b3565b81146105cf575f80fd5b50565b5f813590506105e0816105bc565b92915050565b5f80fd5b5f80fd5b5f8083601f8401126106035761060261041d565b5b8235905067ffffffffffffffff8111156106205761061f6105e6565b5b60208301915083600182028301111561063c5761063b6105ea565b5b9250929050565b5f805f805f6080868803121561065c5761065b610415565b5b5f86013567ffffffffffffffff81111561067957610678610419565b5b6106858882890161052c565b95505060206106968882890161059f565b94505060406106a7888289016105d2565b935050606086013567ffffffffffffffff8111156106c8576106c7610419565b5b6106d4888289016105ee565b92509250509295509295909350565b5f8115159050919050565b6106f7816106e3565b82525050565b5f6020820190506107105f8301846106ee565b92915050565b61071f81610578565b82525050565b5f6020820190506107385f830184610716565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6107708261073e565b61077a8185610748565b935061078a818560208601610758565b61079381610425565b840191505092915050565b6107a7816105b3565b82525050565b5f6060820190508181035f8301526107c58186610766565b90506107d46020830185610716565b6107e1604083018461079e565b949350505050565b5f6020820190508181035f8301526108018184610766565b905092915050565b5f81519050919050565b5f81905092915050565b5f61082782610809565b6108318185610813565b9350610841818560208601610758565b80840191505092915050565b5f819050919050565b5f819050919050565b61087061086b8261084d565b610856565b82525050565b5f610881828561081d565b915061088d828461085f565b6020820191508190509392505050565b5f80fd5b5f80fd5b5f80858511156108b8576108b761089d565b5b838611156108c9576108c86108a1565b5b6001850283019150848603905094509492505050565b5f82905092915050565b5f82821b905092915050565b5f61090083836108df565b8261090b813561084d565b9250602082101561094b576109467fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff836020036008026108e9565b831692505b505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b6109898161084d565b82525050565b5f60ff82169050919050565b6109a48161098f565b82525050565b5f6080820190506109bd5f830187610980565b6109ca602083018661099b565b6109d76040830185610980565b6109e46060830184610980565b9594505050505056fea26469706673582212204100027ab672848b53a429e0e8db596ab7c5b08c317081c37e6231ea927f945c64736f6c63430008190033",
}

// VerifyABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyMetaData.ABI instead.
var VerifyABI = VerifyMetaData.ABI

// VerifyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifyMetaData.Bin instead.
var VerifyBin = VerifyMetaData.Bin

// DeployVerify deploys a new Ethereum contract, binding an instance of Verify to it.
func DeployVerify(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verify, error) {
	parsed, err := VerifyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verify{VerifyCaller: VerifyCaller{contract: contract}, VerifyTransactor: VerifyTransactor{contract: contract}, VerifyFilterer: VerifyFilterer{contract: contract}}, nil
}

// Verify is an auto generated Go binding around an Ethereum contract.
type Verify struct {
	VerifyCaller     // Read-only binding to the contract
	VerifyTransactor // Write-only binding to the contract
	VerifyFilterer   // Log filterer for contract events
}

// VerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifySession struct {
	Contract     *Verify           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyCallerSession struct {
	Contract *VerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyTransactorSession struct {
	Contract     *VerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyRaw struct {
	Contract *Verify // Generic contract binding to access the raw methods on
}

// VerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyCallerRaw struct {
	Contract *VerifyCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyTransactorRaw struct {
	Contract *VerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerify creates a new instance of Verify, bound to a specific deployed contract.
func NewVerify(address common.Address, backend bind.ContractBackend) (*Verify, error) {
	contract, err := bindVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verify{VerifyCaller: VerifyCaller{contract: contract}, VerifyTransactor: VerifyTransactor{contract: contract}, VerifyFilterer: VerifyFilterer{contract: contract}}, nil
}

// NewVerifyCaller creates a new read-only instance of Verify, bound to a specific deployed contract.
func NewVerifyCaller(address common.Address, caller bind.ContractCaller) (*VerifyCaller, error) {
	contract, err := bindVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyCaller{contract: contract}, nil
}

// NewVerifyTransactor creates a new write-only instance of Verify, bound to a specific deployed contract.
func NewVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyTransactor, error) {
	contract, err := bindVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyTransactor{contract: contract}, nil
}

// NewVerifyFilterer creates a new log filterer instance of Verify, bound to a specific deployed contract.
func NewVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyFilterer, error) {
	contract, err := bindVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyFilterer{contract: contract}, nil
}

// bindVerify binds a generic wrapper to an already deployed contract.
func bindVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verify *VerifyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verify.Contract.VerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verify *VerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verify.Contract.VerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verify *VerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verify.Contract.VerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verify *VerifyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verify *VerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verify *VerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verify.Contract.contract.Transact(opts, method, params...)
}

// Address is a free data retrieval call binding the contract method 0xabbddf3b.
//
// Solidity: function Address(string id, address participant, uint256 nonce, bytes sig) pure returns(address)
func (_Verify *VerifyCaller) Address(opts *bind.CallOpts, id string, participant common.Address, nonce *big.Int, sig []byte) (common.Address, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "Address", id, participant, nonce, sig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Address is a free data retrieval call binding the contract method 0xabbddf3b.
//
// Solidity: function Address(string id, address participant, uint256 nonce, bytes sig) pure returns(address)
func (_Verify *VerifySession) Address(id string, participant common.Address, nonce *big.Int, sig []byte) (common.Address, error) {
	return _Verify.Contract.Address(&_Verify.CallOpts, id, participant, nonce, sig)
}

// Address is a free data retrieval call binding the contract method 0xabbddf3b.
//
// Solidity: function Address(string id, address participant, uint256 nonce, bytes sig) pure returns(address)
func (_Verify *VerifyCallerSession) Address(id string, participant common.Address, nonce *big.Int, sig []byte) (common.Address, error) {
	return _Verify.Contract.Address(&_Verify.CallOpts, id, participant, nonce, sig)
}

// MatchSender is a free data retrieval call binding the contract method 0x406b5523.
//
// Solidity: function MatchSender(string id, address participant, uint256 nonce, bytes sig) view returns(bool)
func (_Verify *VerifyCaller) MatchSender(opts *bind.CallOpts, id string, participant common.Address, nonce *big.Int, sig []byte) (bool, error) {
	var out []interface{}
	err := _Verify.contract.Call(opts, &out, "MatchSender", id, participant, nonce, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MatchSender is a free data retrieval call binding the contract method 0x406b5523.
//
// Solidity: function MatchSender(string id, address participant, uint256 nonce, bytes sig) view returns(bool)
func (_Verify *VerifySession) MatchSender(id string, participant common.Address, nonce *big.Int, sig []byte) (bool, error) {
	return _Verify.Contract.MatchSender(&_Verify.CallOpts, id, participant, nonce, sig)
}

// MatchSender is a free data retrieval call binding the contract method 0x406b5523.
//
// Solidity: function MatchSender(string id, address participant, uint256 nonce, bytes sig) view returns(bool)
func (_Verify *VerifyCallerSession) MatchSender(id string, participant common.Address, nonce *big.Int, sig []byte) (bool, error) {
	return _Verify.Contract.MatchSender(&_Verify.CallOpts, id, participant, nonce, sig)
}
