package evmbehaviour

import (
	"context"
	"crypto/ecdsa"
	"evm-behaviour/bridge"
	"evm-behaviour/child"
	"evm-behaviour/consensus"
	"evm-behaviour/mock"
	"evm-behaviour/native"
	"evm-behaviour/root"
	"flag"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/ecodia/golang-awaitility/awaitility"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Proof struct {
	Data     [][32]byte // the proof himself
	Metadata map[string]interface{}
}

type Wallet struct {
	Address    common.Address
	Balance    *big.Int
	PrivateKey *ecdsa.PrivateKey
}

type L2StateSyncedEvent struct {
	ID       *big.Int
	Sender   common.Address
	Receiver common.Address
	Data     []byte
}

var childClient *ethclient.Client
var rootClient *ethclient.Client

var l1Erc20Address common.Address
var l1Erc20Contract *mock.MockERC20

var l1Erc20PredicateAddress common.Address
var l1Erc20PredicateContract *root.RootERC20Predicate

var l1CheckpointManagerAddress common.Address
var l1CheckpointManagerContract *consensus.CheckpointManager

var l1ExitHelperAddress common.Address
var l1ExitHelperContract *bridge.ExitHelper

var l2Erc20Address common.Address
var l2Erc20Contract *native.NativeERC20

var l2Erc20PredicateAddress common.Address
var l2Erc20PredicateContract *child.ChildERC20Predicate

var receiver *Wallet
var sender *Wallet

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
		sender = nil
		receiver = nil

		var err error

		rootClient, err = ethclient.Dial("http://localhost:8545")
		if err != nil {
			panic(err)
		}

		childClient, err = ethclient.Dial("http://localhost:40002")
		if err != nil {
			panic(err)
		}

		l1Erc20Address = common.HexToAddress("0x947a581B2713F58A8145201DA41BCb6aAE90196B")
		l1Erc20Contract, err = mock.NewMockERC20(l1Erc20Address, rootClient)
		if err != nil {
			panic(err)
		}

		l1Erc20PredicateAddress = common.HexToAddress("0xA1DFe8536732EB98BBCA36A7f97C72e3395EaB8E")
		l1Erc20PredicateContract, err = root.NewRootERC20Predicate(l1Erc20PredicateAddress, rootClient)
		if err != nil {
			panic(err)
		}

		l1ExitHelperAddress = common.HexToAddress("0xaCB3Eb2f3c167B56410F0351B6C6EBac9256f553")
		l1ExitHelperContract, err = bridge.NewExitHelper(l1CheckpointManagerAddress, rootClient)
		if err != nil {
			panic(err)
		}

		l1CheckpointManagerAddress = common.HexToAddress("0x22C246401ed6e52C525644659C5304aed63516C7")
		l1CheckpointManagerContract, err = consensus.NewCheckpointManager(l1CheckpointManagerAddress, rootClient)
		if err != nil {
			panic(err)
		}

		l2Erc20Address = common.HexToAddress("0x0000000000000000000000000000000000001010")
		l2Erc20Contract, err = native.NewNativeERC20(l2Erc20Address, childClient)
		if err != nil {
			panic(err)
		}

		l2Erc20PredicateAddress = common.HexToAddress("0x0000000000000000000000000000000000001003")
		l2Erc20PredicateContract, err = child.NewChildERC20Predicate(l2Erc20PredicateAddress, childClient)
		if err != nil {
			panic(err)
		}
	})
}

func bridgeImxToL2(amount int64) error {
	ctx := context.TODO()

	txOpts, err := newTransactionOptions(ctx, rootClient)
	if err != nil {
		return err
	}

	if _, err = l1Erc20Contract.Approve(txOpts, l1Erc20PredicateAddress, big.NewInt(amount)); err != nil {
		return err
	}

	txOpts, err = newTransactionOptions(ctx, rootClient)
	if err != nil {
		return err
	}

	if _, err = l1Erc20PredicateContract.DepositTo(txOpts, l1Erc20Address, sender.Address, big.NewInt(amount)); err != nil {
		return err
	}

	awaitility.Await(15*time.Second, 10*time.Minute, func() bool {
		println("Waiting for bridge deposit...")

		newBalance, err := l2Erc20Contract.BalanceOf(&bind.CallOpts{Context: ctx}, sender.Address)

		return err == nil && newBalance != nil && newBalance.Cmp(sender.Balance) > 0
	})

	return nil
}

func transferImx(
	amount int64,
) error {
	ctx := context.TODO()

	txOpts, err := newTransactionOptions(ctx, childClient)
	if err != nil {
		return err
	}

	tx, err := l2Erc20Contract.Transfer(txOpts, receiver.Address, big.NewInt(amount))
	if err != nil {
		return err
	}

	awaitility.Await(1*time.Second, 30*time.Second, func() bool {
		println("Waiting for transfer...")

		receipt, err := childClient.TransactionReceipt(ctx, tx.Hash())

		return err == nil && receipt != nil && receipt.Status == 1
	})

	return nil
}

func newTransactionOptions(
	ctx context.Context,
	client *ethclient.Client,
) (*bind.TransactOpts, error) {
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(ctx, sender.Address)
	if err != nil {
		return nil, err
	}

	latestBlock, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	feeHistory, err := client.FeeHistory(ctx, 1, new(big.Int).SetUint64(latestBlock), []float64{95, 99})
	if err != nil {
		return nil, err
	}

	return &bind.TransactOpts{
		Nonce: new(big.Int).SetUint64(nonce),
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(t, types.NewLondonSigner(chainID), sender.PrivateKey)
		},
		From:      sender.Address,
		GasTipCap: new(big.Int).Div(feeHistory.BaseFee[0], big.NewInt(2)),
		GasFeeCap: feeHistory.BaseFee[0],
	}, nil
}

func loadL2Receiver(receiverHex string) error {
	ctx := context.TODO()

	receiverAddress := common.HexToAddress(receiverHex)

	balance, err := l2Balance(ctx, receiverAddress)
	if err != nil {
		return err
	}

	receiver = &Wallet{
		Address: receiverAddress,
		Balance: balance,
	}

	return nil
}

func assertL2ReceiverBalance(amount int64) error {
	ctx := context.TODO()

	newBalance, err := l2Balance(ctx, receiver.Address)
	if err != nil {
		return err
	}

	expected := receiver.Balance.Int64() + amount

	actual := newBalance.Int64()

	if actual != expected {
		return fmt.Errorf(
			"Expected %s to have a balance of %d, instead found %d",
			receiver.Address.Hex(),
			expected,
			actual,
		)
	}

	return nil
}

func l1Balance(ctx context.Context, address common.Address) (*big.Int, error) {
	balance, err := l1Erc20Contract.BalanceOf(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func l2Balance(ctx context.Context, address common.Address) (*big.Int, error) {
	balance, err := l2Erc20Contract.BalanceOf(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func withdrawImx(amount int64) error {
	ctx := context.TODO()

	txOpts, err := newTransactionOptions(ctx, childClient)
	if err != nil {
		return err
	}

	tx, err := l2Erc20PredicateContract.WithdrawTo(txOpts, l2Erc20Address, sender.Address, big.NewInt(amount))
	if err != nil {
		return err
	}

	var receipt *types.Receipt
	awaitility.Await(1*time.Second, 30*time.Second, func() bool {
		println("Waiting for withdraw...")

		receipt, err = childClient.TransactionReceipt(ctx, tx.Hash())

		return err == nil && receipt != nil && receipt.Status == 1
	})

	/* TODO:
	 * 1. Wait for checkpoint
	 * 2. Generate exit proof
	 * 3. Call exit
	 * 4. Check balance of sender in L1
	 */

	return nil
}

func assertL1SenderBalance(amount int64) error {
	ctx := context.TODO()

	balance, err := l1Balance(ctx, sender.Address)
	if err != nil {
		return nil
	}

	expected := balance.Int64() + amount
	awaitility.Await(10*time.Second, 15*time.Minute, func() bool {
		println("Waiting for withdraw...")

		newBalance, err := l1Balance(ctx, sender.Address)

		return err == nil && newBalance != nil && newBalance.Int64() == expected
	})

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// clean the state before every scenario
		privateKey, err := crypto.HexToECDSA("6362c4eba37ef8268804deca9f3b89ad28479f802181a15f6ce6a128fba3f3d6")
		if err != nil {
			return ctx, err
		}

		publicKey := privateKey.Public().(*ecdsa.PublicKey)
		address := crypto.PubkeyToAddress(*publicKey)

		balance, err := l2Balance(ctx, address)
		if err != nil {
			return ctx, err
		}

		sender = &Wallet{
			Address:    address,
			Balance:    balance,
			PrivateKey: privateKey,
		}

		receiver = nil

		return ctx, nil
	})

	ctx.Step(`^receiver (\w+) has a certain balance on L2$`, loadL2Receiver)
	ctx.Step(`^sender bridges (\d+) IMX from L1 to receiver on L2$`, bridgeImxToL2)
	ctx.Step(`^sender transfers (\d+) IMX to receiver on L2$`, transferImx)
	ctx.Step(`^sender withdraws (\d+) IMX to L1$`, withdrawImx)
	ctx.Step(`^the receiver balance on L2 should increase by (\d+)$`, assertL2ReceiverBalance)
	ctx.Step(`^sender IMX balance on L1 should increase by (\d+)$`, assertL1SenderBalance)
}
