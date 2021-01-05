package main

import (
	"fmt"
	"os"
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
	const size = 32
	file, err := os.Open("/dev/random")
	if err != nil {
		panic("随机文件打开出错")
	}
	buf := make([]byte, size)
	n, err := file.ReadAt(buf, size)
	if err != nil || n != size {
		panic("随机文件读取出错")
	}
	file.Close()
	rst := [size]byte{}
	copy(rst[:], buf)
	return PriKey(rst)
}
