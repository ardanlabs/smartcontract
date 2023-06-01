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
	deployerAcc = iota
	winnerAcc
	loser1Acc
	loser2Acc
	numAccounts
)

func TestBankSingle(t *testing.T) {
	ctx := context.Background()

	backend, err := ethereum.CreateSimulatedBackend(numAccounts, true, big.NewInt(100))
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer backend.Close()

	// =========================================================================

	deployer, err := ethereum.NewClient(backend, backend.PrivateKeys[deployerAcc])
	if err != nil {
		t.Fatalf("unable to create deployerAcc: %s", err)
	}

	winner, err := ethereum.NewClient(backend, backend.PrivateKeys[winnerAcc])
	if err != nil {
		t.Fatalf("unable to create winnerAcc: %s", err)
	}

	loser1, err := ethereum.NewClient(backend, backend.PrivateKeys[loser1Acc])
	if err != nil {
		t.Fatalf("unable to create loser1Acc: %s", err)
	}

	loser2, err := ethereum.NewClient(backend, backend.PrivateKeys[loser2Acc])
	if err != nil {
		t.Fatalf("unable to create loser2Acc: %s", err)
	}

	callOpts, err := deployer.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	// =========================================================================

	const gasLimit = 2000000
	const valueGwei = 0.0

	var testBank *bank.Bank

	// =========================================================================

	t.Run("deploy bank", func(t *testing.T) {
		deployTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		contractID, tx, _, err := bank.DeployBank(deployTranOpts, deployer.Backend)
		if err != nil {
			t.Fatalf("unable to deploy bank: %s", err)
		}

		if _, err := deployer.WaitMined(ctx, tx); err != nil {
			t.Fatalf("waiting for deploy: %s", err)
		}

		testBank, err = bank.NewBank(contractID, deployer.Backend)
		if err != nil {
			t.Fatalf("unable to create a bank: %s", err)
		}
	})

	// =========================================================================

	t.Run("check owner matches", func(t *testing.T) {
		owner, err := testBank.Owner(callOpts)
		if err != nil {
			t.Fatalf("unable to get account owner: %s", err)
		}

		if owner != deployer.Address() {
			t.Fatalf("Retrieved owner doesn't match expectations: %v != %v", owner, deployer.Address())
		}
	})

	// =========================================================================

	t.Run("check deposit", func(t *testing.T) {
		initialBalance, err := testBank.Balance(callOpts)
		if err != nil {
			t.Fatalf("should get the initial balance: %s", err)
		}

		depositTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for deposit: %s", err)
		}

		depositTranOpts.Value = big.NewInt(10)
		tx, err := testBank.Deposit(depositTranOpts)
		if err != nil {
			t.Fatalf("should be able to deposit money: %s", err)
		}

		if _, err := deployer.WaitMined(ctx, tx); err != nil {
			t.Fatalf("waiting for deposit: %s", err)
		}

		postDepositBalance, err := testBank.Balance(callOpts)
		if err != nil {
			t.Fatalf("unable to get balance after deposit: %s", err)
		}

		expectedBalance := initialBalance.Add(initialBalance, depositTranOpts.Value)
		if postDepositBalance.Cmp(expectedBalance) != 0 {
			t.Fatalf("wrong balance, got %v  exp %v", postDepositBalance, expectedBalance)
		}
	})

	// =========================================================================

	t.Run("check withdraw", func(t *testing.T) {
		initialBalance, err := testBank.Balance(callOpts)
		if err != nil {
			t.Fatalf("should get the initial balance: %s", err)
		}

		withdrawTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for withdraw: %s", err)
		}

		withdrawTranOpts.Value = big.NewInt(10)
		tx, err := testBank.Withdraw(withdrawTranOpts)
		if err != nil {
			t.Fatalf("unable be able to withdraw money: %s", err)
		}

		if _, err := deployer.WaitMined(ctx, tx); err != nil {
			t.Fatalf("waiting for withdraw: %s", err)
		}

		postWithdrawBalance, err := testBank.Balance(callOpts)
		if err != nil {
			t.Fatalf("should get balance after withdraw: %s", err)
		}

		expectedBalance := initialBalance.Sub(initialBalance, withdrawTranOpts.Value)
		if postWithdrawBalance.Cmp(expectedBalance) != 0 {
			t.Fatalf("wrong balance, got %v  exp %v", postWithdrawBalance, expectedBalance)
		}
	})

	// =========================================================================

	t.Run("check reconciliation", func(t *testing.T) {
		reconcileTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for reconcile: %s", err)
		}

		losers := []common.Address{loser1.Address(), loser2.Address()}
		antWei := big.NewInt(1000)
		gameFeeWei := big.NewInt(10)

		tx, err := testBank.Reconcile(reconcileTranOpts, winner.Address(), losers, antWei, gameFeeWei)
		if err != nil {
			t.Fatalf("unable to reconcile: %s", err)
		}

		if _, err := deployer.WaitMined(ctx, tx); err != nil {
			t.Fatalf("waiting for reconcile: %s", err)
		}
	})

	// =========================================================================

	t.Run("check version", func(t *testing.T) {
		version, err := testBank.Version(callOpts)
		if err != nil {
			t.Fatalf("error getting version: %s", err)
		}

		const expectedVersion = "0.1.0"
		if version != expectedVersion {
			t.Fatalf("wrong version. got %s, exp %s\n", version, expectedVersion)
		}
	})
}
