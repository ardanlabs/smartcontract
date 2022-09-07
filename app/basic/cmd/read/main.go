package main

import (
	"context"
	"fmt"
	"os"

	store "github.com/ardanlabs/smartcontract/app/basic/contract/go"
	"github.com/ardanlabs/smartcontract/foundation/smart/contract"
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

	client, err := contract.NewClient(ctx, contract.NetworkLocalhost, keyStoreFile, passPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress:", client.Address())

	// =========================================================================

	contractID := os.Getenv("CONTRACT_ID")
	if contractID == "" {
		return fmt.Errorf("need to export the CONTRACT_ID")
	}
	fmt.Println("contractID:", contractID)

	storeCon, err := store.NewStore(common.HexToAddress(contractID), client.ContractBackend())
	if err != nil {
		return fmt.Errorf("new contract: %w", err)
	}

	version, err := storeCon.Version(nil)
	if err != nil {
		return err
	}
	fmt.Println("version:", version)

	// =========================================================================

	var key [32]byte
	copy(key[:], []byte("name"))

	var result [32]byte
	result, err = storeCon.Items(nil, key)
	if err != nil {
		return err
	}

	fmt.Println("value:", string(result[:]))

	return nil
}
