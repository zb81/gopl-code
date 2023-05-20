package main

import (
	"fmt"
	"time"
)

// 0-255 每个数值包含的 1 的位数（八位数值）
var pc = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1)
	}
	return
}()

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] + // 最低 8 位
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))]) // 最高 8 位
}

func main() {
	arr := make([]int, 10000000)
	for i := range arr {
		arr[i] = i + 1
	}
	start := time.Now()
	for _, v := range arr {
		PopCount(uint64(v))
	}
	fmt.Printf("pc cost: %d ms\n", time.Since(start).Milliseconds()) // 1ms

	start = time.Now()
	for _, v := range arr {
		PopCount2(uint64(v))
	}
	fmt.Printf("loop cost: %d ms\n", time.Since(start).Milliseconds()) // 48ms

}

func PopCount2(x uint64) int {
	count := 0
	for x != 0 {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}
