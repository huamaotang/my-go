package main

import "fmt"

func main() {
	m := make(map[int]*int)
	// for range中的短变量声明，只会声明一次，后面只是赋值。因为m变量的值为指针值，即&v，用的是同一个地址，所以最后v的值为3。
	// 因此所有的值都为最后一次循环的值，即为3
	for k, v := range []int{0, 1, 2, 3} {
		fmt.Println(&k, &v)
		nv := v
		fmt.Println(&nv)
		m[k] = &nv
		//m[k] = &v
	}
	// 输出结果，map无序；
	for k, v := range m {
		fmt.Printf("k:%d v:%d\n", k, *v)
	}
}
