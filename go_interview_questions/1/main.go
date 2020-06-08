package main

import "fmt"

// defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic。
func main() {
	fmt.Println(deferCall(false))
	//fmt.Println(deferCall(false))
}

// defer在函数结束时，按照后进先出的顺序执行；
// panic为内置函数，用来终止当前的控制流，当deferCall调用panic时，函数将停止执行后续的普通语句，但是之前的defer函数调用仍然会正常执行
// 然后再返回到deferCall函数的调用者
func deferCall(b bool) string {
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()
	defer func() {
		fmt.Println("defer 3")
	}()
	if b {
		fmt.Println("defer call ok")
		return "ok"
	} else {
		//fmt.Println("defer call fail")
		panic("触发异常")
	}
}
