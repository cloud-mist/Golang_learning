// 标准输入中，输出有重复的行，次数和行内容
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// make创建一个map,键可以任意类型，值能用`==`比较，此处是键字符串，值整数。map类似py中的dict
	input := bufio.NewScanner(os.Stdin)

	// 计数，读取一行，该行字符串数+1
	for input.Scan() {
		counts[input.Text()]++
	}

	// 输出所有计数大于1的行
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
