package e2_mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

// 为了使MyType满足MyInterface,定义了接口需要实现的所有方法，并另外包含了一个不属于接口的额外方法
type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called.")
}
func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (my MyType) MethodNotInInterface(){
	fmt.Println("MethodNotInInterface called.")
}

// 如果一个类型包含接口声明中的所有方法，那么它可以在任何需要接口的地方使用。
