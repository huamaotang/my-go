package main

import "fmt"

// 基于类型int创建了新类型MyInt1
type MyInt1 int

// 创建了int32类型的别名为MyInt2
type MyInt2 = int32

func main() {
	p := new(struct {
		name string
	})
	// 访问成员变量的两种方式
	p.name = "a"
	(*p).name = "b"
	fmt.Println(p.name, (*p).name)

	var i int32
	i = 5
	// 类型不一致，不能进行赋值，可使用强类型转换
	var i1 MyInt1
	i1 = MyInt1(i)
	var i2 MyInt2
	i2 = i
	fmt.Println(i1, i2)
}
