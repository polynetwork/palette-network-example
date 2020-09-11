package main

import (
	"math/big"
	"time"

	"github.com/ipfs/go-log"
	"github.com/palettechain/testtool/config"
	"github.com/palettechain/testtool/utils"
)

const cfgPath = "/Users/dylen/software/palette/testtool/config.toml"

var (
	logger = log.Logger("geth")
	cfg    = config.GenerateConfig(cfgPath)
	c      = utils.NewPaletteClient(cfg.Rpc, cfg.Admin.KeyFile, cfg.Admin.Passphrase)
)

func main() {
	testMint()
	// testTotalSupply(c)
	// testDecimals()
}

func testMint() {
	var amount int64 = 10000

	data := new(big.Int).Mul(big.NewInt(amount), utils.OnePLT)
	hash, err := c.PLTMint(data)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("mint hash %s", hash.Hex())

	time.Sleep(7 * time.Second)

	if err := c.DumpEventLog(hash); err != nil {
		logger.Fatal(err)
	}
}

func testTotalSupply() {
	amount, err := c.PLTTotalSupply()
	if err != nil {
		logger.Fatal(err)
	}

	supply := new(big.Int).Div(amount, utils.OnePLT)
	logger.Infof("PLT total supply is %s", supply.String())
}

func testDecimals() {
	decimals, err := c.PLTDecimals()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("PLT decimals %d", decimals)
}
