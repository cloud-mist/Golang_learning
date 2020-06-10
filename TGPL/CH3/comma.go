//comma 逗号   每三个字符插入一个逗号
package main

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
func main() {
	a := "djkfdlfjdklfjdlk"
	b := "1231231544864"
	fmt.Println(comma(a))
	fmt.Println(comma(b))
}
