package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/smartcontract/app/basic/contract/go/basic"
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

	contractIDBytes, err := os.ReadFile("zarf/ethereum/basic.cid")
	if err != nil {
		return fmt.Errorf("importing basic.cid file: %w", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return errors.New("need to export the basic.cid file")
	}
	fmt.Println("contractID:", contractID)

	storeCon, err := basic.NewBasic(common.HexToAddress(contractID), clt.Backend)
	if err != nil {
		return fmt.Errorf("new contract: %w", err)
	}

	version, err := storeCon.Version(nil)
	if err != nil {
		return err
	}
	fmt.Println("version:", version)

	// =========================================================================

	key := "bill"
	result, err := storeCon.Items(nil, key)
	if err != nil {
		return err
	}

	fmt.Println("\nRead Value")
	fmt.Println("----------------------------------------------------")
	fmt.Println("value:", result)

	return nil
}
