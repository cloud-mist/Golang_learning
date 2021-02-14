// for..range 遍历数组,更安全
package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "xi"}

	for index,note := range notes {
		fmt.Println(index, note)
	}
}
/*
0 do
1 re
2 mi
3 fa
4 so
5 la
6 xi
*/
