# https://goethereumbook.org/smart-contract-deploy/
#
# The environment has a single account for now with this information.
#    Passkey: 123
#    Public address of the key:   0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
#    Path of the secret key file: node/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd

# ==============================================================================
# These commands build the basic smart contract, deploy the contract, and run the
# play code to test it is working.

basic-build:
	solc --abi basic/contracts/store/src/store.sol -o basic/contracts/store/abi --overwrite
	solc --bin basic/contracts/store/src/store.sol -o basic/contracts/store/bin --overwrite
	abigen --bin=basic/contracts/store/bin/store.bin --abi=basic/contracts/store/abi/store.abi --pkg=store --out=basic/contracts/store/store.go

basic-deploy:
	go run programs/basic/apps/deploy/main.go

basic-play:
	go run programs/basic/apps/play/main.go

# ==============================================================================
# These commands start the Ethereum node and provide examples of attaching
# directly with potential commands to try, and creating a new account if necessary.

geth-init:
	geth --datadir "node/" init "node/genesis.json"

geth:
	geth --rpc.allow-unprotected-txs --cache 512 --ipcpath ~/Library/Ethereum/geth.ipc --networkid 12345 --datadir "node/" --nodiscover --mine --miner.threads 4 --miner.noverify --maxpeers 0 --unlock 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd --password node/password

geth-new: geth-init geth

geth-down:
	kill -INT $(shell ps | grep "geth " | grep -v grep | sed -n 1,1p | cut -c1-5)

geth-attach:
	geth attach --datadir ~/Library/Ethereum
# 	eth.getBalance("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd")
# 	eth.getBalance("0xF01813E4B85e178A83e29B8E7bF26BD830a25f32")
# 	eth.getBalance("0xdd6B972ffcc631a62CAE1BB9d80b7ff429c8ebA4")

geth-new-account:
	geth --datadir "node/" account new
#	If you run this command, add the new account address to the genesis file.
#	To see the new balances, you need to restart the geth service.

# ==============================================================================
# These commands provide Go related support.

tidy:
	go mod tidy
	go mod vendor
