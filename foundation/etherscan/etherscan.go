// Package etherscan provides Go access to the Etherscan API.
package etherscan

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

/*
// https://docs.etherscan.io/api-endpoints
*/

// These are the current networks the API works against.
const (
	NetworkMainnet = "https://api.etherscan.io/"
	NetworkGoerli  = "https://api-goerli.etherscan.io/"
	NetworkSepolia = "https://api-sepolia.etherscan.io/"
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
	url     string
	headers map[string]string
	client  *http.Client
}

// New constructs a new Etherscan for API access.
func New(network string, options ...func(eth *Etherscan)) *Etherscan {
	gql := Etherscan{
		url:     network,
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

// rawRequest performs the actual execution of a request against the specified network.
func (eth *Etherscan) rawRequest(ctx context.Context, endpoint string, r io.Reader, response interface{}) error {
	// WE NEED TO FIGURE THIS URL THING OUT

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, eth.url, r)
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
		return fmt.Errorf("graphql request error: %w", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("graphql copy error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("graphql op error: status code: %s", resp.Status)
	}

	// FILL IN THE BLANKS

	return nil
}
