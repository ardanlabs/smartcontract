package currency

import (
	"bytes"
	"fmt"
)

// formatTranCostDetails displays the transaction cost details.
func formatTranCostDetails(tcd TransactionDetails) string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nTransaction Details\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	fmt.Fprintf(&b, "hash            : %v\n", tcd.Hash)
	fmt.Fprintf(&b, "nonce           : %v\n", tcd.Nonce)
	fmt.Fprintf(&b, "gas limit       : %v\n", tcd.GasLimit)
	fmt.Fprintf(&b, "gas offer price : %v GWei\n", tcd.GasOfferPriceGWei)
	fmt.Fprintf(&b, "value           : %v GWei\n", tcd.Value)
	fmt.Fprintf(&b, "max gas price   : %v GWei\n", tcd.MaxGasPriceGWei)
	fmt.Fprintf(&b, "max gas price   : %v USD\n", tcd.MaxGasPriceUSD)

	return b.String()
}

// formatReceiptCostDetails displays the receipt cost details.
func formatReceiptCostDetails(rcd ReceiptDetails) string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nReceipt Details\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	fmt.Fprintf(&b, "status          : %v\n", rcd.Status)
	fmt.Fprintf(&b, "gas used        : %v\n", rcd.GasUsed)
	fmt.Fprintf(&b, "gas price       : %v GWei\n", rcd.GasPriceGWei)
	fmt.Fprintf(&b, "gas price       : %v USD\n", rcd.GasPriceUSD)
	fmt.Fprintf(&b, "final gas cost  : %v GWei\n", rcd.FinalCostGWei)
	fmt.Fprintf(&b, "final gas cost  : %v USD\n", rcd.FinalCostUSD)

	return b.String()
}

// formatLogs takes the slice of log information and displays it.
func formatLogs(logs []LogData) string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nLogs\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	for _, log := range logs {
		fmt.Fprintln(&b, log.EventName)
		fmt.Fprintln(&b, log.Data)
	}

	return b.String()
}

// formatBalanceDiff outputs the start and ending balances with difference.
func formatBalanceDiff(bd BalanceDiff) string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nBalance\n")
	fmt.Fprintf(&b, "----------------------------------------------------\n")
	fmt.Fprintf(&b, "balance before  : %v GWei\n", bd.BeforeGWei)
	fmt.Fprintf(&b, "balance after   : %v GWei\n", bd.AfterGWei)
	fmt.Fprintf(&b, "balance diff    : %v GWei\n", bd.DiffGWei)
	fmt.Fprintf(&b, "balance diff    : %v USD\n", bd.DiffUSD)

	return b.String()
}
