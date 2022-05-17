package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"

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
	fmt.Println("fromAddress:", fromAddress)

	// =========================================================================

	balBefore, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatal("BalanceAt: ERROR:", err)
	}

	const gasLimit = 300000
	tran, err := smart.NewTransaction(ctx, gasLimit, privateKey, client)
	if err != nil {
		log.Fatal("NewTransaction: ERROR:", err)
	}

	fmt.Println("transaction:", tran)

	// =========================================================================

	address, tx, _, err := store.DeployStore(tran, client)
	if err != nil {
		log.Fatal("deploy ERROR:", err)
	}

	fmt.Println("tx sent        :", tx.Hash().Hex())
	fmt.Println("tx gas price   :", smart.Wei2Eth(tx.GasPrice()))
	fmt.Println("tx cost        :", smart.Wei2Eth(tx.Cost()))

	receipt, err := client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		log.Fatal("set item ERROR:", err)
	}

	fmt.Println("tx gas allowed :", tx.Gas())
	fmt.Println("tx gas used    :", receipt.GasUsed)
	fmt.Println("tx status      :", receipt.Status)

	if receipt.Status == 0 {
		log.Fatal("Transaction Failed, check gas numbers.")
	}

	fmt.Println("Contract Address :", address.Hex())

	if err := os.WriteFile("contract.env", []byte(address.Hex()), 0666); err != nil {
		log.Fatal("cannot write 'contract.env' ERROR: ", err)
	}

	// =========================================================================

	balAfter, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatal("balance at ERROR:", err)
	}

	fmt.Println("balance before   :", smart.Wei2Eth(balBefore))
	fmt.Println("balance after    :", smart.Wei2Eth(balAfter))
}
