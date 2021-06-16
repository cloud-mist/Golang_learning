package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

/*
➜  E2 (master) ✗ go run main.go tnse nates tnse
0 tnse
1 nates
2 tnse
*/
