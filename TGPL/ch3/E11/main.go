// E11 支持浮点数，可选正负号
package main

import (
	"fmt"
	"strings"
)

// TODO: 还未完成  <17-06-21, cloud mist> //
func comma(s string) string {
	n := strings.Index(s, ".")
	if n != -1 {

	} else {

	}
	return res
}

func main() {
	s1 := "1234567890.1111"
	s2 := "-123"
	fmt.Println(comma(s1))
	fmt.Println(comma(s2))
}
