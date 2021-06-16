// 修改dup2,打印重复行所在文件名
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	// 看后面有没有跟文件参数。没有：用标准输入，有：打开
	for _, arg := range files {
		f, err := os.Open(arg)
		// os.Open 参数1：被打开的文件，第二个err
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, arg)
		f.Close()
	}

}

func countLines(f *os.File, counts map[string]int, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Println(arg)
		}

	}
}
