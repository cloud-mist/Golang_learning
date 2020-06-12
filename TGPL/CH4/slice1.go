package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "mar", 4: "april", 5: "May", 6: "June", 
					7: "July",8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)		//[June July Aug]
	for _, s := range summer{
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n",s)
			}
		}
	}
	summer2 := summer[:5]   // extend a slice
	fmt.Println(summer2)	// [June July Aug Sep Oct]
}
