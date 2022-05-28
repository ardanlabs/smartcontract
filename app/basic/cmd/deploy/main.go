package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ardanlabs/smartcontract/app/basic/contracts/store"
	"github.com/ardanlabs/smartcontract/foundation/smartcontract/smart"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()

	client, err := smart.Connect(ctx, smart.NetworkLocalhost, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress:", client.Account)

	// =========================================================================

	startingBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		return err
	}
	defer client.DisplayBalanceSheet(ctx, startingBalance)

	// =========================================================================

	const gasLimit = 300000
	const valueGwei = 0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := store.DeployStore(tranOpts, client.ContractBackend())
	if err != nil {
		return err
	}
	client.DisplayTransaction(tx)

	if err := os.WriteFile("zarf/smart/basic.env", []byte(address.Hex()), 0666); err != nil {
		return err
	}

	receipt, err := client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	client.DisplayTransactionReceipt(receipt, tx)

	return nil
}
