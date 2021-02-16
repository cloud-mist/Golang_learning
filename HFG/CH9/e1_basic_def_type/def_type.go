// 基于基础类型来定义类型，为了让值的用途更加明确
// Package main provides ...
package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	carFuel := Gallons(10.0)
	var busFuel Liters
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)
}
