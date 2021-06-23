// 两个sha256的不同bit位的数量
package main

import (
	"crypto/sha256"
	"fmt"
)

// 循环逐位比较
func main() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))
	var res int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			res++
		}
	}
	fmt.Println(res)
}
