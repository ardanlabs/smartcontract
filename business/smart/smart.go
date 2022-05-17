package smart

import (
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
