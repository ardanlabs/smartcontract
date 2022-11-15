package store_test

import (
	"testing"

	"github.com/ardanlabs/smartcontract/app/basic/contract/go/store"
	"github.com/divergencetech/ethier/ethtest"
)

const deployer = 0

func TestStore(t *testing.T) {
	sim, err := ethtest.NewSimulatedBackend(1)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}

	contractID, _, _, err := store.DeployStore(sim.Acc(deployer), sim)
	if err != nil {
		t.Fatalf("unable to deploy store: %s", err)
	}

	store, err := store.NewStore(contractID, sim)
	if err != nil {
		t.Fatalf("error creating store: %s", err)
	}

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("brianna"))

	if _, err := store.SetItem(sim.Acc(deployer), key, value); err != nil {
		t.Fatalf("should be able to set item: %s", err)
	}

	item, err := store.Items(nil, key)
	if err != nil {
		t.Fatalf("should be able to retrieve item: %s", err)
	}

	if item != value {
		t.Fatalf("wrong value, got %s  exp %s", string(item[:]), string(value[:]))
	}
}
