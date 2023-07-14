package evmbehaviour

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client
var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		Name:                 "eip1559",
		Options:              &o,
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		actual = nil
		expected = nil

		var err error
		client, err = ethclient.Dial("http://localhost:40002")
		if err != nil {
			panic(err)
		}
	})
}

func sendDynamicFeeTransaction(recipientHex string) error {
	ctx := context.TODO()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA("0c17200f35ddfacfc5f5d7458cc744956efa34ae9a6e63ab1f714d1437fbcb7a")
	if err != nil {
		return err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKey)

	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		return err
	}

	recipient := common.HexToAddress(recipientHex)

	latestBlock, err := client.BlockNumber(ctx)
	if err != nil {
		return err
	}

	feeHistory, err := client.FeeHistory(ctx, 1, new(big.Int).SetUint64(latestBlock), []float64{95, 99})
	if err != nil {
		return err
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		To:        &recipient,
		Value:     big.NewInt(100000000),
		Gas:       21000,
		ChainID:   chainID,
		GasTipCap: new(big.Int).Div(feeHistory.BaseFee[0], big.NewInt(2)),
		GasFeeCap: feeHistory.BaseFee[0],
	})

	expected, err = types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		return err
	}

	err = client.SendTransaction(ctx, expected)
	if err != nil {
		return err
	}

	return nil
}

func getTransactionByHash() error {
	ctx := context.TODO()

	var err error

	actual, _, err = client.TransactionByHash(ctx, expected.Hash())
	if err != nil {
		return err
	}

	if actual == nil {
		return fmt.Errorf("Failed to retrieve transaction %s", expected.Hash())
	}

	return nil
}

func assertTransaction() error {
	if actual.Hash() != expected.Hash() {
		return fmt.Errorf("Expected hash to be [%s] but was [%s]", expected.Hash(), actual.Hash())
	}

	if actual.ChainId().Cmp(expected.ChainId()) != 0 {
		return fmt.Errorf("Expected chain ID to be [%d] but was [%d]", expected.ChainId(), actual.ChainId())
	}

	if actual.Nonce() != expected.Nonce() {
		return fmt.Errorf("Expected nonce to be [%d] but was [%d]", expected.Nonce(), actual.Nonce())
	}

	if actual.To().Hex() != expected.To().Hex() {
		return fmt.Errorf("Expected recipient to be [%s] but was [%s]", expected.To(), actual.To())
	}

	if actual.GasFeeCap().Cmp(expected.GasFeeCap()) != 0 {
		return fmt.Errorf("Expected gas fee cap to be [%s] but was [%s]", expected.GasFeeCap(), actual.GasFeeCap())
	}

	if actual.GasTipCap().Cmp(expected.GasTipCap()) != 0 {
		return fmt.Errorf("Expected gas tip cap to be [%s] but was [%s]", expected.GasTipCap(), actual.GasTipCap())
	}

	return nil
}

func assertReceipt() error {
	ctx := context.TODO()

	var err error
	var receipt *types.Receipt
	deadline := time.Now().Add(10 * time.Second)
	for {
		receipt, err = client.TransactionReceipt(ctx, actual.Hash())
		if err == nil && receipt != nil || time.Now().After(deadline) {
			break
		}
	}

	if err != nil {
		return err
	}

	if receipt.Status != 1 {
		return fmt.Errorf("Expected transaction [%s] to be successful", actual.Hash())
	}

	return nil
}

/*
Given I send a dynamic fee transaction

	When I call eth_getTransactionByHash
	Then I get all the details for my transaction
*/
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// clean the state before every scenario
		expected = nil
		actual = nil

		return ctx, nil
	})

	ctx.Step(`^I send a dynamic fee transaction to (\w+)$`, sendDynamicFeeTransaction)
	ctx.Step(`^I call eth_getTransactionByHash$`, getTransactionByHash)
	ctx.Step(`^I get all the details for my transaction$`, assertTransaction)
	ctx.Step(`^the transaction receipt should have a status of (\d)`, assertReceipt)

	ctx.Step(`^I bridge (\d+) IMX to (\w+) on L2$`, bridgeImxToL2)
	ctx.Step(`^(\w+) sends (\d+) IMX to (\w+) on L2 using EIP-1559$`, transferImxWithEip1559)
	ctx.Step(`^the transaction on chain should match what was submitted$`, assertTransaction)
}
