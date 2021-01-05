package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// bytes, err := hex.DecodeString("6e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0")
	// if err != nil {
	// 	panic("")
	// }

	pubKeyHex := "f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315"
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		panic("")
	}
	for _, num := range pubKeyBytes {
		fmt.Printf("%02x ", num)
	}
	fmt.Println()
	fmt.Println(len(pubKeyBytes))
	fmt.Println(hex.EncodeToString(pubKeyBytes))
	// rst := crypto.Keccak256Hash(bytes)
	// fmt.Println(rst)
	sig := make([]byte, crypto.SignatureLength)
	fmt.Println(crypto.Ecrecover(pubKeyBytes, sig))
}
