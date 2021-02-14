// 油漆计算器
package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error){
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area/10.0, nil
}

func main() {
	var amount, total float64
	amount, err := paintNeeded(4.2, 3.0)	
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%.2f liters needed\n", amount)	// 1.26 liters needed
	}
	total += amount
	amount,err = paintNeeded(5.2, 3.5)	
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%.2f liters needed\n", amount)	// 1.82 liters needed
	}
	total += amount
	fmt.Printf("Total: %.2f liters\n", total)	// Total: 3.08 liters
}


