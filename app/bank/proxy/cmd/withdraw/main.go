package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go/bank"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ownerStoreFile    = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	account1StoreFile = "zarf/ethereum/keystore/UTC--2022-05-13T16-57-20.203544000Z--8e113078adf6888b7ba84967f299f29aece24c55"
	account2StoreFile = "zarf/ethereum/keystore/UTC--2022-05-13T16-59-42.277071000Z--0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"

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

	withdrawTarget := os.Getenv("WITHDRAW_TARGET")
	var ethAccount string

	// Validate the deposit target is valid.
	switch withdrawTarget {
	case "owner":
		ethAccount = ownerStoreFile
	case "account1":
		ethAccount = account1StoreFile
	case "account2":
		ethAccount = account2StoreFile
	default:
		return fmt.Errorf("invalid DEPOSIT_TARGET, must be one of: owner, account1, account2")
	}

	backend, err := ethereum.CreateDialedBackend(ctx, ethereum.NetworkLocalhost)
	if err != nil {
		return err
	}
	defer backend.Close()

	privateKey, err := ethereum.PrivateKeyByKeyFile(ethAccount, passPhrase)
	if err != nil {
		return err
	}

	clt, err := ethereum.NewClient(backend, privateKey)
	if err != nil {
		return err
	}

	fmt.Println("\nInput Values")
	fmt.Println("----------------------------------------------------")
	fmt.Println("fromAddress:", clt.Address())

	// =========================================================================

	converter, err := currency.NewConverter(bank.BankMetaData.ABI, coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter(bank.BankMetaData.ABI)
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("oneETHToUSD:", oneETHToUSD)
	fmt.Println("oneUSDToETH:", oneUSDToETH)

	// =========================================================================

	startingBalance, err := clt.Balance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, dErr := clt.Balance(ctx)
		if dErr != nil {
			err = dErr
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	const gasLimit = 1600000
	const gasPriceGwei = 39.576
	const valueGwei = 0.0
	tranOpts, err := clt.NewTransactOpts(ctx, gasLimit, currency.GWei2Wei(big.NewFloat(gasPriceGwei)), big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	contractIDBytes, err := os.ReadFile("zarf/ethereum/bank.cid")
	if err != nil {
		return fmt.Errorf("importing bank.cid file: %w", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return errors.New("need to export the bank.cid file")
	}
	fmt.Println("contractID:", contractID)

	proxyContract, err := bank.NewBank(common.HexToAddress(contractID), clt.Backend)
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	tx, err := proxyContract.Withdraw(tranOpts)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	// =========================================================================

	receipt, err := clt.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
