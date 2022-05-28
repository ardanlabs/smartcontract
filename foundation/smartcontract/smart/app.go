package smart

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// DisplayBalanceSheet is a convenience function for application level support.
func (c *Client) DisplayBalanceSheet(ctx context.Context, startingBalance *big.Int) {
	endingBalance, err := c.CurrentBalance(ctx)
	if err != nil {
		return
	}
	diff, err := CalculateBalanceDiff(ctx, startingBalance, endingBalance)
	if err != nil {
		return
	}
	PrintBalanceDiff(diff)
}

// DisplayTransaction is a convenience function for application level support.
func (c *Client) DisplayTransaction(tx *types.Transaction) {
	tcd := CalculateTranCostDetails(tx)
	PrintTranCostDetails(tcd)
}

// DisplayTransactionReceipt is a convenience function for application level support.
func (c *Client) DisplayTransactionReceipt(receipt *types.Receipt, tx *types.Transaction) {
	rcd := CalculateReceiptCostDetails(receipt, tx)
	PrintReceiptCostDetails(rcd)
	PrintLogs(ExtractLogs(receipt))
}
