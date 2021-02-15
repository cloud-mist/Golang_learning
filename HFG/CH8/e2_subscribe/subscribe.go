package main

import "fmt"

func main() {
	var subscriber struct {
		name string
		rate float64
		active bool
	}
	subscriber.name = "Aman Singh"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Printf("Name:%s\nrate:%v\nactive:%v", subscriber.name, subscriber.rate, subscriber.active)
}
/*
Name:Aman Singh
rate:4.99
active:true
*/
