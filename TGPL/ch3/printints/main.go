// ints2String

package main

import (
	"bytes"
	"fmt"
)

func ints2String(values []int) string {
	// 用于字节slice的缓存
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(ints2String([]int{1, 2, 3}))
}

//  [1, 2, 3]
