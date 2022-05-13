package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ardanlabs/smartcontract/pkg/smart"
	"github.com/ardanlabs/smartcontract/programs/basic/contracts/store"
)

const contractID = "0x87A061ED19dcA76EC5B01643b054f5eae2730a85"

func main() {
	client, privateKey, err := smart.Connect()
	if err != nil {
		log.Fatal("dial ERROR:", err)
	}

	address := common.HexToAddress(contractID)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal("NewStore ERROR:", err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal("version ERROR:", err)
	}
	fmt.Println("version:", version)

	// =========================================================================

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
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

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("ale"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx sent:", tx.Hash().Hex())

	// There is a delay from the time we set to the time we see. This
	// includes changes.

	var result [32]byte
	for {
		result, err = instance.Items(nil, key)
		if err != nil {
			log.Fatal("Items ERROR:", err)
		}

		if string(result[:]) == string(value[:]) {
			break
		}

		time.Sleep(time.Second)
	}

	fmt.Println("value:", string(result[:]))
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
