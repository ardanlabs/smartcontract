package currency

import (
	"bytes"
	"fmt"
	"math/big"

	ethUnit "github.com/DeOne4eg/eth-unit-converter"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// These represents defaults if the API is not working.
// Values represent pricing on August 27th, 2022.
var (
	defaultOneETHToUSD = big.NewFloat(1503.280164057658)
	defaultOneUSDToETH = big.NewFloat(0.000665206530956729)
)

// Converter maintains the information required for ETH conversions.
type Converter struct {
	coinMarketCapKey string
	oneETHToUSD      *big.Float
	oneUSDToETH      *big.Float
	oneGWeiToUSD     *big.Float
	oneUSDToGWei     *big.Float
}

// NewConverter constructs a conversion for working with ETH and USD.
func NewConverter(coinMarketCapKey string) (*Converter, error) {
	oneETHToUSD, err := captureETH2USD(coinMarketCapKey)
	if err != nil {
		return nil, fmt.Errorf("captureETH2USD: %w", err)
	}

	oneUSDToETH, err := captureUSD2ETH(coinMarketCapKey)
	if err != nil {
		return nil, fmt.Errorf("captureUSD2ETH: %w", err)
	}

	oneGWeiToUSD := big.NewFloat(0).SetPrec(1024).Mul(oneETHToUSD, big.NewFloat(0.000000001))
	oneUSDToGWei := big.NewFloat(0).SetPrec(1024).Mul(oneUSDToETH, big.NewFloat(1000000000))

	return &Converter{
		coinMarketCapKey: coinMarketCapKey,
		oneETHToUSD:      oneETHToUSD,
		oneUSDToETH:      oneUSDToETH,
		oneGWeiToUSD:     oneGWeiToUSD,
		oneUSDToGWei:     oneUSDToGWei,
	}, nil
}

// NewDefaultConverter can be used if the API is failing to set reasonable
// defaults. Also good for tests.
func NewDefaultConverter() *Converter {
	oneGWeiToUSD := big.NewFloat(0).SetPrec(1024).Mul(defaultOneETHToUSD, big.NewFloat(0.000000001))
	oneUSDToGWei := big.NewFloat(0).SetPrec(1024).Mul(defaultOneUSDToETH, big.NewFloat(1000000000))

	return &Converter{
		oneETHToUSD:  defaultOneETHToUSD,
		oneUSDToETH:  defaultOneUSDToETH,
		oneGWeiToUSD: oneGWeiToUSD,
		oneUSDToGWei: oneUSDToGWei,
	}
}

// Values returns the currency values being used.
func (c *Converter) Values() (oneETHToUSD *big.Float, oneUSDToETH *big.Float) {
	return c.oneETHToUSD, c.oneUSDToETH
}

// Wei2USD converts Wei to USD.
func (c *Converter) Wei2USD(amountWei *big.Int) string {
	unit := ethUnit.NewWei(amountWei)
	gWeiAmount := unit.GWei()

	return c.GWei2USD(gWeiAmount)
}

// GWei2USD converts GWei to USD.
func (c *Converter) GWei2USD(amountGWei *big.Float) string {
	cost := big.NewFloat(0).SetPrec(1024).Mul(amountGWei, c.oneGWeiToUSD)
	costFloat, _ := cost.Float64()

	return fmt.Sprintf("%.2f", costFloat)
}

// USD2Wei converts USD to Wei.
func (c *Converter) USD2Wei(amountUSD *big.Float) *big.Int {
	gwei := big.NewFloat(0).SetPrec(1024).Mul(amountUSD, c.oneUSDToGWei)
	v, _ := big.NewFloat(0).SetPrec(1024).Mul(gwei, big.NewFloat(1e9)).Int64()

	return big.NewInt(0).SetInt64(v)
}

// USD2GWei converts USD to GWei.
func (c *Converter) USD2GWei(amountUSD *big.Float) *big.Float {
	return big.NewFloat(0).SetPrec(1024).Mul(amountUSD, c.oneUSDToGWei)
}

// CalculateTransactionDetails performs calculations on the transaction.
func (c *Converter) CalculateTransactionDetails(tx *types.Transaction) TransactionDetails {
	return TransactionDetails{
		Hash:              tx.Hash().Hex(),
		Nonce:             tx.Nonce(),
		GasLimit:          tx.Gas(),
		GasOfferPriceGWei: Wei2GWei(tx.GasPrice()).String(),
		Value:             Wei2GWei(tx.Value()).String(),
		MaxGasPriceGWei:   Wei2GWei(tx.Cost()).String(),
		MaxGasPriceUSD:    c.Wei2USD(tx.Cost()),
	}
}

// CalculateReceiptDetails performs calculations on the receipt.
func (c *Converter) CalculateReceiptDetails(receipt *types.Receipt, gasPrice *big.Int) ReceiptDetails {
	cost := big.NewInt(0).Mul(big.NewInt(int64(receipt.GasUsed)), gasPrice)

	return ReceiptDetails{
		Status:        receipt.Status,
		GasUsed:       receipt.GasUsed,
		GasPriceGWei:  Wei2GWei(gasPrice).String(),
		GasPriceUSD:   c.Wei2USD(gasPrice),
		FinalCostGWei: Wei2GWei(cost).String(),
		FinalCostUSD:  c.Wei2USD(cost),
	}
}

// CalculateBalanceDiff performs calculations on the starting and ending balance.
func (c *Converter) CalculateBalanceDiff(startingBalance *big.Int, endingBalance *big.Int) (BalanceDiff, error) {
	cost := big.NewInt(0).Sub(startingBalance, endingBalance)

	return BalanceDiff{
		BeforeGWei: Wei2GWei(startingBalance).String(),
		AfterGWei:  Wei2GWei(endingBalance).String(),
		DiffGWei:   Wei2GWei(cost).String(),
		DiffUSD:    c.Wei2USD(cost),
	}, nil
}

// FmtBalanceSheet produces a easy to read format of the starting and ending
// balance for the connected account.
func (c *Converter) FmtBalanceSheet(startingBalance *big.Int, endingBalance *big.Int) string {
	diff, err := c.CalculateBalanceDiff(startingBalance, endingBalance)
	if err != nil {
		return ""
	}

	return formatBalanceDiff(diff)
}

// FmtTransaction produces a easy to read format of the specified transaction.
func (c *Converter) FmtTransaction(tx *types.Transaction) string {
	tcd := c.CalculateTransactionDetails(tx)

	return formatTranCostDetails(tcd)
}

// FmtTransactionReceipt produces a easy to read format of the specified receipt.
func (c *Converter) FmtTransactionReceipt(receipt *types.Receipt, gasPrice *big.Int) string {
	rcd := c.CalculateReceiptDetails(receipt, gasPrice)

	var b bytes.Buffer

	b.WriteString(formatReceiptCostDetails(rcd))
	b.WriteString(formatLogs(extractLogs(receipt)))

	return b.String()
}

// =============================================================================

// extractLogs extracts the logging events from the receipt.
func extractLogs(receipt *types.Receipt) []string {
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
