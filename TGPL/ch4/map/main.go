// map的简单使用
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1.创建空的map
	m1 := make(map[string]int)
	m2 := map[string]int{}

	// 2.插入
	m1["bob"] = 10
	m1["xm"] = 20
	m1["tstts"] = 1
	m2["tstts"] = 1

	// 3.删除
	delete(m1, "bob") // 即使不存在也不会报错

	a := m1["bsetnse"]
	fmt.Println(a) // 0
	fmt.Println("-------------------------------")
	// 4.1 简单遍历
	for k, v := range m1 {
		fmt.Println(k, v)
		// xm 20
		// tstts 1
	}
	fmt.Println("-------------------------------")

	// 4.2 按名字顺序遍历,

	var names []string
	for k := range m1 {
		names = append(names, k)
	}

	// 排序
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name, m1[name])
	}
	// tstts 1
	// xm 20

}
