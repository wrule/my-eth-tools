package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/dev/random")
	if err != nil {
		panic("打开文件出错")
	}
	buf := make([]byte, 32)
	n, err := file.Read(buf)
	if err != nil {
		panic("读取文件出错")
	}
	fmt.Println(n)
	for _, num := range buf {
		fmt.Printf("%02x ", num)
	}
	fmt.Println()
	file.Close()
	fmt.Println("你好，世界")
}
