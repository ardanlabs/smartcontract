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
)

// SimplecoinMetaData contains all meta data concerning the Simplecoin contract.
var SimplecoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"EventFrozenAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"EventTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizedAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowanceAmount\",\"type\":\"uint256\"}],\"name\":\"Authorize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"CoinBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"freeze\",\"type\":\"bool\"}],\"name\":\"FreezeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FrozenAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001d8338038062001d8383398181016040528101906200003791906200022c565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000a960008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682620000b060201b60201c565b50620002f6565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146200010957600080fd5b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546200015a91906200028d565b925050819055508173ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a4583604051620001e09190620002d9565b60405180910390a35050565b600080fd5b6000819050919050565b6200020681620001f1565b81146200021257600080fd5b50565b6000815190506200022681620001fb565b92915050565b600060208284031215620002455762000244620001ec565b5b6000620002558482850162000215565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200029a82620001f1565b9150620002a783620001f1565b9250828201905080821115620002c257620002c16200025e565b5b92915050565b620002d381620001f1565b82525050565b6000602082019050620002f06000830184620002c8565b92915050565b611a7d80620003066000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638bdf1534116100665780638bdf153414610130578063ab31471e14610160578063b4a99a4e14610190578063c0d84ce5146101ae578063d16a7a4b146101ca57610093565b80630f6798a5146100985780636386a433146100b457806369ca02dd146100e45780636dcdd00f14610100575b600080fd5b6100b260048036038101906100ad919061112c565b6101e6565b005b6100ce60048036038101906100c9919061112c565b61031d565b6040516100db9190611187565b60405180910390f35b6100fe60048036038101906100f9919061112c565b6103aa565b005b61011a600480360381019061011591906111a2565b61058c565b60405161012791906111f1565b60405180910390f35b61014a6004803603810190610145919061120c565b6105b1565b60405161015791906111f1565b60405180910390f35b61017a6004803603810190610175919061120c565b6105c9565b6040516101879190611187565b60405180910390f35b6101986105e9565b6040516101a59190611248565b60405180910390f35b6101c860048036038101906101c39190611263565b61060d565b005b6101e460048036038101906101df91906112e2565b610883565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461023e57600080fd5b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461028d9190611351565b925050819055508173ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a458360405161031191906111f1565b60405180910390a35050565b600081600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001905092915050565b60006103b733848461096f565b90508060000151156104045780602001516040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103fb9190611415565b60405180910390fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104539190611437565b9250508190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104a99190611351565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6104da83610bf3565b6104e385610d7b565b6104ec33610d7b565b6040516020016104fe93929190611519565b60405160208183030381529060405260405161051a9190611415565b60405180910390a18273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a458460405161057f91906111f1565b60405180910390a3505050565b6002602052816000526040600020602052806000526040600020600091509150505481565b60016020528060005260406000206000915090505481565b60036020528060005260406000206000915054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061061a84848461096f565b90508060000151156106675780602001516040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065e9190611415565b60405180910390fd5b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106b69190611437565b9250508190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461070c9190611351565b9250508190555081600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461079f9190611437565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6107d083610bf3565b6107d985610d7b565b6107e287610d7b565b6040516020016107f49392919061159d565b6040516020818303038152906040526040516108109190611415565b60405180910390a18273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a458460405161087591906111f1565b60405180910390a350505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108db57600080fd5b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f3ed8718a0f403576721b15da1095392a64cc33255a766b58bf4d5a0e116bcb1d82826040516109639291906115fb565b60405180910390a15050565b610977611077565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036109f0576109e96040518060400160405280601f81526020017f63616e27742073656e64206d6f6e657920746f20616464726573732030783000815250610f3e565b9050610bec565b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610b0357610afc6040518060400160405280601781526020017f696e737566666963656e742066756e6473202062616c3a000000000000000000815250610ab8600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610bf3565b6040518060400160405280600a81526020017f2020616d6f756e743a2000000000000000000000000000000000000000000000815250610af786610bf3565b610f65565b9050610bec565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610b8e9190611351565b11610be157610bda6040518060400160405280601081526020017f696e76616c696420616d6f756e743a2000000000000000000000000000000000815250610bd584610bf3565b610fb4565b9050610bec565b610be9610ffd565b90505b9392505050565b606060008203610c3a576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610d76565b600082905060005b60008214610c6c578080610c5590611624565b915050600a82610c65919061169b565b9150610c42565b60008167ffffffffffffffff811115610c8857610c876116cc565b5b6040519080825280601f01601f191660200182016040528015610cba5781602001600182028036833780820191505090505b50905060008290505b60008614610d6e57600181610cd89190611437565b90506000600a8088610cea919061169b565b610cf491906116fb565b87610cff9190611437565b6030610d0b9190611762565b905060008160f81b905080848481518110610d2957610d28611797565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610d65919061169b565b97505050610cc3565b819450505050505b919050565b60606000602867ffffffffffffffff811115610d9a57610d996116cc565b5b6040519080825280601f01601f191660200182016040528015610dcc5781602001600182028036833780820191505090505b50905060005b6014811015610f34576000816013610dea9190611437565b6008610df691906116fb565b6002610e0291906118f9565b8573ffffffffffffffffffffffffffffffffffffffff16610e23919061169b565b60f81b9050600060108260f81c610e3a9190611944565b60f81b905060008160f81c6010610e519190611975565b8360f81c610e5f91906119b0565b60f81b9050610e6d82611031565b85856002610e7b91906116fb565b81518110610e8c57610e8b611797565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610ec481611031565b856001866002610ed491906116fb565b610ede9190611351565b81518110610eef57610eee611797565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610f2c90611624565b915050610dd2565b5080915050919050565b610f46611077565b6040518060400160405280600115158152602001838152509050919050565b610f6d611077565b604051806040016040528060011515815260200186868686604051602001610f9894939291906119e5565b6040516020818303038152906040528152509050949350505050565b610fbc611077565b60405180604001604052806001151581526020018484604051602001610fe3929190611a23565b604051602081830303815290604052815250905092915050565b611005611077565b604051806040016040528060001515815260200160405180602001604052806000815250815250905090565b6000600a8260f81c60ff16101561105c5760308260f81c6110529190611762565b60f81b9050611072565b60578260f81c61106c9190611762565b60f81b90505b919050565b6040518060400160405280600015158152602001606081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006110c382611098565b9050919050565b6110d3816110b8565b81146110de57600080fd5b50565b6000813590506110f0816110ca565b92915050565b6000819050919050565b611109816110f6565b811461111457600080fd5b50565b60008135905061112681611100565b92915050565b6000806040838503121561114357611142611093565b5b6000611151858286016110e1565b925050602061116285828601611117565b9150509250929050565b60008115159050919050565b6111818161116c565b82525050565b600060208201905061119c6000830184611178565b92915050565b600080604083850312156111b9576111b8611093565b5b60006111c7858286016110e1565b92505060206111d8858286016110e1565b9150509250929050565b6111eb816110f6565b82525050565b600060208201905061120660008301846111e2565b92915050565b60006020828403121561122257611221611093565b5b6000611230848285016110e1565b91505092915050565b611242816110b8565b82525050565b600060208201905061125d6000830184611239565b92915050565b60008060006060848603121561127c5761127b611093565b5b600061128a868287016110e1565b935050602061129b868287016110e1565b92505060406112ac86828701611117565b9150509250925092565b6112bf8161116c565b81146112ca57600080fd5b50565b6000813590506112dc816112b6565b92915050565b600080604083850312156112f9576112f8611093565b5b6000611307858286016110e1565b9250506020611318858286016112cd565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061135c826110f6565b9150611367836110f6565b925082820190508082111561137f5761137e611322565b5b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156113bf5780820151818401526020810190506113a4565b60008484015250505050565b6000601f19601f8301169050919050565b60006113e782611385565b6113f18185611390565b93506114018185602086016113a1565b61140a816113cb565b840191505092915050565b6000602082019050818103600083015261142f81846113dc565b905092915050565b6000611442826110f6565b915061144d836110f6565b925082820390508181111561146557611464611322565b5b92915050565b7f7472616e73666572656420000000000000000000000000000000000000000000815250565b600081905092915050565b60006114a782611385565b6114b18185611491565b93506114c18185602086016113a1565b80840191505092915050565b7f20746f2000000000000000000000000000000000000000000000000000000000815250565b7f2066726f6d200000000000000000000000000000000000000000000000000000815250565b60006115248261146b565b600b82019150611534828661149c565b915061153f826114cd565b60048201915061154f828561149c565b915061155a826114f3565b60068201915061156a828461149c565b9150819050949350505050565b7f7472616e7366657265642066726f6d2000000000000000000000000000000000815250565b60006115a882611577565b6010820191506115b8828661149c565b91506115c3826114cd565b6004820191506115d3828561149c565b91506115de826114f3565b6006820191506115ee828461149c565b9150819050949350505050565b60006040820190506116106000830185611239565b61161d6020830184611178565b9392505050565b600061162f826110f6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361166157611660611322565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006116a6826110f6565b91506116b1836110f6565b9250826116c1576116c061166c565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000611706826110f6565b9150611711836110f6565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561174a57611749611322565b5b828202905092915050565b600060ff82169050919050565b600061176d82611755565b915061177883611755565b9250828201905060ff81111561179157611790611322565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008160011c9050919050565b6000808291508390505b600185111561181d578086048111156117f9576117f8611322565b5b60018516156118085780820291505b8081029050611816856117c6565b94506117dd565b94509492505050565b60008261183657600190506118f2565b8161184457600090506118f2565b816001811461185a576002811461186457611893565b60019150506118f2565b60ff84111561187657611875611322565b5b8360020a91508482111561188d5761188c611322565b5b506118f2565b5060208310610133831016604e8410600b84101617156118c85782820a9050838111156118c3576118c2611322565b5b6118f2565b6118d584848460016117d3565b925090508184048111156118ec576118eb611322565b5b81810290505b9392505050565b6000611904826110f6565b915061190f836110f6565b925061193c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611826565b905092915050565b600061194f82611755565b915061195a83611755565b92508261196a5761196961166c565b5b828204905092915050565b600061198082611755565b915061198b83611755565b92508160ff04831182151516156119a5576119a4611322565b5b828202905092915050565b60006119bb82611755565b91506119c683611755565b9250828203905060ff8111156119df576119de611322565b5b92915050565b60006119f1828761149c565b91506119fd828661149c565b9150611a09828561149c565b9150611a15828461149c565b915081905095945050505050565b6000611a2f828561149c565b9150611a3b828461149c565b9150819050939250505056fea26469706673582212204fcb89261f518fc99d0a4a56194dab983bebb0c6c7049bc8eecc4058300463f064736f6c63430008100033",
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
	parsed, err := abi.JSON(strings.NewReader(SimplecoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// Allowance is a free data retrieval call binding the contract method 0x6dcdd00f.
//
// Solidity: function Allowance(address , address ) view returns(uint256)
func (_Simplecoin *SimplecoinCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Simplecoin.contract.Call(opts, &out, "Allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x6dcdd00f.
//
// Solidity: function Allowance(address , address ) view returns(uint256)
func (_Simplecoin *SimplecoinSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Simplecoin.Contract.Allowance(&_Simplecoin.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0x6dcdd00f.
//
// Solidity: function Allowance(address , address ) view returns(uint256)
func (_Simplecoin *SimplecoinCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Simplecoin.Contract.Allowance(&_Simplecoin.CallOpts, arg0, arg1)
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

// Authorize is a paid mutator transaction binding the contract method 0x6386a433.
//
// Solidity: function Authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_Simplecoin *SimplecoinTransactor) Authorize(opts *bind.TransactOpts, authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.contract.Transact(opts, "Authorize", authorizedAccount, allowanceAmount)
}

// Authorize is a paid mutator transaction binding the contract method 0x6386a433.
//
// Solidity: function Authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_Simplecoin *SimplecoinSession) Authorize(authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.Authorize(&_Simplecoin.TransactOpts, authorizedAccount, allowanceAmount)
}

// Authorize is a paid mutator transaction binding the contract method 0x6386a433.
//
// Solidity: function Authorize(address authorizedAccount, uint256 allowanceAmount) returns(bool)
func (_Simplecoin *SimplecoinTransactorSession) Authorize(authorizedAccount common.Address, allowanceAmount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.Authorize(&_Simplecoin.TransactOpts, authorizedAccount, allowanceAmount)
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

// TransferFrom is a paid mutator transaction binding the contract method 0xc0d84ce5.
//
// Solidity: function TransferFrom(address from, address to, uint256 amount) returns()
func (_Simplecoin *SimplecoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.contract.Transact(opts, "TransferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xc0d84ce5.
//
// Solidity: function TransferFrom(address from, address to, uint256 amount) returns()
func (_Simplecoin *SimplecoinSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.TransferFrom(&_Simplecoin.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xc0d84ce5.
//
// Solidity: function TransferFrom(address from, address to, uint256 amount) returns()
func (_Simplecoin *SimplecoinTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simplecoin.Contract.TransferFrom(&_Simplecoin.TransactOpts, from, to, amount)
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
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventTransfer is a free log retrieval operation binding the contract event 0x64af04e507117cfe690abc373ad27eb3c6a0b3202af0972ff3dc089e501e3a45.
//
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 value)
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
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 value)
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
// Solidity: event EventTransfer(address indexed from, address indexed to, uint256 value)
func (_Simplecoin *SimplecoinFilterer) ParseEventTransfer(log types.Log) (*SimplecoinEventTransfer, error) {
	event := new(SimplecoinEventTransfer)
	if err := _Simplecoin.contract.UnpackLog(event, "EventTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}