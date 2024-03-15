package ethereum

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

// DialedBackend represents a dialied connection to an ethereum node.
type DialedBackend struct {
	*ethclient.Client
	network string
	chainID *big.Int
}

// CreateDialedBackend constructs an ethereum client value for the specified
// network and attempts to establish a connection.
func CreateDialedBackend(ctx context.Context, network string) (*DialedBackend, error) {
	client, err := ethclient.Dial(network)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("capture chain id: %w", err)
	}

	b := DialedBackend{
		Client:  client,
		network: network,
		chainID: chainID,
	}

	return &b, nil
}

// Network returns the network the backend is connected to.
func (b *DialedBackend) Network() string {
	return b.network
}

// ChainID returns the chain id the backend is connected to.
func (b *DialedBackend) ChainID() *big.Int {
	return b.chainID
}

// =============================================================================

// SimulatedBackend represents a simulated connection to an ethereum node.
type SimulatedBackend struct {
	simulated.Client
	*simulated.Backend
	AutoCommit  bool
	PrivateKeys []*ecdsa.PrivateKey
	network     string
	chainID     *big.Int
}

// CreateSimulatedBackend constructs a simulated backend and a set of private keys
// registered to the backend with a balance of 100 ETH. Use these private keys
// with the NewSimulation call to get an Ethereum API value.
func CreateSimulatedBackend(numAccounts int, autoCommit bool, accountBalance *big.Int) (*SimulatedBackend, error) {
	keys := make([]*ecdsa.PrivateKey, numAccounts)
	alloc := make(types.GenesisAlloc)

	for i := 0; i < numAccounts; i++ {
		privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			return nil, fmt.Errorf("unable to generate private key: %w", err)
		}

		keys[i] = privateKey

		alloc[crypto.PubkeyToAddress(privateKey.PublicKey)] = types.Account{
			Balance: big.NewInt(0).Mul(accountBalance, big.NewInt(1e18)),
		}
	}

	gasLimit := uint64(9223372036854775807)
	backend := simulated.NewBackend(alloc, simulated.WithBlockGasLimit(gasLimit))

	// Setting the clock 5.15 minutes into the past to deal with bugs using
	// the simulated clock. Transactions will not be committed otherwise.
	now := time.Since(time.Date(1970, time.January, 1, 0, 5, 15, 0, time.UTC))
	backend.AdjustTime(now)

	backend.Commit()

	b := SimulatedBackend{
		Client:      backend.Client(),
		Backend:     backend,
		AutoCommit:  autoCommit,
		PrivateKeys: keys,
		network:     "simulated",
		chainID:     big.NewInt(1337),
	}

	return &b, nil
}

// Network returns the network the backend is connected to.
func (b *SimulatedBackend) Network() string {
	return b.network
}

// ChainID returns the chain id the backend is connected to.
func (b *SimulatedBackend) ChainID() *big.Int {
	return b.chainID
}

// SendTransaction functions pipes its parameters to the embedded backend and
// also calls Commit() if sb.AutoCommit==true.
func (b *SimulatedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if err := b.Client.SendTransaction(ctx, tx); err != nil {
		return err
	}

	if b.AutoCommit {
		b.Commit()
	}

	return nil
}

// SetTime creates a time shift to the simulated clock. It can only be called
// on empty blocks.
func (b *SimulatedBackend) SetTime(t time.Time) {
	b.AdjustTime(time.Since(t))
	b.Commit()
}
