// 定义方法
package main

import "fmt"

type MyType string

// 接收器m, sayHi()方法名
func (m MyType) sayHi() {
	fmt.Println("Hi")
}
func main() {
	// 方法被定义在某个类型，它就能被该类型的任何值调用。
	value := MyType("a MyType value")
	value.sayHi()
	anotherValue := MyType("another Value")
	anotherValue.sayHi()

}
