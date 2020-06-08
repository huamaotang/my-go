package main

import "fmt"

const (
	A = iota + 1
	B
	C

	D = iota
	F
)

func main() {
	fmt.Println(A, B, C, D, F)
}
