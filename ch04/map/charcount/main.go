package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // Unicode 字符计数
	var utflen [utf8.UTFMax + 1]int // 统计不同字节长度字符的个数
	invalid := 0                    // 无效的 UTF8 编码字符个数

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		// EOF -> ctrl + D
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("charcount error: %v\n", err)
		}
		// 如果输入的是无效的 UTF-8 编码字符，返回 unicode.ReplacementChar 表示无效字符，并且编码长度 n 是 1
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
