package main

import "fmt"

func main() {
	//var q = [3]int{1, 2, 3}
	var r = [...]int{1, 2, 3}
	for i, v := range r {
		fmt.Printf("%d\t%d\n", i, v)
	}
	//r[3] = 2
	//for i, v := range r {
	//	fmt.Printf("%d\t%d\n", i, v)
	//}
}
