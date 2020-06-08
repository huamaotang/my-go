package main

import "fmt"

func main() {
	a := 1
	b := 1.9
	// 类型不一致，不能进行运算，编译出错
	fmt.Println(a + int(b)) // 2

	c := [4]int{1, 2, 3, 4}
	t := c[2:4:4]
	// [i:j:k] i左包含，j右不包含；k为数组c的容量；t的长度为[2,3)，容量为[2,4),即2
	fmt.Println(len(t), cap(t), t) // 1 2 [3]

	d := [2]int{1, 2}
	f := [3]int{1, 2}
	fmt.Println(d, f)
	// 类型不一致，不能进行比较
	//fmt.Println(d == f)
}
