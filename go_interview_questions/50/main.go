package main

import "fmt"

type Config struct {
	Names []string
}

func change(c Config) {
	c.Names[0] = "c"
}

func main() {
	// 调用函数chang()为值传递，但是函数内的类型成员变量c.Names可以看作是指向c.Names的指针
	c := Config{Names: []string{"a", "b"}}
	change(c)
	fmt.Println(c)

	// Go语言中，不需写明break，因为case完成会默认执行break语句；关键字fallthrough执行时，无需匹配后续判断逻辑，直接执行下一个case；
	// case可以匹配多个值
	f := func(i int32) bool {
		switch i {
		case 1:
			fallthrough
		case 3:
			fallthrough
		case 2, 4:
			return true
		}
		return false
	}
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
}
