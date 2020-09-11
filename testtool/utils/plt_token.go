package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"math/big"

	"github.com/ethereum/go-ethereum/contracts/native/plt"
)

var (
	logger     = log.Logger("palette")
	PLTAddress = common.HexToAddress(native.PLTContractAddress)
)

func init() {
	plt.InitABI()
}

func (c *PaletteClient) PLTMint(amount *big.Int) (common.Hash, error) {
	adminAddr := c.AdminAddress()
	payload, err := utils.PackMethod(plt.ABI, plt.MethodMint, adminAddr, amount)
	if err != nil {
		return common.Hash{}, err
	}

	nonce, err := c.GetNonce(adminAddr.Hex())
	if err != nil {
		return common.Hash{}, err
	}
	tx := types.NewTransaction(
		nonce,
		PLTAddress,
		big.NewInt(0),
		GasNormal,
		big.NewInt(GasPrice),
		payload,
	)

	signedTx, err := c.SignTransaction(tx)
	if err != nil {
		return common.Hash{}, err
	}
	return c.SendRawTransaction(signedTx)
}

func (c *PaletteClient) PLTTotalSupply() (*big.Int, error) {
	payload, err := utils.PackMethod(plt.ABI, plt.MethodTotalSupply)
	if err != nil {
		return nil, err
	}

	raw, err := c.CallContract(c.AdminAddress(), PLTAddress, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to get total supply: [%v]", err)
	}

	supply := new(big.Int).SetBytes(raw)
	return supply, nil
}

func (c *PaletteClient) PLTDecimals() (uint64, error) {
	payload, err := utils.PackMethod(plt.ABI, plt.MethodDecimals)
	if err != nil {
		return 0, err
	}

	raw, err := c.CallContract(c.AdminAddress(), PLTAddress, payload)
	if err != nil {
		return 0, fmt.Errorf("failed to get decimal: [%v]", err)
	}

	decimal := new(big.Int).SetBytes(raw).Uint64()
	return decimal, nil
}
