package book_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	scbook "github.com/ardanlabs/smartcontract/app/book/contract/go/book"
	"github.com/ardanlabs/smartcontract/app/book/pkg/book"
	"github.com/ethereum/go-ethereum/common"
)

// These variables provide some static GWei to play with.
var (
	oneUSD    = big.NewFloat(662_833.00)
	tenUSD    = big.NewFloat(0).Mul(oneUSD, big.NewFloat(10))
	twentyUSD = big.NewFloat(0).Mul(oneUSD, big.NewFloat(20))
)

// We need a string for the bet id.
var betID = "1234"

var (
	backend      *ethereum.SimulatedBackend
	ownerClt     *ethereum.Client
	player1Clt   *ethereum.Client
	player2Clt   *ethereum.Client
	moderatorClt *ethereum.Client
)

// =============================================================================

func TestMain(m *testing.M) {
	var err error
	backend, err = ethereum.CreateSimulatedBackend(4, true, big.NewInt(100))
	if err != nil {
		fmt.Println("create backend", err)
		os.Exit(1)
	}
	defer backend.Close()

	ownerClt, err = ethereum.NewClient(backend, backend.PrivateKeys[0])
	if err != nil {
		fmt.Println("create ownerClt client", err)
		os.Exit(1)
	}

	player1Clt, err = ethereum.NewClient(backend, backend.PrivateKeys[1])
	if err != nil {
		fmt.Println("create player1Clt client", err)
		os.Exit(1)
	}

	player2Clt, err = ethereum.NewClient(backend, backend.PrivateKeys[2])
	if err != nil {
		fmt.Println("create player2Clt client", err)
		os.Exit(1)
	}

	moderatorClt, err = ethereum.NewClient(backend, backend.PrivateKeys[3])
	if err != nil {
		fmt.Println("create moderatorClt client", err)
		os.Exit(1)
	}

	m.Run()
}

// =============================================================================

func Test_DepositWithdraw(t *testing.T) {
	contractID, err := deployContract()
	if err != nil {
		t.Fatalf("deploying contract: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect player 1 to the smart contract.
	playerClient, err := book.New(ctx, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new book for player1: %s", err)
	}

	// *************************************************************************
	// Deposit process

	// Get the starting balance.
	startingBalance, err := playerClient.EthereumBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// Perform a deposit from the player's wallet.
	depositTx, depositReceipt, err := playerClient.Deposit(ctx, twentyUSD)
	if err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	// Calculate the expected balance by subtracting the amount deposited and the
	// gas fees for the transaction.
	gasCost := big.NewInt(0).Mul(depositTx.GasPrice(), big.NewInt(0).SetUint64(depositReceipt.GasUsed))
	depositWeiAmount := currency.GWei2Wei(twentyUSD)
	expectedBalance := big.NewInt(0).Sub(startingBalance, depositWeiAmount)
	expectedBalance.Sub(expectedBalance, gasCost)

	// Get the updated wallet balance.
	currentBalance, err := playerClient.EthereumBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// The player's wallet balance should match the deposit minus the fees.
	if expectedBalance.Cmp(currentBalance) != 0 {
		t.Fatalf("expecting final balance to be %d; got %d", expectedBalance, currentBalance)
	}

	// *************************************************************************
	// Withdraw process

	// Perform a withdraw to the player's wallet.
	withdrawTx, withdrawReceipt, err := playerClient.Withdraw(ctx)
	if err != nil {
		t.Fatalf("error calling withdraw: %s", err)
	}

	// Calculate the expected balance by adding the amount withdrawn and the
	// gas fees for the transaction.
	gasCost = big.NewInt(0).Mul(withdrawTx.GasPrice(), big.NewInt(0).SetUint64(withdrawReceipt.GasUsed))
	expectedBalance = big.NewInt(0).Add(currentBalance, depositWeiAmount)
	expectedBalance.Sub(expectedBalance, gasCost)

	// Get the updated wallet balance.
	currentBalance, err = playerClient.EthereumBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// The player's wallet balance should match the withdrawal minus the fees.
	if expectedBalance.Cmp(currentBalance) != 0 {
		t.Fatalf("expecting final balance to be %d; got %d", expectedBalance, currentBalance)
	}
}

// =============================================================================

func Test_PlaceBet(t *testing.T) {
	placeBet(t)
}

type bet struct {
	contractID      string
	player1Client   *book.Book
	player2Client   *book.Book
	ownerClient     *book.Book
	moderatorClient *book.Book
	player1Bal      *big.Float
	player2Bal      *big.Float
	placeBet        book.PlaceBet
}

func placeBet(t *testing.T) bet {
	contractID, err := deployContract()
	if err != nil {
		t.Fatalf("deploying contract: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// *************************************************************************
	// Establish books for each of the entities involved

	player1Client, err := book.New(ctx, backend, player1Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new book for player 1: %s", err)
	}

	player2Client, err := book.New(ctx, backend, player2Clt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new book for player 1: %s", err)
	}

	ownerClient, err := book.New(ctx, backend, ownerClt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new book for owner: %s", err)
	}

	moderatorClient, err := book.New(ctx, backend, moderatorClt.PrivateKey(), contractID)
	if err != nil {
		t.Fatalf("error creating new book for moderator: %s", err)
	}

	// *************************************************************************
	// Give player accounts money

	if _, _, err := player1Client.Deposit(ctx, twentyUSD); err != nil {
		t.Fatalf("error making deposit player 1: %s", err)
	}

	if _, _, err := player2Client.Deposit(ctx, twentyUSD); err != nil {
		t.Fatalf("error making deposit player 2: %s", err)
	}

	// *************************************************************************
	// Capture signatures

	signatures := make([][]byte, 2)

	signatures[0], err = player1Client.Sign(betID, 0)
	if err != nil {
		t.Fatalf("signing 1: %s", err)
	}

	signatures[1], err = player2Client.Sign(betID, 0)
	if err != nil {
		t.Fatalf("signing 2: %s", err)
	}

	// *************************************************************************
	// Place a bet

	expiration := time.Date(2022, time.September, 1, 1, 1, 1, 0, time.UTC)

	placeBet := book.PlaceBet{
		AmountBetGWei: tenUSD,
		AmountFeeGWei: oneUSD,
		Expiration:    expiration,
		Moderator:     moderatorClt.Address(),
		Participants:  []common.Address{player1Clt.Address(), player2Clt.Address()},
		Nonces:        []*big.Int{big.NewInt(0), big.NewInt(0)},
		Signatures:    signatures,
	}
	if _, _, err := ownerClient.PlaceBet(ctx, betID, placeBet); err != nil {
		t.Fatalf("error calling place bet: %s", err)
	}

	// *************************************************************************
	// Check balances

	expPlayerBal := big.NewFloat(0).Sub(twentyUSD, big.NewFloat(0).Add(placeBet.AmountBetGWei, placeBet.AmountFeeGWei))
	expOwnerBal := big.NewFloat(0).Mul(placeBet.AmountFeeGWei, big.NewFloat(2))

	if err := player1Client.CheckBalance(ctx, expPlayerBal); err != nil {
		t.Fatal(err)
	}

	if err := player2Client.CheckBalance(ctx, expPlayerBal); err != nil {
		t.Fatal(err)
	}

	if err := ownerClient.CheckBalance(ctx, expOwnerBal); err != nil {
		t.Fatal(err)
	}

	// *************************************************************************
	// Check bet state

	check := book.BetInfo{
		State:         book.StateLive,
		Participants:  placeBet.Participants,
		Moderator:     placeBet.Moderator,
		AmountBetGWei: placeBet.AmountBetGWei,
		Expiration:    placeBet.Expiration,
	}
	checkBetState(t, ctx, ownerClient, betID, check)

	// *************************************************************************
	// Get account balances for future tests

	player1Bal, err := player1Client.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting balance for player 1: %s", err)
	}

	player2Bal, err := player2Client.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting balance for player 2: %s", err)
	}

	// *************************************************************************
	// Return values for other tests

	return bet{
		contractID:      contractID,
		player1Client:   player1Client,
		player2Client:   player2Client,
		ownerClient:     ownerClient,
		moderatorClient: moderatorClient,
		player1Bal:      player1Bal,
		player2Bal:      player2Bal,
		placeBet:        placeBet,
	}
}

// =============================================================================

func Test_Reconcile(t *testing.T) {
	bet := placeBet(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// *************************************************************************
	// Generate a signature for the call

	signature, err := bet.moderatorClient.Sign(betID, 0)
	if err != nil {
		t.Fatalf("signing moderator: %s", err)
	}

	// *************************************************************************
	// Reconcile the bet

	rec := book.ReconcileBet{
		Nonce:     big.NewInt(0),
		Moderator: moderatorClt.Address(),
		Signature: signature,
		Winners:   []common.Address{player1Clt.Address()},
	}
	if _, _, err := bet.ownerClient.ReconcileBet(ctx, betID, rec); err != nil {
		t.Fatalf("error calling reconcile: %s", err)
	}

	// *************************************************************************
	// Check balances

	totalWinnings := big.NewFloat(0).Mul(bet.placeBet.AmountBetGWei, big.NewFloat(2))

	expBal := big.NewFloat(0).Add(bet.player1Bal, totalWinnings)
	if err := bet.player1Client.CheckBalance(ctx, expBal); err != nil {
		t.Fatal(err)
	}

	if err := bet.player2Client.CheckBalance(ctx, bet.player2Bal); err != nil {
		t.Fatal(err)
	}

	// *************************************************************************
	// Check bet state

	check := book.BetInfo{
		State:         book.StateReconciled,
		Moderator:     moderatorClt.Address(),
		AmountBetGWei: big.NewFloat(0),
	}
	checkBetState(t, ctx, bet.ownerClient, betID, check)
}

// =============================================================================

func Test_CancelBetModerator(t *testing.T) {
	bet := placeBet(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// *************************************************************************
	// Generate a signature for the call

	signature, err := bet.moderatorClient.Sign(betID, 0)
	if err != nil {
		t.Fatalf("signing moderator: %s", err)
	}

	// *************************************************************************
	// Cancel the bet

	cbm := book.CancelBetModerator{
		AmountFeeGWei: oneUSD,
		Nonce:         big.NewInt(0),
		Moderator:     moderatorClt.Address(),
		Signature:     signature,
	}
	if _, _, err := bet.ownerClient.CancelBetModerator(ctx, betID, cbm); err != nil {
		t.Fatalf("error calling cancel bet moderator: %s", err)
	}

	// *************************************************************************
	// Check balances

	returnBet := big.NewFloat(0).Sub(bet.placeBet.AmountBetGWei, oneUSD)

	expBal := big.NewFloat(0).Add(bet.player1Bal, returnBet)
	if err := bet.player1Client.CheckBalance(ctx, expBal); err != nil {
		t.Fatal(err)
	}

	expBal = big.NewFloat(0).Add(bet.player2Bal, returnBet)
	if err := bet.player2Client.CheckBalance(ctx, expBal); err != nil {
		t.Fatal(err)
	}

	// *************************************************************************
	// Check bet state

	check := book.BetInfo{
		State:         book.StateCancelled,
		Moderator:     moderatorClt.Address(),
		AmountBetGWei: big.NewFloat(0),
	}
	checkBetState(t, ctx, bet.ownerClient, betID, check)
}

// =============================================================================

func Test_CancelBetParticipants(t *testing.T) {
	bet := placeBet(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// *************************************************************************
	// Capture signatures

	var err error
	signatures := make([][]byte, 2)

	signatures[0], err = bet.player1Client.Sign(betID, 1)
	if err != nil {
		t.Fatalf("signing 1: %s", err)
	}

	signatures[1], err = bet.player2Client.Sign(betID, 1)
	if err != nil {
		t.Fatalf("signing 2: %s", err)
	}

	// *************************************************************************
	// Cancel bet

	cbp := book.CancelBetParticipants{
		AmountFeeGWei: oneUSD,
		Nonces:        []*big.Int{big.NewInt(1), big.NewInt(1)},
		Signatures:    signatures,
	}
	if _, _, err := bet.ownerClient.CancelBetParticipants(ctx, betID, cbp); err != nil {
		t.Fatalf("error calling cancel bet participants: %s", err)
	}

	// *************************************************************************
	// Check balances

	returnBet := big.NewFloat(0).Sub(bet.placeBet.AmountBetGWei, oneUSD)

	expBal := big.NewFloat(0).Add(bet.player1Bal, returnBet)
	if err := bet.player1Client.CheckBalance(ctx, expBal); err != nil {
		t.Fatal(err)
	}

	expBal = big.NewFloat(0).Add(bet.player2Bal, returnBet)
	if err := bet.player2Client.CheckBalance(ctx, expBal); err != nil {
		t.Fatal(err)
	}

	// *************************************************************************
	// Check bet state

	check := book.BetInfo{
		State:         book.StateCancelled,
		Moderator:     moderatorClt.Address(),
		AmountBetGWei: big.NewFloat(0),
	}
	checkBetState(t, ctx, bet.ownerClient, betID, check)
}

// =============================================================================

func Test_CancelBetOwner(t *testing.T) {
	bet := placeBet(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// *************************************************************************
	// Cancel the bet

	cbo := book.CancelBetOwner{
		AmountFeeGWei: oneUSD,
	}
	if _, _, err := bet.ownerClient.CancelBetOwner(ctx, betID, cbo); err != nil {
		t.Fatalf("error calling cancel bet owner: %s", err)
	}

	// *************************************************************************
	// Check balances

	returnBet := big.NewFloat(0).Sub(bet.placeBet.AmountBetGWei, oneUSD)

	expBal := big.NewFloat(0).Add(bet.player1Bal, returnBet)
	if err := bet.player1Client.CheckBalance(ctx, expBal); err != nil {
		t.Fatal(err)
	}

	expBal = big.NewFloat(0).Add(bet.player2Bal, returnBet)
	if err := bet.player2Client.CheckBalance(ctx, expBal); err != nil {
		t.Fatal(err)
	}

	// *************************************************************************
	// Check bet state

	check := book.BetInfo{
		State:         book.StateCancelled,
		Moderator:     moderatorClt.Address(),
		AmountBetGWei: big.NewFloat(0),
	}
	checkBetState(t, ctx, bet.ownerClient, betID, check)
}

// =============================================================================

func Test_CancelBetExpired(t *testing.T) {
	bet := placeBet(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// *************************************************************************
	// Cancel the bet

	if _, _, err := bet.player1Client.CancelBetExpired(ctx, betID); err != nil {
		t.Fatalf("error calling cancel bet owner: %s", err)
	}

	// *************************************************************************
	// Check balances

	expBal := big.NewFloat(0).Add(bet.player1Bal, bet.placeBet.AmountBetGWei)
	if err := bet.player1Client.CheckBalance(ctx, expBal); err != nil {
		t.Fatal(err)
	}

	expBal = big.NewFloat(0).Add(bet.player2Bal, bet.placeBet.AmountBetGWei)
	if err := bet.player2Client.CheckBalance(ctx, expBal); err != nil {
		t.Fatal(err)
	}

	// *************************************************************************
	// Check bet state

	check := book.BetInfo{
		State:         book.StateCancelled,
		Moderator:     moderatorClt.Address(),
		AmountBetGWei: big.NewFloat(0),
	}
	checkBetState(t, ctx, bet.ownerClient, betID, check)
}

// =============================================================================

func deployContract() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Deploying Contract ...")
	defer fmt.Println("Deployed")

	contractID, err := smartContract(ctx)
	if err != nil {
		fmt.Println("error deploying a new contract:", err)
		return "", err
	}

	return contractID, nil
}

func smartContract(ctx context.Context) (string, error) {
	tranOpts, err := ownerClt.NewTransactOpts(ctx, 10_000_000, big.NewInt(0), big.NewFloat(0))
	if err != nil {
		return "", err
	}

	address, tx, _, err := scbook.DeployBook(tranOpts, ownerClt.Backend)
	if err != nil {
		return "", err
	}

	if _, err := ownerClt.WaitMined(ctx, tx); err != nil {
		return "", err
	}

	return address.String(), nil
}

func checkBetState(t *testing.T, ctx context.Context, b *book.Book, betID string, check book.BetInfo) {
	betInfo, err := b.BetDetails(ctx, betID)
	if err != nil {
		t.Fatalf("error getting bet details: %s", err)
	}

	if betInfo.State != check.State {
		t.Errorf("invalid bet state, got %d  exp %d", betInfo.State, book.StateLive)
	}

	if check.Participants != nil {
		if len(betInfo.Participants) != len(check.Participants) {
			t.Errorf("number of participants wrong, got %d  exp %d", len(betInfo.Participants), len(check.Participants))
		}

		for i, part := range check.Participants {
			if part != betInfo.Participants[i] {
				t.Errorf("wrong participant address, got %s  exp %s", betInfo.Participants[i], part)
			}
		}
	}

	if betInfo.Moderator != check.Moderator {
		t.Errorf("wrong moderator address, got %s  exp %s", betInfo.Moderator, check.Moderator)
	}

	if check.AmountBetGWei != nil && (betInfo.AmountBetGWei.Cmp(check.AmountBetGWei) != 0) {
		t.Errorf("wrong amount, got %v  exp %v", currency.GWei2Wei(betInfo.AmountBetGWei), currency.GWei2Wei(check.AmountBetGWei))
	}

	if !check.Expiration.IsZero() && (betInfo.Expiration.UTC() != check.Expiration) {
		t.Errorf("wrong expiration, got %v  exp %v", betInfo.Expiration.UTC(), check.Expiration)
	}
}
