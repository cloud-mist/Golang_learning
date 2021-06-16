package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[:], " ")
	// 因为 Args[0]代表着执行命令本身的名字
}
