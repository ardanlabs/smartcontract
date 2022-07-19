package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"

	scoin "github.com/ardanlabs/smartcontract/app/simplecoin/contract/go"
	"github.com/ardanlabs/smartcontract/foundation/smartcontract/smart"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()

	const rawurl = smart.NetworkLocalhost
	client, err := smart.Connect(ctx, rawurl, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress:", client.Account)

	// =========================================================================

	contract, err := newContract(client)
	if err != nil {
		return err
	}

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

	to := common.HexToAddress("0x8e113078adf6888b7ba84967f299f29aece24c55")

	// =========================================================================

	if rawurl == smart.NetworkLocalhost {
		sink := make(chan *scoin.ScoinEventTransfer, 100)
		if _, err := contract.WatchEventTransfer(nil, sink, []common.Address{client.Account}, []common.Address{to}); err != nil {
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
	client.DisplayTransaction(tx)

	receipt, err := client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	client.DisplayTransactionReceipt(receipt, tx)

	return nil
}

// newContract constructs a SimpleCoin contract.
func newContract(client *smart.Client) (*scoin.Scoin, error) {
	data, err := os.ReadFile("zarf/smart/scoin.env")
	if err != nil {
		return nil, fmt.Errorf("readfile: %w", err)
	}
	contractID := string(data)
	fmt.Println("contractID:", contractID)

	contract, err := scoin.NewScoin(common.HexToAddress(contractID), client.ContractBackend())
	if err != nil {
		return nil, fmt.Errorf("NewScoin: %w", err)
	}

	return contract, nil
}
