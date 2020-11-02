package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
)
func main() {
	//scan()
	//bufRead()
	//fmt.Println(bufScan())
	//sayDemo()
	//intSet()
	//Stringer()
	//interfaceUse()
	//objCopyToInterface()
	Fprintf()
}

type B struct {
	m string
}

func (b *B) Write(x []byte) (int, error) {
	b.m = string(x)
	return len(b.m), nil
}


func Fprintf() {
	var w io.Writer
	fmt.Printf("%#v\n", w)
	w = &B{}
	fmt.Fprintf(w, "%d-%s", 110, "sss")
	fmt.Printf("%#v\n", w)

}

func pathError() {
	err := errors.New("abc")
	err = os.ErrNotExist
	fmt.Println(os.IsNotExist(err))
}

func typeAssert() {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				fmt.Printf("1 %#v\n", err)
			default:
				fmt.Printf("2 %#v\n", err)
			}
		}
	}()
	var a interface{} = int32(1)
	a, ok := a.(int32)
	fmt.Printf("%#v %#v\n", a, ok)

	var x io.Writer
	x = os.Stdin
	ax, ok := x.(error)
	fmt.Printf("%#v %#v\n", ax, ok)

	var y error
	ay, ok := y.(error)
	fmt.Printf("%#v %#v\n", ay, ok)
}

func NewErr() error {
	return errors.New("errors")
}

type data struct {
	x int
}

func objCopyToInterface() {
	var er interface{}
	er = &data{x:1}
	er.(*data).x = 1
	fmt.Println(er)
}

func interfaceUse() {
	var x A
	x = new(M)
	x.Add()
	fmt.Println(x.Get())
	x = new(N)
	x.Add()
	fmt.Println(x.Get())
}

type A interface {
	Add()
	Get() int
}

type M int
type N int

func (m *M) Add() {
	*m += 1
}

func (m *M) Get() int {
	return int(*m)
}

func (n *N) Add() {
	*n += 2
}

func (n *N) Get() int {
	return int(*n)
}

func interfaceValue() {
	var w io.Writer
	w = os.Stdout
	z, _ := w.Write([]byte("tanghuamao"))
	x := io.Writer(os.Stdout)
	fmt.Printf(" \n %#v\n %#v\n %#v\n", w, x, z)
}

func Stringer() {
	s := IntSet{}
	var n fmt.Stringer = &s
	var m fmt.Stringer = &IntSet{}
	fmt.Println(n.String(), m.String(), )
}

func intSet() {
	a := IntSet{}
	x := (&IntSet{}).String()
	y := a.String()
	fmt.Println(x, y)
	b := &IntSet{}
	m := (*b).Int()
	n := IntSet{}.Int()
	o := (*(&IntSet{})).Int()
	fmt.Println(m, n, o)
}

type IntSet struct {
}

func (i *IntSet) String() string {
	return "string"
}

func (i IntSet) Int() string {
	return "int"
}

type Person interface {
	SayHello(name string) string
}

type Student struct {}

func (s *Student) SayHello(name string) string {
	return "Hello Student, " + name
}

type Teacher struct {}

func (s *Teacher) SayHello(name string) string {
	return "Hello Teacher, " + name
}

func say(p Person, name string) string {
	return p.SayHello(name)
}

func sayDemo() {
	fmt.Printf("%s\n%s", say(&Teacher{}, "Sam"), say(&Student{}, "Tom"))
}

func bufScan() string {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanLines)
	var s string
	for input.Scan() {
		s += input.Text()
	}
	return s
}

func bufRead() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入：")
	input, err := inputReader.ReadString(' ')
	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println(input)
}

func scan() {
	var name string
	var age int
	fmt.Println("请输入。。。")
	_, _ = fmt.Scanln(&name, &age)
	fmt.Printf("name1: %s age1: %d\n", name, age)

	_, _ = fmt.Scanf("%s : %d", &name, &age)
	fmt.Printf("name2: %s age2: %d\n", name, age)
}

