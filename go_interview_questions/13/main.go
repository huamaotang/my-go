package main

import "fmt"

func hello(m int) {
	fmt.Println(m)
}

func main() {
	// 打印结果：11 31
	// defer 先进后出；调用函数时，会保存一份副本，如果defer语句中带有变量作为实参，实参变量也被保存在副本中，实参值为11，所以执行时，形参m值为11；
	// 第二次执行defer，m未作为实参传入，m为当前值31
	m := 1
	defer func() {
		fmt.Println(m)
	}()
	m += 10
	defer hello(m)
	m += 20

}
