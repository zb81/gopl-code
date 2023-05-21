package main

import "fmt"

const l = 6

func main() {
	arr := [l]int{1, 2, 3, 4, 5, 6}
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(s *[l]int) {
	for i := 0; i < l/2; i++ {
		(*s)[i], (*s)[l-1-i] = (*s)[l-1-i], (*s)[i]
	}
}
