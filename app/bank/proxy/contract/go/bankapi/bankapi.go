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
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600581526020017f302e332e3000000000000000000000000000000000000000000000000000000081525060019081620000589190620002df565b50620003c6565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620000e157607f821691505b602082108103620000f757620000f662000099565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b60008160020a8302905092915050565b600060088302620001647fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000122565b62000170868362000122565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620001bd620001b7620001b18462000188565b62000192565b62000188565b9050919050565b6000819050919050565b620001d9836200019c565b620001f1620001e882620001c4565b84845462000132565b825550505050565b600090565b62000208620001f9565b62000215818484620001ce565b505050565b5b818110156200023d5762000231600082620001fe565b6001810190506200021b565b5050565b601f8211156200028c576200025681620000fd565b620002618462000112565b8101602085101562000271578190505b62000289620002808562000112565b8301826200021a565b50505b505050565b60008160020a8304905092915050565b6000620002b46000198460080262000291565b1980831691505092915050565b6000620002cf8383620002a1565b9150826002028217905092915050565b620002ea826200005f565b67ffffffffffffffff8111156200030657620003056200006a565b5b620003128254620000c8565b6200031f82828562000241565b600060209050601f83116001811462000357576000841562000342578287015190505b6200034e8582620002c1565b865550620003be565b601f1984166200036786620000fd565b60005b8281101562000391578489015182556001820191506020850194506020810190506200036a565b86831015620003b15784890151620003ad601f891682620002a1565b8355505b6001600288020188555050505b505050505050565b611c8a80620003d66000396000f3fe608060405260043610610072576000357c01000000000000000000000000000000000000000000000000000000009004806357ea89b6146100775780637d7b009914610081578063b4a99a4e146100ac578063bb62860d146100d7578063ed21248c14610102578063fa84fd8e1461010c575b600080fd5b61007f610135565b005b34801561008d57600080fd5b506100966102f3565b6040516100a39190610f8c565b60405180910390f35b3480156100b857600080fd5b506100c1610317565b6040516100ce9190610f8c565b60405180910390f35b3480156100e357600080fd5b506100ec61033d565b6040516100f99190611037565b60405180910390f35b61010a6103cb565b005b34801561011857600080fd5b50610133600480360381019061012e9190611217565b6104ca565b005b60003390506000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054036101bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b3906112e6565b60405180910390fd5b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561024157600080fd5b506000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6102b133610a40565b6102ba83610cc0565b6040516020016102cb9291906113b4565b6040516020818303038152906040526040516102e79190611037565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001805461034a90611434565b80601f016020809104026020016040519081016040528092919081815260200182805461037690611434565b80156103c35780601f10610398576101008083540402835291602001916103c3565b820191906000526020600020905b8154815290600101906020018083116103a657829003601f168201915b505050505081565b34600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461041a9190611494565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61044b33610a40565b610493600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610cc0565b6040516020016104a4929190611514565b6040516020818303038152906040526040516104c09190611037565b60405180910390a1565b600082905060005b845181101561075a5783600360008784815181106104f3576104f2611565565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156106c8577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6105bf6003600088858151811061057757610576611565565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610cc0565b6105c886610cc0565b6040516020016105d99291906115e0565b6040516020818303038152906040526040516105f59190611037565b60405180910390a16003600086838151811061061457610613611565565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826106629190611494565b915060006003600087848151811061067d5761067c611565565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610747565b83826106d49190611494565b915083600360008784815181106106ee576106ed611565565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461073f9190611622565b925050819055505b808061075290611656565b9150506104d2565b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61078584610cc0565b61078e84610cc0565b61079784610cc0565b6040516020016107a993929190611710565b6040516020818303038152906040526040516107c59190611037565b60405180910390a160008103610810576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610807906117ef565b60405180910390fd5b818110156108f4578060036000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108899190611494565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6108ba82610cc0565b6040516020016108ca9190611881565b6040516020818303038152906040526040516108e69190611037565b60405180910390a150610a3a565b81816109009190611622565b905080600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109519190611494565b925050819055508160036000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109c99190611494565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6109fa82610cc0565b610a0384610cc0565b604051602001610a149291906118fe565b604051602081830303815290604052604051610a309190611037565b60405180910390a1505b50505050565b60606000602867ffffffffffffffff811115610a5f57610a5e61109e565b5b6040519080825280601f01601f191660200182016040528015610a915781602001600182028036833780820191505090505b50905060005b6014811015610cb6576000816013610aaf9190611622565b6008610abb919061194f565b6002610ac79190611ac4565b8573ffffffffffffffffffffffffffffffffffffffff16610ae89190611b3e565b7f010000000000000000000000000000000000000000000000000000000000000002905060006010827f01000000000000000000000000000000000000000000000000000000000000009004610b3e9190611b7c565b7f01000000000000000000000000000000000000000000000000000000000000000290506000817f010000000000000000000000000000000000000000000000000000000000000090046010610b949190611bad565b837f01000000000000000000000000000000000000000000000000000000000000009004610bc29190611bea565b7f0100000000000000000000000000000000000000000000000000000000000000029050610bef82610e67565b85856002610bfd919061194f565b81518110610c0e57610c0d611565565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610c4681610e67565b856001866002610c56919061194f565b610c609190611494565b81518110610c7157610c70611565565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610cae90611656565b915050610a97565b5080915050919050565b606060008203610d07576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610e62565b600082905060005b60008214610d39578080610d2290611656565b915050600a82610d329190611b3e565b9150610d0f565b60008167ffffffffffffffff811115610d5557610d5461109e565b5b6040519080825280601f01601f191660200182016040528015610d875781602001600182028036833780820191505090505b50905060008290505b60008614610e5a57600181610da59190611622565b90506000600a8088610db79190611b3e565b610dc1919061194f565b87610dcc9190611622565b6030610dd89190611c1f565b90506000817f010000000000000000000000000000000000000000000000000000000000000002905080848481518110610e1557610e14611565565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610e519190611b3e565b97505050610d90565b819450505050505b919050565b6000600a827f0100000000000000000000000000000000000000000000000000000000000000900460ff161015610ef1576030827f01000000000000000000000000000000000000000000000000000000000000009004610ec89190611c1f565b7f0100000000000000000000000000000000000000000000000000000000000000029050610f46565b6057827f01000000000000000000000000000000000000000000000000000000000000009004610f219190611c1f565b7f01000000000000000000000000000000000000000000000000000000000000000290505b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610f7682610f4b565b9050919050565b610f8681610f6b565b82525050565b6000602082019050610fa16000830184610f7d565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610fe1578082015181840152602081019050610fc6565b60008484015250505050565b6000601f19601f8301169050919050565b600061100982610fa7565b6110138185610fb2565b9350611023818560208601610fc3565b61102c81610fed565b840191505092915050565b600060208201905081810360008301526110518184610ffe565b905092915050565b6000604051905090565b600080fd5b600080fd5b61107681610f6b565b811461108157600080fd5b50565b6000813590506110938161106d565b92915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6110d682610fed565b810181811067ffffffffffffffff821117156110f5576110f461109e565b5b80604052505050565b6000611108611059565b905061111482826110cd565b919050565b600067ffffffffffffffff8211156111345761113361109e565b5b602082029050602081019050919050565b600080fd5b600061115d61115884611119565b6110fe565b905080838252602082019050602084028301858111156111805761117f611145565b5b835b818110156111a957806111958882611084565b845260208401935050602081019050611182565b5050509392505050565b600082601f8301126111c8576111c7611099565b5b81356111d884826020860161114a565b91505092915050565b6000819050919050565b6111f4816111e1565b81146111ff57600080fd5b50565b600081359050611211816111eb565b92915050565b6000806000806080858703121561123157611230611063565b5b600061123f87828801611084565b945050602085013567ffffffffffffffff8111156112605761125f611068565b5b61126c878288016111b3565b935050604061127d87828801611202565b925050606061128e87828801611202565b91505092959194509250565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b60006112d0601283610fb2565b91506112db8261129a565b602082019050919050565b600060208201905081810360008301526112ff816112c3565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081905092915050565b600061134282610fa7565b61134c818561132c565b935061135c818560208601610fc3565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b60006113bf82611306565b6009820191506113cf8285611337565b91506113da82611368565b6009820191506113ea8284611337565b91506113f58261138e565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061144c57607f821691505b60208210810361145f5761145e611405565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061149f826111e1565b91506114aa836111e1565b92508282019050808211156114c2576114c1611465565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b600061151f826114c8565b60088201915061152f8285611337565b915061153a826114ee565b600a8201915061154a8284611337565b91506115558261138e565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f6163636f756e742062616c616e63652000000000000000000000000000000000815250565b7f206973206c657373207468616e20616e74652000000000000000000000000000815250565b60006115eb82611594565b6010820191506115fb8285611337565b9150611606826115ba565b6013820191506116168284611337565b91508190509392505050565b600061162d826111e1565b9150611638836111e1565b92508282039050818111156116505761164f611465565b5b92915050565b6000611661826111e1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361169357611692611465565b5b600182019050919050565b7f616e74655b000000000000000000000000000000000000000000000000000000815250565b7f5d2067616d654665655b00000000000000000000000000000000000000000000815250565b7f5d20706f745b0000000000000000000000000000000000000000000000000000815250565b600061171b8261169e565b60058201915061172b8286611337565b9150611736826116c4565b600a820191506117468285611337565b9150611751826116ea565b6006820191506117618284611337565b915061176c8261138e565b600182019150819050949350505050565b7f6e6f20706f74207761732063726561746564206261736564206f6e206163636f60008201527f756e742062616c616e6365730000000000000000000000000000000000000000602082015250565b60006117d9602c83610fb2565b91506117e48261177d565b604082019050919050565b60006020820190508181036000830152611808816117cc565b9050919050565b7f706f74206c657373207468616e206665653a2077696e6e65725b305d206f776e60008201527f65725b0000000000000000000000000000000000000000000000000000000000602082015250565b600061186b60238361132c565b91506118768261180f565b602382019050919050565b600061188c8261185e565b91506118988284611337565b91506118a38261138e565b60018201915081905092915050565b7f77696e6e65725b00000000000000000000000000000000000000000000000000815250565b7f5d206f776e65725b000000000000000000000000000000000000000000000000815250565b6000611909826118b2565b6007820191506119198285611337565b9150611924826118d8565b6008820191506119348284611337565b915061193f8261138e565b6001820191508190509392505050565b600061195a826111e1565b9150611965836111e1565b9250828202611973816111e1565b9150828204841483151761198a57611989611465565b5b5092915050565b6000600282049050919050565b6000808291508390505b60018511156119e8578086048111156119c4576119c3611465565b5b60018516156119d35780820291505b80810290506119e185611991565b94506119a8565b94509492505050565b600082611a015760019050611abd565b81611a0f5760009050611abd565b8160018114611a255760028114611a2f57611a5e565b6001915050611abd565b60ff841115611a4157611a40611465565b5b8360020a915084821115611a5857611a57611465565b5b50611abd565b5060208310610133831016604e8410600b8410161715611a935782820a905083811115611a8e57611a8d611465565b5b611abd565b611aa0848484600161199e565b92509050818404811115611ab757611ab6611465565b5b81810290505b9392505050565b6000611acf826111e1565b9150611ada836111e1565b9250611b077fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846119f1565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611b49826111e1565b9150611b54836111e1565b925082611b6457611b63611b0f565b5b828204905092915050565b600060ff82169050919050565b6000611b8782611b6f565b9150611b9283611b6f565b925082611ba257611ba1611b0f565b5b828204905092915050565b6000611bb882611b6f565b9150611bc383611b6f565b9250828202611bd181611b6f565b9150808214611be357611be2611465565b5b5092915050565b6000611bf582611b6f565b9150611c0083611b6f565b9250828203905060ff811115611c1957611c18611465565b5b92915050565b6000611c2a82611b6f565b9150611c3583611b6f565b9250828201905060ff811115611c4e57611c4d611465565b5b9291505056fea264697066735822122070f14bf3b5c37ae5cd1a441d0a3aa7635f1a6f89be2a2852b82c3a3f38c86f2064736f6c63430008140033",
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
