package main

import (
	"fmt"

	ecies "github.com/wrule/go-1"
)

func main() {
	priKey, err := ecies.NewPrivateKeyFromHex("f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315")
	if err != nil {
		panic("私钥生成失败")
	}
	fmt.Println(priKey.Hex())
	fmt.Println(priKey.PublicKey.Hex(false))
}

// 6e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0
// 046e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0
