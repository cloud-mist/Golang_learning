//打印标准输入中多次出现的行，以重复次数开头
//引入 if， map， bufio包

package main

import(
	"bufio"
	"fmt"
	"os"
)
//bufio处理输入和输出    Scanner读取输入并将其拆成行或者单词；处理行输入最简单的办法
func main(){
	counts := make(map[string]int)
	//map类似于 py的dict
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
	}
	for line, n := range counts{
		if n>1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}