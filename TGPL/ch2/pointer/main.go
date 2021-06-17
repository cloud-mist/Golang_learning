// 指针使用
package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(*p) // 1

	*p = 2
	fmt.Println(x) // 2
}
