// Basic Operation
// map[K]V   k,v 对应key，value      key必须支持==比较运算符的数据类型
package main

import (
	"fmt"
	"sort"
)

func main() {
	//1.声明
	
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	//2.访问，通过下标
	fmt.Println(ages2["alice"])

	//3.删除
	delete(ages2, "alice")

	//4.遍历所有键值对  ,遍历顺序是随机的
	ages2["abc"] = 12
	for name, age := range ages2 {
		fmt.Printf("%s\t%d\n", name, age)
	}

	//5.如果需要顺序遍历     显式的对其进行排序
	ages := map[string]int{
		"a": 12,
		"v": 6,
		"c": 13,
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

}
