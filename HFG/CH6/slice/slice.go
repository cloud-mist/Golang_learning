// Package main provides ...
package main

import "fmt"

func main() {

	// append 在末尾追加元素
	// 注意：要将append完的返回值赋给传入的那个切片变量，是因为避免不一致行为
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))

	
}
