// 允许将函数作为参数传递给其他函数
package main

import "fmt"

func sayHi() {
	fmt.Println("hi")	
}

// 调用传入的函数两次
func twice(theFunc func())  {
	theFunc()	
	theFunc()
}

func sayBye()  {
	fmt.Println("Bye")	
}

func main() {
	twice(sayHi)
	twice(sayBye)
}
/*
hi
hi
Bye
Bye
*/
