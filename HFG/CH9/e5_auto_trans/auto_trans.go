// 调用的时候， 变量类型会自动转换
package main

import "fmt"

type MyType string

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}
func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("a value")
	pointer := &value
	value.method() // 值类型自动转换为指针
	value.pointerMethod()
	pointer.method() // 指针类型自动转换成值
	pointer.pointerMethod()
}

/*
Method with value receiver
Method with pointer receiver
Method with value receiver
Method with pointer receiver
*/
