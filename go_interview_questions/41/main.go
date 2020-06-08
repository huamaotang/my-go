package main

import "fmt"

func main() {
	s := []int{2: 3, 5, 0: 1}
	// 字面量初始化切片，可以指定索引，没有指定索引的元素会在前一个索引的基础上加一
	fmt.Println(s)

	v := 1
	incr(&v)
	fmt.Println(v)
}

func incr(p *int) int {
	// p为指针变量，指向v；*p++表示取出变量v的值并加一
	*p++
	return *p
}
