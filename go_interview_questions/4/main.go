package main

import "fmt"

func main() {
	// new(T) 返回的是指针*T，而不是T类型，所以不能对指针直接append，需要解引用，重新赋值
	// 可以使用make(T, args)初始化之后使用。或者使用字面量的方式初始化
	list := new([]int)
	s := append(*list, 1)
	fmt.Printf("%#v %v\n", *list, s)

	// 两个slice进行append操作，需要使用...操作符，用法：append(s1, s2)或是append(s1, 1, 2)
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)

	var (
		// 短变量声明方式的限制
		// 1.
		//size    := 1024	已经声明的变量，不能再进行变量声明
		size    = 1024
		maxSize = size * 2
	)
	fmt.Println(size, maxSize)
}
