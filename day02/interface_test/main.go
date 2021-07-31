package main

import "fmt"

// 接口
type speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker类型，方法签名
}

// 结构体
type cat struct{}
type dog struct{}
type people struct{}

// cat的方法
func (c cat) speak() {
	fmt.Println("喵喵喵~~")
}

// dog的方法
func (d dog) speak() {
	fmt.Println("汪汪汪~~")
}

// people的方法
func (p people) speak() {
	fmt.Println("啊啊啊~~")
}

// 只要是定义了speak方法的类型都能被当作speaker传入da函数
func da(x speaker) {
	x.speak()
}

func main() {
	var c cat
	var d dog
	var p people

	da(c)
	da(d)
	da(p)
}
