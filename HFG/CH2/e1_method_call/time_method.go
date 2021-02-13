package main

import (
	"fmt"
	"time"
)

// go的方法有点像其他语言附加到“对象”上的方法
func main() {
	// time包的Time类型，每一个time.Time值有一个返回年份的Year方法
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)	// 2021
}
