// Package main provides map
package main

import "fmt"

func main() {
	// 1.声明
	isPrime := make(map[int]bool)
	isPrime[4] = false	
	isPrime[7] = true
	fmt.Println(isPrime[7])		// true
	fmt.Println(isPrime[4])		// false

	// 2.映射字面量
	ranks := map[string]int{"bronze":3, "silver": 2, "gold": 1}
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])

	// 3.创建空映射
	emptyMap := map[string]float64{}
	fmt.Println(emptyMap)		// map[]

}
