package main

import "fmt"

func main() {
	ranks := make(map[string]int)
	ranks["bronze"] = 3
	rank, ok := ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)		//rank: 3, ok: true

	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)		//rank: 0, ok: false 证明被移除，并不是赋值了0
}
