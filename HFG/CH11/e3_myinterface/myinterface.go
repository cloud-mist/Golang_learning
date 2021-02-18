package main

import (
	"Golang_learning/HFG/CH11/e2_mypkg"
	"fmt"
)

func main() {
	var value e2_mypkg.MyInterface // 声明一个接口类型变量
	value = e2_mypkg.MyType(5) //MyType的值满足MyIntherface,所以我们可以赋值给MyInterface的变量

	value.MethodWithoutParameters()
	value.MethodWithParameter(126.3)
	fmt.Println(value.MethodWithReturnValue())
}

/*
MethodWithoutParameters called.
MethodWithParameter called with 126.3
Hi from MethodWithReturnValue
*/
