package point

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	type T struct {
		Value int32
	}
	m := make(map[int]*int32)
	n := make(map[int]*T)
	for k, v := range []int32{1,3,5}{
		tmp := &T{Value:v}
		n[k] = tmp
		m[k] = &v
	}
	t.Log(m, n)
}

func Test2(t *testing.T) {
	for i := 0; i < 10; i++ {
		j := i
		f2(&j)
	}
}

func f2(x *int) {
	fmt.Println(x)
}