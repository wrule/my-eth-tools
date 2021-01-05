package main

import "fmt"

func main() {
	key := NewPriKey()
	fmt.Println(key.Hex())
}
