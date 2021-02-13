// Package main provides ...
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 查看类型
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("hello world"))
}


/*
 int
 float64
 bool
 string
*/

