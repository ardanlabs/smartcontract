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
)

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"api\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"API\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"losers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"anteWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gameFeeWei\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"api\",\"type\":\"address\"}],\"name\":\"UpgradeAPI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200254e3803806200254e83398181016040528101906200003791906200058e565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fbb62860d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405162000184919062000639565b6000604051808303816000865af19150503d8060008114620001c3576040519150601f19603f3d011682016040523d82523d6000602084013e620001c8565b606091505b50915091507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a62000204846200025e60201b62000a721760201c565b6200021a846200045360201b62000c351760201c565b836040516020016200022f9392919062000737565b6040516020818303038152906040526040516200024d919062000815565b60405180910390a150505062000ced565b60606000602867ffffffffffffffff81111562000280576200027f62000839565b5b6040519080825280601f01601f191660200182016040528015620002b35781602001600182028036833780820191505090505b50905060005b601481101562000449576000816013620002d49190620008a1565b6008620002e29190620008dc565b6002620002f0919062000a7b565b8573ffffffffffffffffffffffffffffffffffffffff1662000313919062000afb565b60f81b9050600060108260f81c6200032c919062000b40565b60f81b905060008160f81c601062000345919062000b78565b8360f81c62000355919062000bbe565b60f81b90506200036b82620004d860201b60201c565b858560026200037b9190620008dc565b815181106200038f576200038e62000bfa565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350620003cf81620004d860201b60201c565b856001866002620003e19190620008dc565b620003ed919062000c29565b8151811062000401576200040062000bfa565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080620004409062000c64565b915050620002b9565b5080915050919050565b606081156200049a576040518060400160405280600481526020017f74727565000000000000000000000000000000000000000000000000000000008152509050620004d3565b6040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525090505b919050565b6000600a8260f81c60ff161015620005075760308260f81c620004fc919062000cb1565b60f81b90506200051f565b60578260f81c62000519919062000cb1565b60f81b90505b919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620005568262000529565b9050919050565b620005688162000549565b81146200057457600080fd5b50565b60008151905062000588816200055d565b92915050565b600060208284031215620005a757620005a662000524565b5b6000620005b78482850162000577565b91505092915050565b600081519050919050565b600081905092915050565b60005b83811015620005f6578082015181840152602081019050620005d9565b60008484015250505050565b60006200060f82620005c0565b6200061b8185620005cb565b93506200062d818560208601620005d6565b80840191505092915050565b600062000647828462000602565b915081905092915050565b7f757067726164655b000000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60006200069b8262000678565b620006a7818562000683565b9350620006b9818560208601620005d6565b80840191505092915050565b7f5d20737563636573735b00000000000000000000000000000000000000000000815250565b7f5d2076657273696f6e5b00000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b6000620007448262000652565b6008820191506200075682866200068e565b91506200076382620006c5565b600a820191506200077582856200068e565b91506200078282620006eb565b600a820191506200079482846200068e565b9150620007a18262000711565b600182019150819050949350505050565b600082825260208201905092915050565b6000601f19601f8301169050919050565b6000620007e18262000678565b620007ed8185620007b2565b9350620007ff818560208601620005d6565b6200080a81620007c3565b840191505092915050565b60006020820190508181036000830152620008318184620007d4565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620008ae8262000868565b9150620008bb8362000868565b9250828203905081811115620008d657620008d562000872565b5b92915050565b6000620008e98262000868565b9150620008f68362000868565b9250828202620009068162000868565b9150828204841483151762000920576200091f62000872565b5b5092915050565b60008160011c9050919050565b6000808291508390505b600185111562000986578086048111156200095e576200095d62000872565b5b60018516156200096e5780820291505b80810290506200097e8562000927565b94506200093e565b94509492505050565b600082620009a1576001905062000a74565b81620009b1576000905062000a74565b8160018114620009ca5760028114620009d55762000a0b565b600191505062000a74565b60ff841115620009ea57620009e962000872565b5b8360020a91508482111562000a045762000a0362000872565b5b5062000a74565b5060208310610133831016604e8410600b841016171562000a455782820a90508381111562000a3f5762000a3e62000872565b5b62000a74565b62000a54848484600162000934565b9250905081840481111562000a6e5762000a6d62000872565b5b81810290505b9392505050565b600062000a888262000868565b915062000a958362000868565b925062000ac47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846200098f565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600062000b088262000868565b915062000b158362000868565b92508262000b285762000b2762000acc565b5b828204905092915050565b600060ff82169050919050565b600062000b4d8262000b33565b915062000b5a8362000b33565b92508262000b6d5762000b6c62000acc565b5b828204905092915050565b600062000b858262000b33565b915062000b928362000b33565b925082820262000ba28162000b33565b915080821462000bb75762000bb662000872565b5b5092915050565b600062000bcb8262000b33565b915062000bd88362000b33565b9250828203905060ff81111562000bf45762000bf362000872565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600062000c368262000868565b915062000c438362000868565b925082820190508082111562000c5e5762000c5d62000872565b5b92915050565b600062000c718262000868565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362000ca65762000ca562000872565b5b600182019050919050565b600062000cbe8262000b33565b915062000ccb8362000b33565b9250828201905060ff81111562000ce75762000ce662000872565b5b92915050565b6118518062000cfd6000396000f3fe6080604052600436106100865760003560e01c8063bb62860d11610059578063bb62860d14610116578063c32ab72e14610141578063e63f341f1461016a578063ed21248c146101a7578063fa84fd8e146101b157610086565b80630ef678871461008b57806357ea89b6146100b65780637d7b0099146100c0578063b4a99a4e146100eb575b600080fd5b34801561009757600080fd5b506100a06101da565b6040516100ad9190610d17565b60405180910390f35b6100be610221565b005b3480156100cc57600080fd5b506100d5610390565b6040516100e29190610d73565b60405180910390f35b3480156100f757600080fd5b506101006103b4565b60405161010d9190610d73565b60405180910390f35b34801561012257600080fd5b5061012b6103da565b6040516101389190610e1e565b60405180910390f35b34801561014d57600080fd5b5061016860048036038101906101639190610e80565b610468565b005b34801561017657600080fd5b50610191600480360381019061018c9190610e80565b610682565b60405161019e9190610d17565b60405180910390f35b6101af610725565b005b3480156101bd57600080fd5b506101d860048036038101906101d39190611021565b610894565b005b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f57ea89b6000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516102eb91906110eb565b600060405180830381855af49150503d8060008114610326576040519150601f19603f3d011682016040523d82523d6000602084013e61032b565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61035982610c35565b604051602001610369919061118a565b6040516020818303038152906040526040516103859190610e1e565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600180546103e7906111ee565b80601f0160208091040260200160405190810160405280929190818152602001828054610413906111ee565b80156104605780601f1061043557610100808354040283529160200191610460565b820191906000526020600020905b81548152906001019060200180831161044357829003601f168201915b505050505081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104c257600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fbb62860d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516105cc91906110eb565b6000604051808303816000865af19150503d8060008114610609576040519150601f19603f3d011682016040523d82523d6000602084013e61060e565b606091505b50915091507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61063d84610a72565b61064684610c35565b8360405160200161065993929190611291565b6040516020818303038152906040526040516106759190610e1e565b60405180910390a1505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106de57600080fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fed21248c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516107ef91906110eb565b600060405180830381855af49150503d806000811461082a576040519150601f19603f3d011682016040523d82523d6000602084013e61082f565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61085d82610c35565b60405160200161086d919061118a565b6040516020818303038152906040526040516108899190610e1e565b60405180910390a150565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108ee57600080fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168585858560405160240161093f94939291906113bc565b6040516020818303038152906040527ffa84fd8e000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516109c991906110eb565b600060405180830381855af49150503d8060008114610a04576040519150601f19603f3d011682016040523d82523d6000602084013e610a09565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610a3782610c35565b604051602001610a47919061118a565b604051602081830303815290604052604051610a639190610e1e565b60405180910390a15050505050565b60606000602867ffffffffffffffff811115610a9157610a90610eb2565b5b6040519080825280601f01601f191660200182016040528015610ac35781602001600182028036833780820191505090505b50905060005b6014811015610c2b576000816013610ae19190611437565b6008610aed919061146b565b6002610af991906115e0565b8573ffffffffffffffffffffffffffffffffffffffff16610b1a919061165a565b60f81b9050600060108260f81c610b319190611698565b60f81b905060008160f81c6010610b4891906116c9565b8360f81c610b569190611706565b60f81b9050610b6482610cb8565b85856002610b72919061146b565b81518110610b8357610b8261173b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610bbb81610cb8565b856001866002610bcb919061146b565b610bd5919061176a565b81518110610be657610be561173b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610c239061179e565b915050610ac9565b5080915050919050565b60608115610c7a576040518060400160405280600481526020017f74727565000000000000000000000000000000000000000000000000000000008152509050610cb3565b6040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525090505b919050565b6000600a8260f81c60ff161015610ce35760308260f81c610cd991906117e6565b60f81b9050610cf9565b60578260f81c610cf391906117e6565b60f81b90505b919050565b6000819050919050565b610d1181610cfe565b82525050565b6000602082019050610d2c6000830184610d08565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d5d82610d32565b9050919050565b610d6d81610d52565b82525050565b6000602082019050610d886000830184610d64565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610dc8578082015181840152602081019050610dad565b60008484015250505050565b6000601f19601f8301169050919050565b6000610df082610d8e565b610dfa8185610d99565b9350610e0a818560208601610daa565b610e1381610dd4565b840191505092915050565b60006020820190508181036000830152610e388184610de5565b905092915050565b6000604051905090565b600080fd5b600080fd5b610e5d81610d52565b8114610e6857600080fd5b50565b600081359050610e7a81610e54565b92915050565b600060208284031215610e9657610e95610e4a565b5b6000610ea484828501610e6b565b91505092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610eea82610dd4565b810181811067ffffffffffffffff82111715610f0957610f08610eb2565b5b80604052505050565b6000610f1c610e40565b9050610f288282610ee1565b919050565b600067ffffffffffffffff821115610f4857610f47610eb2565b5b602082029050602081019050919050565b600080fd5b6000610f71610f6c84610f2d565b610f12565b90508083825260208201905060208402830185811115610f9457610f93610f59565b5b835b81811015610fbd5780610fa98882610e6b565b845260208401935050602081019050610f96565b5050509392505050565b600082601f830112610fdc57610fdb610ead565b5b8135610fec848260208601610f5e565b91505092915050565b610ffe81610cfe565b811461100957600080fd5b50565b60008135905061101b81610ff5565b92915050565b6000806000806080858703121561103b5761103a610e4a565b5b600061104987828801610e6b565b945050602085013567ffffffffffffffff81111561106a57611069610e4f565b5b61107687828801610fc7565b93505060406110878782880161100c565b92505060606110988782880161100c565b91505092959194509250565b600081519050919050565b600081905092915050565b60006110c5826110a4565b6110cf81856110af565b93506110df818560208601610daa565b80840191505092915050565b60006110f782846110ba565b915081905092915050565b7f737563636573735b000000000000000000000000000000000000000000000000815250565b600081905092915050565b600061113e82610d8e565b6111488185611128565b9350611158818560208601610daa565b80840191505092915050565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b600061119582611102565b6008820191506111a58284611133565b91506111b082611164565b60018201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061120657607f821691505b602082108103611219576112186111bf565b5b50919050565b7f757067726164655b000000000000000000000000000000000000000000000000815250565b7f5d20737563636573735b00000000000000000000000000000000000000000000815250565b7f5d2076657273696f6e5b00000000000000000000000000000000000000000000815250565b600061129c8261121f565b6008820191506112ac8286611133565b91506112b782611245565b600a820191506112c78285611133565b91506112d28261126b565b600a820191506112e28284611133565b91506112ed82611164565b600182019150819050949350505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61133381610d52565b82525050565b6000611345838361132a565b60208301905092915050565b6000602082019050919050565b6000611369826112fe565b6113738185611309565b935061137e8361131a565b8060005b838110156113af5781516113968882611339565b97506113a183611351565b925050600181019050611382565b5085935050505092915050565b60006080820190506113d16000830187610d64565b81810360208301526113e3818661135e565b90506113f26040830185610d08565b6113ff6060830184610d08565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061144282610cfe565b915061144d83610cfe565b925082820390508181111561146557611464611408565b5b92915050565b600061147682610cfe565b915061148183610cfe565b925082820261148f81610cfe565b915082820484148315176114a6576114a5611408565b5b5092915050565b60008160011c9050919050565b6000808291508390505b6001851115611504578086048111156114e0576114df611408565b5b60018516156114ef5780820291505b80810290506114fd856114ad565b94506114c4565b94509492505050565b60008261151d57600190506115d9565b8161152b57600090506115d9565b8160018114611541576002811461154b5761157a565b60019150506115d9565b60ff84111561155d5761155c611408565b5b8360020a91508482111561157457611573611408565b5b506115d9565b5060208310610133831016604e8410600b84101617156115af5782820a9050838111156115aa576115a9611408565b5b6115d9565b6115bc84848460016114ba565b925090508184048111156115d3576115d2611408565b5b81810290505b9392505050565b60006115eb82610cfe565b91506115f683610cfe565b92506116237fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461150d565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061166582610cfe565b915061167083610cfe565b9250826116805761167f61162b565b5b828204905092915050565b600060ff82169050919050565b60006116a38261168b565b91506116ae8361168b565b9250826116be576116bd61162b565b5b828204905092915050565b60006116d48261168b565b91506116df8361168b565b92508282026116ed8161168b565b91508082146116ff576116fe611408565b5b5092915050565b60006117118261168b565b915061171c8361168b565b9250828203905060ff81111561173557611734611408565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061177582610cfe565b915061178083610cfe565b925082820190508082111561179857611797611408565b5b92915050565b60006117a982610cfe565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117db576117da611408565b5b600182019050919050565b60006117f18261168b565b91506117fc8361168b565b9250828201905060ff81111561181557611814611408565b5b9291505056fea264697066735822122090ddd53c4f065deba0f53c434e4e281e9b94f132c591d94451d7fe488ff139ed64736f6c63430008110033",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankMetaData.Bin instead.
var BankBin = BankMetaData.Bin

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend, api common.Address) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankBin), backend, api)
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
	parsed, err := abi.JSON(strings.NewReader(BankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// UpgradeAPI is a paid mutator transaction binding the contract method 0xc32ab72e.
//
// Solidity: function UpgradeAPI(address api) returns()
func (_Bank *BankTransactor) UpgradeAPI(opts *bind.TransactOpts, api common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "UpgradeAPI", api)
}

// UpgradeAPI is a paid mutator transaction binding the contract method 0xc32ab72e.
//
// Solidity: function UpgradeAPI(address api) returns()
func (_Bank *BankSession) UpgradeAPI(api common.Address) (*types.Transaction, error) {
	return _Bank.Contract.UpgradeAPI(&_Bank.TransactOpts, api)
}

// UpgradeAPI is a paid mutator transaction binding the contract method 0xc32ab72e.
//
// Solidity: function UpgradeAPI(address api) returns()
func (_Bank *BankTransactorSession) UpgradeAPI(api common.Address) (*types.Transaction, error) {
	return _Bank.Contract.UpgradeAPI(&_Bank.TransactOpts, api)
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
