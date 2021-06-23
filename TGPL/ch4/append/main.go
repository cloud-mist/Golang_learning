package main

import "fmt"

// cap 是当前切片的左端点到数组本身的右端点

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

// len(x) 是slice的长度，cap是底层的容量
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	// 先检测slice底层数组是否有足够的容量来保存新添加的元素
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		// 扩展到2倍大小
		if zcap <= 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		// 把x复制到z里
		copy(z, x) // 目的,源

	}
	// 添加y
	z[len(x)] = y
	return z
}
