package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ardanlabs/smartcontract/app/simplecoin/contracts/scoin"
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

	// =========================================================================

	balance, err := contract.CoinBalance(nil, common.HexToAddress("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd"))
	if err != nil {
		return err
	}
	fmt.Println("balance 0x6327:", smart.Wei2GWei(balance), "GWei")

	balance, err = contract.CoinBalance(nil, common.HexToAddress("0x8e113078adf6888b7ba84967f299f29aece24c55"))
	if err != nil {
		return err
	}
	fmt.Println("balance 0x8e11:", smart.Wei2GWei(balance), "GWei")

	return nil
}

// newContract constructs a SimpleCoin contract.
func newContract(ctx context.Context, client *smart.Client) (*scoin.Scoin, error) {
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
