package main

import (
	"fmt"
	"my_project/study_everyday/20191230/go_interview_questions/common"
)

// init() 知识点
// 1. init()函数是用于程序执行前做包初始化的函数，const、var之后执行，main()之前执行
// 2. 一个包可以包含多个init()函数，一个源文件也可以包含多个init()函数
// 3. 同一个包中多个init()函数的执行顺序没有定义；但是不同包的init函数是根据包导入的依赖关系决定的
// 4. init()函数不能被显示调用、不能被引用（赋值给变量），否则编译出错
// 5. 一个包被引用多次，如A import B，C import B，A import C，B被引用2次，但B包的init()函数只会执行一次
// 6. 引入包不可出现死循环
func init() {
	fmt.Println("main 8")
}

func hello() interface{} {
	return nil
}

func main() {
	// 程序编译时，先执行依赖包的init函数，再执行main包内的init函数
	fmt.Println(common.Get())

	x := hello
	if x == nil {
		fmt.Println("x eq nil")
	} else {
		fmt.Println("x neq nil", x())
	}

}
