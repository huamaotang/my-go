package main

import "fmt"

type Config struct {
	Name string
}

func main() {
	m := map[int32]*Config{
		0: &Config{Name: "a"},
	}
	m[0].Name = "b"
	fmt.Println(m[0].Name)
}
