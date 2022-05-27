package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

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

	contract, err := newContract(ctx, sc.Client)
	if err != nil {
		return err
	}

	version, err := contract.Version(nil)
	if err != nil {
		return err
	}
	fmt.Println("version:", version)

	// =========================================================================

	startingBalance, err := sc.CurrentBalance(ctx)
	if err != nil {
		return err
	}
	defer smart.DisplayBalanceSheet(ctx, sc, startingBalance)

	// =========================================================================

	const gasLimit = 250000
	const valueGwei = 0
	tranOpts, err := sc.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		return err
	}

	// =========================================================================

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("brianna"))

	tx, err := contract.SetItem(tranOpts, key, value)
	if err != nil {
		log.Fatal("SetItem ERROR:", err)
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
func newContract(ctx context.Context, client *ethclient.Client) (*store.Store, error) {
	data, err := os.ReadFile("zarf/smart/basic.env")
	if err != nil {
		return nil, fmt.Errorf("readfile: %w", err)
	}
	contractID := string(data)
	fmt.Println("contractID:", contractID)

	contract, err := store.NewStore(common.HexToAddress(contractID), client)
	if err != nil {
		return nil, fmt.Errorf("NewStore: %w", err)
	}

	return contract, nil
}
