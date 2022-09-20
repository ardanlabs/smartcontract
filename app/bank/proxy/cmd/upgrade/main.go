package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	proxy "github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go"
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

func run() (err error) {
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

	contractIDBytes, err := os.ReadFile("zarf/tmp/bank-proxy/BANK_CID")
	if err != nil {
		return fmt.Errorf("importing BANK_CID file: %w", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return errors.New("need to export the BANK_CID file")
	}
	fmt.Println("contractID:", contractID)

	proxyContract, err := proxy.NewBank(common.HexToAddress(contractID), ethereum.RawClient())
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	apiAddressBytes, err := os.ReadFile("zarf/tmp/bank-proxy/BANK_API_CID")
	if err != nil {
		return fmt.Errorf("importing BANK_API_CID file: %w", err)
	}
	apiAddress := string(apiAddressBytes)
	if apiAddress == "" {
		return errors.New("need to export the BANK_API_CID file")
	}
	fmt.Println("apiAddress:", apiAddress)

	tx, err := proxyContract.UpgradeAPI(tranOpts, common.HexToAddress(apiAddress))
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	// =========================================================================

	receipt, err := ethereum.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
