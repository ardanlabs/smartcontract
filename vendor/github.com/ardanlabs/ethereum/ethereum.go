// Package ethereum higher level Go API's for writing applications and smart
// contracts on the Ethereum blockchain.
package ethereum

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
)

// Set of networks supported by the smart package.
const (
	NetworkHTTPLocalhost = "http://localhost:8545"
	NetworkLocalhost     = "zarf/ethereum/geth.ipc"
	NetworkGoerli        = "https://rpc.ankr.com/eth_goerli"
)

// =============================================================================

// Backend represents behavior for interacting with an ethereum node.
type Backend interface {
	bind.ContractBackend
	bind.DeployBackend
	TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	BalanceAt(ctx context.Context, contract common.Address, blockNumber *big.Int) (*big.Int, error)
	Network() string
	ChainID() *big.Int
}

// Client provides an API for working with smart contracts.
type Client struct {
	Backend
	address    common.Address
	privateKey *ecdsa.PrivateKey
}

// NewClient provides an API for accessing an Ethereum node to perform blockchain
// related operations. It required the private key you want to use for this
// instance when accessing the node.
func NewClient(backend Backend, privateKey *ecdsa.PrivateKey) (*Client, error) {
	clt := Client{
		Backend:    backend,
		address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		privateKey: privateKey,
	}

	return &clt, nil
}

// =============================================================================

// Address returns the current address calculated from the private key.
func (clt *Client) Address() common.Address {
	return clt.address
}

// Network returns the network information for the connected network.
func (clt *Client) Network() string {
	return clt.Backend.Network()
}

// ChainID returns the chain information for the connected network.
func (clt *Client) ChainID() int {
	return int(clt.Backend.ChainID().Int64())
}

// PrivateKey returns the private key being used.
func (clt *Client) PrivateKey() *ecdsa.PrivateKey {
	return clt.privateKey
}

// =============================================================================

// NewCallOpts constructs a new CallOpts which is used to call contract methods
// that does not require a transaction.
func (clt *Client) NewCallOpts(ctx context.Context) (*bind.CallOpts, error) {
	call := bind.CallOpts{
		Pending: true,
		From:    clt.address,
		Context: ctx,
	}

	return &call, nil
}

// NewTransactOpts constructs a new TransactOpts which is the collection of
// authorization data required to create a valid Ethereum transaction. If the
// gasLimit is set to 0, an estimate will be made for the amount of gas needed.
// If the gasPrice is set to 0, then the connected geth service  is consulted
// for the suggested gas price.
func (clt *Client) NewTransactOpts(ctx context.Context, gasLimit uint64, gasPrice *big.Int, valueGWei *big.Float) (*bind.TransactOpts, error) {
	nonce, err := clt.PendingNonceAt(ctx, clt.address)
	if err != nil {
		return nil, fmt.Errorf("retrieving next nonce: %w", err)
	}

	if gasPrice == nil || gasPrice.Cmp(big.NewInt(0)) == 0 {
		gasPrice, err = clt.SuggestGasPrice(ctx)
		if err != nil {
			return nil, fmt.Errorf("retrieving suggested gas price: %w", err)
		}
	}

	tranOpts, err := bind.NewKeyedTransactorWithChainID(clt.privateKey, clt.Backend.ChainID())
	if err != nil {
		return nil, fmt.Errorf("keying transaction: %w", err)
	}

	// This will convert the GWei value to Wei.
	gwe2Wei := big.NewInt(0)
	big.NewFloat(0).SetPrec(1024).Mul(valueGWei, big.NewFloat(1e9)).Int(gwe2Wei)

	tranOpts.Nonce = big.NewInt(0).SetUint64(nonce)
	tranOpts.Value = gwe2Wei
	tranOpts.GasLimit = gasLimit // The maximum amount of Gas you are willing to pay for.
	tranOpts.GasPrice = gasPrice // What you will agree to pay per unit of gas.

	return tranOpts, nil
}

// WaitMined will wait for the transaction to be minded and return a receipt.
func (clt *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, clt.Backend, tx)
	if err != nil {
		return nil, fmt.Errorf("waiting for tx to be mined: %w", err)
	}

	if receipt.Status == 0 {
		if err := clt.extractError(ctx, tx); err != nil {
			return nil, fmt.Errorf("extracting tx error: %w", err)
		}
	}

	return receipt, nil
}

// SendTransaction sends a transaction to the specified address for the
// specified amount. The function will wait for the transaction to be mined
// based on the timeout value specified in the context.
func (clt *Client) SendTransaction(ctx context.Context, address common.Address, value *big.Int, gasLimit uint64) error {
	nonce, err := clt.PendingNonceAt(ctx, clt.address)
	if err != nil {
		return fmt.Errorf("retrieving next nonce: %w", err)
	}

	gasPrice, err := clt.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("retrieving suggested gas price: %w", err)
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &address,
		Value:    value,
		Data:     nil,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(clt.Backend.ChainID()), clt.privateKey)
	if err != nil {
		return fmt.Errorf("signing transaction: %w", err)
	}

	if err := clt.Backend.SendTransaction(ctx, signedTx); err != nil {
		return fmt.Errorf("signing transaction: %w", err)
	}

	if _, err := clt.WaitMined(ctx, signedTx); err != nil {
		return fmt.Errorf("timedout waiting: %w", err)
	}

	return nil
}

// =============================================================================

// Balance retrieves the current balance for the client account.
func (clt *Client) Balance(ctx context.Context) (wei *big.Int, err error) {
	return clt.BalanceAt(ctx, clt.address, nil)
}

// BaseFee calculates the base fee from the block for this receipt.
func (clt *Client) BaseFee(receipt *types.Receipt) (wei *big.Int) {
	client, isClient := clt.Backend.(*DialedBackend)
	if !isClient {
		return big.NewInt(0)
	}

	block, err := client.BlockByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		return big.NewInt(0)
	}

	return block.BaseFee()
}

// =============================================================================

// extractError checks the failed transaction for the error message.
func (clt *Client) extractError(ctx context.Context, tx *types.Transaction) error {
	msg := ethereum.CallMsg{
		From:     clt.address,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	_, err := clt.CallContract(ctx, msg, nil)
	return err
}
