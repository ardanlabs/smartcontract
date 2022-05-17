package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ardanlabs/smartcontract/business/smart"
)

func main() {
	ctx := context.Background()

	client, privateKey, err := smart.Connect()
	if err != nil {
		log.Fatal("Connect:", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("address:", fromAddress.String())

	// =========================================================================

	store, contractID, err := smart.NewStore(ctx, client)
	if err != nil {
		log.Fatal("NewStore:", err)
	}
	fmt.Println("contractID:", contractID)

	version, err := store.Version(nil)
	if err != nil {
		log.Fatal("version:", err)
	}
	fmt.Println("version:", version)

	// =========================================================================

	balBefore, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatal("BalanceAt:", err)
	}

	const gasLimit = 30000
	tran, err := smart.NewTransaction(ctx, gasLimit, privateKey, client)
	if err != nil {
		log.Fatal("NewTransaction:", err)
	}

	fmt.Println("transaction:", tran)

	// =========================================================================

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("tommy"))

	tx, err := store.SetItem(tran, key, value)
	if err != nil {
		log.Fatal("set item ERROR:", err)
	}

	fmt.Println("tx sent        :", tx.Hash().Hex())
	fmt.Println("tx gas units   :", tx.Gas())
	fmt.Println("tx gas price   :", smart.Wei2Eth(tx.GasPrice()))
	fmt.Println("tx cost        :", smart.Wei2Eth(tx.Cost()))

	// =========================================================================

	balAfter, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatal("balance at ERROR:", err)
	}

	fmt.Println("balance before :", smart.Wei2Eth(balBefore))
	fmt.Println("balance after  :", smart.Wei2Eth(balAfter))
}
