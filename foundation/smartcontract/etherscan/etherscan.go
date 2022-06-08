// Package etherscan provides Go access to the Etherscan API.
// Documentation: https://docs.etherscan.io/api-endpoints
package etherscan

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"time"
)

// These are the current networks the API works against.
const (
	NetworkMainnet = "api.etherscan.io"
	NetworkGoerli  = "api-goerli.etherscan.io"
	NetworkSepolia = "api-sepolia.etherscan.io"
)

// This provides a default client configuration, but it's recommended
// this is replaced by the user with application specific settings using
// the WithClient function at the time a GraphQL is constructed.
var defaultClient = http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// Etherscan represents support for accessing the Etherscan API.
type Etherscan struct {
	apiKey  string
	network string
	headers map[string]string
	client  *http.Client
}

// New constructs a new Etherscan for API access.
func New(apiKey string, options ...func(eth *Etherscan)) *Etherscan {
	gql := Etherscan{
		apiKey:  apiKey,
		network: NetworkMainnet,
		headers: make(map[string]string),
		client:  &defaultClient,
	}

	for _, option := range options {
		option(&gql)
	}

	return &gql
}

// WithClient adds a custom client for processing requests. It's recommend
// to not use the default client and provide your own.
func WithClient(client *http.Client) func(eth *Etherscan) {
	return func(eth *Etherscan) {
		eth.client = client
	}
}

// WithHeader adds a key/value pair to the request header for all calls made to
// the host. This is for things like authentication or application specific needs.
// These headers are already included:
// "Cache-Control": "no-cache", "Content-Type": "application/json", "Accept": "application/json"
func WithHeader(key string, value string) func(eth *Etherscan) {
	return func(eth *Etherscan) {
		if key != "" {
			eth.headers[key] = value
		}
	}
}

// WithNetwork changes the default network which is Mainnet.
func WithNetwork(network string) func(eth *Etherscan) {
	return func(eth *Etherscan) {
		eth.network = network
	}
}

// =============================================================================

// Gas represents the result from the gastracker/gasoracle call. All
// values are in Wei.
type Gas struct {
	LastBlock       int64
	SafeGasPrice    *big.Int
	ProposeGasPrice *big.Int
	FastGasPrice    *big.Int
	SuggestBaseFee  *big.Int
}

// GasTracker returns the current gas in GWei that should be proposed on the
// configured network.
func (eth *Etherscan) GasTracker(ctx context.Context) (Gas, error) {
	var gas struct {
		Status  string
		Message string
		Result  struct {
			LastBlock       string
			SafeGasPrice    string
			ProposeGasPrice string
			FastGasPrice    string
			SuggestBaseFee  string
		}
	}

	if err := eth.send(ctx, "gastracker", "gasoracle", &gas); err != nil {
		return Gas{}, err
	}

	lastBlock, err := strconv.ParseInt(gas.Result.LastBlock, 10, 64)
	if err != nil {
		return Gas{}, fmt.Errorf("parse last block: %w", err)
	}

	safeGasPrice, err := strconv.ParseFloat(gas.Result.SafeGasPrice, 64)
	if err != nil {
		return Gas{}, fmt.Errorf("safe gas price: %w", err)
	}

	bigSafeGasPrice := big.NewFloat(safeGasPrice)
	bigSafeGasPrice.Mul(bigSafeGasPrice, big.NewFloat(1e9))
	bigSafeGasPriceInt := new(big.Int)
	bigSafeGasPrice.Int(bigSafeGasPriceInt)

	proposeGasPrice, err := strconv.ParseFloat(gas.Result.ProposeGasPrice, 64)
	if err != nil {
		return Gas{}, fmt.Errorf("propose gas price: %w", err)
	}

	bigproposeGasPrice := big.NewFloat(proposeGasPrice)
	bigproposeGasPrice.Mul(bigproposeGasPrice, big.NewFloat(1e9))
	bigproposeGasPriceInt := new(big.Int)
	bigproposeGasPrice.Int(bigproposeGasPriceInt)

	fastGasPrice, err := strconv.ParseFloat(gas.Result.FastGasPrice, 64)
	if err != nil {
		return Gas{}, fmt.Errorf("fast gas price: %w", err)
	}

	bigfastGasPrice := big.NewFloat(fastGasPrice)
	bigfastGasPrice.Mul(bigfastGasPrice, big.NewFloat(1e9))
	bigfastGasPriceInt := new(big.Int)
	bigfastGasPrice.Int(bigfastGasPriceInt)

	suggestBaseFee, err := strconv.ParseFloat(gas.Result.SuggestBaseFee, 64)
	if err != nil {
		return Gas{}, fmt.Errorf("suggest gas price: %w", err)
	}

	bigsuggestBaseFee := big.NewFloat(suggestBaseFee)
	bigsuggestBaseFee.Mul(bigsuggestBaseFee, big.NewFloat(1e9))
	bigsuggestBaseFeeInt := new(big.Int)
	bigsuggestBaseFee.Int(bigsuggestBaseFeeInt)

	result := Gas{
		LastBlock:       lastBlock,
		SafeGasPrice:    bigSafeGasPriceInt,
		ProposeGasPrice: bigproposeGasPriceInt,
		FastGasPrice:    bigfastGasPriceInt,
		SuggestBaseFee:  bigsuggestBaseFeeInt,
	}

	return result, nil
}

// =============================================================================

// send performs the actual execution of a request against the specified network.
func (eth *Etherscan) send(ctx context.Context, module string, action string, response interface{}) error {
	url := fmt.Sprintf("https://%s/api?module=%s&action=%s&apikey=%s", eth.network, module, action, eth.apiKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("graphql create request error: %w", err)
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	for key, value := range eth.headers {
		req.Header.Set(key, value)
	}

	resp, err := eth.client.Do(req)
	if err != nil {
		return fmt.Errorf("error: client: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: status code: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(response)
}
