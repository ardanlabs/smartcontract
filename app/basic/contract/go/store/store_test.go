package store_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/smartcontract/app/basic/contract/go/store"
)

func TestStore(t *testing.T) {
	ctx := context.Background()

	sim, err := ethereum.CreateSimulation(1, true)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer sim.Close()

	eth := ethereum.NewSimulation(sim, sim.PrivateKeys[0])

	const gasLimit = 1600000
	const valueGwei = 0.0
	tranOpts, err := eth.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	contractID, _, _, err := store.DeployStore(tranOpts, eth.ContractBackend())
	if err != nil {
		t.Fatalf("unable to deploy store: %s", err)
	}

	testStore, err := store.NewStore(contractID, eth.ContractBackend())
	if err != nil {
		t.Fatalf("error creating store: %s", err)
	}

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("brianna"))

	tranOpts, err = eth.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for setitem: %s", err)
	}

	if _, err := testStore.SetItem(tranOpts, key, value); err != nil {
		t.Fatalf("should be able to set item: %s", err)
	}

	item, err := testStore.Items(nil, key)
	if err != nil {
		t.Fatalf("should be able to retrieve item: %s", err)
	}

	if item != value {
		t.Fatalf("wrong value, got %s  exp %s", string(item[:]), string(value[:]))
	}
}
