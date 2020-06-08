package main

import "fmt"

func main() {
	// 1 5 1
	// 知识点：defer、返回值执行顺序
	// 1. 返回值 = xxx
	// 2. 调用defer函数，查看返回值是否被作为参数传入
	// 3. 空的return
	fmt.Println(f1(), f2(), f3())
}

func f1() (r int) {
	// r = 0
	// defer语句为匿名函数且r未作为参数传入，则返回值被修改
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	// 匿名函数内外的变量r是两个变量，所以互不影响
	defer func(r int) {
		fmt.Println(r) // 0
		r = r + 5
		fmt.Println(r) // 5
	}(r)
	return 1
}
