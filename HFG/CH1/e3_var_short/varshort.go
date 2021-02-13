package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1.短变量声明 := 
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "damon"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity,"sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	// 2.类型转换
	myInt := 2
	fmt.Println(reflect.TypeOf(myInt))				// int
	fmt.Println(reflect.TypeOf(float64(myInt)))		// float64

}

// 变量、函数、类型的名称以大写字母开头，是可以导出，可以从当前包外的包访问它。后面单词遵循驼峰规则
