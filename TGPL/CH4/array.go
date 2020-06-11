package main

import "fmt"

func main() {
	var a [3]int= [3]int{1,2,3}
	var r [3]int = [3]int{1,2}
	fmt.Println(r[2])
	for _,v := range a{
		fmt.Printf("%d\n", v)
	}
	q := [...]int{1,2,3} //数组的长度根据根据初始化个数指定
	fmt.Printf("%T\n",q)
}
