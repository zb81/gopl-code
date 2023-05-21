package main

import (
	"fmt"
	"sort"
)

/*
一个 map 就是一个哈希表的引用。
map 中所有的 key 都有相同的类型，所有的 value 也有相同的类型。
key 必须是支持 == 比较运算符的数据类型。
*/

func main() {
	// 可以用 make 函数创建一个 map
	//ages := make(map[string]int)

	// 也可以用 map 字面值创建：
	ages := map[string]int{
		"alice":   30,
		"charlie": 28,
	}

	// 通过 key 对应的下标语法访问元素
	ages["alice"] = 32
	fmt.Println(ages["alice"])

	// 使用 delete 函数删除元素
	delete(ages, "alice")

	// 所有这些操作是安全的，即使这些元素不在 map 中也没关系；
	// 如果查找失败，将返回 value 类型对应的零值
	ages["bob"] = ages["bob"] + 1
	// 可以简写：
	ages["bob"]++

	// map 中的元素不是一个变量，不能对 map 的元素进行取址操作：
	// map 可能随元素数量增长而分配更大的内存空间，可能导致之前的地址无效
	//_ = &ages["bob"] // Cannot take the address of 'ages["bob"]'

	// 可以使用 range 风格的 for 循环实现遍历：
	// 迭代顺序是不确定的
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	// 如果需要按照顺序，必须显式地对 key 排序：sort 包的 Strings 函数
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	// 向 map 存数据前必须先创建 map

	// 判断元素是否真的存在：
	if _, ok := ages["zb"]; !ok {
		fmt.Println("zb is not exist")
	}

}

// 必须通过一个循环来判断两个 map 是否包含相同的 key 和 value
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	// y 中不存在 k，或两个 map 中 k 对应的 value 不相等
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
