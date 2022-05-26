package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	rawurl := smart.NetworkLocalhost

	client, privateKey, err := smart.Connect(rawurl)
	if err != nil {
		return err
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("address:", fromAddress.String())

	// =========================================================================

	contract, err := newContract(ctx, client)
	if err != nil {
		return err
	}

	// =========================================================================

	startingBalance, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		return err
	}
	defer smart.PrintBalanceDiff(ctx, startingBalance, fromAddress, client)

	// =========================================================================

	const gasLimit = 300000
	tran, err := smart.NewTransaction(ctx, gasLimit, privateKey, client)
	if err != nil {
		return err
	}

	to := common.HexToAddress("0x8e113078adf6888b7ba84967f299f29aece24c55")

	// =========================================================================

	if rawurl == smart.NetworkLocalhost {
		sink := make(chan *scoin.ScoinTransfer, 100)
		if _, err := contract.WatchTransfer(nil, sink, []common.Address{fromAddress}, []common.Address{to}); err != nil {
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

	tx, err := contract.Transfer(tran, to, big.NewInt(100))
	if err != nil {
		return err
	}
	smart.PrintTransaction(tx)

	receipt, err := smart.WaitMined(ctx, tx, fromAddress, client)
	if err != nil {
		return err
	}
	smart.PrintTransactionReceipt(receipt, tx)

	return nil
}

// newContract constructs a SimpleCoin contract.
func newContract(ctx context.Context, client *ethclient.Client) (*scoin.Scoin, error) {
	data, err := os.ReadFile("contract.env")
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
