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

	// The data to be signed and hashed.
	data := []byte("hello world")

	stampedData, err := stamp(data)
	if err != nil {
		return fmt.Errorf("stamp: %w", err)
	}

	// Convert the stamped data into a 32-byte array for input into the smart contract.
	hashedMessage := [32]byte{}
	copy(hashedMessage[:], stampedData)

	// Sign the message with the private key.
	signedMessage, err := Sign(data, ownerKey)
	if err != nil {
		return fmt.Errorf("signing message: %w", err)
	}

	// Retrieve the signature's v, r, s values.
	v, r, s := toSignatureValues(signedMessage)

	// The address we want to match the signature against.
	matchAddress := ethereum.Address()

	// Retrieve the address via ecrecover in the smart contract.
	extracted, err := verifyCon.AddressFromMessage(callOpts, hashedMessage, v, r, s)
	if err != nil {
		return fmt.Errorf("address from message: %w", err)
	}
	if extracted != matchAddress {
		return fmt.Errorf("addresses do not match, got[%s]; want[%s]", extracted, matchAddress)
	}

	// Verify that the address provided matches the hashedMessage + signature via
	// ecrecover in the smart contract.
	result, err := verifyCon.VerifySignature(callOpts, hashedMessage, v, r, s, matchAddress)
	if err != nil {
		return fmt.Errorf("verifying signature: %w", err)
	}
	if !result {
		return fmt.Errorf("smart contract verification failed")
	}

	fmt.Println("\nResults")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Extracted address:", extracted)
	fmt.Println("Address verified:", result)

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
