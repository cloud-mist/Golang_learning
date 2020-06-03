// Echo2 prints its command-line arguments.
// 展示  在某种数据类型区间（range） 上遍历， 如字符串或者切片

package main

import(
	"fmt"
	"os"
)

func main(){
	s, sep := "", ""
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}