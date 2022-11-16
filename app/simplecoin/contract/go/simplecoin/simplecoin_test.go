package simplecoin_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/smartcontract/app/simplecoin/contract/go/simplecoin"
)

const (
	ownerAcc = iota
	secondAcc
	numAccounts
	gasLimit              = 1600000
	valueGwei             = 0.0
	startingBalanceAmount = 1000
)

func TestSimpleCoin(t *testing.T) {
	ctx := context.Background()

	//
	// Setup the simulation backend and create/bind the account
	//

	sim, err := ethereum.CreateSimulation(numAccounts, true)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer sim.Close()

	deployer := ethereum.NewSimulation(sim, sim.PrivateKeys[ownerAcc])
	second := ethereum.NewSimulation(sim, sim.PrivateKeys[secondAcc])

	//
	// Deploy the contract
	//

	deployTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	startingBalance := big.NewInt(startingBalanceAmount)
	contractID, tx, _, err := simplecoin.DeploySimplecoin(deployTranOpts, sim, startingBalance)
	if err != nil {
		t.Fatalf("unable to deploy simplecoin: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for deploy: %s", err)
	}

	//
	// Get the Simplecoin object
	//

	scoin, err := simplecoin.NewSimplecoin(contractID, sim)
	if err != nil {
		t.Fatalf("unable to create a simplecoin: %s", err)
	}

	//
	// Check the deployer coin balance and make sure it matches the value we created with
	//

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

	//
	// Transfer coins from the deployer account to the second account and check that the balances are correct
	//

	transferTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	transferAmount := big.NewInt(100)
	if _, err = scoin.Transfer(transferTranOpts, second.Address(), transferAmount); err != nil {
		t.Fatalf("should be able to transfer money from deployer to account: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}

	postTransferBalance, err := scoin.CoinBalance(callOpts, deployer.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of deployer: %s", err)
	}

	exp := startingBalance.Sub(startingBalance, transferAmount)
	if postTransferBalance.Cmp(exp) != 0 {
		t.Fatalf("should have the expected deployer balance, got %v  exp %v", postTransferBalance, exp)
	}

	secondBalance, err := scoin.CoinBalance(callOpts, second.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	if secondBalance.Cmp(transferAmount) != 0 {
		t.Fatalf("should have the expected account balance, got %v  exp %v", secondBalance, transferAmount)
	}

	//
	// Transfer coins back from the second account to the first account and check the balances are correct
	//

	preTransferFromBalanceDeployer, err := scoin.CoinBalance(callOpts, deployer.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	preTransferFromBalanceSecond, err := scoin.CoinBalance(callOpts, second.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	fmt.Printf("tfd: %v tfs: %v\n", preTransferFromBalanceDeployer, preTransferFromBalanceSecond)

	transferFromTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	transferFromAmount := big.NewInt(10)
	if _, err = scoin.TransferFrom(transferFromTranOpts, second.Address(), deployer.Address(), transferFromAmount); err != nil {
		t.Fatalf("should be able to transfer money from deployer to account: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}

	transferFromBalanceDeployer, err := scoin.CoinBalance(callOpts, deployer.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	transferFromBalanceSecond, err := scoin.CoinBalance(callOpts, second.Address())
	if err != nil {
		t.Fatalf("should be able to get balance of account: %s", err)
	}

	fmt.Printf("tfd: %v tfs: %v\n", transferFromBalanceDeployer, transferFromBalanceSecond)

	//
	// Get the Allowance (TODO: Figure out what these even means/are)
	//

	_, err = scoin.Allowance(callOpts, deployer.Address(), second.Address())
	if err != nil {
		t.Fatalf("unable be able to get allowences: %s", err)
	}

	//
	// Check that the owner returned by the coin interface is what we expect it to be
	//

	owner, err := scoin.Owner(callOpts)
	if err != nil {
		t.Fatalf("unable to get account owner: %s", err)
	}

	if owner != deployer.Address() {
		t.Fatalf("Retrieved owner doesn't match expectations: %v != %v", owner, deployer.Address())
	}

	//
	// Get the Authorization (TODO: Figure out what these even means/are)
	//

	allowanceTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	allowanceAmount := big.NewInt(startingBalanceAmount)
	_, err = scoin.Authorize(allowanceTranOpts, deployer.Address(), allowanceAmount)

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}

	//
	// Check if the second account is frozen (it shouldn't be)
	//

	frozen, err := scoin.FrozenAccount(callOpts, second.Address())
	if err != nil {
		t.Fatalf("unable be able to get allowences: %s", err)
	}

	if frozen {
		t.Fatal("Account shouldn't be frozen")
	}

	//
	// Freeze the second account
	//

	freezeTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	_, err = scoin.FreezeAccount(freezeTranOpts, second.Address(), true)
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}

	//
	// Check if the second account is frozen (it should be)
	//

	frozen, err = scoin.FrozenAccount(callOpts, second.Address())
	if err != nil {
		t.Fatalf("unable be able to get allowences: %s", err)
	}

	if !frozen {
		t.Fatal("Account should be frozen")
	}

	//
	// Unfreeze the second account
	//

	thawTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	_, err = scoin.FreezeAccount(thawTranOpts, second.Address(), false)
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}

	//
	// Check if the second account is frozen (it shouldn't be)
	//

	frozen, err = scoin.FrozenAccount(callOpts, second.Address())
	if err != nil {
		t.Fatalf("unable be able to get allowences: %s", err)
	}

	if frozen {
		t.Fatal("Account shouldn't be frozen")
	}

	//
	// Mint 1000 coins (TODO: figure out what this does)
	//

	mintTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	mintedAmount := big.NewInt(1000)
	_, err = scoin.Mint(mintTranOpts, second.Address(), mintedAmount)
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	if _, err := deployer.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for transfer: %s", err)
	}
}
