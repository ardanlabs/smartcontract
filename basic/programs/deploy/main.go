package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"sc/contracts/store"
)

func main() {
	client, err := ethclient.Dial("/Users/bill/Library/Ethereum/geth.ipc")
	if err != nil {
		log.Fatal("dial ERROR:", err)
	}

	chain, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("chain id ERROR:", err)
	}

	fmt.Println("chain:", chain)

	privateKey, err := privateKey()
	if err != nil {
		log.Fatal("capture private key ERROR:", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("address:", fromAddress.String())

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("pending nonce at ERROR:", err)
	}

	fmt.Println("next nonce:", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("suggest gas price ERROR:", err)
	}

	fmt.Println("gas price:", gasPrice)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	fmt.Println("deploy store contract:", auth)

	address, tx, instance, err := store.DeployStore(auth, client)
	if err != nil {
		log.Fatal("deploy ERROR:", err)
	}

	fmt.Println("address:", address.Hex())   // Capture this information.
	fmt.Println("tx hash:", tx.Hash().Hex()) // Capture this information.

	_ = instance
}

func privateKey() (*ecdsa.PrivateKey, error) {
	inPath := "node/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	password := "123"

	data, err := ioutil.ReadFile(inPath)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}
