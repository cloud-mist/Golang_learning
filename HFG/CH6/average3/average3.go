// 使用可变长参数函数计算平均值
package main

import "fmt"

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
func main() {
	fmt.Println(average(100, 50))						//75
	fmt.Println(average(90.2, 32.23, 323.1, 323.1))		//192.15750000000003
	
}
