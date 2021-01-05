package main

import (
	"crypto/elliptic"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

func hexKey(prv string) *ecies.PrivateKey {
	key, err := crypto.HexToECDSA(prv)
	if err != nil {
		panic(err)
	}
	return ecies.ImportECDSA(key)
}

func main() {
	priKey := hexKey("f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315")
	pubkey := elliptic.Marshal(crypto.S256(), priKey.X, priKey.Y)
	fmt.Println(hex.EncodeToString(pubkey))
}
