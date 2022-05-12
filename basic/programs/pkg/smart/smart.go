package smart

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Connect() (*ethclient.Client, *ecdsa.PrivateKey, error) {
	client, err := ethclient.Dial("/Users/bill/Library/Ethereum/geth.ipc")
	if err != nil {
		return nil, nil, fmt.Errorf("ethclient.Dial: %w", err)
	}

	privateKey, err := privateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("privateKey: %w", err)
	}

	return client, privateKey, nil
}

func privateKey() (*ecdsa.PrivateKey, error) {
	inPath := "node/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	password := "123"

	data, err := ioutil.ReadFile(inPath)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}
