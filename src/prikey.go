package main

import (
	"fmt"
)

// PriKey 私钥
type PriKey [32]byte

// Bytes 获取字节切片
func (me PriKey) Bytes() []byte {
	return me[:]
}

// Hex 获取十六进制私钥字符串
func (me PriKey) Hex() string {
	rst := ""
	for _, num := range me.Bytes() {
		rst += fmt.Sprintf("%02x", num)
	}
	return rst
}

// NewPriKey 构造函数
func NewPriKey() PriKey {
	rst := [32]byte{}
	copy(rst[:], GetRand256())
	return PriKey(rst)
}
