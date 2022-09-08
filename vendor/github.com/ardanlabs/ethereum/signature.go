package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

// ZeroHash represents a hash code of zeros.
const ZeroHash string = "0x0000000000000000000000000000000000000000000000000000000000000000"

// ardanID is an arbitrary number for signing messages. This will make it
// clear that the signature comes from the Ardan blockchain.
// Ethereum and Bitcoin do this as well, but they use the value of 27.
const ethID = 27

// =============================================================================

// PrivateKeyByKeyFile opens a key file for the private key.
func PrivateKeyByKeyFile(keyFile string, passPhrase string) (*ecdsa.PrivateKey, error) {
	data, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("read key file: %w", err)
	}

	key, err := keystore.DecryptKey(data, passPhrase)
	if err != nil {
		return nil, fmt.Errorf("decrypt key file: %w", err)
	}

	return key.PrivateKey, nil
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

	// Add the Ethereum ID to the last byte represents v.
	sig[crypto.RecoveryIDOffset] += ethID

	return sig, nil
}

// FromAddress extracts the address for the account that signed the data.
func FromAddress(value any, signature string) (string, error) {

	// Perform a basic check that the signature is formatted properly.
	sig, err := verifySignature(signature)
	if err != nil {
		return "", fmt.Errorf("validating signature: %w", err)
	}

	// NOTE: If the same exact data for the given signature is not provided
	// we will get the wrong from address for this transaction. There is no
	// way to check this on the node since we don't have a copy of the public
	// key used. The public key is being extracted from the data and signature.

	// Prepare the data for public key extraction.
	data, err := stamp(value)
	if err != nil {
		return "", err
	}

	// Decrement the Ethereum ID from the V value.
	sig[64] = sig[64] - ethID

	// Capture the public key associated with this data and signature.
	publicKey, err := crypto.SigToPub(data, sig)
	if err != nil {
		return "", err
	}

	// Extract the account address from the public key.
	return crypto.PubkeyToAddress(*publicKey).String(), nil
}

// =============================================================================

// verifySignature verifies the signature conforms to our standards.
func verifySignature(signature string) ([]byte, error) {

	// Convert the signature to a 65 bytes.
	sig, err := hex.DecodeString(signature[2:])
	if err != nil {
		return nil, err
	}

	// Convert the string to the underlying slice of bytes and extract
	// the [R|S|V] values from the signature.
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])
	v := new(big.Int).SetBytes([]byte{sig[64]})

	// Check the recovery id is either 0 or 1.
	uintV := v.Uint64() - ethID
	if uintV != 0 && uintV != 1 {
		return nil, errors.New("invalid recovery id")
	}

	// Check the signature values are valid.
	if !crypto.ValidateSignatureValues(byte(uintV), r, s, false) {
		return nil, errors.New("invalid signature values")
	}

	return sig, nil
}

// stamp returns a hash of 32 bytes that represents this data with
// the Ardan stamp embedded into the final hash.
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
