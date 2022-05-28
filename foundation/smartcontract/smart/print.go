package smart

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// TranCostDetails provides details about a transaction and it's cost.
type TranCostDetails struct {
	Hash              string
	GasLimit          uint64
	GasOfferPriceGWei string
	Value             string
	MaxGasPriceGWei   string
	MaxGasPriceUSD    string
}

// CalculateTranCostDetails performs calculations on the transaction.
func CalculateTranCostDetails(tx *types.Transaction) TranCostDetails {
	return TranCostDetails{
		Hash:              tx.Hash().Hex(),
		GasLimit:          tx.Gas(),
		GasOfferPriceGWei: Wei2GWei(tx.GasPrice()),
		Value:             Wei2GWei(tx.Value()),
		MaxGasPriceGWei:   Wei2GWei(tx.Cost()),
		MaxGasPriceUSD:    Wei2USD(tx.Cost()),
	}
}

// PrintTranCostDetails displays the transaction cost details.
func PrintTranCostDetails(tcd TranCostDetails) {
	fmt.Println("\nTransaction Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("sent            :", tcd.Hash)
	fmt.Println("gas limit       :", tcd.GasLimit)
	fmt.Println("gas offer price :", tcd.GasOfferPriceGWei, "GWei")
	fmt.Println("value           :", tcd.Value, "GWei")
	fmt.Println("max gas price   :", tcd.MaxGasPriceGWei, "GWei")
	fmt.Println("max gas price   :", tcd.MaxGasPriceUSD, "USD")
}

// =============================================================================

// ReceiptCostDetails performs calculations on the receipt.
type ReceiptCostDetails struct {
	Status        uint64
	GasUsed       uint64
	GasPrice      string
	FinalCostGWei string
	FinalCostUSD  string
}

// CalculateReceiptCostDetails performs calculations on the receipt.
func CalculateReceiptCostDetails(receipt *types.Receipt, tx *types.Transaction) ReceiptCostDetails {
	cost := big.NewInt(0).Mul(big.NewInt(int64(receipt.GasUsed)), tx.GasPrice())

	return ReceiptCostDetails{
		Status:        receipt.Status,
		GasUsed:       receipt.GasUsed,
		GasPrice:      Wei2GWei(tx.GasPrice()),
		FinalCostGWei: Wei2GWei(cost),
		FinalCostUSD:  Wei2USD(cost),
	}
}

// PrintReceiptCostDetails displays the receipt cost details.
func PrintReceiptCostDetails(rcd ReceiptCostDetails) {
	fmt.Println("\nReceipt Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("status          :", rcd.Status)
	fmt.Println("gas used        :", rcd.GasUsed)
	fmt.Println("gas price       :", rcd.GasPrice)
	fmt.Println("final gas cost  :", rcd.FinalCostGWei, "GWei")
	fmt.Println("final gas cost  :", rcd.FinalCostUSD, "USD")
}

// =============================================================================

// ExtractLogs extracts the logging events from the receipt.
func ExtractLogs(receipt *types.Receipt) []string {
	var logs []string

	if len(receipt.Logs) > 0 {

		// We have a particular log event that if we find, we can separate
		// from the rest of the events.
		topicLog := crypto.Keccak256Hash([]byte("Log(string)"))

		// Iterate over the logs and separate.
		for _, v := range receipt.Logs {
			if v.Topics[0] == topicLog {
				l := v.Data[63]
				logs = append(logs, string(v.Data[64:64+l]))
			}
		}
	}

	return logs
}

// PrintLogs takes the slice of log information and displays it.
func PrintLogs(logs []string) {
	fmt.Println("\nLogs")
	fmt.Println("----------------------------------------------------")
	for _, log := range logs {
		fmt.Println(log)
	}
}

// =============================================================================

// BalanceDiff performs calculations on the starting and ending balance.
type BalanceDiff struct {
	BeforeGWei string
	AfterGWei  string
	DiffGWei   string
	DiffUSD    string
}

// CalculateBalanceDiff performs calculations on the starting and ending balance.
func CalculateBalanceDiff(ctx context.Context, startingBalance *big.Int, endingBalance *big.Int) (BalanceDiff, error) {
	cost := big.NewInt(0).Sub(startingBalance, endingBalance)

	return BalanceDiff{
		BeforeGWei: Wei2GWei(startingBalance),
		AfterGWei:  Wei2GWei(endingBalance),
		DiffGWei:   Wei2GWei(cost),
		DiffUSD:    Wei2USD(cost),
	}, nil
}

// PrintBalanceDiff outputs the start and ending balances with difference.
func PrintBalanceDiff(bd BalanceDiff) error {
	fmt.Println("\nBalance")
	fmt.Println("----------------------------------------------------")
	fmt.Println("balance before  :", bd.BeforeGWei, "GWei")
	fmt.Println("balance after   :", bd.AfterGWei, "GWei")
	fmt.Println("balance diff    :", bd.DiffGWei, "GWei")
	fmt.Println("balance diff    :", bd.DiffUSD, "USD")

	return nil
}
