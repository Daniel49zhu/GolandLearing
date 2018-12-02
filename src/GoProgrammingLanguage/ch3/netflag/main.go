package main

import (
	"fmt"
	. "net"
)

//iota从0开始 每次加一
const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
	fmt.Printf("%d,%d,%d,%d,%d", KiB, MiB, GiB, TiB, PiB)
}
