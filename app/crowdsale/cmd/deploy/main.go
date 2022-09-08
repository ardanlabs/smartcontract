package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	crowd "github.com/ardanlabs/smartcontract/app/crowdsale/contract/go"
	"github.com/ardanlabs/smartcontract/foundation/blockchain/currency"
	"github.com/ardanlabs/smartcontract/foundation/blockchain/ethereum"
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

	const gasLimit = 1600000
	const valueGwei = 0.0
	tranOpts, err := ethereum.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	startDate := big.NewInt(2003526559)
	endDate := big.NewInt(2003526600)
	weiTokenPrice := big.NewInt(2000000000000000)
	etherInvestmentObjective := big.NewInt(15000)

	address, tx, _, err := crowd.DeployCrowd(tranOpts, ethereum.RawClient(), startDate, endDate, weiTokenPrice, etherInvestmentObjective)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())
	fmt.Printf("export CONTRACT_ID=%s\n", address.Hex())

	// =========================================================================

	ethCopy, err := ethereum.Copy(ctx)
	if err != nil {
		return err
	}

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")
	log.Root().SetHandler(log.StdoutHandler)

	receipt, err := ethCopy.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
