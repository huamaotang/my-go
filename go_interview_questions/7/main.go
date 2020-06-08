package main

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main() {
	// 字符串连接 https://segmentfault.com/a/1190000012978989
	// 1.在已有字符串数组的场合，使用strings.Join()能有比较好的性能
	// 2.在一些性能要求较高的场合，尽量使用buffer.WriteString()获得更好的性能
	s := "abc" + "123"
	s = fmt.Sprintf("%s%s", "abc", "123")
	s = strings.Join([]string{"abc", "123"}, "")
	var buf bytes.Buffer
	buf.WriteString("abc")
	buf.WriteString("123")
	s = buf.String()
	fmt.Println(s)
	// 0 2 zz zz 5
	fmt.Println(x, y, z, k, p)

	// nil只能赋值给指针、chan、func、interface、map、slice，数组也不行；error是一种内置的接口类型
	//var y [1]int
	//y[0] = 1
	//y = nil
	//var x0 = nil
	var x1 interface{} = nil
	//var x2 string = nil
	var x3 error = nil
	fmt.Println(x1, x3)
}
