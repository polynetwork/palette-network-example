package utils

import (
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/rpc"
)

type PaletteClient struct {
	*rpc.Client
	key *keystore.Key
}

func NewPaletteClient(
	url string,
	keyFile string,
	passphrase string,
) *PaletteClient {
	client, err := rpc.Dial(url)
	if err != nil {
		panic(fmt.Errorf("failed to dial geth rpc: [%v]", err))
	}

	keyJson, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic(fmt.Errorf("failed to read file: [%v]", err))
	}

	key, err := keystore.DecryptKey(keyJson, passphrase)
	if err != nil {
		panic(fmt.Errorf("failed to decrypt keyjson: [%v]", err))
	}

	return &PaletteClient{
		Client: client,
		key:    key,
	}
}
