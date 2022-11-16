package simplecoin_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/smartcontract/app/simplecoin/contract/go/simplecoin"
)

const (
	ownerAcc = iota
	secondAcc
	numAccounts
)

func TestSimpleCoin(t *testing.T) {
	ctx := context.Background()

	sim, err := ethereum.CreateSimulation(numAccounts, true)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer sim.Close()

	deployer := ethereum.NewSimulation(sim, sim.PrivateKeys[ownerAcc])
	second := ethereum.NewSimulation(sim, sim.PrivateKeys[secondAcc])

	const gasLimit = 1600000
	const valueGwei = 0.0
	const startingBalanceAmount = 1000

	deployTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	startingBalance := big.NewInt(startingBalanceAmount)
	contractID, tx, _, err := simplecoin.DeploySimplecoin(deployTranOpts, sim, startingBalance)
	if err != nil {
		t.Fatalf("unable to deploy simplecoin: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for deploy: %s", err)
	}

	scoin, err := simplecoin.NewSimplecoin(contractID, sim)
	if err != nil {
		t.Fatalf("unable to create a simplecoin: %s", err)
	}

	callOpts, err := deployer.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	initialBalance, err := scoin.CoinBalance(callOpts, deployer.Address())
	if err != nil {
		t.Fatalf("unable be able to get initial balance of deployer: %s", err)
	}

	if initialBalance.Cmp(startingBalance) != 0 {
		t.Fatalf("should have the expected starting balance, got %v  exp %v", initialBalance, startingBalance)
	}

	transferTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	transferAmount := big.NewInt(100)
	if _, err = scoin.Transfer(transferTranOpts, second.Address(), transferAmount); err != nil {
		t.Fatalf("should be able to transfer money from deployer to account: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}

	postTransferBalance, err := scoin.CoinBalance(callOpts, deployer.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of deployer: %s", err)
	}

	exp := startingBalance.Sub(startingBalance, transferAmount)
	if postTransferBalance.Cmp(exp) != 0 {
		t.Fatalf("should have the expected deployer balance, got %v  exp %v", postTransferBalance, exp)
	}

	secondBalance, err := scoin.CoinBalance(callOpts, second.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	if secondBalance.Cmp(transferAmount) != 0 {
		t.Fatalf("should have the expected account balance, got %v  exp %v", secondBalance, transferAmount)
	}
}
