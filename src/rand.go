package main

import (
	"crypto/rand"
)

// GetRandN 获取随机字节切片
func GetRandN(size int) []byte {
	rst := make([]byte, size)
	n, err := rand.Read(rst)
	if err != nil || n != size {
		panic(err)
	}
	return rst
}

// GetRand128 s
func GetRand128() []byte {
	return GetRandN(128 / 8)
}

// GetRand256 s
func GetRand256() []byte {
	return GetRandN(256 / 8)
}

// GetRand512 s
func GetRand512() []byte {
	return GetRandN(512 / 8)
}
