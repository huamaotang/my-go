package _map

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	// 并不是完全随机，而是不指定。。。
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	var name string
	for _, v := range m {
		name = v
		break
	}
	fmt.Println(name)
}

func TestMapEmpty(t *testing.T) {
	var a map[int]string

	if m, ok := a[0]; !ok {
		t.Log("a 1 not ok")
	} else {
		t.Log(m)
	}
	if a == nil {
		t.Log("a is nil")
	} else {
		t.Log("a is not nil")
	}

	b := make(map[int]string)
	if m, ok := a[0]; !ok {
		t.Log("a 1 not ok")
	} else {
		t.Log(m)
	}
	if b == nil {
		t.Log("b is nil")
	} else {
		t.Log("b is not nil")
	}
	t.Log(len(a), len(b))
}
