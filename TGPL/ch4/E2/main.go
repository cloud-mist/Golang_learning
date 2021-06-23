// 1.可以命令行指定额外的编码方案
// 2.标准输入
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var fl1 = flag.String("sha512", " ", "512 code")
var fl2 = flag.String("sha384", " ", "384 code")

func main() {
	flag.Parse()
	if len(os.Args) <= 1 {
		var s []byte
		fmt.Scan(&s)
		res := sha256.Sum256(s)
		fmt.Println(res)
	} else {
		if os.Args[1] == "-sha512" {
			res := sha512.Sum512([]byte(os.Args[2]))
			fmt.Println(res)
		} else {
			res := sha512.Sum384([]byte(os.Args[2]))
			fmt.Println(res)
		}
	}
}
