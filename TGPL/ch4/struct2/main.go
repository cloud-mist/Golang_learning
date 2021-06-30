package main

func main() {
	// v1
	// 圆心和半径
	type Circle struct {
		X, Y, Radius int
	}
	// 增加了S,径向辐条
	type Wheel struct {
		X, Y, Radius, Spokes int
	}

	// wheel变量
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	// v2
	//这两个有相似之处,所以把相同的独立出来,但访问成员繁琐
	type Point struct {
		X, Y int
	}
	type Circ struct {
		Center Point
		Radius int
	}

	type Whe struct {
		Cir    Circ
		Spokes int
	}
	var w1 Whe
	w1.Cir.Center.X = 8
	w1.Cir.Center.Y = 8
	w1.Cir.Radius = 5
	w1.Spokes = 20

	// v3
	// Go特性: 只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。
	// 匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针

	type Circ2 struct {
		Point
		Radius int
	}
	type Whe2 struct {
		Circ2
		Spokes int
	}

	var w2 Whe2
	w2.X = 8
	w2.Y = 8
	w2.Radius = 5
	w2.Spokes = 20

}
