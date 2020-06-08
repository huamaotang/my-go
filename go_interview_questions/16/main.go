package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	fmt.Println(len(c), cap(c))

	//
	var m map[string]string
	m = make(map[string]string)
	m["a"] = "1"
	if v, ok := m["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}
}
