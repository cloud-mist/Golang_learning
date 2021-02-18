//error类型是一个接口，声明一个接口的error类型意味，如果它具有返回string的Error方法，就满足error接口
// 所以我们能定义自己类型并且用在任何需要error值的地方
package main

import "fmt"

type ComedyError string
func (c ComedyError) Error() string {
	return string(c)	
}

func main() {
	var err error
	err = ComedyError("what's a programmer's favorite berr?")
	fmt.Println(err)
}

