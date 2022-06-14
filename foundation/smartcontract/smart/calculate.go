package smart

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// TransactionDetails provides details about a transaction and it's cost.
type TransactionDetails struct {
	Hash              string
	GasLimit          uint64
	GasOfferPriceGWei string
	Value             string
	MaxGasPriceGWei   string
	MaxGasPriceUSD    string
}

// CalculateTranactionDetails performs calculations on the transaction.
func (*Client) CalculateTransactionDetails(tx *types.Transaction) TransactionDetails {
	return TransactionDetails{
		Hash:              tx.Hash().Hex(),
		GasLimit:          tx.Gas(),
		GasOfferPriceGWei: Wei2GWei(tx.GasPrice()),
		Value:             Wei2GWei(tx.Value()),
		MaxGasPriceGWei:   Wei2GWei(tx.Cost()),
		MaxGasPriceUSD:    Wei2USD(tx.Cost()),
	}
}

// =============================================================================

// ReceiptDetails provides details about a receipt and it's cost.
type ReceiptDetails struct {
	Status        uint64
	GasUsed       uint64
	GasPriceGWei  string
	GasPriceUSD   string
	FinalCostGWei string
	FinalCostUSD  string
}

// CalculateReceiptDetails performs calculations on the receipt.
func (c *Client) CalculateReceiptDetails(receipt *types.Receipt, tx *types.Transaction) ReceiptDetails {
	cost := big.NewInt(0).Mul(big.NewInt(int64(receipt.GasUsed)), tx.GasPrice())

	return ReceiptDetails{
		Status:        receipt.Status,
		GasUsed:       receipt.GasUsed,
		GasPriceGWei:  Wei2GWei(tx.GasPrice()),
		GasPriceUSD:   Wei2USD(tx.GasPrice()),
		FinalCostGWei: Wei2GWei(cost),
		FinalCostUSD:  Wei2USD(cost),
	}
}

// ExtractLogs extracts the logging events from the receipt.
func (*Client) ExtractLogs(receipt *types.Receipt) []string {
	var logs []string

	if len(receipt.Logs) > 0 {

		// We have a particular log event that if we find, we can separate
		// from the rest of the events.
		topicLog := crypto.Keccak256Hash([]byte("EventLog(string)"))

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

// =============================================================================

// MainnetReceiptDetails provides details about a transaction and it's cost based
// on the Mainnet costs.
type MainnetReceiptDetails struct {
	GasUsed       uint64
	GasPriceGWei  string
	GasPriceUSD   string
	FinalCostGWei string
	FinalCostUSD  string
}

// CalculateMainnetReceiptDetails performs calculations on the receipt using Mainnet pricing.
func (c *Client) CalculateMainnetReceiptDetails(receipt *types.Receipt, tx *types.Transaction) MainnetReceiptDetails {
	if c.etherscan == nil {
		return MainnetReceiptDetails{}
	}

	gas, err := c.etherscan.GasTracker(context.Background())
	if err != nil {
		return MainnetReceiptDetails{}
	}

	cost := big.NewInt(0).Mul(big.NewInt(int64(receipt.GasUsed)), gas.ProposeGasPrice)

	return MainnetReceiptDetails{
		GasUsed:       receipt.GasUsed,
		GasPriceGWei:  Wei2GWei(gas.ProposeGasPrice),
		GasPriceUSD:   Wei2USD(gas.ProposeGasPrice),
		FinalCostGWei: Wei2GWei(cost),
		FinalCostUSD:  Wei2USD(cost),
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
func (*Client) CalculateBalanceDiff(ctx context.Context, startingBalance *big.Int, endingBalance *big.Int) (BalanceDiff, error) {
	cost := big.NewInt(0).Sub(startingBalance, endingBalance)

	return BalanceDiff{
		BeforeGWei: Wei2GWei(startingBalance),
		AfterGWei:  Wei2GWei(endingBalance),
		DiffGWei:   Wei2GWei(cost),
		DiffUSD:    Wei2USD(cost),
	}, nil
}
