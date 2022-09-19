package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	proxy "github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ownerStoreFile   = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
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

	const gasLimit = 1600000
	const valueGwei = 0.0
	tranOpts, err := ethereum.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	contractIDBytes, err := os.ReadFile("zarf/tmp/.BANK_CID")
	if err != nil {
		return fmt.Errorf("importing BANK_CID: %v\n", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return fmt.Errorf("need to export the BANK_CID")
	}
	fmt.Println("contractID:", contractID)

	proxyContract, err := proxy.NewBank(common.HexToAddress(contractID), ethereum.RawClient())
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	// Sets the winner ID.
	winnerID := os.Getenv("RECONCILE_WINNER_ID")
	if winnerID == "" {
		return fmt.Errorf("need to export the RECONCILE_WINNER_ID")
	}
	fmt.Println("winnerID:", winnerID)

	// Sets the losers IDs.
	losersID := os.Getenv("RECONCILE_LOSERS_ID")
	if losersID == "" {
		return fmt.Errorf("need to export the RECONCILE_LOSERS_ID")
	}
	fmt.Println("losersID:", losersID)

	var losers []common.Address
	for _, id := range strings.Split(losersID, ",") {
		losers = append(losers, common.HexToAddress(id))
	}

	// Sets the anteWei.
	anteWei := os.Getenv("RECONCILE_ANTE_WEI")
	if anteWei == "" {
		return fmt.Errorf("need to export the RECONCILE_ANTE_WEI")
	}
	fmt.Println("anteWei:", anteWei)

	anteWeiInt, err := strconv.Atoi(anteWei)
	if err != nil {
		return fmt.Errorf("unable to parse anteWei into int: %v", err)
	}

	// Sets the gameFeeWei.
	gameFeeWei := os.Getenv("RECONCILE_GAME_FEE_WEI")
	if gameFeeWei == "" {
		return fmt.Errorf("need to export the RECONCILE_GAME_FEE_WEI")
	}
	fmt.Println("gameFeeWei:", gameFeeWei)

	gameFeeWeiInt, err := strconv.Atoi(gameFeeWei)
	if err != nil {
		return fmt.Errorf("unable to parse gameFeeWeiInt into int: %v", err)
	}

	tx, err := proxyContract.Reconcile(tranOpts, common.HexToAddress(winnerID), losers, big.NewInt(int64(anteWeiInt)), big.NewInt(int64(gameFeeWeiInt)))
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
