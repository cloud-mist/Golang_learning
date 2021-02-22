// 一级函数
package main

import "fmt"

func sayHi() {
	fmt.Println("Hi!")
}
func main() {
	var myFunc func() // 声明一个可以储存一个函数的变量
	myFunc = sayHi
	myFunc() // 调用储存在变量中的函数
}

// Hi!
