package bank_test

import (
	"math/big"
	"testing"

	"github.com/ardanlabs/smartcontract/app/bank/single/contract/go/bank"
	"github.com/divergencetech/ethier/ethtest"
)

const deployer = 0

func TestBank(t *testing.T) {
	sim, err := ethtest.NewSimulatedBackend(1)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}

	contractID, _, _, err := bank.DeployBank(sim.Acc(deployer), sim)
	if err != nil {
		t.Fatalf("unable to deploy bank: %s", err)
	}

	bank, err := bank.NewBank(contractID, sim)
	if err != nil {
		t.Fatalf("unable to create a bank: %s", err)
	}

	balanceBefore, err := bank.Balance(sim.CallFrom(deployer))
	if err != nil {
		t.Fatalf("should get the balance: %s", err)
	}

	to := sim.Acc(deployer)
	to.Value = big.NewInt(10)
	if _, err = bank.Deposit(to); err != nil {
		t.Fatalf("should be able to deposit money: %s", err)
	}

	expBal, err := bank.Balance(sim.CallFrom(deployer))
	if err != nil {
		t.Fatalf("should get balance after deposit: %s", err)
	}

	gotBal := balanceBefore.Add(balanceBefore, to.Value)
	if expBal.Cmp(gotBal) != 0 {
		t.Fatalf("wrong balance, got %v  exp %v", gotBal, expBal)
	}
}
