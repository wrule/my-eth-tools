package main

import (
	"encoding/hex"
	"fmt"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

func main() {
	bytes, err := hex.DecodeString("6e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0")
	if err != nil {
		panic("")
	}
	rst := ethCrypto.Keccak256Hash(bytes)
	fmt.Println(rst)
}
