package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ardanlabs/smartcontract/app/basic/contracts/store"
	"github.com/ardanlabs/smartcontract/business/smart"
)

func main() {
	ctx := context.Background()

	client, privateKey, err := smart.Connect()
	if err != nil {
		log.Fatal("Connect: ERROR:", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("address:", fromAddress.String())

	// =========================================================================

	store, contractID, err := newStore(ctx, client)
	if err != nil {
		log.Fatal("newStore: ERROR:", err)
	}
	fmt.Println("contractID:", contractID)

	version, err := store.Version(nil)
	if err != nil {
		log.Fatal("version: ERROR:", err)
	}
	fmt.Println("version:", version)

	// =========================================================================

	balBefore, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatal("BalanceAt ERROR:", err)
	}

	const gasLimit = 100000
	tran, err := smart.NewTransaction(ctx, gasLimit, privateKey, client)
	if err != nil {
		log.Fatal("NewTransaction: ERROR:", err)
	}

	fmt.Println("transaction:", tran)

	// =========================================================================

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("brianna"))

	tx, err := store.SetItem(tran, key, value)
	if err != nil {
		log.Fatal("SetItem ERROR:", err)
	}

	fmt.Println("tx sent        :", tx.Hash().Hex())
	fmt.Println("tx gas price   :", smart.Wei2Eth(tx.GasPrice()))
	fmt.Println("tx cost        :", smart.Wei2Eth(tx.Cost()))
	fmt.Println("tx gas allowed :", tx.Gas())

	ctx, cancel := context.WithTimeout(ctx, time.Second*14)
	defer cancel()

	receipt, err := smart.CheckReceipt(ctx, tx.Hash(), client)
	if err != nil {
		log.Fatal("CheckReceipt ERROR:", err)
	}

	fmt.Println("tx gas used    :", receipt.GasUsed)
	fmt.Println("tx status      :", receipt.Status)

	// =========================================================================

	balAfter, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatal("balance at ERROR:", err)
	}

	fmt.Println("balance before :", smart.Wei2Eth(balBefore))
	fmt.Println("balance after  :", smart.Wei2Eth(balAfter))
}

// newStore constructs a Store value for smart contract API access.
func newStore(ctx context.Context, client *ethclient.Client) (*store.Store, string, error) {
	data, err := os.ReadFile("contract.env")
	if err != nil {
		return nil, "", fmt.Errorf("readfile: %w", err)
	}
	contractID := string(data)

	contract := common.HexToAddress(contractID)
	store, err := store.NewStore(contract, client)
	if err != nil {
		return nil, "", fmt.Errorf("NewStore: %w", err)
	}

	return store, contractID, nil
}
