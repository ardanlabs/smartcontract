package basic_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/smartcontract/app/basic/contract/go/basic"
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

	contractID, tx, _, err := basic.DeployBasic(tranOpts, clt.Backend)
	if err != nil {
		t.Fatalf("unable to deploy basic: %s", err)
	}

	if _, err := clt.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for deploy: %s", err)
	}

	testBasic, err := basic.NewBasic(contractID, clt.Backend)
	if err != nil {
		t.Fatalf("error creating basic: %s", err)
	}

	// =========================================================================

	key := "bill"
	value := big.NewInt(1_000_000)

	tranOpts, err = clt.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for setitem: %s", err)
	}

	tx, err = testBasic.SetItem(tranOpts, key, value)
	if err != nil {
		t.Fatalf("should be able to set item: %s", err)
	}

	if _, err := clt.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for setitem: %s", err)
	}

	// =========================================================================

	item, err := testBasic.Items(nil, key)
	if err != nil {
		t.Fatalf("should be able to retrieve item: %s", err)
	}

	if item.Cmp(value) != 0 {
		t.Fatalf("wrong value, got %s  exp %s", item, value)
	}
}