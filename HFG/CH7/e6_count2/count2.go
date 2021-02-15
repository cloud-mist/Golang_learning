package main

import (
	"Golang_learning/HFG/CH7/e1_datefile"
	"fmt"
	"log"
)

func main() {
	lines, err := e1_datefile.GetStrings("../e2_count/votes.txt")	
	if err != nil {
		log.Fatal()
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
//Votes for Amber Draham: 3
//Votes for Brian Martin: 2
