// Package main provides  同步goroutine和channel的例子
package main

import "fmt"

// 创建两个channel并将他们传递给两个新goroutine中的函数
// main goroutine从这些channel 接收值并打印

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go abc(channel1)
	go def(channel2)

	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()
}
// adbecf  改善：指定了打印的顺序
