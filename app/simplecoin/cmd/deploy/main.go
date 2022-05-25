package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/crypto"

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

	const gasLimit = 3000000
	tran, err := smart.NewTransaction(ctx, gasLimit, privateKey, client)
	if err != nil {
		return err
	}

	// =========================================================================

	// Start the contract by giving the account deploying 10k smart coins.
	address, tx, _, err := scoin.DeployScoin(tran, client, big.NewInt(10000))
	if err != nil {
		return err
	}
	smart.PrintTransaction(tx)

	if err := os.WriteFile("contract.env", []byte(address.Hex()), 0666); err != nil {
		log.Fatal("cannot write 'contract.env' ERROR: ", err)
	}

	receipt, err := smart.WaitMined(ctx, tx, fromAddress, client)
	if err != nil {
		return err
	}
	smart.PrintTransactionReceipt(receipt, tx, client)

	return nil
}
