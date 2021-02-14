// Package main provides ...
package main

import (
	"Golang_learning/HFG/CH5/e5_datefile"
	"fmt"
	"log"
)

func main() {
	// 加载date.txt,解析包含的数，并保存到数组
	numbers, err := e5_datefile.GetFloats("../e4_readfile/date.txt")
	if err != nil {
		log.Fatal()
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/sampleCount)
}
