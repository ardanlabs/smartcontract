package verify_test

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/smartcontract/app/signature/contract/go/verify"
	"github.com/divergencetech/ethier/ethtest"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	deployer     = 0
	keyStoreFile = "../../../../../zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase   = "123"
)

func TestVerify(t *testing.T) {
	sim, err := ethtest.NewSimulatedBackend(1)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}

	contractID, _, _, err := verify.DeployVerify(sim.Acc(deployer), sim)
	if err != nil {
		t.Fatalf("unable to deploy verify: %s", err)
	}

	verify, err := verify.NewVerify(contractID, sim)
	if err != nil {
		t.Fatalf("unable to create a verify: %s", err)
	}

	id := "asdjh1231"
	participant := sim.Addr(deployer)
	nonce := big.NewInt(1)

	bytes, err := encodeForSolidity(id, participant, nonce)
	if err != nil {
		t.Fatalf("should be able to encode data: %s", err)
	}

	signature, err := ethereum.SignBytes(bytes, sim.PrivateKey(deployer))
	if err != nil {
		t.Fatalf("should be able to sign the message: %s", err)
	}

	sig, err := hex.DecodeString(signature[2:])
	if err != nil {
		t.Fatalf("should be able to decode the signature: %s", err)
	}

	matched, err := verify.MatchSender(sim.CallFrom(deployer), id, participant, nonce, sig)
	if err != nil {
		t.Fatalf("should be able to match the sender: %s", err)
	}

	if !matched {
		t.Fatal("Match failed!")
	}
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

	return bytes, nil
}
