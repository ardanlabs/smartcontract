package bank_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/bank/proxy/contract/go/bank"
	"github.com/ardanlabs/smartcontract/app/simplecoin/contract/go/simplecoin"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	deployerAcc = iota
	winnerAcc
	loser1Acc
	loser2Acc
	numAccounts
)

func TestBankProxy(t *testing.T) {
	ctx := context.Background()
	var testBank *bank.Bank
	converter := currency.NewDefaultConverter(simplecoin.SimplecoinMetaData.ABI)
	var contractID common.Address

	sim, err := ethereum.CreateSimulation(numAccounts, true)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer sim.Close()

	deployer := ethereum.NewSimulation(sim, sim.PrivateKeys[deployerAcc])
	winner := ethereum.NewSimulation(sim, sim.PrivateKeys[winnerAcc])
	loser1 := ethereum.NewSimulation(sim, sim.PrivateKeys[loser1Acc])
	loser2 := ethereum.NewSimulation(sim, sim.PrivateKeys[loser2Acc])

	callOpts, err := deployer.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	const gasLimit = 1700000
	const valueGwei = 0.0

	t.Run("deploy bank", func(t *testing.T) {
		deployTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		var tx *types.Transaction
		contractID, tx, _, err = bank.DeployBank(deployTranOpts, deployer.ContractBackend())
		if err != nil {
			t.Fatalf("unable to deploy bank: %s", err)
		}

		receipt, err := deployer.WaitMined(ctx, tx)
		if err != nil {
			t.Fatalf("waiting for deploy: %s", err)
		}

		t.Logf("Transfer\n%s", converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

		testBank, err = bank.NewBank(contractID, sim)
		if err != nil {
			t.Fatalf("unable to create a bank: %s", err)
		}
	})

	t.Run("set contract", func(t *testing.T) {
		setContractTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		tx, err := testBank.SetContract(setContractTranOpts, contractID)
		if err != nil {
			t.Fatalf("unable to set contract: %s", err)
		}

		receipt, err := deployer.WaitMined(ctx, tx)
		if err != nil {
			t.Fatalf("waiting for set contract: %s", err)
		}

		t.Logf("Transfer\n%s", converter.FmtTransactionReceipt(receipt, tx.GasPrice()))
	})

	t.Run("check owner matches", func(t *testing.T) {
		owner, err := testBank.Owner(callOpts)
		if err != nil {
			t.Fatalf("unable to get account owner: %s", err)
		}

		if owner != deployer.Address() {
			t.Fatalf("Retrieved owner doesn't match expectations: %v != %v", owner, deployer.Address())
		}
	})

	t.Run("check deposit", func(t *testing.T) {
		initialBalance, err := testBank.Balance(callOpts)
		if err != nil {
			t.Fatalf("should get the initial balance: %s", err)
		}

		t.Logf("Initial balance: %v", initialBalance)

		depositTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for deposit: %s", err)
		}

		depositTranOpts.Value = big.NewInt(10)
		tx, err := testBank.Deposit(depositTranOpts)
		if err != nil {
			t.Fatalf("should be able to deposit money: %s", err)
		}

		receipt, err := deployer.WaitMined(ctx, tx)
		if err != nil {
			t.Fatalf("waiting for deposit: %s", err)
		}

		t.Logf("Transfer\n%s", converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

		postDepositBalance, err := testBank.Balance(callOpts)
		if err != nil {
			t.Fatalf("unable to get balance after deposit: %s", err)
		}

		t.Logf("post deposit balance: %v", postDepositBalance)

		expectedBalance := initialBalance.Add(initialBalance, depositTranOpts.Value)
		if postDepositBalance.Cmp(expectedBalance) != 0 {
			t.Fatalf("wrong balance, got %v  exp %v", postDepositBalance, expectedBalance)
		}
	})

	t.Run("check withdraw", func(t *testing.T) {
		initialBalance, err := testBank.Balance(callOpts)
		if err != nil {
			t.Fatalf("should get the initial balance: %s", err)
		}

		withdrawTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for withdraw: %s", err)
		}

		withdrawTranOpts.Value = big.NewInt(10)
		tx, err := testBank.Withdraw(withdrawTranOpts)
		if err != nil {
			t.Fatalf("unable be able to withdraw money: %s", err)
		}

		receipt, err := deployer.WaitMined(ctx, tx)
		if err != nil {
			t.Fatalf("waiting for withdraw: %s", err)
		}

		t.Logf("Transfer\n%s", converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

		postWithdrawBalance, err := testBank.Balance(callOpts)
		if err != nil {
			t.Fatalf("should get balance after withdraw: %s", err)
		}

		expectedBalance := initialBalance.Sub(initialBalance, withdrawTranOpts.Value)
		if postWithdrawBalance.Cmp(expectedBalance) != 0 {
			t.Fatalf("wrong balance, got %v  exp %v", postWithdrawBalance, expectedBalance)
		}
	})

	t.Run("check reconciliation", func(t *testing.T) {
		reconcileTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
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

		receipt, err := deployer.WaitMined(ctx, tx)
		if err != nil {
			t.Fatalf("waiting for reconcile: %s", err)
		}

		t.Logf("Transfer\n%s", converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	})

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
