package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")	
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}


type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")	
}
	

// 1.代表任何含有MakeSound方法的类型，Whistle和Horn类型都满足
type NoiseMaker interface {
	MakeSound()
}


func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()

	toy = Horn("Toyco Blaster")
	toy.MakeSound()


	var noiseMaker NoiseMaker = Robot("Botco") // 定义一个接口类型变量，将满足接口类型值赋给它
	noiseMaker.MakeSound() // 调用接口中的方法

	// 使用类型断言取回具体类型的值，调用那儿个类型上的方法，但这方法不属于接口
	var robot Robot = noiseMaker.(Robot) // *** 使用类型断言取回具体类型
	robot.Walk() // 调用在具体类型上的方法

		
}
//Tweet!
//Honk!
