package main

import (
	"crypto/sha256"
	"fmt"
)

//escreva uma função que conte o numero de de bits
//diferentes em dois hashes SHA256
//  Veja PopCount na seção 2.6.2

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func CountDiffBits(h1, h2 [sha256.Size]byte) int {
	count := 0

	for i, byteH1 := range h1 {
		byteH2 := h2[i]
		count += int(pc[byteH1^byteH2])
	}

	return count
}

func main() {
	h1 := sha256.Sum256([]byte("hello"))
	h2 := sha256.Sum256([]byte("world"))
	fmt.Println(CountDiffBits(h1, h2))
}
