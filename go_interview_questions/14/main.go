package main

import "fmt"

func main() {
	str := "incre"
	// 知识点：golang中的字符串是只读的
	//str[0] = "x"
	fmt.Println(str)

	// 知识点：指针，
	m := 1
	fmt.Println(inCre(&m))

	// 知识点：可变函数
	fmt.Println(add(1))
	fmt.Println(add(1, 2, 3))
	fmt.Println(add([]int{1, 2, 3}...))
}

func inCre(m *int) int {
	// 变量m指向的是main函数的变量m的地址，修改的值也就是main函数中变量m的值
	*m++
	return *m
}

func add(params ...int) int {
	var sum int
	for _, v := range params {
		sum += v
	}
	return sum
}
