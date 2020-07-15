package main

import "fmt"

func main() {
	ch := make(chan struct{})

	a := 1

	go func() {
		a = 2
		ch <- struct{}{}
	}()

	<- ch
	fmt.Println(a)
}
