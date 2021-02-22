// Package main provides ...
package main

import "fmt"

func greeting(myChannel chan string) { // 将channel作为参数
	myChannel <- "hi"				// 通过channel发送一个值
}

func main() {
	myChannel := make(chan string)	// 创建channel
	go greeting(myChannel)			// 将channel传递给goroutine中运行的函数
	recivedValue := <-myChannel		//从channel接收值
	fmt.Println(recivedValue)		
}
