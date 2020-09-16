package main

import (
	"math/big"
	"time"

	"github.com/ipfs/go-log"
	"github.com/palettechain/testtool/config"
	"github.com/palettechain/testtool/utils"
)

var logger = log.Logger("geth")

const cfgPath = "/Users/dylen/software/palette/testtool/config.toml"

func main() {

	cfg := config.GenerateConfig(cfgPath)

	// create wrap rpc client
	client := utils.NewPaletteClient(cfg)

	src := client.Admin.Address.Hex()
	dst := client.TestAccounts[0].Address.Hex()
	amount := utils.OneEth

	// get balance before transfer
	srcBalanceBeforeTransfer, err := client.Balance(src)
	if err != nil {
		logger.Fatal(err)
	}
	dstBalanceBeforeTransfer, err := client.Balance(dst)
	if err != nil {
		logger.Fatal(err)
	}

	// get account nonce
	nonce, err := client.GetNonce(src)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("src nonce %d", nonce)

	// build and send transfer transaction
	tx := utils.TransferETH(nonce, dst, amount)
	signedTx, err := client.SignTransaction(client.Admin.PrivateKey, tx)
	if err != nil {
		logger.Fatal(err)
	}
	hash, err := client.SendRawTransaction(signedTx)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("transfer hash %s", hash.Hex())

	// waiting for block sealing
	time.Sleep(7 * time.Second)

	// check balance
	srcBalanceAfterTransfer, err := client.Balance(src)
	if err != nil {
		logger.Fatal(err)
	}
	dstBalanceAfterTransfer, err := client.Balance(dst)
	if err != nil {
		logger.Fatal(err)
	}
	srcBalanceDec := new(big.Int).Sub(srcBalanceBeforeTransfer, srcBalanceAfterTransfer)
	dstBalanceAdd := new(big.Int).Sub(dstBalanceAfterTransfer, dstBalanceBeforeTransfer)
	intAmount := new(big.Int).Div(amount, utils.OneEth)

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
