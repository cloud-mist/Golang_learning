package main

import "fmt"

func main() {
	// 类型转换：相同的类型才可以计算，比较

	price := 100
	fmt.Println("Price is", price, "dollars.")

	taxRate := 0.08
	tax := float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars")

	total := float64(price) + tax
	fmt.Println("Total cost is", total, "dollars")

	availableFunds := 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))
}

/*
Price is 100 dollars.
Tax is 8 dollars
Total cost is 108 dollars
120 dollars available.
Within budget? true
*/
