package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
	// %g 浮点数,根据情况选用%e或%g
}

// result
// boiling point = 212F or 100C
