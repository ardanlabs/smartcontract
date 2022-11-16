package bank_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/smartcontract/app/bank/single/contract/go/bank"
	"github.com/ethereum/go-ethereum/common"
)

const (
	delpoyerAcc = iota
	winnerAcc
	loser1Acc
	loser2Acc
	numAccounts
)

func TestBank(t *testing.T) {
	ctx := context.Background()

	sim, err := ethereum.CreateSimulation(numAccounts, true)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer sim.Close()

	deployer := ethereum.NewSimulation(sim, sim.PrivateKeys[delpoyerAcc])
	winner := ethereum.NewSimulation(sim, sim.PrivateKeys[winnerAcc])
	loser1 := ethereum.NewSimulation(sim, sim.PrivateKeys[loser1Acc])
	loser2 := ethereum.NewSimulation(sim, sim.PrivateKeys[loser2Acc])

	const gasLimit = 1700000
	const valueGwei = 0.0

	deployTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	contractID, tx, _, err := bank.DeployBank(deployTranOpts, deployer.ContractBackend())
	if err != nil {
		t.Fatalf("unable to deploy bank: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for deploy: %s", err)
	}

	testBank, err := bank.NewBank(contractID, sim)
	if err != nil {
		t.Fatalf("unable to create a bank: %s", err)
	}

	callOpts, err := deployer.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	owner, err := testBank.Owner(callOpts)
	if err != nil {
		t.Fatalf("unable to get account owner: %s", err)
	}

	if owner != deployer.Address() {
		t.Fatalf("Retrieved owner doesn't match expectations: %v != %v", owner, deployer.Address())
	}

	initialBalance, err := testBank.Balance(callOpts)
	if err != nil {
		t.Fatalf("should get the initial balance: %s", err)
	}

	depositTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deposit: %s", err)
	}

	depositTranOpts.Value = big.NewInt(10)
	if _, err = testBank.Deposit(depositTranOpts); err != nil {
		t.Fatalf("should be able to deposit money: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for deposit: %s", err)
	}

	postDepositBalance, err := testBank.Balance(callOpts)
	if err != nil {
		t.Fatalf("unable to get balance after deposit: %s", err)
	}

	gotBal := initialBalance.Add(initialBalance, depositTranOpts.Value)
	if postDepositBalance.Cmp(gotBal) != 0 {
		t.Fatalf("wrong balance, got %v  exp %v", gotBal, postDepositBalance)
	}

	withdrawTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for withdraw: %s", err)
	}

	withdrawTranOpts.Value = big.NewInt(10)
	if _, err = testBank.Withdraw(withdrawTranOpts); err != nil {
		t.Fatalf("unable be able to withdraw money: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for withdraw: %s", err)
	}

	postWithdrawBalance, err := testBank.Balance(callOpts)
	if err != nil {
		t.Fatalf("should get balance after withdraw: %s", err)
	}

	gotBal = postDepositBalance.Sub(postDepositBalance, withdrawTranOpts.Value)
	if postWithdrawBalance.Cmp(gotBal) != 0 {
		t.Fatalf("wrong balance, got %v  exp %v", gotBal, postWithdrawBalance)
	}

	reconcileTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for reconcile: %s", err)
	}

	losers := []common.Address{loser1.Address(), loser2.Address()}
	antWei := big.NewInt(1000)
	gameFeeWei := big.NewInt(10)

	_, err = testBank.Reconcile(reconcileTranOpts, winner.Address(), losers, antWei, gameFeeWei)
	if err != nil {
		t.Fatalf("unable to reconcile: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for reconcile: %s", err)
	}

	version, err := testBank.Version(callOpts)
	if err != nil {
		t.Fatalf("error getting version: %s", err)
	}

	const expectedVersion = "0.1.0"
	if version != expectedVersion {
		t.Fatalf("wrong version. got %s, exp %s\n", version, expectedVersion)
	}
}
