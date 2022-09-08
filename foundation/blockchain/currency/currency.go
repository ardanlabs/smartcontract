// Package currency provides information on the price of ETH.
package currency

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	ethUnit "github.com/DeOne4eg/eth-unit-converter"
)

// Wei2GWei converts the wei unit into a GWei for display.
func Wei2GWei(amountWei *big.Int) *big.Float {
	unit := ethUnit.NewWei(amountWei)
	return unit.GWei()
}

// GWei2Wei converts the wei unit into a GWei for display.
func GWei2Wei(amountGWei *big.Float) *big.Int {
	unit := ethUnit.NewGWei(amountGWei)
	return unit.Wei()
}

// =============================================================================

// captureETH2USD calls the coinmarketcap API to get the current price in USD of 1 ETH.
func captureETH2USD(coinMarketCapKey string) (*big.Float, error) {
	url := "https://pro-api.coinmarketcap.com/v2/tools/price-conversion?amount=1&symbol=ETH&convert=USD"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Add("X-CMC_PRO_API_KEY", coinMarketCapKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}
	defer resp.Body.Close()

	var result ResponseETH2USD
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("performing request: %s", result.Status.ErrorMessage)
	}

	return big.NewFloat(result.Data[0].Quote.USD.Price), nil
}

// captureETH2USD calls the coinmarketcap API to get the current price in ETH of 1 USD.
func captureUSD2ETH(coinMarketCapKey string) (*big.Float, error) {
	url := "https://pro-api.coinmarketcap.com/v2/tools/price-conversion?amount=1&symbol=USD&convert=ETH"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Add("X-CMC_PRO_API_KEY", coinMarketCapKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}
	defer resp.Body.Close()

	var result ResponseUSD2ETH
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("performing request: %s", result.Status.ErrorMessage)
	}

	return big.NewFloat(result.Data[0].Quote.ETH.Price), nil
}
