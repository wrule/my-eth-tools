package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"io/ioutil"
	"strings"

	"golang.org/x/crypto/pbkdf2"
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
	hashBits := BytesToBits(hashBytes)
	bits = append(bits, hashBits[:checkLen]...)
	return bits
}

// BIP39GetWords s
func BIP39GetWords(bytes []byte) []string {
	bits := BIP39AddCheckInfo(bytes)
	segNum := len(bits) / 11
	modNum := len(bits) % 11
	if modNum != 0 {
		panic("位数不是11的倍数")
	}
	rst := []string{}
	words := ReadWordListEn()
	for i := 0; i < segNum; i++ {
		start := i * 11
		bitSeg := bits[start : start+11]
		num := 0
		num += bitSeg[0] * 1024
		num += bitSeg[1] * 512
		num += bitSeg[2] * 256
		num += bitSeg[3] * 128
		num += bitSeg[4] * 64
		num += bitSeg[5] * 32
		num += bitSeg[6] * 16
		num += bitSeg[7] * 8
		num += bitSeg[8] * 4
		num += bitSeg[9] * 2
		num += bitSeg[10] * 1
		rst = append(rst, words[num])
	}
	return rst
}

// BIP39GetSeed s
func BIP39GetSeed(words string, pwd string) []byte {
	return pbkdf2.Key([]byte(words), []byte("mnemonic"+pwd), 2048, 64, sha512.New)
}
