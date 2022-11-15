package bank

import (
	"fmt"
	"testing"

	"github.com/divergencetech/ethier/ethtest"
)

const (
	deployer = iota
	account1
	account2
	account3
	account4
	numAccounts
)

func TestBank(t *testing.T) {
	sim, err := ethtest.NewSimulatedBackend(numAccounts)
	if err != nil {
		t.Fatalf("Something went wrong settup up the simulated backend: %s", err)
	}

	_, _, contract, err := DeployBank(sim.Acc(deployer), sim /*,,, [constructor arguments]*/)
	if err != nil {
		t.Fatalf("DeployBank error %v", err)
	}

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("brianna"))

	t.Run("Test Bank", func(t *testing.T) {
		balance, err := contract.Balance(sim.CallFrom(deployer))
		if err != nil {
			t.Fatalf("Error getting balance before deposit: %s", err)
		}

		fmt.Printf("balance before deposit: %v\n", balance)

		// proxyContract, err := NewBank(sim.Addr(deployer), sim)
		// if err != nil {
		// 	t.Fatalf("new proxy connection: %s", err)
		// }

		_, err = contract.Deposit(sim.Acc(deployer))
		if err != nil {
			t.Fatalf("Error setting test data: %s", err)
		}

		balance, err = contract.Balance(sim.CallFrom(deployer))
		if err != nil {
			t.Fatalf("Error getting balance after deposit: %s", err)
		}

		fmt.Printf("balance after deposit: %v\n", balance)

	})
}
