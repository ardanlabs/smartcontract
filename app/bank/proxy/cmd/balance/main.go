package main

import (
	"context"
	"fmt"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	proxy "github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go/proxy"
	"github.com/ethereum/go-ethereum/common"
	"os"
)

const (
	ownerStoreFile   = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase       = "123" // All three accounts use the same passphrase
	coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() (err error) {
	ctx := context.Background()

	ethereum, err := ethereum.New(ctx, ethereum.NetworkLocalhost, ownerStoreFile, passPhrase)
	if err != nil {
		return err
	}

	fmt.Println("\nInput Values")
	fmt.Println("----------------------------------------------------")
	fmt.Println("fromAddress:", ethereum.Address())

	// =========================================================================

	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter()
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("oneETHToUSD:", oneETHToUSD)
	fmt.Println("oneUSDToETH:", oneUSDToETH)

	// =========================================================================

	startingBalance, err := ethereum.Balance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, dErr := ethereum.Balance(ctx)
		if dErr != nil {
			err = dErr
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	tranOpts, err := ethereum.NewCallOpts(ctx)
	if err != nil {
		return err
	}

	// =========================================================================

	contractID := os.Getenv("PROXY_CONTRACT_ID")
	if contractID == "" {
		return fmt.Errorf("need to export the PROXY_CONTRACT_ID")
	}
	fmt.Println("contractID:", contractID)

	proxyContract, err := proxy.NewBankproxy(common.HexToAddress(contractID), ethereum.RawClient())
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	balance, err := proxyContract.Balance(tranOpts)
	if err != nil {
		return err
	}
	fmt.Printf("account balance: %v", balance)

	return nil
}
