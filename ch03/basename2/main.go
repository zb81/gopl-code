package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c/hello.go"))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // 如果没有 "/"，返回 -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
