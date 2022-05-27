package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	sc, err := smart.Connect(ctx, smart.NetworkLocalhost, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress:", sc.Account)

	// =========================================================================

	startingBalance, err := sc.CurrentBalance(ctx)
	if err != nil {
		return err
	}
	defer smart.DisplayBalanceSheet(ctx, sc, startingBalance)

	// =========================================================================

	const gasLimit = 300000
	const valueGwei = 0
	tranOpts, err := sc.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := store.DeployStore(tranOpts, sc.Client)
	if err != nil {
		return err
	}
	smart.DisplayTransaction(tx)

	if err := os.WriteFile("zarf/smart/basic.env", []byte(address.Hex()), 0666); err != nil {
		return err
	}

	receipt, err := sc.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	smart.DisplayTransactionReceipt(receipt, tx)

	return nil
}
