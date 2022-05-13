# https://goethereumbook.org/smart-contract-deploy/
#
# The environment has a single account for now with this information.
#    Passkey: 123
#    Public address of the key:   0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
#    Path of the secret key file: zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd

# ==============================================================================
# These commands build the basic smart contract, deploy the contract, and run the
# play code to test it is working.

basic-build:
	solc --abi app/basic/contracts/store/src/store.sol -o app/basic/contracts/store/abi --overwrite
	solc --bin app/basic/contracts/store/src/store.sol -o app/basic/contracts/store/bin --overwrite
	abigen --bin=app/basic/contracts/store/bin/Store.bin --abi=app/basic/contracts/store/abi/Store.abi --pkg=store --out=app/basic/contracts/store/store.go

basic-deploy:
	go run app/basic/cmd/deploy/main.go

basic-play:
	go run app/basic/cmd/play/main.go

# ==============================================================================
# These commands start the Ethereum node and provide examples of attaching
# directly with potential commands to try, and creating a new account if necessary.

geth-up:
	geth --dev --rpc.allow-unprotected-txs --verbosity 4 --cache 512 --ipcpath zarf/ethereum/geth.ipc --networkid 12345 --datadir "zarf/ethereum/" --unlock 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd --password zarf/ethereum/password

geth-down:
	kill -INT $(shell ps | grep "geth " | grep -v grep | sed -n 1,1p | cut -c1-5)

geth-attach:
	geth attach --datadir zarf/ethereum/
# 	eth.getBalance("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd")
# 	eth.getBalance("0x8e113078adf6888b7ba84967f299f29aece24c55")
# 	eth.getBalance("0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7")
#   personal.newAccount()
#   eth.sendTransaction({from:eth.coinbase, to:"address", value: web3.toWei(0.05, "ether")})

geth-new-account:
	geth --datadir zarf/ethereum/ account new
#	If you run this command, add the new account address to the genesis file.
#	To see the new balances, you need to restart the geth service.

# ==============================================================================
# These commands provide Go related support.

tidy:
	go mod tidy
	go mod vendor

# ==============================================================================
# These commands install geth, abigen, and solc using Homebrew.

# https://geth.ethereum.org/docs/install-and-build/installing-geth
install-geth:
	brew update
	brew tap ethereum/ethereum
	brew install ethereum

# https://docs.soliditylang.org/en/v0.8.11/installing-solidity.html
install-solc:
	brew update
	brew tap ethereum/ethereum
	brew install solidity
