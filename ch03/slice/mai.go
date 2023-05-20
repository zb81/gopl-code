package main

import "fmt"

/*
一个字符串是包含只读字节的数组，一旦创建，是不可变的。

一个字节 slice 的元素则可以自由地修改。

字符串和字节 slice 之间可以相互转换
*/
func main() {
	s := "abc，世界"
	b := []byte(s)
	fmt.Println(b)
	s2 := string(b)
	fmt.Println(s2)
}
