// 将看起来是系统路径的前缀删除，将看似是文件类型的后缀删除。
package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "a/b/c/abc.go"
	b := "a.b.go"
	c := "abc"
	fmt.Println(basename1(a))
	fmt.Println(basename1(b))
	fmt.Println(basename1(c))
	fmt.Println(basename2(a))
}

// v1:纯手工
func basename1(s string) string {
	// 1.将最后一个/的后面保留
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 将最后一个.前面的东西保留
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// v2：使用库函数
func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

/* RES
abc
a.b
abc
abc
*/
