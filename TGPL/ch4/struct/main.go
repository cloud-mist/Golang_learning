// struct的简单使用
package main

import "time"

// 1.声明一个Employee的命名的结构体类型
// 如果结构体成员名字是以大写字母开头的，那么该成员就是导出的,一个结构体可能同时包含导出和未导出的成员。
type Employee struct {
	ID            int
	Name, Address string // 类型相同的可以合并
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func main() {
	// 一个变量,类型是Employee
	var dilbert Employee

	// 2.访问,通过`.`
	dilbert.Salary -= 5000
	// 通过指针,对成员去地址，通过指针访问
	pos := &dilbert.Position
	*pos = "senior" + *pos

}
