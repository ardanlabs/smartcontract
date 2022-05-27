package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

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

	const rawurl = smart.NetworkLocalhost
	sc, err := smart.Connect(ctx, rawurl, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress:", sc.Account)

	// =========================================================================

	contract, err := newContract(ctx, sc.Client)
	if err != nil {
		return err
	}

	// =========================================================================

	startingBalance, err := sc.Client.BalanceAt(ctx, sc.Account, nil)
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

	to := common.HexToAddress("0x8e113078adf6888b7ba84967f299f29aece24c55")

	// =========================================================================

	if rawurl == smart.NetworkLocalhost {
		sink := make(chan *scoin.ScoinTransfer, 100)
		if _, err := contract.WatchTransfer(nil, sink, []common.Address{sc.Account}, []common.Address{to}); err != nil {
			return err
		}

		go func() {
			event := <-sink
			fmt.Println("\nEvents")
			fmt.Println("----------------------------------------------------")
			fmt.Println("tx event", event)
		}()
	}

	// =========================================================================

	tx, err := contract.Transfer(tranOpts, to, big.NewInt(100))
	if err != nil {
		return err
	}
	smart.DisplayTransaction(tx)

	receipt, err := sc.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	smart.DisplayTransactionReceipt(receipt, tx)

	return nil
}

// newContract constructs a SimpleCoin contract.
func newContract(ctx context.Context, client *ethclient.Client) (*scoin.Scoin, error) {
	data, err := os.ReadFile("zarf/smart/scoin.env")
	if err != nil {
		return nil, fmt.Errorf("readfile: %w", err)
	}
	contractID := string(data)
	fmt.Println("contractID:", contractID)

	contract, err := scoin.NewScoin(common.HexToAddress(contractID), client)
	if err != nil {
		return nil, fmt.Errorf("NewScoin: %w", err)
	}

	return contract, nil
}
