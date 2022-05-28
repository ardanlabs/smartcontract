package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"

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

	contract, err := newContract(ctx, client)
	if err != nil {
		return err
	}

	version, err := contract.Version(nil)
	if err != nil {
		return err
	}
	fmt.Println("version:", version)

	// =========================================================================

	var key [32]byte
	copy(key[:], []byte("name"))

	var result [32]byte
	result, err = contract.Items(nil, key)
	if err != nil {
		return err
	}

	fmt.Println("value:", string(result[:]))

	return nil
}

// newContract constructs a SimpleCoin contract.
func newContract(ctx context.Context, client *smart.Client) (*store.Store, error) {
	data, err := os.ReadFile("zarf/smart/basic.env")
	if err != nil {
		return nil, fmt.Errorf("readfile: %w", err)
	}
	contractID := string(data)
	fmt.Println("contractID:", contractID)

	contract, err := store.NewStore(common.HexToAddress(contractID), client.ContractBackend())
	if err != nil {
		return nil, fmt.Errorf("NewStore: %w", err)
	}

	return contract, nil
}
