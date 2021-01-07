package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	text := ReadWordListZh()
	fmt.Println(len(text))

	h := sha256.New()
	h.Write([]byte("hello world"))
	fmt.Printf("%x\n", h.Sum(nil))
	return

	// byte256 := GetRand256()
	// fmt.Println(hex.EncodeToString(byte256))

	// return
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
