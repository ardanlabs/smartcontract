package book

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// BetInfo represents information about a bet inside the contract.
type BetInfo struct {
	State         int
	Participants  []common.Address
	Moderator     common.Address
	AmountBetGWei *big.Float
	Expiration    time.Time
}

// PlaceBet represents the input required to place a bet.
type PlaceBet struct {
	AmountBetGWei *big.Float
	AmountFeeGWei *big.Float
	Expiration    time.Time
	Moderator     common.Address
	Participants  []common.Address
	Nonces        []*big.Int
	Signatures    [][]byte
}

// Validate verifies the new bet value is properly initialized.
func (pb PlaceBet) Validate() error {
	return nil
}

// ReconcileBet represents the input required to reconcile a bet.
type ReconcileBet struct {
	Nonce     *big.Int
	Moderator common.Address
	Signature []byte
	Winners   []common.Address
}

// Validate verifies the reconcile bet value is properly initialized.
func (rb ReconcileBet) Validate() error {
	return nil
}

// CancelBetModerator represents the input required to cancel a bet by the moderator.
type CancelBetModerator struct {
	AmountFeeGWei *big.Float
	Nonce         *big.Int
	Moderator     common.Address
	Signature     []byte
}

// Validate verifies the reconcile bet value is properly initialized.
func (cbm CancelBetModerator) Validate() error {
	return nil
}

// CancelBetParticipants represents the input required to cancel a bet by the participants.
type CancelBetParticipants struct {
	AmountFeeGWei *big.Float
	Nonces        []*big.Int
	Signatures    [][]byte
}

// Validate verifies the reconcile bet value is properly initialized.
func (cbp CancelBetParticipants) Validate() error {
	return nil
}

// CancelBetOwner represents the input required to cancel a bet by the contract owner.
type CancelBetOwner struct {
	AmountFeeGWei *big.Float
}

// Validate verifies the reconcile bet value is properly initialized.
func (cbo CancelBetOwner) Validate() error {
	return nil
}
