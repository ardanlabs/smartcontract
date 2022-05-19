package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ardanlabs/smartcontract/app/simplecoin/contracts/scoin"
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

	coin, contractID, err := newScoin(ctx, client)
	if err != nil {
		log.Fatal("newScoin: ERROR:", err)
	}
	fmt.Println("contractID:", contractID)

	// =========================================================================

	balBefore, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatal("BalanceAt ERROR:", err)
	}

	const gasLimit = 300000
	tran, err := smart.NewTransaction(ctx, gasLimit, privateKey, client)
	if err != nil {
		log.Fatal("NewTransaction: ERROR:", err)
	}

	fmt.Println("transaction:", tran)

	to := common.HexToAddress("0x8e113078adf6888b7ba84967f299f29aece24c55")

	// =========================================================================

	sink := make(chan *scoin.ScoinTransfer, 100)
	sub, err := coin.WatchTransfer(nil, sink, []common.Address{fromAddress}, []common.Address{to})
	if err != nil {
		log.Fatal("WatchTransfer: ERROR:", err)
	}
	fmt.Println("sub:", sub)

	go func() {
		event := <-sink
		fmt.Println("=======================================")
		fmt.Println("tx event", event)
		fmt.Println("=======================================")
	}()

	// =========================================================================

	tx, err := coin.Transfer(tran, to, big.NewInt(100))
	if err != nil {
		log.Fatal("Transfer ERROR:", err)
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

// newScoin constructs a SimpleCoin value for smart contract API access.
func newScoin(ctx context.Context, client *ethclient.Client) (*scoin.Scoin, string, error) {
	data, err := os.ReadFile("contract.env")
	if err != nil {
		return nil, "", fmt.Errorf("readfile: %w", err)
	}
	contractID := string(data)

	contract := common.HexToAddress(contractID)
	scoin, err := scoin.NewScoin(contract, client)
	if err != nil {
		return nil, "", fmt.Errorf("NewScoin: %w", err)
	}

	return scoin, contractID, nil
}
