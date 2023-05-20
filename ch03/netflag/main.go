package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // 1
	FlagBroadcast                      // 2
	FlagLoopback                       // 4
	FlagPointToPoint                   // 8
	FlagMulticast                      // 16
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}
func TurnDown(v *Flags) {
	*v &^= FlagUp
}
func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}
func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	var v = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
