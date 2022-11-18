package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/simplecoin/contract/go/simplecoin"
	"github.com/ethereum/go-ethereum/common"
)

const (
	keyStoreFile     = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase       = "123"
	coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	backend, err := ethereum.CreateDialedBackend(ctx, ethereum.NetworkLocalhost)
	if err != nil {
		return err
	}
	defer backend.Close()

	privateKey, err := ethereum.PrivateKeyByKeyFile(keyStoreFile, passPhrase)
	if err != nil {
		return err
	}

	clt, err := ethereum.NewClient(backend, privateKey)
	if err != nil {
		return err
	}

	fmt.Println("\nInput Values")
	fmt.Println("----------------------------------------------------")
	fmt.Println("fromAddress:", clt.Address())

	// =========================================================================

	contractIDBytes, err := os.ReadFile("zarf/ethereum/simplecoin.cid")
	if err != nil {
		return fmt.Errorf("importing simplecoin.cid file: %w", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return errors.New("need to export the simplecoin.cid file")
	}
	fmt.Println("contractID:", contractID)

	scoinCon, err := simplecoin.NewSimplecoin(common.HexToAddress(contractID), clt.Backend)
	if err != nil {
		return fmt.Errorf("new contract: %w", err)
	}

	// =========================================================================

	balance, err := scoinCon.CoinBalance(nil, common.HexToAddress("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd"))
	if err != nil {
		return err
	}
	fmt.Println("balance 0x6327:", currency.Wei2GWei(balance), "GWei")

	balance, err = scoinCon.CoinBalance(nil, common.HexToAddress("0x8e113078adf6888b7ba84967f299f29aece24c55"))
	if err != nil {
		return err
	}
	fmt.Println("balance 0x8e11:", currency.Wei2GWei(balance), "GWei")

	return nil
}
