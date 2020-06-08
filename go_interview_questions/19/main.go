package main

import "fmt"

type Person struct {
	age int
}

func main() {
	//

	person := &Person{28}

	// 1. 赋值的副本为值28，所以输出28
	defer fmt.Println(person.age)

	// 2. 赋值的副本为指针地址，输出时，成员变量age的值已经被修改为29
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3. defer函数在person.age = 29执行之后执行，所以输出29
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}
