package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (b *Point) Init(x, y int) {
	b.X = x
	b.Y = y
}

func main() {

	// var p Point // 先声明再赋值
	// p.X = 10
	// p.Y = 20
	// p := Point{X: 10, Y: 20} // 直接声明并赋值
	// fmt.Printf("%d %d", p.X, p.Y)

	// 方法、接收器
	var p Point
	p.Init(4, 5)
	fmt.Printf("%d %d", p.X, p.Y)
}
