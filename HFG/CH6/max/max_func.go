// 可变长参数的使用,必须是最后一个参数，在其类型前加...
package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)	 // 一个非常小的值
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
func main() {
	fmt.Println(maximum(71.8, 34.3, 89.2))			// 89.2
	fmt.Println(maximum(90.2, 23.2, 43.1, 43.2))	// 90.2
	
}
