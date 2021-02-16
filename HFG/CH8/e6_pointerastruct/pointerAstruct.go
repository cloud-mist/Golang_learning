// 指针可以节省空间
package main

import (
	"Golang_learning/HFG/CH8/e7_magazine"
	"fmt"
)

// 获取指针作为参数
func printInfo(s *e7_magazine.Subscriber) {
	fmt.Println("name:", s.Name)	
	fmt.Println("rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}

func defaultSubscriber(name string) *e7_magazine.Subscriber {
	var s e7_magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s  // 返回一个指向subscirber的指针
}

func applyDiscount(s *e7_magazine.Subscriber) {
	s.Rate = 4.99	
}

func main() {
	subscriber1 := defaultSubscriber("aman") // subscriber1是一个指针
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("beth")
	printInfo(subscriber2)

	employee := e7_magazine.Employee{Name: "Joy", Salary: 60000}
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	address := e7_magazine.Address{Street: "123 Oak St", 
								   City: "Omaha", 
								   State: "NE", 
								   PostalCode: "68111"}
	subscriber0 := e7_magazine.Subscriber{Name: "zzz"}
	subscriber0.HomeAddress = address
	fmt.Println(subscriber0.HomeAddress)
}
/*
name: aman
rate: 4.99
Active? true
name: beth
rate: 5.99
Active? true
Joy
60000
{123 Oak St Omaha NE 68111}
*/


