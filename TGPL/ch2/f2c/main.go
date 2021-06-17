// 用函数封装
package main

import "fmt"

func f2c(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	//
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, f2c(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, f2c(boilingF))
}

/*
	32F = 0C
	212F = 100C
*/
