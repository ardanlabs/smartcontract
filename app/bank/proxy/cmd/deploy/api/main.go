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
	"github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go/bankapi"
	"github.com/ethereum/go-ethereum/common"
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

	backend, err := ethereum.CreateDialedBackend(ctx, ethereum.NetworkLocalhost)
	if err != nil {
		return err
	}
	defer backend.Close()

	privateKey, err := ethereum.PrivateKeyByKeyFile(keyStoreFile, passPhrase)
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

	address, tx, _, err := bankapi.DeployBankapi(tranOpts, clt.Backend)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())

	// =========================================================================

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")
	log.Root().SetHandler(log.StdoutHandler)

	receipt, err := clt.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))
	log.Root().SetHandler(log.DiscardHandler())

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

	fmt.Println("\nSet This Contract To Bank")
	fmt.Println("----------------------------------------------------")
	fmt.Println("bank id         :", contractID)
	fmt.Println("contract id     :", address.Hex())

	bankContract, err := bank.NewBank(common.HexToAddress(contractID), clt.Backend)
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	tranOpts.Nonce = big.NewInt(0).Add(tranOpts.Nonce, big.NewInt(1))

	tx, err = bankContract.SetContract(tranOpts, common.HexToAddress(address.Hex()))
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	// =========================================================================

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")
	log.Root().SetHandler(log.StdoutHandler)

	receipt, err = clt.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))
	log.Root().SetHandler(log.DiscardHandler())

	// =========================================================================

	callOpts, err := clt.NewCallOpts(ctx)
	if err != nil {
		return err
	}

	version, err := bankContract.Version(callOpts)
	if err != nil {
		return err
	}

	api, err := bankContract.API(callOpts)
	if err != nil {
		return err
	}

	fmt.Println("\nValidate Version and API")
	fmt.Println("----------------------------------------------------")
	fmt.Println("version         :", version)
	fmt.Println("api             :", api)

	return nil
}
