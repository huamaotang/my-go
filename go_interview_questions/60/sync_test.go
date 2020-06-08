package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNil(t *testing.T) {
	// nil不是关键字
	nil := 1
	fmt.Println(nil)
}

func TestForArray(t *testing.T) {
	for k := range (*[3]int)(nil) {
		fmt.Println(k)
	}
}

type data2 struct {
	// sync.Mutex作为结构体成员变量，相关方法必须使用指针接收者，否则会导致锁机制失效
	// 或者以*sync.Mutex作为结构体成员变量
	m sync.Mutex
}

// 如果接受者为值类型，锁机制失效，必须为指针类型
func (d data2) say2(s string) {
	d.m.Lock()
	defer d.m.Unlock()

	fmt.Println(s)
	time.Sleep(time.Second * 5)
}

func TestSay2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	d := data2{m: sync.Mutex{}}

	go func() {
		defer wg.Done()
		d.say2("read")
	}()

	go func() {
		defer wg.Done()
		d.say2("write")
	}()

	wg.Wait()
}

type data1 struct {
	*sync.Mutex
}

// 如果接受者为值类型，锁机制失效，必须为指针类型
func (d data1) say1(s string) {
	d.Lock()
	defer d.Unlock()

	fmt.Println(s)
	time.Sleep(time.Second * 5)
}

func TestSay1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	d := data1{&sync.Mutex{}}

	go func() {
		defer wg.Done()
		d.say1("read")
	}()

	go func() {
		defer wg.Done()
		d.say1("write")
	}()

	wg.Wait()
}

//
type data struct {
	sync.Mutex
}

func (d *data) say(s string) {
	d.Lock()
	defer d.Unlock()

	fmt.Println(s)
	time.Sleep(time.Second * 5)
}

func TestSync(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	var d data

	go func() {
		defer wg.Done()
		d.say("read")
	}()

	go func() {
		defer wg.Done()
		d.say("write")
	}()

	wg.Wait()
}
