package main

import "fmt"

func main() {
	// 知识点：变量隐藏；短变量声明，如果左边存在多个变量，只需要一个变量是未声明的，会对已声明的变量进行赋值
	// 但是如果出现作用域，会出现变量隐藏
	x, y := 1, 1
	{
		x = 2
		z := 1
		y, z := 2, 2
		fmt.Println(x, y, z)
	}
	fmt.Println(x, y)
}
