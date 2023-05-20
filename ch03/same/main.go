package main

import "fmt"

func main() {
	fmt.Println(isShuffledEachOther("你好，世界", "世界，你好")) // expect `true`
}

func isShuffledEachOther(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	return true
}
