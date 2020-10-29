package main

import (
	"fmt"
	"strconv"
)

type Ner interface {
	a()
	b(int)
	c(int) string
}

type N struct {}

func (n *N) a()  {
	fmt.Println("a")
}

func (n *N) b(i int)  {
	fmt.Println("b", i)
}

func (n *N) c(i int) string  {
	fmt.Println("c", i)
	return "c_" + strconv.Itoa(i)
}

func main() {
	var n Ner
	n = &N{}
	n.a()
	n.b(1)
	fmt.Println(n.c(2))
}
