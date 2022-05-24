package smart

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

// These values are used to calculate ETH in USD.
var (
	AmountWei = big.NewInt(10_000_000)
	PriceWei  = big.NewFloat(0.00000002)
)

// Hardcoding these for now since all the apps will use this same information
// and the information is static.
const (
	// geth = "http://localhost:8545"
	geth           = "zarf/ethereum/geth.ipc"
	primaryAccount = "zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	passPhrase     = "123"
)

// Connect provides boilerplate for connecting to the geth service using
// an IPC socket created by the geth service on startup.
func Connect() (*ethclient.Client, *ecdsa.PrivateKey, error) {
	client, err := ethclient.Dial(geth)
	if err != nil {
		return nil, nil, fmt.Errorf("DialConnect: %w", err)
	}

	privateKey, err := privateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("privateKey: %w", err)
	}

	return client, privateKey, nil
}

// NewTransaction constructs a new transaction for executing function that cost money.
func NewTransaction(ctx context.Context, gasLimit uint64, pk *ecdsa.PrivateKey, client *ethclient.Client) (*bind.TransactOpts, error) {
	address := crypto.PubkeyToAddress(pk.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("PendingNonceAt: %w", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("ChainID: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("SuggestGasPrice: %w", err)
	}

	tran, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return nil, fmt.Errorf("NewKeyedTransactorWithChainID: %w", err)
	}

	tran.Nonce = big.NewInt(int64(nonce))
	tran.Value = big.NewInt(0) // in wei
	tran.GasLimit = gasLimit   // The maximum amount of Gas a user can consume to conduct this transaction.
	tran.GasPrice = gasPrice   // What you are willing to pay per unit to complete this transaction.

	return tran, nil
}

// WaitMined will wait for the transaction to be minded and return a receipt.
func WaitMined(ctx context.Context, tx *types.Transaction, fromAddress common.Address, client *ethclient.Client) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return nil, err
	}

	if receipt.Status == 0 {
		err := ExtractError(client, tx, fromAddress)
		return nil, err
	}

	return receipt, nil
}

// ExtractError checks the failed transaction for the error message.
func ExtractError(client *ethclient.Client, tx *types.Transaction, fromAddress common.Address) error {
	msg := ethereum.CallMsg{
		From:     fromAddress,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	_, err := client.CallContract(context.Background(), msg, nil)
	return err
}

// PrintTransaction outputs the transaction details.
func PrintTransaction(tx *types.Transaction) {
	fmt.Println("tx sent        :", tx.Hash().Hex())
	fmt.Println("tx gas Eth     :", Wei2Eth(tx.GasPrice()))
	fmt.Println("tx gas allowed :", tx.Gas())
	fmt.Println("tx value Eth   :", Wei2Eth(tx.Value()))
	fmt.Println("tx max Eth     :", Wei2Eth(tx.Cost()), " // gas * gasPrice + value")
	fmt.Println("tx max USD     :", USDCost(tx.Cost()))
}

// PrintTransactionReceipt outputs the transaction receipt.
func PrintTransactionReceipt(receipt *types.Receipt, tx *types.Transaction) {
	cost := big.NewInt(0).Mul(big.NewInt(int64(receipt.GasUsed)), tx.GasPrice())

	fmt.Println("tx gas used    :", receipt.GasUsed)
	fmt.Println("tx cost eth    :", Wei2Eth(cost), " // gasUsed * gasPrice + value")
	fmt.Println("tx cost USD    :", USDCost(cost))
	fmt.Println("tx status      :", receipt.Status)

	topic := crypto.Keccak256Hash([]byte("Log(string)"))
	if len(receipt.Logs) > 0 {
		fmt.Println("================ LOGS =================")
		for _, v := range receipt.Logs {
			if v.Topics[0] == topic {
				l := v.Data[63]
				fmt.Println(string(v.Data[64 : 64+l]))
			}
		}
		fmt.Println("=======================================")
	}
}

// PrintBalanceDiff outputs the start and ending balances with difference.
func PrintBalanceDiff(ctx context.Context, startingBalance *big.Int, fromAddress common.Address, client *ethclient.Client) error {
	endingBalance, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		return err
	}

	cost := big.NewInt(0).Sub(startingBalance, endingBalance)

	fmt.Println("bal before Eth :", Wei2Eth(startingBalance))
	fmt.Println("bal after  Eth :", Wei2Eth(endingBalance))
	fmt.Println("bal diff   Eth :", Wei2Eth(cost))
	fmt.Println("diff cost  USD :", USDCost(cost))

	return nil
}

// PrintBaseFee calculates the base fee for the latest block.
func PrintBaseFee(client *ethclient.Client) {
	bn, _ := client.BlockNumber(context.Background())
	bignumBn := big.NewInt(0).SetUint64(bn)
	blk, _ := client.BlockByNumber(context.Background(), bignumBn)
	baseFee := misc.CalcBaseFee(params.RopstenChainConfig, blk.Header())

	fmt.Println("base fee       :", Wei2Eth(baseFee))
}

// USDCost converts Wei to USD.
func USDCost(amount *big.Int) string {
	units := big.NewFloat(0).SetInt(big.NewInt(0).Div(amount, AmountWei))
	ans, _ := big.NewFloat(0).Mul(units, PriceWei).Float64()
	return fmt.Sprintf("$%.2f", ans)
}

// Wei2Eth converts the wei unit into a Eth for display.
func Wei2Eth(amount *big.Int) string {
	compact_amount := big.NewInt(0)
	reminder := big.NewInt(0)
	divisor := big.NewInt(1e18)
	compact_amount.QuoRem(amount, divisor, reminder)
	return fmt.Sprintf("%s.%018s", compact_amount.String(), reminder.String())
}

func privateKey() (*ecdsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(primaryAccount)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(data, passPhrase)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}
