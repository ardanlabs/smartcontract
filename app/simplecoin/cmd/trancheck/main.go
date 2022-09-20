package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	scoin "github.com/ardanlabs/smartcontract/app/simplecoin/contract/go"
	"github.com/ethereum/go-ethereum/common"
)

const (
	keyStoreFile     = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase       = "123"
	coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	ethereum, err := ethereum.New(ctx, ethereum.NetworkLocalhost, keyStoreFile, passPhrase)
	if err != nil {
		return err
	}

	fmt.Println("\nInput Values")
	fmt.Println("----------------------------------------------------")
	fmt.Println("fromAddress:", ethereum.Address())

	// =========================================================================

	contractIDBytes, err := os.ReadFile("zarf/tmp/simplecoin/SCOIN_CID")
	if err != nil {
		return fmt.Errorf("importing SCOIN_CID: %v\n", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return fmt.Errorf("need to export the SCOIN_CID")
	}
	fmt.Println("contractID:", contractID)

	scoinCon, err := scoin.NewScoin(common.HexToAddress(contractID), ethereum.RawClient())
	if err != nil {
		return fmt.Errorf("new contract: %w", err)
	}

	// =========================================================================

	balance, err := scoinCon.CoinBalance(nil, common.HexToAddress("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd"))
	if err != nil {
		return err
	}
	fmt.Println("balance 0x6327:", currency.Wei2GWei(balance), "GWei")

	balance, err = scoinCon.CoinBalance(nil, common.HexToAddress("0x8e113078adf6888b7ba84967f299f29aece24c55"))
	if err != nil {
		return err
	}
	fmt.Println("balance 0x8e11:", currency.Wei2GWei(balance), "GWei")

	return nil
}
