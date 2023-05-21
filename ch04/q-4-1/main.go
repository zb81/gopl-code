package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(DiffCount(c1, c2))
}

func DiffCount(c1, c2 [32]byte) int {
	count := 0
	for i, _ := range c1 {
		count += PopCount(c1[i], c2[i])
	}
	return count
}

// PopCount 计算每字节不同的 bit 数
func PopCount(x, y byte) int {
	count := 0
	for i := 0; i < 8; i++ {
		if x%2 != y%2 {
			count++
		}
		x >>= 1
		y >>= 1
	}
	return count
}
