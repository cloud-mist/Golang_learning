// iota,第一个初始化为0，后面依次+1
package main

import "fmt"

type Weekday int

const (
	Sun Weekday = iota // 0
	Mon                // 1
	Tue
	Wed
	Tus
	Fri
	Sat
)

func main() {
	fmt.Println(Wed) // 3
	fmt.Println(Sat) // 6
}
