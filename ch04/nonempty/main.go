package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)            // `["one" "three" "three"]`

}

// 在原有内存空间修改元素
// 输入的 slice 和输出的 slice 共享一个底层数组，原来的数据将可能会被覆盖
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
