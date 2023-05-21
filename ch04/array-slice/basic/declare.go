package main

import "fmt"

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func main() {
	//var q = [3]int{1, 2, 3}
	//var r = [...]int{1, 2, 3} // type of r: [3]int
	//
	//for i, v := range r {
	//	fmt.Printf("%d\t%d\n", i, v)
	//}
	//r[3] = 2
	//for i, v := range r {
	//	fmt.Printf("%d\t%d\n", i, v)
	//}

	/*
		数组的长度是数组类型的一个组成部分。数组的长度必须是常量表达式，数组长度需要在编译阶段确定。
		初始化索引的顺序无所谓
	*/
	//q := [3]int{1,2,3}
	//q = [4]int{1,2,3,4} // compile error

	/*
		初始化时，可以指定一个索引和对应值列表
	*/
	//symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	//fmt.Println(RMB, symbol[RMB])

	/*
		如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的
		只有当两个数组的所有元素都是相等时，数组才是相等
	*/
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // true false false
	//d := [3]int{1,2}
	//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}
