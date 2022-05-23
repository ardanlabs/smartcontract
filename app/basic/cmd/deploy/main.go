package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ardanlabs/smartcontract/app/basic/contracts/store"
	"github.com/ardanlabs/smartcontract/business/smart"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()

	client, privateKey, err := smart.Connect()
	if err != nil {
		return err
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("fromAddress:", fromAddress)

	// =========================================================================

	startingBalance, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		return err
	}
	defer smart.PrintBalanceDiff(ctx, startingBalance, fromAddress, client)

	// =========================================================================

	const gasLimit = 300000
	tranOpts, err := smart.NewTransaction(ctx, gasLimit, privateKey, client)
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := store.DeployStore(tranOpts, client)
	if err != nil {
		return err
	}
	smart.PrintTransaction(tx)

	if err := os.WriteFile("contract.env", []byte(address.Hex()), 0666); err != nil {
		return err
	}

	receipt, err := smart.WaitMined(ctx, tx, fromAddress, client)
	if err != nil {
		return err
	}
	smart.PrintTransactionReceipt(receipt, tx)

	return nil
}
