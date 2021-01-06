package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	priKey, err := crypto.HexToECDSA("f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315")
	if err != nil {
		panic(err)
	}
	priKeyBytes := crypto.FromECDSA(priKey)
	fmt.Println(hex.EncodeToString(priKeyBytes))
	pubKey := priKey.PublicKey
	pubKeyBytes := crypto.FromECDSAPub(&pubKey)
	fmt.Println(hex.EncodeToString(pubKeyBytes[1:]))
	adrBytes := crypto.Keccak256(pubKeyBytes[1:])
	fmt.Println(hex.EncodeToString(adrBytes[len(adrBytes)-20:]))
	// 6e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0
}
