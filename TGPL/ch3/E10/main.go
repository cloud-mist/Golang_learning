// 非递归版本

package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	start := n % 3
	var res bytes.Buffer
	if n <= 3 {
		return s
	}
	i := start
	var t int
	for ; i < n; i += 3 {
		res.WriteString(s[t:i])
		res.WriteByte(',')
		t = i
	}
	res.WriteString(s[t:])

	if res.String()[0] == ',' {
		return res.String()[1:]
	} else {
		return res.String()
	}
}

func main() {
	s := "123456789012345"
	fmt.Println(comma(s))
}
