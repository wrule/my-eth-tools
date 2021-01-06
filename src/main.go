package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	priKeyHash := "796c823671b118258b53ef6056fd1f9fc96d125600f348f75f397b2000267fe8"
	// 创建私钥对象
	priKey, err := crypto.HexToECDSA(priKeyHash)
	if err != nil {
		panic(err)
	}
	priKeyBytes := crypto.FromECDSA(priKey)
	fmt.Printf("私钥为: %s\n", hex.EncodeToString(priKeyBytes))

	pubKey := priKey.Public().(*ecdsa.PublicKey)
	// 获取公钥并去除头部0x04
	pubKeyBytes := crypto.FromECDSAPub(pubKey)[1:]
	fmt.Printf("公钥为: %s\n", hex.EncodeToString(pubKeyBytes))

	// 获取地址
	addr := crypto.PubkeyToAddress(*pubKey)
	fmt.Printf("地址为: %s\n", addr.Hex())
}
