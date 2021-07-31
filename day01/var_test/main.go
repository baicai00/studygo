package main

import "fmt"

// func test(a int, b float32) {
// 	fmt.Printf("test:%d %f", a, b)
// }

func main() {
	// fmt.Println("hello world")
	// test(4, 10.3)

	// 字符串 = 测试
	// var str string
	// str = "hello world" // =赋值的变量需先被声明
	// fmt.Printf("%s", str)

	// := 测试
	// str := "hello world" // =:
	// fmt.Printf("%s", str)

	// 指针
	var a int = 10
	ptr := &a
	*ptr = *ptr + 100
	fmt.Printf("%d", a)
}
