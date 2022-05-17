package smart

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ardanlabs/smartcontract/app/basic/contracts/store"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Hardcoding these for now since all the apps will use this same information
// and the information is static.
const (
	ethereum       = "http://localhost:8545"
	primaryAccount = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase     = "123"
)

// Connect provides boilerplate for connecting to the geth service using
// an IPC socket created by the geth service on startup.
func Connect() (*ethclient.Client, *ecdsa.PrivateKey, error) {
	client, err := ethclient.Dial(ethereum)
	if err != nil {
		return nil, nil, fmt.Errorf("ethclient.Dial: %w", err)
	}

	privateKey, err := privateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("privateKey: %w", err)
	}

	return client, privateKey, nil
}

// NewTransaction constructs a new transaction for executing function that cost money.
func NewTransaction(ctx context.Context, gasLimit uint64, pk *ecdsa.PrivateKey, client *ethclient.Client) (*bind.TransactOpts, error) {
	address := crypto.PubkeyToAddress(pk.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("PendingNonceAt: ERROR: %w", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("ChainID: ERROR: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("SuggestGasPrice: ERROR: %w", err)
	}

	tran, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return nil, fmt.Errorf("NewKeyedTransactorWithChainID: ERROR: %w", err)
	}

	tran.Nonce = big.NewInt(int64(nonce))
	tran.Value = big.NewInt(0) // in wei
	tran.GasLimit = gasLimit   // The maximum amount of Gas a user can consume to conduct this transaction.
	tran.GasPrice = gasPrice   // What you are willing to pay per unit to complete this transaction.

	return tran, nil
}

// NewStore constructs a Store value for smart contract API access.
func NewStore(ctx context.Context, client *ethclient.Client) (*store.Store, string, error) {
	data, err := os.ReadFile("contract.env")
	if err != nil {
		return nil, "", fmt.Errorf("SuggestGasPrice: ERROR: %w", err)
	}
	contractID := string(data)

	contract := common.HexToAddress(contractID)
	store, err := store.NewStore(contract, client)
	if err != nil {
		return nil, "", fmt.Errorf("NewStore: ERROR: %w", err)
	}

	return store, contractID, nil
}

// Wei2Eth converts the wei unit into a Eth for display.
func Wei2Eth(amount *big.Int) string {
	compact_amount := big.NewInt(0)
	reminder := big.NewInt(0)
	divisor := big.NewInt(1e18)
	compact_amount.QuoRem(amount, divisor, reminder)
	return fmt.Sprintf("%v.%018s", compact_amount.String(), reminder.String())
}

func privateKey() (*ecdsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(primaryAccount)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(data, passPhrase)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}
