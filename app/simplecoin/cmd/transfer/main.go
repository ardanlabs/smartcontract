package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/simplecoin/contract/go/simplecoin"
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

	eth, err := ethereum.New(ctx, ethereum.NetworkLocalhost, keyStoreFile, passPhrase)
	if err != nil {
		return err
	}

	fmt.Println("\nInput Values")
	fmt.Println("----------------------------------------------------")
	fmt.Println("fromAddress:", eth.Address())

	// =========================================================================

	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter()
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("oneETHToUSD:", oneETHToUSD)
	fmt.Println("oneUSDToETH:", oneUSDToETH)

	// =========================================================================

	contractIDBytes, err := os.ReadFile("zarf/ethereum/simplecoin.cid")
	if err != nil {
		return fmt.Errorf("importing simplecoin.cid file: %w", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return errors.New("need to export the simplecoin.cid file")
	}
	fmt.Println("contractID:", contractID)

	scoinCon, err := simplecoin.NewSimplecoin(common.HexToAddress(contractID), eth.RawClient())
	if err != nil {
		return fmt.Errorf("new contract: %w", err)
	}

	// =========================================================================

	startingBalance, err := eth.Balance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, dErr := eth.Balance(ctx)
		if dErr != nil {
			err = dErr
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	const gasLimit = 300000
	const valueGwei = 0
	tranOpts, err := eth.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	to := common.HexToAddress("0x8e113078adf6888b7ba84967f299f29aece24c55")

	// =========================================================================

	sink := make(chan *simplecoin.SimplecoinEventTransfer, 100)
	if _, err := scoinCon.WatchEventTransfer(nil, sink, []common.Address{eth.Address()}, []common.Address{to}); err != nil {
		return err
	}

	go func() {
		event := <-sink
		fmt.Println("\nEvents")
		fmt.Println("----------------------------------------------------")
		fmt.Println("tx event", event)
	}()

	// =========================================================================

	tx, err := scoinCon.Transfer(tranOpts, to, big.NewInt(100))
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := eth.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
