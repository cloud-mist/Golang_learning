package main

import (
	"fmt"
	"log"
)
type OverheatError float64
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %.2f degrees!", o)
}

func CheckTemperature(actual float64, safe float64)  error { // 指定函数返回原生error值
	excess := actual - safe	
	if excess > 0 {
		 return OverheatError(excess)
	}
	return nil
}

func main() {
	var err error = CheckTemperature(121.3, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
/*
2021/02/18 17:59:19 Overheating by 21.30 degrees!
exit status 1
*/
