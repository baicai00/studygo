package main

import "fmt"

func main() {
	// 一维数组赋值、比较、遍历
	// a := [...]int{2, 3, 4}
	// b := [3]int{2, 3, 4}
	// fmt.Println(a == b)   // 判断数组是否相等
	// for k, v := range a { // 数组遍历
	// 	fmt.Println(k, v)
	// }

	// 切片
	a := [...]int{1, 2, 3, 4, 5}
	b := a[1:]
	for k, v := range b {
		fmt.Println(k, v)
	}
}
