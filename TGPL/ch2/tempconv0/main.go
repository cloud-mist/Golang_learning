// 明明类型
package main

import "fmt"

// 虽然有相同底层类型，但是是不同的数据类型
type Cel float64 // 摄氏
type Fah float64 // 华氏

const (
	Absolute0C Cel = -273.15 // 绝对零度
	FreezingC  Cel = 0       // 结冰温度
	BoilingC   Cel = 100     // 沸水温度
)

func C2F(c Cel) Fah {
	return Fah(c*9/5 + 32)
}

func F2C(f Fah) Cel {
	return Cel((f - 32) * 5 / 9)
}

// 参数c出现在函数名前，声明的是cel类型的一个名叫String的方法
func (c Cel) String() string {
	return fmt.Sprintf("%gC", c)
}
func main() {
	c := F2C(212.0)
	fmt.Println(c.String()) // 100C
	// 许多类型会定义一个String方法，会优先使用该类型对应的方法返回的结果打印
}
