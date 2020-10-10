package main

import (
	"fmt"
)

type A struct {
	X, Y int
}

type B struct {
	M, N int
}

type C struct {
	A
	B
}

func main() {
	fmt.Println("slice")

	a := []int{1}
	b := append(a, 1)
	fmt.Println(a, cap(a), b, cap(b))

	c := 1
	d := c
	c = 2
	fmt.Println(c, d)

	e := []int{1}
	f := e
	fmt.Printf("%p, %p\n", e, f)
	e = []int{2}
	fmt.Printf("%p, %p\n", e, f)

	g := e[:0]
	fmt.Println(len(g), cap(g), g)

	h := map[int]string{}
	delete(h, 1)

	i := fmt.Sprintf("%q", []string{"a", "b"})
	fmt.Printf("%#v, %[1]T\n", i)

	j := &C{}
	j.X = 1
	j.M = 2
	fmt.Println(j)
}
