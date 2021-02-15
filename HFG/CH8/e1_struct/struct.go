// Package main provides Struct使用
package main

import "fmt"

func main() {
	// 1. 声明
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)

	// 2. 赋值。获取值
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println(myStruct.toggle) // true
}

// struct { number float64; word string; toggle bool }{number:0, word:"", toggle:false}
