package _select

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- "two"
	}()
	// select阻塞，main协程等待case执行，第二个协程1s后往信道写入数据，第二个case执行，输出two
	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	}
}

func TestTwo(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "one"
	}()
	go func() {
		ch2 <- "two"
	}()
	time.Sleep(time.Second * 1)
	// select阻塞，1s后两个协程都已经注备好了，随机选择一个case执行
	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	}
}

func TestThree(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "one"
	}()
	go func() {
		ch2 <- "two"
	}()
	// select阻塞，等待main协程执行case，case将执行先写数据的信道
	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	}
}

func TestFour(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "one"
	}()
	go func() {
		ch2 <- "two"
	}()
	//time.Sleep(1 * time.Second)
	// select阻塞main协程，执行case时，如果两个信道上还没有值，则会执行default语句
	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	default:
		fmt.Println("default")
	}
}

func TestFive(t *testing.T) {
	// nil channel
	var ch chan int

	go func() {
		ch <- 1
	}()
	// 信道为nil，该分支会被忽略，相当于select{}，阻塞main协程；调度协程，在nil的信道上操作，报错；
	select {
	case i := <-ch:
		fmt.Println(i)
	}
}

func TestSix(t *testing.T) {
	// nil channel
	var ch chan int

	go func() {
		ch <- 1
	}()
	// 信道为nil，该分支会被忽略，相当于select{}，阻塞main协程；调度协程，在nil的信道上操作，报错；
	// 使用default分支，可以避免报错
	select {
	case i := <-ch:
		fmt.Println(i)
	default:
		fmt.Println("default")
	}
}

func TestSeven(t *testing.T) {
	var ch chan int
	ch = make(chan int)

	go func() {
		time.Sleep(10 * time.Second)
		ch <- 1
	}()
	// 添加超时时间

	select {
	case i := <-ch:
		fmt.Println(i)
	case <-time.After(2 * time.Second):
		fmt.Println("no case ok")
	}
}

func TestEight(t *testing.T) {
	ch := make(chan string)
	ch1 := make(chan string)

	go func() {
		fmt.Println(<-ch)
	}()
	go func() {
		fmt.Println(<-ch1)
	}()

	// 切记select阻塞main协程，无缓冲信道ch为同步操作，写信道会阻塞，直到信道被读
	select {
	case ch <- "hi":
		fmt.Println("write ch")
	case ch1 <- "hi1":
		fmt.Println("write ch1")
	case <-time.After(1 * time.Second):
		fmt.Println("no case ok")
	}
}

func TestTen(t *testing.T) {
	run(context.Background())
}

func run(ctx context.Context) {
	ch := make(chan int)
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		}
	}()

	select {
	case <- ch:
		fmt.Println("all done")
	case <- time.After(10*time.Second):
		fmt.Println("over 10s")
	}
}
