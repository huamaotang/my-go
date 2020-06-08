package main

import "fmt"

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type work struct {
	i int
}

func (w *work) ShowA() int {
	fmt.Println("show a")
	return w.i + 1
}

func (w *work) ShowB() int {
	fmt.Println("show b")
	return w.i + 2
}

func main() {
	// nil切片 与nil相等，表示不存在的切片
	var s1 []int
	// 空切片 []int{}，表示一个空的集合；不是一个类型，而是值，不能直接var声明，如：var s2 []int{}，需要赋值
	var s2 = []int{}

	if s1 == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	if s2 == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	// 知识点：接口
	// 一种类型实现多个接口，接口变量a、b调用各自的方法ShowA()、ShowB()方法
	w := new(work)
	w.i = 1
	fmt.Println(w.ShowA(), w.ShowB())
	// 知识点：接口的静态类型
	// a、b具有相同的动态类型和动态值，分别是 work 和 {3}；a的静态类型是A，b的静态类型是B，仅包含所属静态类型的方法
	var a A = w
	var b B = w
	fmt.Println(a.ShowA(), b.ShowB())
	//fmt.Println(a.ShowB(), b.ShowA())
	var c A = w
	// type assertion，去除掉静态类型，只保留动态类型和动态值
	d, ok := c.(*work)
	fmt.Println(d.ShowA(), d.ShowB(), ok)

}
