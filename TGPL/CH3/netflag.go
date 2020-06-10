//一个位运算的复习过程 和自带的iota
package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLookback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadCast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp // v=10001
	fmt.Printf("%b %t\n", v, IsUp(v))    //true
	TurnDown(&v)                         //v = 10000
	fmt.Printf("%b %t\n", v, IsUp(v))    //false
	SetBroadCast(&v)                     //10010
	fmt.Printf("%b %t\n", v, IsUp(v))    //false
	fmt.Printf("%b %t\n", v, IsCast(v))  // 10010 true
}

