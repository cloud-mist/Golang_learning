// Package main provides ...
// 和函数一起使用已定义类型
package main

import "fmt"

type part struct{
	description string
	count int
}

// 声明一个part类型的参数
func showInfo(p part) {
	fmt.Println("Description: ", p.description)
	fmt.Println("count: ", p.count)
}

// 声明一个part类型的返回值
func minOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func main() {
	var bolts part	
	bolts.description = "Hex"
	bolts.count = 24
	showInfo(bolts)

	fmt.Println()

	p := minOrder("Hex bolts")
	fmt.Print(p.description, p.count)
}
/*
Description:  Hex
count:  24

Hex bolts 100
*/

