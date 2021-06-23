// 重复行只打印一次
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !s[line] {
			s[line] = true
			fmt.Println(line)
		}
	}
}
