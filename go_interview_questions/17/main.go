package main

import "fmt"

func increaseA(m int) int {
	// n = 0
	var n int
	// res = n, res = 0
	defer func() {
		n = m + 1
	}()
	// res = 0, n = 1
	// return
	return n
}

func increaseB(m int) (r int) {
	var n int
	// n = 0, r = n = 0
	defer func() {
		n = m + 1
	}()
	// n = 1, r = 0
	// return
	return n
}

func increaseC(m int) (r int) {
	// r = 0, r = r = 0
	defer func() {
		r = m + 1
	}()
	// r = 1
	// return
	return r
}

func main() {
	// return 并不是原子操作，而是先赋值，再返回值
	fmt.Println(increaseA(1))
	fmt.Println(increaseB(1))
	fmt.Println(increaseC(1))
}
