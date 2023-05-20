package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(intToString(100) == "100")

	fmt.Println(strconv.ParseInt("11", 10, 64))
}

/*
将一个整数转为字符串：
 1. fmt.Sprintf("%d", x)
 2. strconv.Itoa()
*/
func intToString(n int) string {
	//return fmt.Sprintf("%d", n)
	return strconv.Itoa(n)
}
