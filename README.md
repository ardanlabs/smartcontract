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

If this is successful, the following output should be similar to.

```
go run app/basic/cmd/deploy/main.go
fromAddress: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd

Transaction Details
----------------------------------------------------
tx sent            : 0xafc894991575d94d1cbc004784e6d6bd4201e56fedfbe417e250f2b0a9c624ef
tx gas offer price : 0.431556798 GWei
tx gas limit       : 300000
tx value           : 0.0 GWei
tx max price       : 129467.39400000 GWei (Gas Offer Price * Max Gas Allowed)
tx max price       : 0.33014085 USD

waiting for transaction to be mined...

Receipt Details
----------------------------------------------------
re status          : 1
re gas used        : 299044
final cost         : 129054.471101112 GWei (Gas Offer Price * Gas Used)
final cost         : 0.32908770 USD

Balance
----------------------------------------------------
balance before     : 115792089237316195423570985008687907853269984665640564039457583061627.121056760 GWei
balance after      : 115792089237316195423570985008687907853269984665640564039457583061619.525339160 GWei
balance diff price : 7.595717600 GWei
balance diff price : 0.00001785 USD
```

Each time you deploy the contract, you will get a new contract ID. Every code change needs to be mined into a new block and therefore the API moves. Be aware what version of the API you are using.

## Executing Smart Contract API

To validate everything is working, run the follow make command.

```
$ make basic-write
```

When you run this command, attempt to update the contracts map with the key `name` and value `brianna`. You should see the following output if everything is working correctly.

```
address: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
contractID: 0x46D9cE545007E4E694016b3959D4dC11D96F3F2b
version: 1.1

Transaction Details
----------------------------------------------------
tx sent            : 0x1d4581b3ec85ae60b331f8d21752e1805ad86118577d1965b06cfc2b407fcb95
tx gas offer price : 0.431553329 GWei
tx gas limit       : 250000
tx value           : 0.0 GWei
tx max price       : 107888.332250000 GWei (Gas Offer Price * Max Gas Allowed)
tx max price       : 0.27511440 USD

waiting for transaction to be mined...

Receipt Details
----------------------------------------------------
re status          : 1
re gas used        : 45795
final cost         : 19762.984701555 GWei (Gas Offer Price * Gas Used)
final cost         : 0.05039310 USD

Logs
----------------------------------------------------

Balance
----------------------------------------------------
balance before     : 115792089237316195423570985008687907853269984665640564039457583061619.525339160 GWei
balance after      : 115792089237316195423570985008687907853269984665640564039457583061618.499347980 GWei
balance diff price : 1.25991180 GWei
balance diff price : 0.00000255 USD
```

To see if the value was written to the map for that key, run the follow make command.

```
$ make basic-read
```

You should see the following output if everything is working correctly.

```
go run app/basic/cmd/read/main.go
address: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
contractID: 0x1da2185Ff21DE4291E8690407A4de5459E09de16
version: 1.1
value: brianna
```

## What's Next

Check out the `simplecoin` app which is a more complext smart contract. This is a work in progress.
The steps are similar: `scoin-build`, `scoin-deploy`, `scoin-transfer` and `scoin-trancheck`.
