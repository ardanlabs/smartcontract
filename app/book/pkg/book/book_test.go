package book_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	scbook "github.com/ardanlabs/smartcontract/app/book/contract/go/book"
	"github.com/ardanlabs/smartcontract/app/book/pkg/book"
	"github.com/ethereum/go-ethereum/common"
)

const (
	OwnerAddress    = "0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd"
	OwnerKeyPath    = "../../../../zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	OwnerPassPhrase = "123"

	Player1Address    = "0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1KeyPath    = "../../../../zarf/ethereum/keystore/UTC--2022-05-13T16-59-42.277071000Z--0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1PassPhrase = "123"

	Player2Address    = "0x8e113078adf6888b7ba84967f299f29aece24c55"
	Player2KeyPath    = "../../../../zarf/ethereum/keystore/UTC--2022-05-13T16-57-20.203544000Z--8e113078adf6888b7ba84967f299f29aece24c55"
	Player2PassPhrase = "123"

	ModeratorAddress    = "0x40CFaB8ab694937d644764A3f58237be4c568458"
	ModeratorKeyPath    = "../../../../zarf/ethereum/keystore/UTC--2022-09-29T16-18-17.064954000Z--40cfab8ab694937d644764a3f58237be4c568458"
	ModeratorPassPhrase = "123"
)

// These variables provide some static GWei to play with.
var (
	oneUSD    = big.NewFloat(662_833.00)
	tenUSD    = big.NewFloat(0).Mul(oneUSD, big.NewFloat(10))
	twentyUSD = big.NewFloat(0).Mul(oneUSD, big.NewFloat(20))
	fiftyUSD  = big.NewFloat(0).Mul(oneUSD, big.NewFloat(50))
)

// We need a string for the bet id.
var betID = "1234"

// =============================================================================

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ethereum, err := ethereum.New(ctx, ethereum.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Adding money to player 1 account")

	// Add money to this account.
	if err := ethereum.SendTransaction(ctx, common.HexToAddress(Player1Address), currency.GWei2Wei(fiftyUSD), 21000); err != nil {
		fmt.Println("Player1Address:", err)
		os.Exit(1)
	}

	fmt.Println("Adding money to player 2 account")

	// Add money to this account.
	if err := ethereum.SendTransaction(ctx, common.HexToAddress(Player2Address), currency.GWei2Wei(fiftyUSD), 21000); err != nil {
		fmt.Println("Player2Address:", err)
		os.Exit(1)
	}

	m.Run()
}

// =============================================================================

func Test_DepositWithdraw(t *testing.T) {
	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect player 1 to the smart contract.
	playerClient, err := book.New(ctx, ethereum.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new book for owner: %s", err)
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
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// *************************************************************************
	// Establish books for each of the entities involved

	player1Client, err := book.New(ctx, ethereum.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new book for player 1: %s", err)
	}

	player2Client, err := book.New(ctx, ethereum.NetworkHTTPLocalhost, Player2KeyPath, Player2PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new book for player 1: %s", err)
	}

	ownerClient, err := book.New(ctx, ethereum.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new book for owner: %s", err)
	}

	moderatorClient, err := book.New(ctx, ethereum.NetworkHTTPLocalhost, ModeratorKeyPath, ModeratorPassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new book for moderator: %s", err)
	}

	// *************************************************************************
	// Give player accounts money

	if _, _, err := player1Client.Deposit(ctx, twentyUSD); err != nil {
		t.Fatalf("error making deposit player 1: %s", err)
	}

	if _, _, err := player2Client.Deposit(ctx, twentyUSD); err != nil {
		t.Fatalf("error making deposit player 1: %s", err)
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
		Moderator:     ModeratorAddress,
		Participants:  []string{Player1Address, Player2Address},
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
		Moderator: ModeratorAddress,
		Signature: signature,
		Winners:   []string{Player1Address},
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
		Moderator:     ModeratorAddress,
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
		AmountBetGWei: big.NewFloat(0),
	}
	checkBetState(t, ctx, bet.ownerClient, betID, check)
}

// =============================================================================

func deployContract() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("*** Deploying Contract ***")

	contractID, err := smartContract(ctx)
	if err != nil {
		fmt.Println("error deploying a new contract:", err)
		return "", err
	}

	return contractID, nil
}

func smartContract(ctx context.Context) (string, error) {
	ethereum, err := ethereum.New(ctx, ethereum.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase)
	if err != nil {
		return "", err
	}

	tranOpts, err := ethereum.NewTransactOpts(ctx, 5_000_000, big.NewFloat(0))
	if err != nil {
		return "", err
	}

	address, tx, _, err := scbook.DeployBook(tranOpts, ethereum.RawClient())
	if err != nil {
		return "", err
	}

	if _, err := ethereum.WaitMined(ctx, tx); err != nil {
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
			if !strings.EqualFold(part, betInfo.Participants[i]) {
				t.Errorf("wrong participant address, got %s  exp %s", betInfo.Participants[i], part)
			}
		}
	}

	if check.Moderator != "" && !strings.EqualFold(betInfo.Moderator, check.Moderator) {
		t.Errorf("wrong moderator address, got %s  exp %s", betInfo.Moderator, check.Moderator)
	}

	if check.AmountBetGWei != nil && (betInfo.AmountBetGWei.Cmp(check.AmountBetGWei) != 0) {
		t.Errorf("wrong amount, got %v  exp %v", currency.GWei2Wei(betInfo.AmountBetGWei), currency.GWei2Wei(check.AmountBetGWei))
	}

	if !check.Expiration.IsZero() && (betInfo.Expiration.UTC() != check.Expiration) {
		t.Errorf("wrong expiration, got %v  exp %v", betInfo.Expiration.UTC(), check.Expiration)
	}
}
