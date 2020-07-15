package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1

	go func() {
		a = 2
	}()
	a = 3
	time.Sleep(time.Second*1)
	fmt.Println(a)
}
