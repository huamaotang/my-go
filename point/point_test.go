package point

import "testing"

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