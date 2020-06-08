package main

import (
	"fmt"
	"reflect"
)

func main() {
	// https://golang.org/ref/spec#Comparison_operators
	// 结构体的比较
	// 1.结构体只能比较是否相等，不能比较大小
	// 2.相同类型的结构体才能进行比较，结构体是否相同，不仅需要成员变量名称一致，还需要成员变量的顺序一致
	// 3.如果struct的所有成员变量都是可比较的，则可以通过==或!=进行比较。比较时逐个项进行比较，如果每一项都相等，则两个结构体相等，否则不相等
	sn1 := struct {
		name string
		age  int
	}{name: "t1", age: 10}
	sn2 := struct {
		name string
		age  int
	}{name: "t1", age: 10}
	// 结构体成员变量的顺序不一致，也不能比较
	//sn3 := struct {
	//	age  int
	//	name string
	//}{name: "t1", age: 10}
	//fmt.Println(sn1 == sn3)

	fmt.Println(sn1 == sn2)

	sm1 := struct {
		m   map[string]string
		age int
	}{m: map[string]string{"name": "t1"}, age: 10}
	sm2 := struct {
		m   map[string]string
		age int
	}{m: map[string]string{"name": "t2"}, age: 10}
	fmt.Println(sm1.age == sm2.age)
	fmt.Println(reflect.DeepEqual(sm1.m, sm2.m))
}
