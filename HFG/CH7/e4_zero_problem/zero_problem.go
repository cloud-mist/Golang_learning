// 测试值是否赋值，还是赋值了是零值。访问映射键可选获取第二个布尔类型的值，若赋过会返回true
package main

import "fmt"

func status(name string) {
	grades := map[string]float64{"alma": 0, "pohit": 86.5}
	_, ok := grades[name]
	if !ok {
		fmt.Printf("no grade recorded for %s.\n", name)
	} else {
		fmt.Printf("%s is falling!\n", name)
	}
}
func main() {
	status("alma")	
	status("carl")
}
//alma is falling!
//no grade recorded for carl.
