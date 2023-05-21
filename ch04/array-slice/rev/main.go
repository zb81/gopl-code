package main

import "fmt"

func main() {
	//a := [...]int{0, 1, 2, 3, 4, 5}
	//reverse(a[:])
	//fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	// 将 s 向左旋转两位
	/*
		1. 先将前两位反转		-> [1, 0, 2, 3, 4, 5]
		2. 再将后四位反转		-> [1, 0, 5, 4, 3, 2]
		3. 最后将整个切片反转	-> [2, 3, 4, 5, 0, 1]
	*/
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)
}

func reverse(s []int) {
	// 首尾互换
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
