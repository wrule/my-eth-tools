package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

// ReadWordList 读取助记词词典列表
func ReadWordList(filePath string) []string {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	rst := []string{}
	for _, item := range strings.Split(string(bytes), "\n") {
		itemTrimed := strings.TrimSpace(item)
		if len(itemTrimed) > 0 {
			rst = append(rst, itemTrimed)
		}
	}
	if len(rst) != 2048 {
		panic("词典内单词数量不为2048")
	}
	return rst
}

// ReadWordListEn 读取英文助记词词典列表
func ReadWordListEn() []string {
	return ReadWordList("bip39wordlist-en.txt")
}

// ReadWordListZh 读取中文助记词词典列表
func ReadWordListZh() []string {
	return ReadWordList("bip39wordlist-zh.txt")
}

// BIP39AddCheckInfo s
func BIP39AddCheckInfo(bytes []byte) []int {
	bits := BytesToBits(bytes)
	checkLen := len(bits) / 32
	h := sha256.New()
	h.Write(bytes)
	hashBytes := h.Sum(nil)
	fmt.Println(hex.EncodeToString(hashBytes))
	hashBits := BytesToBits(hashBytes)
	bits = append(bits, hashBits[:checkLen]...)
	return bits
}
