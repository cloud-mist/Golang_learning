// 类型定义可以放到函数中，但为了函数适用，所以一般定义在函数外。
package main

import "fmt"

// 定义part类型，struct是基础类型
type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	var porsche car // 定义一个car类型的变量
	porsche.name = "porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name: ", porsche.name)
	fmt.Println("topSpeed: ", porsche.topSpeed)
}

/*
Name:  porsche 911 R
topSpeed:  323
*/
