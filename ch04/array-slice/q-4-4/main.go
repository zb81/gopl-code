package main

import "fmt"

/*
编写一个rotate函数，通过一次循环完成旋转。
*/

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(x, 5)
	fmt.Println(x)
}

func rotate(s []int, n int) {
	n = n % len(s)
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
