package main

import "fmt"

// 声明全局变量
// 必须用var声明，短变量声明仅限局部变量
//sss := ""
var s string
var ss = ""

type people struct{}

func (p *people) ShowA() {
	fmt.Println("p show a")
	// 调用内部类型方法
	p.ShowB()
}

func (p *people) ShowB() {
	fmt.Println("p show b")
}

type teacher struct {
	people
}

func (t *teacher) ShowB() {
	fmt.Println("t show b")
}

func main() {
	fmt.Println(ss)
	// %d 表示输出十进制数字，+表示输出数值的符号
	i, j := -5, 5
	fmt.Printf("%+d %+d\n", i, j)

	// 知识点：结构体嵌套
	// people称为内部类型，teacher称为外部类型；通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像外部类型自己的一样。
	// 外部类型可以定义自己的属性、方法，甚至可以定义与内部类型相同的方法，这样内部类型就会被『屏蔽』。
	x := &teacher{}
	x.ShowB()
	x.ShowA()
}
