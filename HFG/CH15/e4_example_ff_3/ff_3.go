// 函数的参数和返回值是其类型的一部分
package main

import "fmt"

func sayHi() {
	fmt.Println("hi")
}
func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}
func main() {
	var greeterFunc func()              // 这个变量将保存没有返回值和参数的函数
	var mathFunc func(int, int) float64 // 这个变量保存有两个int参数和一个f64返回值的函数
	greeterFunc = sayHi
	mathFunc = divide
	greeterFunc()
	fmt.Println(mathFunc(5, 2))

}

//hi
//2.5
