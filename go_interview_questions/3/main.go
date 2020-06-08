package main

import "fmt"

func main() {
	// 第一题
	s1 := make([]int, 5)
	s1 = append(s1, 1, 2, 3)
	// [0 0 0 0 0 1 2 3]
	// append是在原slice上添加元素，而不是覆盖
	fmt.Println(s1)
	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2)

	// 第二题 new(T) make(T, args)的区别
	// 都是Go内置函数，用来分配内存，但适用的类型不同
	// new(T) 会为T分配已置零的内存空间，并返回地址（指针），即类型为*T的值。换句话说，返回一个指针，该指针指向新分配的、类型为T的零值。适用于值类型，如数组、结构体
	// make(T, args) 返回初始化之后的T类型的值，这个值并不是T类型的零值，也不是指针*T，是经过初始化之后的T的引用。make()只适用于slice、map、channel

}
