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
	Bin: "0x60806040523480156200001157600080fd5b50604051620028883803806200288883398181016040528101906200003791906200058e565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fbb62860d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405162000184919062000639565b6000604051808303816000865af19150503d8060008114620001c3576040519150601f19603f3d011682016040523d82523d6000602084013e620001c8565b606091505b50915091507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a62000204846200025e60201b62000a561760201c565b6200021a846200045360201b62000c191760201c565b836040516020016200022f9392919062000737565b6040516020818303038152906040526040516200024d919062000815565b60405180910390a150505062000cff565b60606000602867ffffffffffffffff81111562000280576200027f62000839565b5b6040519080825280601f01601f191660200182016040528015620002b35781602001600182028036833780820191505090505b50905060005b601481101562000449576000816013620002d49190620008a1565b6008620002e29190620008dc565b6002620002f0919062000a91565b8573ffffffffffffffffffffffffffffffffffffffff1662000313919062000b11565b60f81b9050600060108260f81c6200032c919062000b56565b60f81b905060008160f81c601062000345919062000b8e565b8360f81c62000355919062000bd0565b60f81b90506200036b82620004d860201b60201c565b858560026200037b9190620008dc565b815181106200038f576200038e62000c0c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350620003cf81620004d860201b60201c565b856001866002620003e19190620008dc565b620003ed919062000c3b565b8151811062000401576200040062000c0c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080620004409062000c76565b915050620002b9565b5080915050919050565b606081156200049a576040518060400160405280600481526020017f74727565000000000000000000000000000000000000000000000000000000008152509050620004d3565b6040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525090505b919050565b6000600a8260f81c60ff161015620005075760308260f81c620004fc919062000cc3565b60f81b90506200051f565b60578260f81c62000519919062000cc3565b60f81b90505b919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620005568262000529565b9050919050565b620005688162000549565b81146200057457600080fd5b50565b60008151905062000588816200055d565b92915050565b600060208284031215620005a757620005a662000524565b5b6000620005b78482850162000577565b91505092915050565b600081519050919050565b600081905092915050565b60005b83811015620005f6578082015181840152602081019050620005d9565b60008484015250505050565b60006200060f82620005c0565b6200061b8185620005cb565b93506200062d818560208601620005d6565b80840191505092915050565b600062000647828462000602565b915081905092915050565b7f757067726164655b000000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60006200069b8262000678565b620006a7818562000683565b9350620006b9818560208601620005d6565b80840191505092915050565b7f5d20737563636573735b00000000000000000000000000000000000000000000815250565b7f5d2076657273696f6e5b00000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b6000620007448262000652565b6008820191506200075682866200068e565b91506200076382620006c5565b600a820191506200077582856200068e565b91506200078282620006eb565b600a820191506200079482846200068e565b9150620007a18262000711565b600182019150819050949350505050565b600082825260208201905092915050565b6000601f19601f8301169050919050565b6000620007e18262000678565b620007ed8185620007b2565b9350620007ff818560208601620005d6565b6200080a81620007c3565b840191505092915050565b60006020820190508181036000830152620008318184620007d4565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620008ae8262000868565b9150620008bb8362000868565b9250828203905081811115620008d657620008d562000872565b5b92915050565b6000620008e98262000868565b9150620008f68362000868565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161562000932576200093162000872565b5b828202905092915050565b60008160011c9050919050565b6000808291508390505b60018511156200099c5780860481111562000974576200097362000872565b5b6001851615620009845780820291505b808102905062000994856200093d565b945062000954565b94509492505050565b600082620009b7576001905062000a8a565b81620009c7576000905062000a8a565b8160018114620009e05760028114620009eb5762000a21565b600191505062000a8a565b60ff84111562000a0057620009ff62000872565b5b8360020a91508482111562000a1a5762000a1962000872565b5b5062000a8a565b5060208310610133831016604e8410600b841016171562000a5b5782820a90508381111562000a555762000a5462000872565b5b62000a8a565b62000a6a84848460016200094a565b9250905081840481111562000a845762000a8362000872565b5b81810290505b9392505050565b600062000a9e8262000868565b915062000aab8362000868565b925062000ada7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484620009a5565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600062000b1e8262000868565b915062000b2b8362000868565b92508262000b3e5762000b3d62000ae2565b5b828204905092915050565b600060ff82169050919050565b600062000b638262000b49565b915062000b708362000b49565b92508262000b835762000b8262000ae2565b5b828204905092915050565b600062000b9b8262000b49565b915062000ba88362000b49565b92508160ff048311821515161562000bc55762000bc462000872565b5b828202905092915050565b600062000bdd8262000b49565b915062000bea8362000b49565b9250828203905060ff81111562000c065762000c0562000872565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600062000c488262000868565b915062000c558362000868565b925082820190508082111562000c705762000c6f62000872565b5b92915050565b600062000c838262000868565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362000cb85762000cb762000872565b5b600182019050919050565b600062000cd08262000b49565b915062000cdd8362000b49565b9250828201905060ff81111562000cf95762000cf862000872565b5b92915050565b611b798062000d0f6000396000f3fe6080604052600436106100865760003560e01c8063bb62860d11610059578063bb62860d14610116578063c32ab72e14610141578063e63f341f1461016a578063ed21248c146101a7578063fa84fd8e146101b157610086565b80630ef678871461008b57806357ea89b6146100b65780637d7b0099146100c0578063b4a99a4e146100eb575b600080fd5b34801561009757600080fd5b506100a06101da565b6040516100ad9190610e83565b60405180910390f35b6100be610221565b005b3480156100cc57600080fd5b506100d56103e4565b6040516100e29190610edf565b60405180910390f35b3480156100f757600080fd5b50610100610408565b60405161010d9190610edf565b60405180910390f35b34801561012257600080fd5b5061012b61042e565b6040516101389190610f8a565b60405180910390f35b34801561014d57600080fd5b5061016860048036038101906101639190610fec565b6104bc565b005b34801561017657600080fd5b50610191600480360381019061018c9190610fec565b6106d6565b60405161019e9190610e83565b60405180910390f35b6101af610779565b005b3480156101bd57600080fd5b506101d860048036038101906101d3919061118d565b610878565b005b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003390506000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054036102a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029f9061125c565b60405180910390fd5b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610332573d6000803e3d6000fd5b506000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103a233610a56565b6103ab83610c9c565b6040516020016103bc92919061132a565b6040516020818303038152906040526040516103d89190610f8a565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001805461043b906113aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610467906113aa565b80156104b45780601f10610489576101008083540402835291602001916104b4565b820191906000526020600020905b81548152906001019060200180831161049757829003601f168201915b505050505081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461051657600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fbb62860d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516106209190611422565b6000604051808303816000865af19150503d806000811461065d576040519150601f19603f3d011682016040523d82523d6000602084013e610662565b606091505b50915091507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61069184610a56565b61069a84610c19565b836040516020016106ad939291906114ab565b6040516020818303038152906040526040516106c99190610f8a565b60405180910390a1505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461073257600080fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b34600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107c89190611547565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6107f933610a56565b610841600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c9c565b6040516020016108529291906115c7565b60405160208183030381529060405260405161086e9190610f8a565b60405180910390a1565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108d257600080fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168585858560405160240161092394939291906116d6565b6040516020818303038152906040527ffa84fd8e000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516109ad9190611422565b600060405180830381855af49150503d80600081146109e8576040519150601f19603f3d011682016040523d82523d6000602084013e6109ed565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610a1b82610c19565b604051602001610a2b9190611748565b604051602081830303815290604052604051610a479190610f8a565b60405180910390a15050505050565b60606000602867ffffffffffffffff811115610a7557610a7461101e565b5b6040519080825280601f01601f191660200182016040528015610aa75781602001600182028036833780820191505090505b50905060005b6014811015610c0f576000816013610ac5919061177d565b6008610ad191906117b1565b6002610add919061193e565b8573ffffffffffffffffffffffffffffffffffffffff16610afe91906119b8565b60f81b9050600060108260f81c610b1591906119f6565b60f81b905060008160f81c6010610b2c9190611a27565b8360f81c610b3a9190611a62565b60f81b9050610b4882610e24565b85856002610b5691906117b1565b81518110610b6757610b66611a97565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610b9f81610e24565b856001866002610baf91906117b1565b610bb99190611547565b81518110610bca57610bc9611a97565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610c0790611ac6565b915050610aad565b5080915050919050565b60608115610c5e576040518060400160405280600481526020017f74727565000000000000000000000000000000000000000000000000000000008152509050610c97565b6040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525090505b919050565b606060008203610ce3576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610e1f565b600082905060005b60008214610d15578080610cfe90611ac6565b915050600a82610d0e91906119b8565b9150610ceb565b60008167ffffffffffffffff811115610d3157610d3061101e565b5b6040519080825280601f01601f191660200182016040528015610d635781602001600182028036833780820191505090505b50905060008290505b60008614610e1757600181610d81919061177d565b90506000600a8088610d9391906119b8565b610d9d91906117b1565b87610da8919061177d565b6030610db49190611b0e565b905060008160f81b905080848481518110610dd257610dd1611a97565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610e0e91906119b8565b97505050610d6c565b819450505050505b919050565b6000600a8260f81c60ff161015610e4f5760308260f81c610e459190611b0e565b60f81b9050610e65565b60578260f81c610e5f9190611b0e565b60f81b90505b919050565b6000819050919050565b610e7d81610e6a565b82525050565b6000602082019050610e986000830184610e74565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ec982610e9e565b9050919050565b610ed981610ebe565b82525050565b6000602082019050610ef46000830184610ed0565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610f34578082015181840152602081019050610f19565b60008484015250505050565b6000601f19601f8301169050919050565b6000610f5c82610efa565b610f668185610f05565b9350610f76818560208601610f16565b610f7f81610f40565b840191505092915050565b60006020820190508181036000830152610fa48184610f51565b905092915050565b6000604051905090565b600080fd5b600080fd5b610fc981610ebe565b8114610fd457600080fd5b50565b600081359050610fe681610fc0565b92915050565b60006020828403121561100257611001610fb6565b5b600061101084828501610fd7565b91505092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61105682610f40565b810181811067ffffffffffffffff821117156110755761107461101e565b5b80604052505050565b6000611088610fac565b9050611094828261104d565b919050565b600067ffffffffffffffff8211156110b4576110b361101e565b5b602082029050602081019050919050565b600080fd5b60006110dd6110d884611099565b61107e565b90508083825260208201905060208402830185811115611100576110ff6110c5565b5b835b8181101561112957806111158882610fd7565b845260208401935050602081019050611102565b5050509392505050565b600082601f83011261114857611147611019565b5b81356111588482602086016110ca565b91505092915050565b61116a81610e6a565b811461117557600080fd5b50565b60008135905061118781611161565b92915050565b600080600080608085870312156111a7576111a6610fb6565b5b60006111b587828801610fd7565b945050602085013567ffffffffffffffff8111156111d6576111d5610fbb565b5b6111e287828801611133565b93505060406111f387828801611178565b925050606061120487828801611178565b91505092959194509250565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000611246601283610f05565b915061125182611210565b602082019050919050565b6000602082019050818103600083015261127581611239565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081905092915050565b60006112b882610efa565b6112c281856112a2565b93506112d2818560208601610f16565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b60006113358261127c565b60098201915061134582856112ad565b9150611350826112de565b60098201915061136082846112ad565b915061136b82611304565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806113c257607f821691505b6020821081036113d5576113d461137b565b5b50919050565b600081519050919050565b600081905092915050565b60006113fc826113db565b61140681856113e6565b9350611416818560208601610f16565b80840191505092915050565b600061142e82846113f1565b915081905092915050565b7f757067726164655b000000000000000000000000000000000000000000000000815250565b7f5d20737563636573735b00000000000000000000000000000000000000000000815250565b7f5d2076657273696f6e5b00000000000000000000000000000000000000000000815250565b60006114b682611439565b6008820191506114c682866112ad565b91506114d18261145f565b600a820191506114e182856112ad565b91506114ec82611485565b600a820191506114fc82846112ad565b915061150782611304565b600182019150819050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061155282610e6a565b915061155d83610e6a565b925082820190508082111561157557611574611518565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b60006115d28261157b565b6008820191506115e282856112ad565b91506115ed826115a1565b600a820191506115fd82846112ad565b915061160882611304565b6001820191508190509392505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61164d81610ebe565b82525050565b600061165f8383611644565b60208301905092915050565b6000602082019050919050565b600061168382611618565b61168d8185611623565b935061169883611634565b8060005b838110156116c95781516116b08882611653565b97506116bb8361166b565b92505060018101905061169c565b5085935050505092915050565b60006080820190506116eb6000830187610ed0565b81810360208301526116fd8186611678565b905061170c6040830185610e74565b6117196060830184610e74565b95945050505050565b7f7265636f6e63696c655b00000000000000000000000000000000000000000000815250565b600061175382611722565b600a8201915061176382846112ad565b915061176e82611304565b60018201915081905092915050565b600061178882610e6a565b915061179383610e6a565b92508282039050818111156117ab576117aa611518565b5b92915050565b60006117bc82610e6a565b91506117c783610e6a565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611800576117ff611518565b5b828202905092915050565b60008160011c9050919050565b6000808291508390505b60018511156118625780860481111561183e5761183d611518565b5b600185161561184d5780820291505b808102905061185b8561180b565b9450611822565b94509492505050565b60008261187b5760019050611937565b816118895760009050611937565b816001811461189f57600281146118a9576118d8565b6001915050611937565b60ff8411156118bb576118ba611518565b5b8360020a9150848211156118d2576118d1611518565b5b50611937565b5060208310610133831016604e8410600b841016171561190d5782820a90508381111561190857611907611518565b5b611937565b61191a8484846001611818565b9250905081840481111561193157611930611518565b5b81810290505b9392505050565b600061194982610e6a565b915061195483610e6a565b92506119817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461186b565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006119c382610e6a565b91506119ce83610e6a565b9250826119de576119dd611989565b5b828204905092915050565b600060ff82169050919050565b6000611a01826119e9565b9150611a0c836119e9565b925082611a1c57611a1b611989565b5b828204905092915050565b6000611a32826119e9565b9150611a3d836119e9565b92508160ff0483118215151615611a5757611a56611518565b5b828202905092915050565b6000611a6d826119e9565b9150611a78836119e9565b9250828203905060ff811115611a9157611a90611518565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611ad182610e6a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b0357611b02611518565b5b600182019050919050565b6000611b19826119e9565b9150611b24836119e9565b9250828201905060ff811115611b3d57611b3c611518565b5b9291505056fea2646970667358221220c8c0f1d727a446e0feb4ff689a466fd97f225c57d7f5efae4e5f6f16bf5058e264736f6c63430008100033",
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
