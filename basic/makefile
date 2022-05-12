# https://goethereumbook.org/smart-contract-deploy/

run-play:
	go run programs/play/main.go

run-deploy:
	go run programs/deploy/main.go

build:
	solc --abi contracts/store/src/store.sol -o contracts/store/abi --overwrite
	solc --bin contracts/store/src/store.sol -o contracts/store/bin --overwrite
	abigen --bin=contracts/store/bin/store.bin --abi=contracts/store/abi/store.abi --pkg=store --out=contracts/store/store.go

# ==============================================================================

run:
	geth --datadir "node/" init "node/genesis.json"
	geth --rpc.allow-unprotected-txs --cache 512 --ipcpath ~/Library/Ethereum/geth.ipc --networkid 12345 --datadir "node/" --nodiscover --mine --miner.threads 4 --miner.noverify --maxpeers 0 --unlock 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd --password node/password

# eth.getBalance("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd")
# eth.getBalance("0xF01813E4B85e178A83e29B8E7bF26BD830a25f32")
# eth.getBalance("0xdd6B972ffcc631a62CAE1BB9d80b7ff429c8ebA4")
attach:
	geth attach --datadir ~/Library/Ethereum

init:
#   Passkey: 123
#   Public address of the key:   0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
#   Path of the secret key file: node/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd
#   geth --datadir "node/" account new
#   Update genesis with this new account.
#   geth --datadir "node/" init "node/genesis.json"	

# ==============================================================================

tidy:
	go mod tidy
	go mod vendor

size:
	du node/