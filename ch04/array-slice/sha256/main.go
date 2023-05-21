package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	/*
		当调用一个函数时，函数的每个调用参数将会赋值给函数内部的参数变量，所以函数参数变量接收的是一个复制的副本
		函数参数传递大的数组类型将是低效的，并且不能直接修改调用时原始的数组变量。

		可以显式地传入一个数组指针，使得函数通过指针对数组的任何修改都可以直接反馈到调用者。
	*/

}

func zero(ptr *[32]byte) {
	for i := range ptr { // 会自动将 ptr 转换成 *ptr
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}
