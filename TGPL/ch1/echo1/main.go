// 输出命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string // 会隐式初始化，赋予其类型的零值
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
