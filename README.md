# Smart Contract

This is a work in progress to understand how to write, debug, and maintain Ethereum smart contracts. This repository is expected to have several apps that help to discover and understand the Ethereum and smart contract development. Please look at the makefile for more details and help.

The following will guide you in deploying the basic smart contract to a local Ethereum service and execute its API.

## Installing Tooling

You must install Ethereum and Solidity into your development environment. If you are using MacOS or Linux, the makefile has brew commands to help you with this. If you can't use brew, there are links to documents in the makefile.

makefile
```
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
```

## Running Ethereum

Ethereum is configured to run in a developer mode that will allow for the deployment and execution of smart contracts. Use the following `make` command to start the Ethereum service. Make sure you start Ethereum in a dedicated terminal session.

```
$ make geth-up
```

To stop the Ethereum service, use the down command.

```
$ make geth-down
```

## Building Smart Contract

The smart contract can be found in `app/basic/contracts/src/store.sol` and is a smart contract written in the Ethereum programming language called Solidity. It will be helpful to install the Solidity extension by Juan Blanco. To compile and build the smart contract, run the following make command.

```
$ make basic-build
```

That will produce an `abi`, `bin`, and `go` file inside the `app/basic/contracts/store` directory. Any time code for the smart contract is changed, it must be built again for the Go applications.

## Deploying Smart Contract

With the Ethereum service up and running, run the following make command to deploy the smart contract.

```
$ make basic-deploy
```

If this is successful, the follow output should result.

```
go run app/basic/cmd/deploy/main.go
address: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
next nonce: 8
gas price: 417115737
deploy store contract: &{0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd 8 0x436a4c0 0 417115737 <nil> <nil> 300000 context.Background false}
address : 0x0BA052bAeb8925Ac8b480a291F75Ff0dD2c4297c
tx hash : 0x55555675c0035271998f6ca98ae2d3c3a8346c50aee0be7178bd71a1cfc0549c
run this: export CONTRACT="0x0BA052bAeb8925Ac8b480a291F75Ff0dD2c4297c"
```

Make sure you execute the export command in the terminal.

```
export CONTRACT="0x0BA052bAeb8925Ac8b480a291F75Ff0dD2c4297c"
```

Each time you deploy the contract, you will get a new contract ID. Every code change needs to be mined into a new block and therefore the API moves. Be aware what version of the API you are using.

## Executing Smart Contract API

To validate everything is working, run the follow make command.

```
$ make basic-play
```

When you run this command, attempt to update the contracts map with the key `name` and value `ale`. You should see the following output if everything is working correctly.

```
go run app/basic/cmd/play/main.go
contractID: 0x0BA052bAeb8925Ac8b480a291F75Ff0dD2c4297c
version: 1.1
address: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
next nonce: 9
gas price: 373867576
tx sent: 0x119a8a44ddb429ef5652261d07cd74b6bf85dcee629d2adde6cd210b26ebaff4
value: ale
```

In this case, a transaction is sent and the value returned is `ale`.
