package verify_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/smartcontract/app/signature/contract/go/verify"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	deployerAcc = iota
	numAccounts
	keyStoreFile = "../../../../../zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase   = "123"
)

func TestVerify(t *testing.T) {
	ctx := context.Background()

	backend, err := ethereum.CreateSimulatedBackend(numAccounts, true, big.NewInt(100))
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer backend.Close()

	deployer, err := ethereum.NewClient(backend, backend.PrivateKeys[deployerAcc])
	if err != nil {
		t.Fatalf("unable to create deployerAcc: %s", err)
	}

	callOpts, err := deployer.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	// =========================================================================

	const gasLimit = 1700000
	const valueGwei = 0.0

	var testVerify *verify.Verify

	// =========================================================================

	t.Run("deploy verify", func(t *testing.T) {
		deployTranOpts, err := deployer.NewTransactOpts(ctx, gasLimit, big.NewInt(0), big.NewFloat(valueGwei))
		if err != nil {
			t.Fatalf("unable to create transaction opts for verify: %s", err)
		}

		contractID, tx, _, err := verify.DeployVerify(deployTranOpts, deployer.Backend)
		if err != nil {
			t.Fatalf("unable to deploy verify: %s", err)
		}

		if _, err := deployer.WaitMined(ctx, tx); err != nil {
			t.Fatalf("waiting for deploy: %s", err)
		}

		testVerify, err = verify.NewVerify(contractID, deployer.Backend)
		if err != nil {
			t.Fatalf("unable to create a verify: %s", err)
		}
	})

	// =========================================================================

	t.Run("match sender", func(t *testing.T) {
		id := "asdjh1231"
		nonce := big.NewInt(1)

		bytes, err := encodeForSolidity(id, deployer.Address(), nonce)
		if err != nil {
			t.Fatalf("should be able to encode data: %s", err)
		}

		signature, err := ethereum.SignBytes(bytes, deployer.PrivateKey())
		if err != nil {
			t.Fatalf("should be able to sign the message: %s", err)
		}

		sig, err := hex.DecodeString(signature[2:])
		if err != nil {
			t.Fatalf("should be able to decode the signature: %s", err)
		}

		addr, err := testVerify.Address(callOpts, id, deployer.Address(), nonce, sig)
		if err != nil {
			t.Fatalf("should be able to get address from signature: %s", err)
		}

		if addr.Hex() != deployer.Address().Hex() {
			t.Log("got:", addr.Hex())
			t.Log("exp:", deployer.Address().Hex())
			t.Fatalf("should be able to match the addresses: %s", err)
		}

		matched, err := testVerify.MatchSender(callOpts, id, deployer.Address(), nonce, sig)
		if err != nil {
			t.Fatalf("should be able to match the sender: %s", err)
		}

		if !matched {
			t.Fatal("Match failed!")
		}
	})
}

// encodeForSolidity will take the arguments and pack them into a byte array that
// conforms to the solidity abi.encode function.
func encodeForSolidity(id string, participant common.Address, nonce *big.Int) ([]byte, error) {
	String, _ := abi.NewType("string", "", nil)
	Address, _ := abi.NewType("address", "", nil)
	Uint, _ := abi.NewType("uint", "", nil)

	arguments := abi.Arguments{
		{
			Type: String,
		},
		{
			Type: Address,
		},
		{
			Type: Uint,
		},
	}

	bytes, err := arguments.Pack(id, participant, nonce)
	if err != nil {
		return nil, fmt.Errorf("arguments pack: %w", err)
	}

	bytes = crypto.Keccak256(bytes)

	return bytes, nil
}
