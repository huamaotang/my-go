package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func main() {
	//scan()
	//bufRead()
	//fmt.Println(bufScan())
	//sayDemo()
	//intSet()
	//Stringer()
	interfaceValue()
}

func interfaceValue() {
	var w io.Writer
	w = os.Stdout
	x := io.Writer(os.Stdout)
	fmt.Println(w, x)
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

