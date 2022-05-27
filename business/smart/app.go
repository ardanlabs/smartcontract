package smart

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// DisplayBalanceSheet is a convenience function for application level support.
func DisplayBalanceSheet(ctx context.Context, sc *SmartContract, startingBalance *big.Int) {
	endingBalance, err := sc.CurrentBalance(ctx)
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
func DisplayTransaction(tx *types.Transaction) {
	tcd := CalculateTranCostDetails(tx)
	PrintTranCostDetails(tcd)
}

// DisplayTransactionReceipt is a convenience function for application level support.
func DisplayTransactionReceipt(receipt *types.Receipt, tx *types.Transaction) {
	rcd := CalculateReceiptCostDetails(receipt, tx)
	PrintReceiptCostDetails(rcd)
	PrintLogs(ExtractLogs(receipt))
}
