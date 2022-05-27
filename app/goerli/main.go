package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ardanlabs/smartcontract/business/smart"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()

	sc, err := smart.Connect(ctx, smart.NetworkLocalhost, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress:", sc.Account)

	balance, err := sc.CurrentBalance(ctx)
	if err != nil {
		return err
	}

	fmt.Println("balance :", smart.Wei2GWei(balance), "GWei")

	return nil
}
