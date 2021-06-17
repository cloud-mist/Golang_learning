// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// 1,命令行参数名, 2.默认值， 3.提示信息
// 用法 命令行用-n ，-s

func main() {
	flag.Parse()
	// 更新标志参数值
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
