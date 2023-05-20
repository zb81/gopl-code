package main

func main() {
	// 因为 Go 用 UTF8 编码，可以将 Unicode 码点写到字符串面值中

	/*
		可以通过十六进制或八进制转义在字符串面值中包含任意的字节
		十六进制：\xhh 包含两个十六禁止的 x，大小写都可以
		八进制：\ooo 包含三个八进制的 o 数字，不能超过 \377 （对应一个字节的范围，255）
	*/

	/*
		原生字符串面值：`...` ，使用反引号，没有转义操作；
	*/
}