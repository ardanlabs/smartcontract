package ethereum

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// SimulatedBackend represents an Ethereum node running for simulation purposes.
type SimulatedBackend struct {
	*backends.SimulatedBackend
	AutoCommit  bool
	PrivateKeys []*ecdsa.PrivateKey
}

// CreateSimulation constructs a simulated backend and a set of private keys
// registered to the backend. You must call NewAccount to get a working
// Ethereum value for that key.
func CreateSimulation(numAccounts int, autoCommit bool) (*SimulatedBackend, error) {
	keys := make([]*ecdsa.PrivateKey, numAccounts)
	alloc := make(core.GenesisAlloc)

	for i := 0; i < numAccounts; i++ {
		privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			return nil, fmt.Errorf("unable to generate private key: %w", err)
		}

		keys[i] = privateKey

		alloc[crypto.PubkeyToAddress(privateKey.PublicKey)] = core.GenesisAccount{
			Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18)),
		}
	}

	client := backends.NewSimulatedBackend(alloc, 3e7)
	client.AdjustTime(365 * 24 * time.Hour)
	client.Commit()

	sb := SimulatedBackend{
		SimulatedBackend: client,
		AutoCommit:       autoCommit,
		PrivateKeys:      keys,
	}

	return &sb, nil
}

// SendTransaction functions pipes its parameters to the embedded backend and
// also calls Commit() if sb.AutoCommit==true.
func (sb *SimulatedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if err := sb.SimulatedBackend.SendTransaction(ctx, tx); err != nil {
		return err
	}
	if sb.AutoCommit {
		sb.SimulatedBackend.Commit()
	}
	return nil
}
