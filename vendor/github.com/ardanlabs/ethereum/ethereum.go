// Package ethereum higher level Go API's for writing applications and smart
// contracts on the Ethereum blockchain.
package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
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

// client combines all the behavior a client needs.
type client interface {
	bind.ContractBackend
	bind.DeployBackend
	TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// Ethereum provides an API for working with smart contracts.
type Ethereum struct {
	network    string
	keyFile    string
	passPhrase string
	address    common.Address
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	client     client
}

// New provides an API for accessing an Ethereum node to perform blockchain
// related operations.
func New(ctx context.Context, network string, keyFile string, passPhrase string) (*Ethereum, error) {
	ethClient, err := ethclient.Dial(network)
	if err != nil {
		return nil, fmt.Errorf("dial network: %w", err)
	}

	privateKey, err := PrivateKeyByKeyFile(keyFile, passPhrase)
	if err != nil {
		return nil, fmt.Errorf("extract private key: %w", err)
	}

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("capture chain id: %w", err)
	}

	eth := Ethereum{
		network:    network,
		keyFile:    keyFile,
		passPhrase: passPhrase,
		address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		privateKey: privateKey,
		chainID:    chainID,
		client:     ethClient,
	}

	return &eth, nil
}

// NewSimulation provides an API for accessing an Ethereum node to perform blockchain
// related operations against the specified simulated backend and private key
// that was registered by the CreateSimulation function.
func NewSimulation(simBack *SimulatedBackend, pk *ecdsa.PrivateKey) *Ethereum {
	return &Ethereum{
		network:    "simulation",
		address:    crypto.PubkeyToAddress(pk.PublicKey),
		privateKey: pk,
		chainID:    big.NewInt(1337),
		client:     simBack,
	}
}

// Copy creates a new client connection copying the client's settings.
func (eth *Ethereum) Copy(ctx context.Context) (*Ethereum, error) {
	if _, isClient := eth.client.(*ethclient.Client); !isClient {
		return eth, nil
	}

	return New(ctx, eth.network, eth.keyFile, eth.passPhrase)
}

// =============================================================================

// RawClient returns the raw ethereum client API.
func (eth *Ethereum) RawClient() *ethclient.Client {
	client, isClient := eth.client.(*ethclient.Client)
	if !isClient {
		return nil
	}

	return client
}

// RawSimulation returns the raw ethereum simulated backend API.
func (eth *Ethereum) RawSimulation() *backends.SimulatedBackend {
	client, isSim := eth.client.(*backends.SimulatedBackend)
	if !isSim {
		return nil
	}

	return client
}

// ContractBackend returns the client where a contract backend is needed.
func (eth *Ethereum) ContractBackend() bind.ContractBackend {
	return eth.client
}

// DeployBackend returns the client where a deploy backend is needed.
func (eth *Ethereum) DeployBackend() bind.DeployBackend {
	return eth.client
}

// Address returns the current address calculated from the private key.
func (eth *Ethereum) Address() common.Address {
	return eth.address
}

// Network returns the network information.
func (eth *Ethereum) Network() string {
	return eth.network
}

// ChainID returns the chain information for the connected network.
func (eth *Ethereum) ChainID() int {
	return int(eth.chainID.Int64())
}

// PrivateKey returns the private key being used.
func (eth *Ethereum) PrivateKey() *ecdsa.PrivateKey {
	return eth.privateKey
}

// =============================================================================

// NewCallOpts constructs a new CallOpts which is used to call contract methods
// that does not require a transaction.
func (eth *Ethereum) NewCallOpts(ctx context.Context) (*bind.CallOpts, error) {
	call := bind.CallOpts{
		Pending: true,
		From:    eth.address,
		Context: ctx,
	}

	return &call, nil
}

// NewTransactOpts constructs a new TransactOpts which is the collection of
// authorization data required to create a valid Ethereum transaction. If the
// gasLimit is set to 0, an estimate will be made for the amount of gas needed.
func (eth *Ethereum) NewTransactOpts(ctx context.Context, gasLimit uint64, valueGWei *big.Float) (*bind.TransactOpts, error) {
	nonce, err := eth.client.PendingNonceAt(ctx, eth.address)
	if err != nil {
		return nil, fmt.Errorf("retrieving next nonce: %w", err)
	}

	gasPrice, err := eth.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieving suggested gas price: %w", err)
	}

	tranOpts, err := bind.NewKeyedTransactorWithChainID(eth.privateKey, eth.chainID)
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
func (eth *Ethereum) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, eth.client, tx)
	if err != nil {
		return nil, fmt.Errorf("waiting for tx to be mined: %w", err)
	}

	if receipt.Status == 0 {
		if err := eth.extractError(ctx, tx); err != nil {
			return nil, fmt.Errorf("extracting tx error: %w", err)
		}
	}

	return receipt, nil
}

// SendTransaction sends a transaction to the specified address for the
// specified amount. The function will wait for the transaction to be mined
// based on the timeout value specified in the context.
func (eth *Ethereum) SendTransaction(ctx context.Context, address common.Address, value *big.Int, gasLimit uint64) error {
	nonce, err := eth.client.PendingNonceAt(ctx, eth.address)
	if err != nil {
		return fmt.Errorf("retrieving next nonce: %w", err)
	}

	gasPrice, err := eth.client.SuggestGasPrice(ctx)
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

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(eth.chainID), eth.privateKey)
	if err != nil {
		return fmt.Errorf("signing transaction: %w", err)
	}

	if err := eth.client.SendTransaction(ctx, signedTx); err != nil {
		return fmt.Errorf("signing transaction: %w", err)
	}

	if _, err := eth.WaitMined(ctx, signedTx); err != nil {
		return fmt.Errorf("timedout waiting: %w", err)
	}

	return nil
}

// =============================================================================

// TransactionByHash returns a transaction value for the specified transaction hash.
func (eth *Ethereum) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	return eth.client.TransactionByHash(ctx, txHash)
}

// TransactionReceipt returns a receipt value for the specified transaction hash.
func (eth *Ethereum) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return eth.client.TransactionReceipt(ctx, txHash)
}

// Balance retrieves the current balance for the client account.
func (eth *Ethereum) Balance(ctx context.Context) (wei *big.Int, err error) {
	return eth.BalanceAt(ctx, eth.address)
}

// BalanceAt retrieves the current balance for the specified account.
func (eth *Ethereum) BalanceAt(ctx context.Context, address common.Address) (wei *big.Int, err error) {
	client, isClient := eth.client.(*ethclient.Client)
	if !isClient {
		return big.NewInt(0), errors.New("running simulated backend")
	}

	return client.BalanceAt(ctx, address, nil)
}

// BaseFee calculates the base fee from the block for this receipt.
func (eth *Ethereum) BaseFee(receipt *types.Receipt) (wei *big.Int) {
	client, isClient := eth.client.(*ethclient.Client)
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
func (eth *Ethereum) extractError(ctx context.Context, tx *types.Transaction) error {
	msg := ethereum.CallMsg{
		From:     eth.address,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	_, err := eth.client.CallContract(ctx, msg, nil)
	return err
}
