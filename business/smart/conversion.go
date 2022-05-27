package smart

import (
	"fmt"
	"math/big"
)

// These values are used to calculate Wei values in both GWei and USD.
var (
	GWeiConv  = big.NewInt(1_000_000_000)
	GWeiPrice = big.NewFloat(0.00000255) // 5/26/22 at 7pm ET
)

// Wei2USD converts Wei to USD.
func Wei2USD(amount *big.Int) string {

	// Convert the amount in wei to gwei.
	gWei := big.NewInt(0)
	reminder := big.NewInt(0)
	gWei.QuoRem(amount, GWeiConv, reminder)
	gWeiAmount := new(big.Float).SetInt(gWei)

	// Multiple the current price of GWei to the USD.
	cost := big.NewFloat(0).Mul(gWeiAmount, GWeiPrice)
	costFloat, _ := cost.Float64()
	return fmt.Sprintf("%.8f", costFloat)
}

// Wei2GWei converts the wei unit into a GWei for display.
func Wei2GWei(amount *big.Int) string {
	compact_amount := big.NewInt(0)
	reminder := big.NewInt(0)
	divisor := big.NewInt(1e9)
	compact_amount.QuoRem(amount, divisor, reminder)
	return fmt.Sprintf("%s.%s", compact_amount.String(), reminder.String())
}
