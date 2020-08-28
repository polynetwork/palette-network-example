package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ipfs/go-log"
)

var (
	logger = log.Logger("geth")
	One, _ = new(big.Int).SetString("1000000000000000000", 10)
)

const (
	NetworkID = 10
	GasLimit  = 21000
	GasPrice  = 0

	RpcURL       = "http://localhost:22000"
	Keystore     = `/Users/dylen/software/palette/node0/data/keystore/`
	KeyFile      = Keystore + `UTC--2020-08-26T09-55-25.770341000Z--57a259e0bcd61dffdd205a5cd046be9068e832dd`
	Passphrase   = `111111`
	AdminAddress = `0x57A259e0BcD61dFfdd205a5Cd046be9068E832dd`
	TestAddress  = `0x8409f65FD78a03edd654671a9d15c6E9962C07c9`
)

func init() {
	if err := log.SetLogLevel("*", "DEBUG"); err != nil {
		panic(fmt.Sprintf("failed to initialize logger: [%v]", err))
	}
}

func main() {

	src := AdminAddress
	dst := TestAddress
	amount := One

	// create wrap rpc client
	client := getClient(RpcURL, KeyFile, Passphrase)

	// get balance before transfer
	srcBalanceBeforeTransfer := client.Balance(src)
	dstBalanceBeforeTransfer := client.Balance(dst)

	// get account nonce
	nonce := client.GetNonce(src)
	logger.Infof("src nonce %d", nonce)

	// build and send transfer transaction
	tx := transferETH(nonce, dst, amount)
	signedTx := client.SignTransaction(tx)
	hash := client.SendRawTransaction(signedTx)
	logger.Infof("transfer hash %s", hash.Hex())

	// waiting for block sealing
	time.Sleep(10 * time.Second)

	// check balance
	srcBalanceAfterTransfer := client.Balance(src)
	dstBalanceAfterTransfer := client.Balance(dst)
	srcBalanceDec := new(big.Int).Sub(srcBalanceBeforeTransfer, srcBalanceAfterTransfer)
	dstBalanceAdd := new(big.Int).Sub(dstBalanceAfterTransfer, dstBalanceBeforeTransfer)
	intAmount := new(big.Int).Div(amount, One)

	if srcBalanceDec.Cmp(intAmount) != 0 || dstBalanceAdd.Cmp(intAmount) != 0 {
		logger.Fatalf("transfer amount error, "+
			"src balance before transfer %s, src balance after transfer %s,"+
			"dst balance before transfer %s, dst balance after transfer %s",
			srcBalanceBeforeTransfer, srcBalanceAfterTransfer,
			dstBalanceBeforeTransfer, dstBalanceAfterTransfer,
		)
	}

	logger.Infof("transfer success!")
}

type wrapClient struct {
	*rpc.Client
	key *keystore.Key
}

func (c *wrapClient) Balance(address string) *big.Int {
	var raw string

	if err := c.Call(
		&raw,
		"eth_getBalance",
		address,
		"latest",
	); err != nil {
		logger.Fatal(err)
	}

	src, _ := new(big.Int).SetString(raw, 0)
	data := new(big.Int).Div(src, One)

	return data
}

func (c *wrapClient) GetNonce(address string) uint64 {
	var raw string

	if err := c.Call(
		&raw,
		"eth_getTransactionCount",
		address,
		"latest",
	); err != nil {
		logger.Fatal(err)
	}

	without0xStr := strings.Replace(raw, "0x", "", -1)
	bigNonce, _ := new(big.Int).SetString(without0xStr, 16)
	return bigNonce.Uint64()
}

func (c *wrapClient) SignTransaction(tx *types.Transaction) string {

	signer := types.HomesteadSigner{}
	signedTx, err := types.SignTx(
		tx,
		signer,
		c.key.PrivateKey,
	)
	if err != nil {
		logger.Fatalf("failed to sign tx: [%v]", err)
	}

	bz, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		logger.Fatalf("failed to rlp encode bytes: [%v]", err)
	}
	return "0x" + hex.EncodeToString(bz)
}

func (c *wrapClient) SendRawTransaction(signedTx string) common.Hash {
	var result common.Hash
	if err := c.Client.Call(&result, "eth_sendRawTransaction", signedTx); err != nil {
		logger.Fatalf("failed to send raw transaction: [%v]", err)
	}

	return result
}

func transferETH(nonce uint64, toAddress string, value *big.Int) *types.Transaction {
	return types.NewTransaction(
		nonce,
		common.HexToAddress(toAddress),
		value,
		GasLimit,
		big.NewInt(GasPrice),
		nil,
	)
}

func getClient(
	url string,
	keyFile string,
	passphrase string,
) *wrapClient {
	client, err := rpc.Dial(url)
	if err != nil {
		panic(err)
	}

	keyJson, err := ioutil.ReadFile(keyFile)
	if err != nil {
		logger.Fatalf("failed to read file: [%v]", err)
	}

	key, err := keystore.DecryptKey(keyJson, passphrase)
	if err != nil {
		logger.Fatalf("failed to decrypt keyjson: [%v]", err)
	}

	return &wrapClient{
		Client: client,
		key:    key,
	}
}
