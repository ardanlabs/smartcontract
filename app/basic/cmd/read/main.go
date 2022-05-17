package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ardanlabs/smartcontract/business/smart"
)

func main() {
	ctx := context.Background()

	client, privateKey, err := smart.Connect()
	if err != nil {
		log.Fatal("Connect:", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("address:", fromAddress.String())

	// =========================================================================

	store, contractID, err := smart.NewStore(ctx, client)
	if err != nil {
		log.Fatal("NewStore:", err)
	}
	fmt.Println("contractID:", contractID)

	version, err := store.Version(nil)
	if err != nil {
		log.Fatal("version:", err)
	}
	fmt.Println("version:", version)

	// =========================================================================

	var key [32]byte
	copy(key[:], []byte("name"))

	var result [32]byte
	result, err = store.Items(nil, key)
	if err != nil {
		log.Fatal("Items ERROR:", err)
	}

	fmt.Println("value:", string(result[:]))
}
