package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ardanlabs/smartcontract/business/smart"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()

	client, _, err := smart.Connect(smart.NetworkGoerli)
	if err != nil {
		return err
	}

	balance, err := client.BalanceAt(ctx, common.HexToAddress("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd"), nil)
	if err != nil {
		return err
	}

	fmt.Println("balance before :", smart.Wei2Eth(balance), "ETH")

	return nil
}
