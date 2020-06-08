package main

import (
	"fmt"
)

func main() {
	// cap() 适用于数组、slice、channel，而不是用于map
	s := []int{1, 2}
	fmt.Println(cap(s))
	a := [2]int{1, 2}
	fmt.Println(cap(a))
	ch := make(chan int)
	fmt.Println("start", cap(ch))
	go func() {
		ch <- 1
		fmt.Println("write", cap(ch))
	}()
	fmt.Println("read", cap(ch))
	<-ch
	fmt.Println("end", cap(ch))

	// 当且仅当接口的动态值和动态类型都为nil时，接口类型值才为nil
	var i interface{}
	fmt.Println(i == nil)

	// 删除map不存在的键值对时，不会报错，相当于没有任何作用；获取不存在的键值对时，返回值类型的零值
	var m map[string]int
	delete(m, "x")
	fmt.Println(m["x"])

}
