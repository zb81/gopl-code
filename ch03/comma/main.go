package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1298387.1208"))
}

//func comma(s string) string {
//	n := len(s)
//	if n <= 3 {
//		return s
//	}
//	return comma(s[:n-3]) + "," + s[n-3:]
//}

//func comma(s string) string {
//	var buf bytes.Buffer
//	for i := 0; i < len(s); i++ {
//		if i > 0 && (len(s)-i)%3 == 0 {
//			buf.WriteByte(',')
//		}
//		buf.WriteByte(s[i])
//	}
//	return buf.String()
//}

func comma(s string) string {
	var buf bytes.Buffer

	// 如果以 +/- 号开头，设置 start 为 1，将 +/- 号写入 buf
	var start int
	if s[0] == '+' || s[0] == '-' {
		start = 1
		buf.WriteByte(s[0])
	} else {
		start = 0
	}

	// 获取小数点的索引，根据 start 和 end 获取整数部分
	end := strings.Index(s, ".")
	ints := s[start:end]

	// 对整数部分进行 comma 操作
	for i := 0; i < len(ints); i++ {
		if i > 0 && (len(ints)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(ints[i])
	}

	// 拼接小数点及之后的部分
	rest := s[end:]
	for i := 0; i < len(rest); i++ {
		buf.WriteByte(rest[i])
	}

	return buf.String()
}
