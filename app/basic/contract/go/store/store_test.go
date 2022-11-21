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

	backend, err := ethereum.CreateSimulatedBackend(1, true, big.NewInt(100))
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer backend.Close()

	clt, err := ethereum.NewClient(backend, backend.PrivateKeys[0])
	if err != nil {
		t.Fatalf("unable to create ethereum api: %s", err)
	}

	// =========================================================================

	const gasLimit = 1600000
	const valueGwei = 0.0
	tranOpts, err := clt.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	contractID, tx, _, err := store.DeployStore(tranOpts, clt.Backend)
	if err != nil {
		t.Fatalf("unable to deploy store: %s", err)
	}

	if _, err := clt.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for deploy: %s", err)
	}

	testStore, err := store.NewStore(contractID, clt.Backend)
	if err != nil {
		t.Fatalf("error creating store: %s", err)
	}

	// =========================================================================

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("brianna"))

	tranOpts, err = clt.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for setitem: %s", err)
	}

	tx, err = testStore.SetItem(tranOpts, key, value)
	if err != nil {
		t.Fatalf("should be able to set item: %s", err)
	}

	if _, err := clt.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for setitem: %s", err)
	}

	// =========================================================================

	item, err := testStore.Items(nil, key)
	if err != nil {
		t.Fatalf("should be able to retrieve item: %s", err)
	}

	if item != value {
		t.Fatalf("wrong value, got %s  exp %s", string(item[:]), string(value[:]))
	}
}
