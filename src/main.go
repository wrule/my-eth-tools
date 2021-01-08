package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	bytes128, err := hex.DecodeString("0c1e24e5917779d297e14d45f14e1a1a")
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(bytes128))
	words := BIP39GetWords(bytes128)
	fmt.Println(words, len(words))
	seed := BIP39GetSeed("army van defense carry jealous true garbage claim echo media make crunch", "")
	fmt.Println(hex.EncodeToString(seed))
	return
	for {
		priKeyHash := NewPriKey().Hex()
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
		addrHex := addr.Hex()
		fmt.Printf("地址为: %s\n", addrHex)
		if strings.HasPrefix(addrHex, "0x123") {
			break
		}
	}
}
