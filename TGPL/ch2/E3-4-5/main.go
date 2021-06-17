package main

import (
	"Golang_learning/TGPL/ch2/popcount"
	"fmt"
)

func main() {
	fmt.Println(popcount.PopCount(10000))
	fmt.Println(popcount.PopCountE3(10000))
	fmt.Println(popcount.PopCountE4(10000))
	fmt.Println(popcount.PopCountE5(10000))

}
