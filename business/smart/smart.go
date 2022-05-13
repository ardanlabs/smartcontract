package smart

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Hardcoding these for now since all the apps will use this same information
// and the information is static.
const (
	ethereum   = "zarf/ethereum"
	keyStore   = ethereum + "/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase = "123"
)

// Connect provides boilerplate for connecting to the geth service using
// an IPC socket created by the geth service on startup.
func Connect() (*ethclient.Client, *ecdsa.PrivateKey, error) {
	client, err := ethclient.Dial(ethereum + "/geth.ipc")
	if err != nil {
		return nil, nil, fmt.Errorf("ethclient.Dial: %w", err)
	}

	privateKey, err := privateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("privateKey: %w", err)
	}

	return client, privateKey, nil
}

func Contract(ctx context.Context, client *ethclient.Client) error {
	number, err := client.BlockNumber(ctx)
	if err != nil {
		return err
	}

	b, err := client.BlockByNumber(ctx, big.NewInt(int64(number)))
	if err != nil {
		return err
	}

	txs := b.Transactions()

	fmt.Println("**************", len(txs))

	for _, tx := range txs {
		fmt.Println("HASH", tx.Hash())
	}

	return nil
}

func privateKey() (*ecdsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(keyStore)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(data, passPhrase)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}
