package main

import (
	"context"
	"fmt"
	"log"
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

	client, _, err := smart.Connect()
	if err != nil {
		return err
	}

	// =========================================================================

	scoin, contractID, err := newScoin(ctx, client)
	if err != nil {
		return err
	}
	fmt.Println("contractID:", contractID)

	// =========================================================================

	balance, err := scoin.CoinBalance(nil, common.HexToAddress("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd"))
	if err != nil {
		return err
	}
	fmt.Println("balance 0x6327:", smart.Wei2Eth(balance))

	balance, err = scoin.CoinBalance(nil, common.HexToAddress("0x8e113078adf6888b7ba84967f299f29aece24c55"))
	if err != nil {
		return err
	}
	fmt.Println("balance 0x8e11:", smart.Wei2Eth(balance))

	return nil
}

// newScoin constructs a SimpleCoin value for smart contract API access.
func newScoin(ctx context.Context, client *ethclient.Client) (*scoin.Scoin, string, error) {
	data, err := os.ReadFile("contract.env")
	if err != nil {
		return nil, "", fmt.Errorf("readfile: %w", err)
	}
	contractID := string(data)

	contract := common.HexToAddress(contractID)
	scoin, err := scoin.NewScoin(contract, client)
	if err != nil {
		return nil, "", fmt.Errorf("NewScoin: %w", err)
	}

	return scoin, contractID, nil
}
