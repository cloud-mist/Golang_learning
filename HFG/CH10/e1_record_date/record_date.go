// 1.setter方法：判断值是否有效，有效才设置到struct字段
// 2.为了不让用字段访问，进而设置无效值，所以移入另一个包，并不开放访问即可。
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

// 未导出的struct字段可以被相同包导出的函数或方法访问。
func main() {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
// TODO:  后面的几个小节跳过了  <17-02-21, cloud mist> //
