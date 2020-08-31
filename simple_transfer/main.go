package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ipfs/go-log"
)

const (
	NetworkID = 10
	GasLimit  = 21000
	GasPrice  = 0
)

var (
	logger  = log.Logger("geth")
	Eth1, _ = new(big.Int).SetString("1000000000000000000", 10)

	config = &Config{Alloc: new(AllocAccount)}
)

type Config struct {
	Rpc         string
	Alloc       *AllocAccount
	TestAddress string
}

type AllocAccount struct {
	KeyFile    string
	Address    string
	Passphrase string
}

func init() {
	if err := log.SetLogLevel("*", "DEBUG"); err != nil {
		panic(fmt.Sprintf("failed to initialize logger: [%v]", err))
	}

	if _, err := toml.DecodeFile("./config.toml", config); err != nil {
		panic(fmt.Sprintf("failed to decode config file: [%v]", err))
	}
}

func main() {

	src := config.Alloc.Address
	dst := config.TestAddress
	amount := Eth1

	// create wrap rpc client
	client := getClient(config.Rpc, config.Alloc.KeyFile, config.Alloc.Passphrase)

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
	intAmount := new(big.Int).Div(amount, Eth1)

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
	data := new(big.Int).Div(src, Eth1)

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
		logger.Fatalf("failed to dial geth rpc: [%v]", err)
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
