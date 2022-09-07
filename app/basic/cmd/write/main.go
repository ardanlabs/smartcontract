package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	store "github.com/ardanlabs/smartcontract/app/basic/contract/go"
	"github.com/ardanlabs/smartcontract/foundation/smart/contract"
	"github.com/ardanlabs/smartcontract/foundation/smart/currency"
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

	contractID := os.Getenv("CONTRACT_ID")
	if contractID == "" {
		return fmt.Errorf("need to export the CONTRACT_ID")
	}
	fmt.Println("contractID:", contractID)

	storeCon, err := store.NewStore(common.HexToAddress(contractID), client.ContractBackend())
	if err != nil {
		return fmt.Errorf("new contract: %w", err)
	}

	version, err := storeCon.Version(nil)
	if err != nil {
		return err
	}
	fmt.Println("version:", version)

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

	const gasLimit = 250000
	const valueGwei = 0
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("brianna"))

	tx, err := storeCon.SetItem(tranOpts, key, value)
	if err != nil {
		log.Fatal("SetItem ERROR:", err)
	}
	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
