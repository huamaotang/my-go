package main

import "fmt"

type Person struct {
	age int
}

func main() {
	// 29 28 28
	person := &Person{28}

	// 1. 实参28被复制副本，执行时，输出28
	defer fmt.Println(person.age)

	// 2. 指针地址被复制副本，成员变量age的值为28
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3. 后进先出；指针地址修改，之后的地址成员变量age的值为29
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}
