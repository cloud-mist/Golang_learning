// 声明变量

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1.先声明，后赋值
	var quantity int
	var length, width float64
	var customerName string

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "damon cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	// 2.声明的时候同时赋值，可以省略变量类型
	var name = "shawn"
	fmt.Println(reflect.TypeOf(name))

	// 3.声明不赋值，是该类型对应的零值，int 0;float64 0;bool false;string 空字符串;
	var myInt int
	var myBool bool
	var myString string
	fmt.Println(myBool, myString, myInt)
}

/*		
damon cole
has ordered 4 sheets
each with an area of
2.88 square meters

string

*/
