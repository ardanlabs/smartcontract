# https://ethdocs.org/en/latest/contracts-and-transactions/accessing-contracts-and-transactions.html
# https://goethereumbook.org/smart-contract-deploy/
# https://documenter.getpostman.com/view/4117254/ethereum-json-rpc/RVu7CT5J#dd57ef90-f990-037e-5512-4929e7280d7c
#
# The environment has three accounts all using this same passkey (123).
# Geth is started with address 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd and is used as the coinbase address.
# The coinbase address is the account to pay mining rewards to.
# The coinbase address is give a LOT of money to start.
#
# These are examples of what you can do in the attach JS environment.
# 	eth.getBalance("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd") or eth.getBalance(eth.coinbase)
# 	eth.getBalance("0x8e113078adf6888b7ba84967f299f29aece24c55")
# 	eth.getBalance("0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7")
#   eth.sendTransaction({from:eth.coinbase, to:"0x8e113078adf6888b7ba84967f299f29aece24c55", value: web3.toWei(0.05, "ether")})
#   eth.sendTransaction({from:eth.coinbase, to:"0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7", value: web3.toWei(0.05, "ether")})
#   eth.blockNumber
#   eth.getBlockByNumber(8)
#   eth.getTransaction("0xaea41e7c13a7ea627169c74ade4d5ea86664ff1f740cd90e499f3f842656d4ad")
#
# These are examples of what you can do in the HTTP-RPC environment.
# https://documenter.getpostman.com/view/4117254/ethereum-json-rpc/RVu7CT5J
# curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_coinbase", "id":1}' localhost:8545
# curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_getBalance", "params": ["0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "latest"], "id":1}' localhost:8545


# ==============================================================================
# These commands build the basic smart contract, deploy the contract, and run the
# play code to test it is working.

# This will compile the smart contract and produce the binary code. Then with the
# abi and binary code, a Go source code file can be generated for Go API access.
basic-build:
	solc --abi app/basic/contracts/store/src/store.sol -o app/basic/contracts/store/abi --overwrite
	solc --bin app/basic/contracts/store/src/store.sol -o app/basic/contracts/store/bin --overwrite
	abigen --bin=app/basic/contracts/store/bin/Store.bin --abi=app/basic/contracts/store/abi/Store.abi --pkg=store --out=app/basic/contracts/store/store.go

# This will deploy the smart contract to the locally running Ethereum environment.
basic-deploy:
	go run app/basic/cmd/deploy/main.go

# This will execute a simple program to test access to the smart contract API.
basic-write:
	go run app/basic/cmd/write/main.go

# This will execute a simple program to test access to the smart contract API.
basic-read:
	go run app/basic/cmd/read/main.go


# ==============================================================================
# These commands start the Ethereum node and provide examples of attaching
# directly with potential commands to try, and creating a new account if necessary.

# This is start Ethereum in developer mode. Only when a transaction is pending will
# Ethereum mine a block. It provides a minimal environment for development.
geth-up:
	geth --dev --http --allow-insecure-unlock --rpc.allow-unprotected-txs --mine --miner.threads 1 --verbosity 5 --datadir "zarf/ethereum/" --unlock 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd --password zarf/ethereum/password

# This will signal Ethereum to shutdown.
geth-down:
	kill -INT $(shell ps | grep "geth " | grep -v grep | sed -n 1,1p | cut -c1-5)

# This is a JS console environment for making geth related API calls.
geth-attach:
	geth attach --datadir zarf/ethereum/

# This will add a new account to the keystore. The account will have a zero
# balance until you give it some money.
geth-new-account:
	geth --datadir zarf/ethereum/ account new

# This will deposit 1 ETH into the two extra accounts from the coinbase account.
# Do this if you delete the geth folder and start over or if the accounts need money.
geth-deposit:
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x8e113078adf6888b7ba84967f299f29aece24c55", "to":"0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7", "value":"0x1000000000000000000"}], "id":1}' localhost:8545


# ==============================================================================
# These commands provide Go related support.

# This will tidy up the Go dependencies.
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
