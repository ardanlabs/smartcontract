package simplecoin_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/smartcontract/app/simplecoin/contract/go/simplecoin"
)

const (
	ownerAcc = iota
	secondAcc
	numAccounts
)

func TestSimpleCoin(t *testing.T) {
	ctx := context.Background()

	backend, err := ethereum.CreateSimulatedBackend(numAccounts, true, big.NewInt(100))
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer backend.Close()

	deployer, err := ethereum.NewClient(backend, backend.PrivateKeys[ownerAcc])
	if err != nil {
		t.Fatalf("unable to create ownerAcc: %s", err)
	}

	second, err := ethereum.NewClient(backend, backend.PrivateKeys[secondAcc])
	if err != nil {
		t.Fatalf("unable to create secondAcc: %s", err)
	}

	callOpts, err := deployer.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	// =========================================================================

	const gasLimit = 1600000
	const valueGwei = 0.0
	const startingBalanceAmount = 5000

	var scoin *simplecoin.Simplecoin
	startingBalance := big.NewInt(startingBalanceAmount)

	// =========================================================================

	t.Run("deploy simplecoin", func(t *testing.T) {
		deployTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		contractID, tx, _, err := simplecoin.DeploySimplecoin(deployTranOpts, deployer.Backend, startingBalance)
		if err != nil {
			t.Fatalf("unable to deploy simplecoin: %s", err)
		}

		if _, err := deployer.WaitMined(ctx, tx); err != nil {
			t.Fatalf("waiting for deploy: %s", err)
		}

		scoin, err = simplecoin.NewSimplecoin(contractID, deployer.Backend)
		if err != nil {
			t.Fatalf("unable to create a simplecoin: %s", err)
		}
	})

	// =========================================================================

	t.Run("check starting balance", func(t *testing.T) {
		initialBalance, err := scoin.CoinBalance(callOpts, deployer.Address())
		if err != nil {
			t.Fatalf("unable be able to get initial balance of deployer: %s", err)
		}

		if initialBalance.Cmp(startingBalance) != 0 {
			t.Fatalf("should have the expected starting balance, got %v  exp %v", initialBalance, startingBalance)
		}
	})

	// =========================================================================

	t.Run("check transfer", func(t *testing.T) {
		transferTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		transferAmount := big.NewInt(1000)
		tx, err := scoin.Transfer(transferTranOpts, second.Address(), transferAmount)
		if err != nil {
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
	})

	// =========================================================================

	t.Run("check owner", func(t *testing.T) {
		owner, err := scoin.Owner(callOpts)
		if err != nil {
			t.Fatalf("unable to get account owner: %s", err)
		}

		if owner != deployer.Address() {
			t.Fatalf("Retrieved owner doesn't match expectations: %v != %v", owner, deployer.Address())
		}
	})

	// =========================================================================

	t.Run("check non-frozen account", func(t *testing.T) {
		frozen, err := scoin.FrozenAccount(callOpts, second.Address())
		if err != nil {
			t.Fatalf("unable be able to get allowences: %s", err)
		}

		if frozen {
			t.Fatal("Account shouldn't be frozen")
		}
	})

	// =========================================================================

	t.Run("freeze account", func(t *testing.T) {
		freezeTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		tx, err := scoin.FreezeAccount(freezeTranOpts, second.Address(), true)
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		if _, err := deployer.WaitMined(ctx, tx); err != nil {
			t.Fatalf("waiting for transfer: %s", err)
		}
	})

	// =========================================================================

	t.Run("check account was frozen", func(t *testing.T) {
		frozen, err := scoin.FrozenAccount(callOpts, second.Address())
		if err != nil {
			t.Fatalf("unable be able to get allowences: %s", err)
		}

		if !frozen {
			t.Fatal("Account should be frozen")
		}
	})

	// =========================================================================

	t.Run("thaw account", func(t *testing.T) {
		thawTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		tx, err := scoin.FreezeAccount(thawTranOpts, second.Address(), false)
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		if _, err := deployer.WaitMined(ctx, tx); err != nil {
			t.Fatalf("waiting for transfer: %s", err)
		}
	})

	// =========================================================================

	t.Run("check account was thawed", func(t *testing.T) {
		frozen, err := scoin.FrozenAccount(callOpts, second.Address())
		if err != nil {
			t.Fatalf("unable be able to get allowences: %s", err)
		}

		if frozen {
			t.Fatal("Account shouldn't be frozen")
		}
	})

	// =========================================================================

	t.Run("check minting", func(t *testing.T) {
		mintTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		mintedAmount := big.NewInt(1000)
		tx, err := scoin.Mint(mintTranOpts, second.Address(), mintedAmount)
		if err != nil {
			t.Fatalf("unable to create transaction opts for deploy: %s", err)
		}

		if _, err := deployer.WaitMined(ctx, tx); err != nil {
			t.Fatalf("waiting for transfer: %s", err)
		}
	})
}
