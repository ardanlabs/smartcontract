package store

//go:generate ethier gen store.sol

import (
	"testing"

	"github.com/divergencetech/ethier/ethtest"
)

const (
	deployer = iota
	numAccounts
)

func TestMyContract(t *testing.T) {
	sim, err := ethtest.NewSimulatedBackend(numAccounts)
	if err != nil {
		t.Fatalf("Something went wrong settup up the simulated backend: %s", err)
	}

	_, _, contract, err := DeployStore(sim.Acc(deployer), sim /*,,, [constructor arguments]*/)
	if err != nil {
		t.Fatalf("DeployStore error %v", err)
	}

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("brianna"))

	t.Run("protect something sensitive", func(t *testing.T) {
		_, err := contract.SetItem(sim.Acc(deployer), key, value)
		if err != nil {
			t.Fatalf("Error setting test data: %s", err)
		}

		items, err := contract.Items(nil, key)
		if err != nil {
			t.Fatalf("Error getting test data: %s", err)
		}

		if items != value {
			t.Fatalf("Expected %s, got %s", string(value[:]), string(items[:]))
		}

	})
}
