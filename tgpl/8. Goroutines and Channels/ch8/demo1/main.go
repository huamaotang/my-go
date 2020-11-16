package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	selectTwo()
}

func selectTwo() {
	abort := make(chan int)
	select {
	case <-abort:
		fmt.Printf("Launch aborted!\n")
		return
	default:
		// do nothing
	}
}

func selectOne() {
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}

func existBuffer() {
	ch := make(chan int, 1)
	test(ch)
	x := <-ch
	fmt.Println(x)
}

func test(ch chan<- int) {
	time.Sleep(10*time.Second)
	ch <- 1
}

func noBuffer() {
	sayCh := make(chan string)
	nameCh := make(chan string)

	go say(sayCh)
	go name(sayCh, nameCh)
	result(nameCh)
}

func say(sayCh chan<- string) {
	for i := 0; i < 10; i++ {
		sayCh <- "say" + strconv.Itoa(i)
	}
	close(sayCh)
}

func name(ch <-chan string, nameCh chan<- string)  {
	for v := range ch {
		nameCh <- v + "tanghuamao"
	}
	close(nameCh)
}

func result(ch <-chan string)  {
	for v := range ch {
		fmt.Println(v)
	}
}