// 将一个整数值字符串，每隔三个字符插入一个逗号。

// 整数类型

package main

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	// 递归
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	s := "1234567890"
	fmt.Println(comma(s)) // 1,234,567,890
}
