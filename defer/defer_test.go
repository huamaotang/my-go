package _defer

import (
	"errors"
	"fmt"
	"testing"
)

func set() error {
	err := errors.New("err 1")
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	return errors.New("err 2")
}

func set1() error {
	err := errors.New("err 1")
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	err = errors.New("err 2")
	return nil
}

func get() (err error) {
	// return执行顺序：给err赋值；执行defer函数；执行return操作
	err = errors.New("err 1")
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	return errors.New("err 2")
}

func TestDeferErr(t *testing.T) {
	t.Log(set())
	t.Log(set1())
	t.Log(get())
}
