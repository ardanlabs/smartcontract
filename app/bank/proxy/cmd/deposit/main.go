package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	proxy "github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ownerStoreFile    = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	account1StoreFile = "zarf/ethereum/keystore/UTC--2022-05-13T16-57-20.203544000Z--8e113078adf6888b7ba84967f299f29aece24c55"
	account2StoreFile = "zarf/ethereum/keystore/UTC--2022-05-13T16-59-42.277071000Z--0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	account3StoreFile = "zarf/ethereum/keystore/UTC--2022-09-16T16-13-42.375710134Z--7fdfc99999f1760e8dbd75a480b93c7b8386b79a"
	account4StoreFile = "zarf/ethereum/keystore/UTC--2022-09-16T16-13-55.707637523Z--000cf95cb5eb168f57d0befcdf6a201e3e1acea9"

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

	depositAmount := os.Getenv("DEPOSIT_AMOUNT")
	depositTarget := os.Getenv("DEPOSIT_TARGET")
	var ethAccount string

	// Validate the deposit target is valid.
	switch depositTarget {
	case "owner":
		ethAccount = ownerStoreFile
	case "account1":
		ethAccount = account1StoreFile
	case "account2":
		ethAccount = account2StoreFile
	case "account3":
		ethAccount = account3StoreFile
	case "account4":
		ethAccount = account4StoreFile
	default:
		ethAccount = account1StoreFile
	}

	ethereum, err := ethereum.New(ctx, ethereum.NetworkLocalhost, ethAccount, passPhrase)
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

	valueGwei, err := strconv.ParseFloat(depositAmount, 64)
	if err != nil {
		return fmt.Errorf("converting deposit amount to float: %v", err)
	}

	const gasLimit = 1600000
	tranOpts, err := ethereum.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	contractIDBytes, err := os.ReadFile("zarf/tmp/.PROXY_CONTRACT_ID")
	if err != nil {
		return fmt.Errorf("importing PROXY_CONTRACT_ID: %v\n", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return fmt.Errorf("need to export the PROXY_CONTRACT_ID")
	}
	fmt.Println("contractID:", contractID)

	proxyContract, err := proxy.NewBank(common.HexToAddress(contractID), ethereum.RawClient())
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	tx, err := proxyContract.Deposit(tranOpts)
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
