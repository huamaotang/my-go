package main

import "fmt"

type MyInt int

// 接受者的类型可以是任意类型，不仅仅是结构体，任何类型都可以拥有方法；
func (MyInt) test() {
	fmt.Println("test")
}

type T struct {
	name string
}

// 值类型的接受者：Go语言会在代码运行时将接受者的值复制一份。可以获取接受者的成员变量；修改只针对副本，无法修改接受者变量本身
func (t T) bar() {
	// 当接受者为指针类型时，需要先解引用（所以接受者变量不能为nil，否则nil pointer），再复制值给t
}

// 指针类型接受者：由一个结构体的指针组成，修改接受者的任意成员变量，方法结束之后，修改是有效的；类似于其他语言的this、self
func (t *T) foo() {
	// t不能等于nil，nil pointer
	t.name = "foo"
}

type S struct {
	*T
}

func main() {
	//var s = S{}
	//s.bar()
	//fmt.Println(s.name)

	i := MyInt(1)
	i.test()
}
