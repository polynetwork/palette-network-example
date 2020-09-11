package utils

import "math/big"

const (
	NetworkID = 10
	GasLimit  = 21000
	GasPrice  = 0
)

var OneEth, _ = new(big.Int).SetString("1000000000000000000", 10)