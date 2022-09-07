package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	store "github.com/ardanlabs/smartcontract/app/basic/contract/go"
	"github.com/ardanlabs/smartcontract/foundation/smart/contract"
	"github.com/ardanlabs/smartcontract/foundation/smart/currency"
	"github.com/ethereum/go-ethereum/log"
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

func run() (err error) {
	ctx := context.Background()

	client, err := contract.NewClient(ctx, contract.NetworkLocalhost, keyStoreFile, passPhrase)
	if err != nil {
		return err
	}

	fmt.Println("\nInput Values")
	fmt.Println("----------------------------------------------------")
	fmt.Println("fromAddress:", client.Address())

	// =========================================================================

	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter()
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("oneETHToUSD:", oneETHToUSD)
	fmt.Println("oneUSDToETH:", oneUSDToETH)

	// =========================================================================

	startingBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, dErr := client.CurrentBalance(ctx)
		if dErr != nil {
			err = dErr
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	const gasLimit = 1600000
	const valueGwei = 0.0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := store.DeployStore(tranOpts, client.ContractBackend())
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())
	fmt.Printf("export CONTRACT_ID=%s\n", address.Hex())

	// =========================================================================

	clientWait, err := contract.NewClient(ctx, contract.NetworkLocalhost, keyStoreFile, passPhrase)
	if err != nil {
		return err
	}

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")
	log.Root().SetHandler(log.StdoutHandler)

	receipt, err := clientWait.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
