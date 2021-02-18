// Package main provides ...
package main

import (
	"fmt"
	"time"
)
func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}	
}
func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}	
}

func main() {
	// 这样使得a和b和main函数同时运行
	// 这样结果：main goroutine结束后立即停止，但其他还在运行
	go a()
	go b()
	// 所以，我们需要保持main goroutine运行到a,b都运行完。
	// 方法1：先暂停main 1s,输出是混合的
	time.Sleep(time.Second)
	fmt.Println("end main()")
}
