package main

import "fmt"

/*
UTF8 是一个将 Unicode 码点编码为字节序列的变长编码。
使用 1 到 4 个字节来表示每个 Unicode 码点，ASCII 部分字符只使用 1 个字节，常用字符部分使用 2 或 3 个字节

每个符号编码后第一个字节的高端 bit 位用于表示编码总共有多少个子节：
	如果第一个字节的高端 bit 为 0，则表示对应 7bit 的 ASCII 字符
*/

func main() {
	//c := '世'
	//c = '\U00003456'
	//fmt.Println("\xe4\xb8\x96\xe7\x95\x8c")
	//
	//s := "Hello, world!"
	//fmt.Println(HasPrefix(s, "He"))
	//fmt.Println(HasSuffix(s, "d!"))
	//
	//fmt.Println(Contains(s, "o, "))

	//s := "Hello, 世界"
	//fmt.Println(len(s))
	//fmt.Println(utf8.RuneCountInString(s))

	//for i := 0; i < len(s); {
	//	r, size := utf8.DecodeRuneInString(s[i:])
	//	fmt.Printf("%d\t%c\n", i, r)
	//	i += size
	//}

	// 自动隐式解码 UTF8 字符
	//for i, r := range s {
	//	fmt.Printf("%d\t%q\t%d\n", i, r, r)
	//}

	//n := 0
	//for range s {
	//	n++
	//}
	//fmt.Printf("'s' has %d runes\n", n)

	// "program" in Japanese katakana
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r)        // "[30d7 30ed 30b0 30e9 30e0]"
	fmt.Println(string(r))       // "プログラム"
	fmt.Println(string(65))      // "A", not "65"
	fmt.Println(string(0x4eac))  // "京"
	fmt.Println(string(1234567)) // "?"
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[(len(s)-len(suffix)):] == suffix
}
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
