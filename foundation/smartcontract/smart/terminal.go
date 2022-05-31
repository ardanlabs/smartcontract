package smart

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// DisplayBalanceSheet is a convenience function for application level support.
func (c *Client) DisplayBalanceSheet(ctx context.Context, startingBalance *big.Int) {
	endingBalance, err := c.CurrentBalance(ctx)
	if err != nil {
		return
	}
	diff, err := c.CalculateBalanceDiff(ctx, startingBalance, endingBalance)
	if err != nil {
		return
	}
	PrintBalanceDiff(diff)
}

// DisplayTransaction is a convenience function for application level support.
func (c *Client) DisplayTransaction(tx *types.Transaction) {
	tcd := c.CalculateTransactionDetails(tx)
	PrintTranCostDetails(tcd)
}

// DisplayTransactionReceipt is a convenience function for application level support.
func (c *Client) DisplayTransactionReceipt(receipt *types.Receipt, tx *types.Transaction) {
	rcd := c.CalculateReceiptDetails(receipt, tx)
	PrintReceiptCostDetails(rcd)
	mrd := c.CalculateMainnetReceiptDetails(receipt, tx)
	PrintReceiptMainnetDetails(mrd)
	PrintLogs(c.ExtractLogs(receipt))
}

// =============================================================================

// PrintTranCostDetails displays the transaction cost details.
func PrintTranCostDetails(tcd TransactionDetails) {
	fmt.Println("\nTransaction Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("sent            :", tcd.Hash)
	fmt.Println("gas limit       :", tcd.GasLimit)
	fmt.Println("gas offer price :", tcd.GasOfferPriceGWei, "GWei")
	fmt.Println("value           :", tcd.Value, "GWei")
	fmt.Println("max gas price   :", tcd.MaxGasPriceGWei, "GWei")
	fmt.Println("max gas price   :", tcd.MaxGasPriceUSD, "USD")
}

// PrintReceiptCostDetails displays the receipt cost details.
func PrintReceiptCostDetails(rcd ReceiptDetails) {
	fmt.Println("\nReceipt Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("status          :", rcd.Status)
	fmt.Println("gas used        :", rcd.GasUsed)
	fmt.Println("gas price       :", rcd.GasPriceGWei, "GWei")
	fmt.Println("gas price       :", rcd.GasPriceUSD, "USD")
	fmt.Println("final gas cost  :", rcd.FinalCostGWei, "GWei")
	fmt.Println("final gas cost  :", rcd.FinalCostUSD, "USD")
}

// PrintReceiptMainnetDetails displays the receipt cost details.
func PrintReceiptMainnetDetails(rcd MainnetReceiptDetails) {
	fmt.Println("\nReceipt Mainnet Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("gas used        :", rcd.GasUsed)
	fmt.Println("gas price       :", rcd.GasPriceGWei, "GWei")
	fmt.Println("gas price       :", rcd.GasPriceUSD, "USD")
	fmt.Println("final gas cost  :", rcd.FinalCostGWei, "GWei")
	fmt.Println("final gas cost  :", rcd.FinalCostUSD, "USD")
}

// PrintLogs takes the slice of log information and displays it.
func PrintLogs(logs []string) {
	fmt.Println("\nLogs")
	fmt.Println("----------------------------------------------------")
	for _, log := range logs {
		fmt.Println(log)
	}
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
