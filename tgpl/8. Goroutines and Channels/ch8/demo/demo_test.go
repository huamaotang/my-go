package demo

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch := make(chan int, 2)

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("send", x)
		case ch <- i:
			fmt.Println("receive", i)
		default:
			fmt.Println("default")
		}
	}
}


func TestXSelect(t *testing.T) {
	ch := make(chan int, 2)

	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			fmt.Println("receive", i)
		default:
			fmt.Println("default")
		}
	}
}

func TestNil(t *testing.T) {
	var cht chan int
	cht <- 1
}

func TestReadDir(t *testing.T) {
	fmt.Println(ioutil.ReadDir("/Users/mac/go/src/my-go/"))
}

func TestBreak(t *testing.T) {
	x := 10
	l:
		for {
			x--
			if x <= 0 {
				fmt.Println(x)
				break l
			}
			fmt.Println(x)
		}
	fmt.Println(x)
}

func TestBreakX(t *testing.T) {
	x := 10
	for {
		x--
		if x <= 0 {
			fmt.Println(x)
			break
		}
		fmt.Println(x)
	}
	fmt.Println(x)
}

func TestClose(t *testing.T) {
	//ch := make(chan struct{})

}

func TestStdin(t *testing.T) {
	x, err := os.Stdin.Read(make([]byte, 1))
	fmt.Println(x, err)
}

func TestTime(t *testing.T) {
	select {
	case ch, ok := <- time.After(10 * time.Second):
		fmt.Println(ch, ok)
	}
}
