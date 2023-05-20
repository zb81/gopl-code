package main

import "fmt"

func main() {
	fmt.Println(basename("/a/b/a.b.c.go"))
}

// basename 获取一个路径中的文件名。
// 例：/a/b/a.b.c.go -> a.b.c
func basename(s string) string {
	// 丢弃最后一个 '/' 及其之前的内容
	// /a/b/a.b.c.go -> a.b.c.go
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 保留最后一个 '.' 之前的内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
