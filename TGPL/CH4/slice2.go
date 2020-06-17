package main

import "fmt"

//   slice内存技巧
func nonempty(strings []string) []string {
	out := strings[:0]
	//zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
}
