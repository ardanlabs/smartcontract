package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ardanlabs/smartcontract/app/simplecoin/contracts/scoin"
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

	const gasLimit = 3000000
	const valueGwei = 0
	tranOpts, err := sc.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		return err
	}

	// =========================================================================

	// Start the contract by giving the account deploying 10k smart coins.
	address, tx, _, err := scoin.DeployScoin(tranOpts, sc.Client, big.NewInt(10000))
	if err != nil {
		return err
	}
	smart.DisplayTransaction(tx)

	if err := os.WriteFile("zarf/smart/scoin.env", []byte(address.Hex()), 0666); err != nil {
		log.Fatal("cannot write 'contract.env' ERROR: ", err)
	}

	receipt, err := sc.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	smart.DisplayTransactionReceipt(receipt, tx)

	return nil
}
