package utils

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func (c *PaletteClient) Balance(address string) (*big.Int, error) {
	var raw string

	if err := c.Call(
		&raw,
		"eth_getBalance",
		address,
		"latest",
	); err != nil {
		return nil, fmt.Errorf("faild to get balance [%v]", err)
	}

	src, _ := new(big.Int).SetString(raw, 0)
	data := new(big.Int).Div(src, OneEth)

	return data, nil
}

func (c *PaletteClient) GetNonce(address string) (uint64, error) {
	var raw string

	if err := c.Call(
		&raw,
		"eth_getTransactionCount",
		address,
		"latest",
	); err != nil {
		return 0, fmt.Errorf("failed to get nonce: [%v]", err)
	}

	without0xStr := strings.Replace(raw, "0x", "", -1)
	bigNonce, _ := new(big.Int).SetString(without0xStr, 16)
	return bigNonce.Uint64(), nil
}

func (c *PaletteClient) SignTransaction(tx *types.Transaction) (string, error) {

	signer := types.HomesteadSigner{}
	signedTx, err := types.SignTx(
		tx,
		signer,
		c.key.PrivateKey,
	)
	if err != nil {
		return "", fmt.Errorf("failed to sign tx: [%v]", err)
	}

	bz, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to rlp encode bytes: [%v]", err)
	}
	return "0x" + hex.EncodeToString(bz), nil
}

func (c *PaletteClient) SendRawTransaction(signedTx string) (common.Hash, error) {
	var result common.Hash
	if err := c.Client.Call(&result, "eth_sendRawTransaction", signedTx); err != nil {
		return result, fmt.Errorf("failed to send raw transaction: [%v]", err)
	}

	return result, nil
}

func TransferETH(nonce uint64, toAddress string, value *big.Int) *types.Transaction {
	return types.NewTransaction(
		nonce,
		common.HexToAddress(toAddress),
		value,
		GasLimit,
		big.NewInt(GasPrice),
		nil,
	)
}
