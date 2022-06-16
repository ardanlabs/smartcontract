package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	crowd "github.com/ardanlabs/smartcontract/app/crowdsale/contract/go"
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

	const gasLimit = 4000000
	const valueGwei = 0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		return err
	}

	// =========================================================================

	startDate := big.NewInt(2003526559)
	endDate := big.NewInt(2003526600)
	weiTokenPrice := big.NewInt(2000000000000000)
	etherInvestmentObjective := big.NewInt(15000)

	address, tx, _, err := crowd.DeployCrowd(tranOpts, client.ContractBackend(), startDate, endDate, weiTokenPrice, etherInvestmentObjective)
	if err != nil {
		return err
	}
	client.DisplayTransaction(tx)

	os.MkdirAll("zarf/smart/", 0755)
	if err := os.WriteFile("zarf/smart/crowd.env", []byte(address.Hex()), 0666); err != nil {
		log.Fatal("cannot write 'crowd.env' ERROR: ", err)
	}

	receipt, err := client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	client.DisplayTransactionReceipt(receipt, tx)

	return nil
}
