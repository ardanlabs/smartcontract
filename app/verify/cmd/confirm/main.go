package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/verify/contract/go/verify"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	keyStoreFile     = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase       = "123"
	coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

// ZeroHash represents a hash code of zeros.
const ZeroHash string = "0x0000000000000000000000000000000000000000000000000000000000000000"

// ardanID is an arbitrary number for signing messages. This will make it
// clear that the signature comes from the Ardan blockchain.
// Ethereum and Bitcoin do this as well, but they use the value of 27.
const ethID = 27

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() (err error) {
	ctx := context.Background()

	ownerKey, err := ethereum.PrivateKeyByKeyFile(keyStoreFile, passPhrase)
	if err != nil {
		return err
	}

	ethereum, err := ethereum.New(ctx, ethereum.NetworkLocalhost, keyStoreFile, passPhrase)
	if err != nil {
		return err
	}

	fmt.Println("\nInput Values")
	fmt.Println("----------------------------------------------------")
	fmt.Println("fromAddress:", ethereum.Address())

	// =========================================================================

	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter()
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("oneETHToUSD:", oneETHToUSD)
	fmt.Println("oneUSDToETH:", oneUSDToETH)

	// =========================================================================

	callOpts, err := ethereum.NewCallOpts(ctx)
	if err != nil {
		return err
	}

	// =========================================================================

	contractIDBytes, err := os.ReadFile("zarf/ethereum/verify.cid")
	if err != nil {
		return fmt.Errorf("importing verify.cid file: %w", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return errors.New("need to export the verify.cid file")
	}
	fmt.Println("contractID:", contractID)

	verifyCon, err := verify.NewVerify(common.HexToAddress(contractID), ethereum.RawClient())
	if err != nil {
		return fmt.Errorf("new contract: %w", err)
	}

	// =========================================================================

	// Sign the message with the private key.
	signedMessage, err := Sign([]byte("hello world"), ownerKey)
	if err != nil {
		return fmt.Errorf("signing message: %w", err)
	}

	// Convert the []byte signedMessage into a [32]byte array for smart contract input.
	signedHash := [32]byte{}
	copy(signedHash[:], signedMessage)

	// Retrieve the signature's v, r, s values.
	v, r, s := toSignatureValues(signedMessage)

	// Perform a local verification.
	localFoundAddr, err := FromAddress([]byte("hello world"), v, r, s)
	if err != nil {
		return fmt.Errorf("from address: %w", err)
	}

	// The address we want to match the signature against.
	matchAddress := ethereum.Address()

	// Verify the signed hash in the smart contract.
	result, err := verifyCon.VerifyMessage(callOpts, signedHash, v, r, s)
	if err != nil {
		return fmt.Errorf("verify message: %w", err)
	}

	// The hashed message before converting to bytes32.
	fmt.Printf("msg:     %x\n", signedMessage)

	// The hashed message after converting to bytes32.
	fmt.Printf("msg[32]: %x\n", signedHash)

	// RSV signature values extracted from signedMessage.
	fmt.Printf("r:       %x\n", r)
	fmt.Printf("s:       %x\n", s)
	fmt.Printf("v:       %d\n", v)

	// The address we signed with and want to match.
	fmt.Println("target: ", matchAddress)

	// The address found by a local FromAddress call.
	fmt.Println("local:  ", localFoundAddr)

	// The address returned by the smart contract.
	fmt.Println("result: ", result)

	return nil
}

// Sign uses the specified private key to sign the data.
func Sign(value any, privateKey *ecdsa.PrivateKey) ([]byte, error) {

	// Prepare the data for signing.
	data, err := stamp(value)
	if err != nil {
		return nil, err
	}

	// Sign the hash with the private key to produce a signature.
	sig, err := crypto.Sign(data, privateKey)
	if err != nil {
		return nil, err
	}

	// Extract the bytes for the original public key.
	publicKeyOrg := privateKey.Public()
	publicKeyECDSA, ok := publicKeyOrg.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	// Check the public key validates the data and signature.
	rs := sig[:crypto.RecoveryIDOffset]
	if !crypto.VerifySignature(publicKeyBytes, data, rs) {
		return nil, errors.New("invalid signature produced")
	}

	return sig, nil
}

// toSignatureValues converts the signature into the r, s, v values.
func toSignatureValues(sig []byte) (v uint8, r, s [32]byte) {
	copy(r[:], sig[:32])
	copy(s[:], sig[32:64])
	v = uint8(sig[64]) + ethID

	return v, r, s
}

// stamp returns a hash of 32 bytes that represents this data with
// the Ethereum stamp embedded into the final hash.
func stamp(value any) ([]byte, error) {

	// Marshal the data.
	v, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	// Hash the data data into a 32 byte array. This will provide
	// a data length consistency with all data.
	txHash := crypto.Keccak256(v)

	// Convert the stamp into a slice of bytes. This stamp is
	// used so signatures we produce when signing data
	// are always unique to the Ardan blockchain.
	stamp := []byte("\x19Ethereum Signed Message:\n32")

	// Hash the stamp and txHash together in a final 32 byte array
	// that represents the data.
	data := crypto.Keccak256(stamp, txHash)

	return data, nil
}

// FromAddress extracts the address for the account that signed the data.
func FromAddress(value any, v uint8, r, s [32]byte) (string, error) {

	// NOTE: If the same exact data for the given signature is not provided
	// we will get the wrong from address for this transaction. There is no
	// way to check this on the node since we don't have a copy of the public
	// key used. The public key is being extracted from the data and signature.

	// Prepare the data for public key extraction.
	data, err := stamp(value)
	if err != nil {
		return "", err
	}

	// Convert the [R|S|V] format into the original 65 bytes.
	sig := ToSignatureBytes(v, r, s)

	// Capture the public key associated with this data and signature.
	publicKey, err := crypto.SigToPub(data, sig)
	if err != nil {
		return "", err
	}

	// Extract the account address from the public key.
	return crypto.PubkeyToAddress(*publicKey).String(), nil
}

// ToSignatureBytes converts the r, s, v values into a slice of bytes
// with the removal of the ardanID.
func ToSignatureBytes(v uint8, r, s [32]byte) []byte {
	sig := make([]byte, crypto.SignatureLength)
	copy(sig, r[:])
	copy(sig[32:], s[:])
	sig[64] = byte(v - ethID)
	return sig
}
