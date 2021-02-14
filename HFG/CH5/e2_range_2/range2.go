package main

import "fmt"

// 用下划线来忽略index返回值
func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "xi"}

	for _, note := range notes {
		fmt.Println(note)
	}
}
/*
do
re
mi
fa
so
la
xi

*/
