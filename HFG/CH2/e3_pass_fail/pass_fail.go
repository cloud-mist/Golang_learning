package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 评分
func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin) // 读取输入的方法
	input, err := reader.ReadString('\n') // 字符串形式返回用户输入;标记输入的结束，直到按下<Enter>
	// 只有错误不为空的时候，才报告错误
	if err != nil{
		log.Fatal(err)  // 用Fatal函数记录错误并停止程序运行
	}

	/* 需要将input转换成数字。
	* 1.将换行符去掉 TrimSpace:删除字符串开头结尾的所有空白字符
	 * 2.转换成符点数   strconv包的ParseFloat函数
	*/
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil{
		log.Fatal(err)  // 用Fatal函数记录错误并停止程序运行
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)

}
