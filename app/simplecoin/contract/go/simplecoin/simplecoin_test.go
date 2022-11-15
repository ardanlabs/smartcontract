package simplecoin_test

import (
	"math/big"
	"testing"

	"github.com/ardanlabs/smartcontract/app/simplecoin/contract/go/simplecoin"
	"github.com/divergencetech/ethier/ethtest"
)

const (
	deployer = 0
	account  = 1
)

func TestSimpleCoin(t *testing.T) {
	sim, err := ethtest.NewSimulatedBackend(2)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}

	startingBalance := big.NewInt(1000)
	contractID, _, _, err := simplecoin.DeploySimplecoin(sim.Acc(deployer), sim, startingBalance)
	if err != nil {
		t.Fatalf("unable to deploy simplecoin: %s", err)
	}

	scoin, err := simplecoin.NewSimplecoin(contractID, sim)
	if err != nil {
		t.Fatalf("unable to create a simplecoin: %s", err)
	}

	got, err := scoin.CoinBalance(sim.CallFrom(deployer), sim.Addr(deployer))
	if err != nil {
		t.Fatalf("should be able to get balance of deployer: %s", err)
	}

	if got.Cmp(startingBalance) != 0 {
		t.Fatalf("should have the expected starting balance, got %v  exp %v", got, startingBalance)
	}

	transferAmount := big.NewInt(100)
	if _, err = scoin.Transfer(sim.Acc(deployer), sim.Addr(account), transferAmount); err != nil {
		t.Fatalf("should be able to transfer money from deployer to account: %s", err)
	}

	got, err = scoin.CoinBalance(sim.CallFrom(deployer), sim.Addr(deployer))
	if err != nil {
		t.Fatalf("should be able to get balance of deployer: %s", err)
	}

	exp := startingBalance.Sub(startingBalance, transferAmount)
	if got.Cmp(exp) != 0 {
		t.Fatalf("should have the expected deployer balance, got %v  exp %v", got, exp)
	}

	got, err = scoin.CoinBalance(sim.CallFrom(deployer), sim.Addr(account))
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	if got.Cmp(transferAmount) != 0 {
		t.Fatalf("should have the expected account balance, got %v  exp %v", got, transferAmount)
	}
}
