// echo1 prints its command-line arguments.命令行参数
package main

import(
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// 两个string变量隐式初始化为空字符串。
	// 声明时直接初始化。若没被初始化，则会赋予其类型的零值
	for i := 1; i < len(os.Args); i++ {
		// := 短变量声明    ps： 不能写成++i
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}