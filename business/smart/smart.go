package smart

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Set of networks supported by the smart package.
const (
	// geth = "http://localhost:8545"
	NetworkLocalhost = "zarf/ethereum/geth.ipc"
	NetworkGoerli    = "https://rpc.goerli.mudit.blog/"
)

// Harded this here for now just to make life easier.
const (
	PrimaryKeyPath    = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	PrimaryPassPhrase = "123"
)

// =============================================================================

// SmartContract provides an API for working with smart contracts.
type SmartContract struct {
	Network string
	Account common.Address
	Client  *ethclient.Client

	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
}

// Connect provides boilerplate for connecting to the geth service using
// an IPC socket created by the geth service on startup.
func Connect(ctx context.Context, network string, keyPath string, passPhrase string) (*SmartContract, error) {
	client, err := ethclient.Dial(network)
	if err != nil {
		return nil, fmt.Errorf("dial network: %w", err)
	}

	privateKey, err := privateKeyByKeyFile(keyPath, passPhrase)
	if err != nil {
		return nil, fmt.Errorf("extract private key: %w", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("capture chain id: %w", err)
	}

	sc := SmartContract{
		Network: network,
		Account: crypto.PubkeyToAddress(privateKey.PublicKey),
		Client:  client,

		privateKey: privateKey,
		chainID:    chainID,
	}

	return &sc, nil
}

// NewTransaction constructs a new TransactOpts which is the collection of
// authorization data required to create a valid Ethereum transaction.
func (sc *SmartContract) NewTransactOpts(ctx context.Context, gasLimit uint64, valueGwei uint64) (*bind.TransactOpts, error) {
	nonce, err := sc.Client.PendingNonceAt(ctx, sc.Account)
	if err != nil {
		return nil, fmt.Errorf("retrieving next nonce: %w", err)
	}

	gasPrice, err := sc.Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieving suggested gas price: %w", err)
	}

	tranOpts, err := bind.NewKeyedTransactorWithChainID(sc.privateKey, sc.chainID)
	if err != nil {
		return nil, fmt.Errorf("keying transaction: %w", err)
	}

	// The value must be converted from Gwei to Wei.
	valueWei := big.NewInt(0).Mul(big.NewInt(int64(valueGwei)), GWeiConv)

	tranOpts.Nonce = big.NewInt(int64(nonce))
	tranOpts.Value = valueWei
	tranOpts.GasLimit = gasLimit // The maximum amount of Gas you are willing to pay for.
	tranOpts.GasPrice = gasPrice // What you will agree to pay per unit of gas.

	return tranOpts, nil
}

// WaitMined will wait for the transaction to be minded and return a receipt.
func (sc *SmartContract) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, sc.Client, tx)
	if err != nil {
		return nil, fmt.Errorf("waiting for tx to be mined: %w", err)
	}

	if receipt.Status == 0 {
		err := sc.extractError(ctx, tx)
		return nil, fmt.Errorf("extracting tx error: %w", err)
	}

	return receipt, nil
}

// BaseFee calculates the base fee from the block for this receipt.
func (sc *SmartContract) BaseFee(receipt *types.Receipt) *big.Int {
	block, err := sc.Client.BlockByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		return big.NewInt(0)
	}
	return block.BaseFee()
}

// CurrentBalance retrieves the current balance for the account.
func (sc *SmartContract) CurrentBalance(ctx context.Context) (*big.Int, error) {
	balance, err := sc.Client.BalanceAt(ctx, sc.Account, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

// =============================================================================

// extractError checks the failed transaction for the error message.
func (sc *SmartContract) extractError(ctx context.Context, tx *types.Transaction) error {
	msg := ethereum.CallMsg{
		From:     sc.Account,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	_, err := sc.Client.CallContract(ctx, msg, nil)
	return err
}

// privateKeyByKeyFile opens a key file for the private key.
func privateKeyByKeyFile(keyPath string, passPhrase string) (*ecdsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(data, passPhrase)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}
