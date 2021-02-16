package main

import "fmt"

type Number int

func (n *Number) Double() {
	*n *= 2
}

func main() {
	number1 := Number(4)
	fmt.Println("Original value:", number1)
	number1.Double()
	fmt.Println("double :", number1)
}

//Original value: 4
//double : 8
