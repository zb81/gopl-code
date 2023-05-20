package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(strings.Join(os.Args[1:], " "))
	//
	//fmt.Println(os.Args[1:])
	//fmt.Println(os.Args[0])

	for i, v := range os.Args {
		fmt.Printf("%d\t%s\n", i, v)
	}
}
