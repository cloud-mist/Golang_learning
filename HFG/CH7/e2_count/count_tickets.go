// Package main provides count tallies the number of times each line
package main

import (
	"Golang_learning/HFG/CH7/e1_datefile"
	"fmt"
	"log"
)

func main() {
	lines, err := e1_datefile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal()
	}
	var names []string // 保存候选人姓名
	var counts []int   // 保存每个名字出现的次数
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
//Amber Draham: 3
//Brian Martin: 2


