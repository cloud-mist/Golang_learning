package main

import "fmt"
	type subscriber struct{
		name string
		rate float64
		active bool
	}
func main() {
	var subscriber1 subscriber
	subscriber1.name = "Aman"
	fmt.Println("Name1: ", subscriber1.name)

	var subscriber2 subscriber
	subscriber2.name = "Beth"
	fmt.Println("Name2: ", subscriber2.name)

}
//Name1:  Aman
//Name2:  Beth
