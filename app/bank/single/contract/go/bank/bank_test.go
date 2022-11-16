package bank_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/smartcontract/app/bank/single/contract/go/bank"
)

func TestBank(t *testing.T) {
	ctx := context.Background()

	sim, err := ethereum.CreateSimulation(1, true)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer sim.Close()

	eth := ethereum.NewSimulation(sim, sim.PrivateKeys[0])

	const (
		gasLimit  = 1600000
		valueGwei = 0.0
	)

	tranOpts, err := eth.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	contractID, _, _, err := bank.DeployBank(tranOpts, eth.ContractBackend())
	if err != nil {
		t.Fatalf("unable to deploy bank: %s", err)
	}

	testBank, err := bank.NewBank(contractID, sim)
	if err != nil {
		t.Fatalf("unable to create a bank: %s", err)
	}

	callOpts, err := eth.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts for owner: %s", err)
	}

	owner, err := testBank.Owner(callOpts)
	if err != nil {
		t.Fatalf("unable to get account owner: %s", err)
	}

	if owner != eth.Address() {
		t.Fatalf("Retrieved owner doesn't match expectations: %v != %v", owner, eth.Address())
	}

	balanceBefore, err := testBank.Balance(callOpts)
	if err != nil {
		t.Fatalf("should get the balance: %s", err)
	}

	to, err := eth.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deposit: %s", err)
	}

	to.Value = big.NewInt(10)
	if _, err = testBank.Deposit(to); err != nil {
		t.Fatalf("should be able to deposit money: %s", err)
	}

	expBal, err := testBank.Balance(callOpts)
	if err != nil {
		t.Fatalf("should get balance after deposit: %s", err)
	}

	gotBal := balanceBefore.Add(balanceBefore, to.Value)
	if expBal.Cmp(gotBal) != 0 {
		t.Fatalf("wrong balance, got %v  exp %v", gotBal, expBal)
	}
}
