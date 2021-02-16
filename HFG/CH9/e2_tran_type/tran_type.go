// 不同类型转换
package main

import "fmt"

type Liters float64
type Gallons float32

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))

	fmt.Printf("Car fuel: %.1f gallons.", carFuel)
	fmt.Printf("Bus fuel: %.1f liters", busFuel)
}
