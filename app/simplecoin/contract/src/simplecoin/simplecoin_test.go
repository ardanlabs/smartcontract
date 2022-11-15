package simplecoin

//go:generate ethier gen simplecoin.sol error.sol

import (
	"math/big"
	"testing"

	"github.com/divergencetech/ethier/ethtest"
)

const (
	startingBalance = 10000
	transferAmount  = 100
	deployer        = iota
	account1
	numAccounts
)

func TestSimpleCoin(t *testing.T) {
	sim, err := ethtest.NewSimulatedBackend(numAccounts)
	if err != nil {
		t.Fatalf("Something went wrong settup up the simulated backend: %s", err)
	}

	_, _, contract, err := DeploySimpleCoin(sim.Acc(deployer), sim, big.NewInt(startingBalance))
	if err != nil {
		t.Fatalf("DeployBank error %v", err)
	}

	t.Run("Test SimpleCoin", func(t *testing.T) {
		balance, err := contract.CoinBalance(sim.CallFrom(deployer), sim.Addr(deployer))
		if err != nil {
			t.Fatalf("Error getting balance before deposit: %s", err)
		}

		if balance.Cmp(big.NewInt(startingBalance)) != 0 {
			t.Fatalf("Deployer starting balance doesn't match: %v != %v", balance, startingBalance)
		}

		_, err = contract.Transfer(sim.Acc(deployer), sim.Addr(account1), big.NewInt(transferAmount))
		if err != nil {
			t.Fatalf("Error setting test data: %s", err)
		}

		balance, err = contract.CoinBalance(sim.CallFrom(deployer), sim.Addr(deployer))
		if err != nil {
			t.Fatalf("Error getting balance after deposit: %s", err)
		}

		if balance.Cmp(big.NewInt(startingBalance-transferAmount)) != 0 {
			t.Fatalf("Deployer ending balance doesn't match: %v != %v", balance, startingBalance)
		}

		balance, err = contract.CoinBalance(sim.CallFrom(deployer), sim.Addr(account1))
		if err != nil {
			t.Fatalf("Error getting balance after deposit: %s", err)
		}

		if balance.Cmp(big.NewInt(transferAmount)) != 0 {
			t.Fatalf("Account1 ending balance doesn't match: %v != %v", balance, transferAmount)
		}
	})
}
