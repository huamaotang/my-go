package main

import "fmt"

type person struct {
	name string
}

// 可变函数
func hello(num ...int) {
	num[0] = 11
}

func main() {
	// 使用var声明chan变量时，并没有对ch分配空间。
	// var在声明指针、chan、map时，都需要使用make()分配空间，否则无法使用
	var ch chan int
	ch = make(chan int)
	//ch := make(chan int)
	go func() {
		// 写channel是必须带上值
		ch <- 1
	}()
	// 读取channel可以不带变量
	m, ok := <-ch
	fmt.Println(ok, m)
	//
	var n map[person]int
	p := person{"tom"}
	fmt.Println(n, n[p])

	var mx map[string]string
	fmt.Println(mx["a"])
	my := make(map[string]string)
	my["a"] = "1"
	fmt.Println(my)

	s := []int{1, 3}
	hello(s...)
	fmt.Println(s) // [11 3]

	var ss []int
	ss = append(ss, 4, 5)
	fmt.Println(ss)
	ss = make([]int, 2)
	ss[0] = 1
	fmt.Println(ss)
}
