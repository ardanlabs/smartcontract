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

// SignBytes takes the specified bytes and first hashes them using Keccak256 to
// create a 32 byte array of data. Then those bytes are salted with the Ethereum
// stamp and signed using the specified private key. The signature is returned as
// a hexadecimal string.
func SignBytes(bytes []byte, privateKey *ecdsa.PrivateKey) (signature string, err error) {

	// Prepare the data for signing.
	data, err := stamp(bytes)
	if err != nil {
		return "", err
	}

	// Sign the hash with the private key to produce a signature.
	sig, err := crypto.Sign(data[:], privateKey)
	if err != nil {
		return "", err
	}

	// Extract the bytes for the original public key.
	publicKeyOrg := privateKey.Public()
	publicKeyECDSA, ok := publicKeyOrg.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	// Check the public key validates the data and signature.
	rs := sig[:crypto.RecoveryIDOffset]
	if !crypto.VerifySignature(publicKeyBytes, data[:], rs) {
		return "", errors.New("invalid signature produced")
	}

	// Add the Ethereum ID to the last byte represents v.
	sig[crypto.RecoveryIDOffset] += ethID

	return fmt.Sprintf("0x%s", hex.EncodeToString(sig)), nil
}

// SignAny marshals the specified value into JSON to create an array of bytes.
// Then the SignBytes function is used to sign the data.
func SignAny(value any, privateKey *ecdsa.PrivateKey) (signature string, err error) {

	// Convert the value to bytes using the json marshaler.
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return SignBytes(bytes, privateKey)
}

// FromAddressBytes extracts the address for the account that signed the data.
// The signature must be provided as a hexadecimal string.
func FromAddressBytes(bytes []byte, signature string) (string, error) {

	// Perform a basic check that the signature is formatted properly.
	sig, err := verifySignature(signature)
	if err != nil {
		return "", fmt.Errorf("validating signature: %w", err)
	}

	// Prepare the data for public key extraction.
	stampData, err := stamp(bytes)
	if err != nil {
		return "", err
	}

	// Decrement the Ethereum ID from the V value.
	sig[64] = sig[64] - ethID

	// Capture the public key associated with this data and signature.
	publicKey, err := crypto.SigToPub(stampData[:], sig)
	if err != nil {
		return "", err
	}

	// Extract the account address from the public key.
	return crypto.PubkeyToAddress(*publicKey).String(), nil
}

// FromAddressAny marshals the specified value into JSON to create an array of
// bytes. Then uses FromAddressBytes function to extract the address.
func FromAddressAny(value any, signature string) (string, error) {

	// Convert the value to bytes using the json marshaler.
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return FromAddressBytes(bytes, signature)
}

// =============================================================================

// stamp returns a hash of 32 bytes that represents this data with the Ethereum
// stamp embedded into the final hash. This is the data that is signed.
func stamp(b []byte) ([32]byte, error) {

	// Convert the stamp into a slice of bytes. This stamp is
	// used so signatures we produce when signing data
	// are always unique to the Ethereum blockchain.
	stamp := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(b)))

	// Hash the stamp and txHash together in a final 32 byte array
	// that represents the data.
	data := crypto.Keccak256(stamp, b)

	// The final data to be signed will always be a 32 byte array and
	// it's best to return it as such to make this obvious.
	var data32 [32]byte
	copy(data32[:], data)

	return data32, nil
}

// verifySignature verifies the signature conforms to our standards. The
// signature must be provided as a hexadecimal string.
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
