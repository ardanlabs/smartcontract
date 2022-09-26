package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/signature/contract/go/verify"
	"github.com/ethereum/go-ethereum/common"
)

const (
	keyStoreFile     = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase       = "123"
	coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

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

	eth, err := ethereum.New(ctx, ethereum.NetworkLocalhost, keyStoreFile, passPhrase)
	if err != nil {
		return err
	}

	fmt.Println("\nInput Values")
	fmt.Println("----------------------------------------------------")
	fmt.Println("fromAddress:", eth.Address())

	// =========================================================================

	converter, err := currency.NewConverter(coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter()
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("oneETHToUSD:", oneETHToUSD)
	fmt.Println("oneUSDToETH:", oneUSDToETH)

	// =========================================================================

	callOpts, err := eth.NewCallOpts(ctx)
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
	fmt.Println("contractID :", contractID)

	verify, err := verify.NewVerify(common.HexToAddress(contractID), eth.RawClient())
	if err != nil {
		return fmt.Errorf("new contract: %w", err)
	}

	// =========================================================================

	// The data to be signed and hashed.
	value := []byte("hello world")

	// Sign the message with the private key.
	signedMessage, err := ethereum.Sign(value, ownerKey)
	if err != nil {
		return fmt.Errorf("signing message: %w", err)
	}

	// =========================================================================

	data, err := ethereum.ValueToBytes(value)
	if err != nil {
		return fmt.Errorf("getting bytes: %w", err)
	}

	// Retrieve the address via the smart contract Address call.
	sigAddress, err := verify.Address(callOpts, data, signedMessage)
	if err != nil {
		return fmt.Errorf("address from message: %w", err)
	}

	// Retrieve the address via the smart contract Address call.
	matched, err := verify.MatchSender(callOpts, data, signedMessage)
	if err != nil {
		return fmt.Errorf("address from message: %w", err)
	}

	fmt.Println("\nResults")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Keyfile address   :", eth.Address())
	fmt.Println("Signature address :", sigAddress)
	fmt.Println("Match sender      :", matched)

	return nil
}
