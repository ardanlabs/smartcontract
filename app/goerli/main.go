package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ardanlabs/smartcontract/foundation/smartcontract/smart"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()

	client, err := smart.Connect(ctx, smart.NetworkLocalhost, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress:", client.Account)

	balance, err := client.CurrentBalance(ctx)
	if err != nil {
		return err
	}

	fmt.Println("balance :", smart.Wei2GWei(balance), "GWei")

	return nil
}
