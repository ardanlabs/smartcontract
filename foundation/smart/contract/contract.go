// Package contract provides support for executing smart contract APIs.
package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Set of networks supported by the smart package.
const (
	NetworkHTTPLocalhost = "http://localhost:8545"
	NetworkLocalhost     = "zarf/ethereum/geth.ipc"
	NetworkGoerli        = "https://rpc.ankr.com/eth_goerli"
)

// =============================================================================

// Client provides an API for working with smart contracts.
type Client struct {
	network    string
	address    common.Address
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	ethClient  *ethclient.Client
}

// NewClient provides boilerplate for connecting to the geth service using
// an IPC socket created by the geth service on startup.
func NewClient(ctx context.Context, network string, keyPath string, passPhrase string) (*Client, error) {
	ethClient, err := ethclient.Dial(network)
	if err != nil {
		return nil, fmt.Errorf("dial network: %w", err)
	}

	privateKey, err := PrivateKeyByKeyFile(keyPath, passPhrase)
	if err != nil {
		return nil, fmt.Errorf("extract private key: %w", err)
	}

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("capture chain id: %w", err)
	}

	c := Client{
		network: network,
		address: crypto.PubkeyToAddress(privateKey.PublicKey),

		privateKey: privateKey,
		chainID:    chainID,
		ethClient:  ethClient,
	}

	return &c, nil
}

// EthClient returns the raw ethereum client API.
func (c *Client) EthClient() *ethclient.Client {
	return c.ethClient
}

// Address returns the current address calculated from the private key.
func (c *Client) Address() common.Address {
	return c.address
}

// Network returns the network information.
func (c *Client) Network() string {
	return c.network
}

// ChainID returns the chain information for the connected network.
func (c *Client) ChainID() int {
	return int(c.chainID.Int64())
}

// NewCallOpts constructs a new CallOpts which is used to call contract methods
// that does not require a transaction.
func (c *Client) NewCallOpts(ctx context.Context) (*bind.CallOpts, error) {
	call := bind.CallOpts{
		Pending: true,
		From:    c.address,
		Context: ctx,
	}

	return &call, nil
}

// NewTransactOpts constructs a new TransactOpts which is the collection of
// authorization data required to create a valid Ethereum transaction.
func (c *Client) NewTransactOpts(ctx context.Context, gasLimit uint64, valueGWei *big.Float) (*bind.TransactOpts, error) {
	nonce, err := c.ethClient.PendingNonceAt(ctx, c.address)
	if err != nil {
		return nil, fmt.Errorf("retrieving next nonce: %w", err)
	}

	gasPrice, err := c.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieving suggested gas price: %w", err)
	}

	tranOpts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("keying transaction: %w", err)
	}

	// This will convert the GWei value to Wei.
	GWei2Wei := big.NewInt(0)
	big.NewFloat(0).SetPrec(1024).Mul(valueGWei, big.NewFloat(1e9)).Int(GWei2Wei)

	tranOpts.Nonce = big.NewInt(0).SetUint64(nonce)
	tranOpts.Value = GWei2Wei
	tranOpts.GasLimit = gasLimit // The maximum amount of Gas you are willing to pay for.
	tranOpts.GasPrice = gasPrice // What you will agree to pay per unit of gas.

	return tranOpts, nil
}

// WaitMined will wait for the transaction to be minded and return a receipt.
func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, fmt.Errorf("waiting for tx to be mined: %w", err)
	}

	if receipt.Status == 0 {
		if err := c.extractError(ctx, tx); err != nil {
			return nil, fmt.Errorf("extracting tx error: %w", err)
		}
	}

	return receipt, nil
}

// TransactionByHash returns a transaction value for the specified transaction hash.
func (c *Client) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	return c.ethClient.TransactionByHash(ctx, txHash)
}

// TransactionReceipt returns a receipt value for the specified transaction hash.
func (c *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return c.ethClient.TransactionReceipt(ctx, txHash)
}

// BaseFee calculates the base fee from the block for this receipt.
func (c *Client) BaseFee(receipt *types.Receipt) (wei *big.Int) {
	block, err := c.ethClient.BlockByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		return big.NewInt(0)
	}
	return block.BaseFee()
}

// CurrentBalance retrieves the current balance for the account.
func (c *Client) CurrentBalance(ctx context.Context) (wei *big.Int, err error) {
	balance, err := c.ethClient.BalanceAt(ctx, c.address, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

// ContractBackend returns the ethereum client. This is needed for smart
// contract creation and other calls.
func (c *Client) ContractBackend() *ethclient.Client {
	return c.ethClient
}

// =============================================================================

// extractError checks the failed transaction for the error message.
func (c *Client) extractError(ctx context.Context, tx *types.Transaction) error {
	msg := ethereum.CallMsg{
		From:     c.address,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	_, err := c.ethClient.CallContract(ctx, msg, nil)
	return err
}