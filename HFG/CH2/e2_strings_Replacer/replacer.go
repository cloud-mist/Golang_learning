package main

import (
	"fmt"
	"strings"
)

// Replacer 类型。在字符串中搜索子字符串，并将该子字符串出现的地方用另一个字符串替换
func main() {
	broken := "G# r#cks!"	
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)		// Go rocks!
}
