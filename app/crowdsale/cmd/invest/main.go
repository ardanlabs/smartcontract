package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"

	crowd "github.com/ardanlabs/smartcontract/app/crowdsale/contract/go"
	"github.com/ardanlabs/smartcontract/foundation/smartcontract/smart"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ctx := context.Background()

	const rawurl = smart.NetworkLocalhost
	client, err := smart.Connect(ctx, rawurl, smart.PrimaryKeyPath, smart.PrimaryPassPhrase)
	if err != nil {
		return err
	}

	fmt.Println("fromAddress:", client.Account)

	// =========================================================================

	contract, err := newContract(client)
	if err != nil {
		return err
	}

	// =========================================================================

	startingBalance, err := client.CurrentBalance(ctx)
	if err != nil {
		return err
	}
	defer client.DisplayBalanceSheet(ctx, startingBalance)

	// =========================================================================

	const gasLimit = 300000
	const valueGwei = 100
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, valueGwei)
	if err != nil {
		return err
	}

	to := common.HexToAddress("0x6327a38415c53ffb36c11db55ea74cc9cb4976fd")

	// =========================================================================

	if rawurl == smart.NetworkLocalhost {
		sink := make(chan *crowd.CrowdEventInvestment, 100)
		if _, err := contract.WatchEventInvestment(nil, sink, []common.Address{to}); err != nil {
			return err
		}

		go func() {
			event := <-sink
			fmt.Println("\nEvents")
			fmt.Println("----------------------------------------------------")
			fmt.Println("tx event", event)
		}()
	}

	// =========================================================================

	tx, err := contract.Invest(tranOpts)
	if err != nil {
		return err
	}
	client.DisplayTransaction(tx)

	receipt, err := client.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	client.DisplayTransactionReceipt(receipt, tx)

	return nil
}

// newContract constructs a Crowdsale contract.
func newContract(client *smart.Client) (*crowd.Crowd, error) {
	data, err := os.ReadFile("zarf/smart/crowd.env")
	if err != nil {
		return nil, fmt.Errorf("readfile: %w", err)
	}
	contractID := string(data)
	fmt.Println("contractID:", contractID)

	contract, err := crowd.NewCrowd(common.HexToAddress(contractID), client.ContractBackend())
	if err != nil {
		return nil, fmt.Errorf("NewScoin: %w", err)
	}

	return contract, nil
}
