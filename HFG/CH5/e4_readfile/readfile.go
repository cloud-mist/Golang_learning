package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 1.打开文件，若打开出错，报告错误退出
	file, err := os.Open("date.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 2.为文件创建一个扫描器，并读取一行打印一行
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 3.关闭文件，若错误，报告错误并退出
	err = file.Close()
	if err != nil {
		log.Fatal()
	}

	// 4.扫描文件错误，就报告错误退出
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
/*
71.8
56.2
89.5
*/
