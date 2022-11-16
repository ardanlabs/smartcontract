package simplecoin_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/simplecoin/contract/go/simplecoin"
)

const (
	ownerAcc = iota
	secondAcc
	numAccounts
	gasLimit               = 1600000
	valueGwei              = 0.0
	xstartingBalanceAmount = 5000
)

func TestSimpleCoin(t *testing.T) {
	ctx := context.Background()

	converter := currency.NewDefaultConverter(simplecoin.SimplecoinMetaData.ABI)

	// =========================================================================

	sim, err := ethereum.CreateSimulation(numAccounts, true)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer sim.Close()

	deployer := ethereum.NewSimulation(sim, sim.PrivateKeys[ownerAcc])
	second := ethereum.NewSimulation(sim, sim.PrivateKeys[secondAcc])

	// =========================================================================

	deployTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	startingBalance := big.NewInt(3000)
	contractID, tx, _, err := simplecoin.DeploySimplecoin(deployTranOpts, sim, startingBalance)
	if err != nil {
		t.Fatalf("unable to deploy simplecoin: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for deploy: %s", err)
	}

	scoin, err := simplecoin.NewSimplecoin(contractID, sim)
	if err != nil {
		t.Fatalf("unable to create a simplecoin: %s", err)
	}

	// =========================================================================

	callOpts, err := deployer.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	initialBalance, err := scoin.CoinBalance(callOpts, deployer.Address())
	if err != nil {
		t.Fatalf("unable be able to get initial balance of deployer: %s", err)
	}

	if initialBalance.Cmp(startingBalance) != 0 {
		t.Fatalf("should have the expected starting balance, got %v  exp %v", initialBalance, startingBalance)
	}

	// =========================================================================

	transferTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	transferAmount := big.NewInt(1000)
	tx, err = scoin.Transfer(transferTranOpts, second.Address(), transferAmount)
	if err != nil {
		t.Fatalf("should be able to transfer money from deployer to account: %s", err)
	}

	receipt, err := deployer.WaitMined(ctx, tx)
	if err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}

	t.Logf("Transfer\n%s", converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	// =========================================================================

	postTransferBalance, err := scoin.CoinBalance(callOpts, deployer.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of deployer: %s", err)
	}

	exp := startingBalance.Sub(startingBalance, transferAmount)
	if postTransferBalance.Cmp(exp) != 0 {
		t.Fatalf("should have the expected deployer balance, got %v  exp %v", postTransferBalance, exp)
	}

	// =========================================================================

	secondBalance, err := scoin.CoinBalance(callOpts, second.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	if secondBalance.Cmp(transferAmount) != 0 {
		t.Fatalf("should have the expected account balance, got %v  exp %v", secondBalance, transferAmount)
	}

	// =========================================================================

	preTransferFromBalanceDeployer, err := scoin.CoinBalance(callOpts, deployer.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	preTransferFromBalanceSecond, err := scoin.CoinBalance(callOpts, second.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	fmt.Printf("tfd: %v tfs: %v\n", preTransferFromBalanceDeployer, preTransferFromBalanceSecond)

	transferFromBalanceDeployer, err := scoin.CoinBalance(callOpts, deployer.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	transferFromBalanceSecond, err := scoin.CoinBalance(callOpts, second.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	fmt.Printf("tfd: %v tfs: %v\n", transferFromBalanceDeployer, transferFromBalanceSecond)

	// =========================================================================

	owner, err := scoin.Owner(callOpts)
	if err != nil {
		t.Fatalf("unable to get account owner: %s", err)
	}

	if owner != deployer.Address() {
		t.Fatalf("Retrieved owner doesn't match expectations: %v != %v", owner, deployer.Address())
	}

	// =========================================================================

	frozen, err := scoin.FrozenAccount(callOpts, second.Address())
	if err != nil {
		t.Fatalf("unable be able to get allowences: %s", err)
	}

	if frozen {
		t.Fatal("Account shouldn't be frozen")
	}

	// =========================================================================

	freezeTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	tx, err = scoin.FreezeAccount(freezeTranOpts, second.Address(), true)
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	receipt, err = deployer.WaitMined(ctx, tx)
	if err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}

	t.Logf("FreezeAccount\n%s", converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	// =========================================================================

	frozen, err = scoin.FrozenAccount(callOpts, second.Address())
	if err != nil {
		t.Fatalf("unable be able to get allowences: %s", err)
	}

	if !frozen {
		t.Fatal("Account should be frozen")
	}

	// =========================================================================

	thawTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	tx, err = scoin.FreezeAccount(thawTranOpts, second.Address(), false)
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	receipt, err = deployer.WaitMined(ctx, tx)
	if err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}

	t.Logf("FreezeAccount\n%s", converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	// =========================================================================

	frozen, err = scoin.FrozenAccount(callOpts, second.Address())
	if err != nil {
		t.Fatalf("unable be able to get allowences: %s", err)
	}

	if frozen {
		t.Fatal("Account shouldn't be frozen")
	}

	// =========================================================================

	mintTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	mintedAmount := big.NewInt(1000)
	tx, err = scoin.Mint(mintTranOpts, second.Address(), mintedAmount)
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	receipt, err = deployer.WaitMined(ctx, tx)
	if err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}

	t.Logf("Mint\n%s", converter.FmtTransactionReceipt(receipt, tx.GasPrice()))
}
