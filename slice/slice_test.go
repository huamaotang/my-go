package slice

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestFor(t *testing.T) {
	s := []int{1,2,3,4,5}
	m := map[int]struct{}{}
	for k, v := range s {
		if _, ok := m[v]; ok {
			continue
		}
		t.Log(k, v)
		m[v] = struct{}{}
		//s = append(s[:k], s[k + 1:]...)
	}
	for k, v := range s {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		t.Log(k, v)
	}
}

func TestValueCopy(t *testing.T) {
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Printf("v: %d, %p s:%p\n", v, &v, &s[k])
	}
}

func TestBug(t *testing.T) {
	s := []int{1, 2, 3, 4}
	sA := s[:1:2]
	sA = append(sA, 1)
	fmt.Printf("s:%v, %p, sA:%v, %p\n", s, s, sA, sA)
}

func TestAppend(t *testing.T) {
	var s = make([]int, 0, 2)
	for i := 0; i < 16; i++ {
		s = append(s, 1)
		fmt.Printf("times: %d, ptr: %p\n", i+1, s)
	}
}

func TestNilSlice(t *testing.T) {
	var s []int
	sA := make([]int, 0)
	fmt.Println(s == nil, sA == nil)
	fmt.Printf("s: %#v, ptr: %p, len: %d, cap: %d  \nsA:%#v, ptr: %p, len: %d, cap: %d \n", s, &s, len(s), cap(s), sA, &sA, len(sA), cap(sA))
}

func TestCreate(t *testing.T) {
	s := make([]int, 2, 3)
	sA := []int{1, 2}

	fmt.Printf("s: %#v, len: %d, cap: %d  \nsA:%#v, len: %d, cap: %d \n", s, len(s), cap(s), sA, len(sA), cap(sA))
}

func printArray(arr [4]int) {
	fmt.Printf("arr: %p %v\n", &arr, arr)
}

func TestArray(t *testing.T) {
	arrA := [4]int{1, 2, 3, 4}
	var arrB [4]int

	arrB = arrA

	fmt.Printf("arrA: %p %v\n", &arrA, arrA)
	fmt.Printf("arrB: %p %v\n", &arrB, arrB)

	printArray(arrA)
}

func TestStruct(t *testing.T) {
	arr1 := [20]int{}
	a := arr1[2:10]
	ptr := unsafe.Pointer(&a[0])
	fmt.Printf("%p\n", a)
	fmt.Println(arr1, a, &a, ptr)

	var p unsafe.Pointer
	var length int = 1
	var s1 = struct {
		addr unsafe.Pointer
		len  int
		cap  int
	}{p, length, length}
	s := *(*[]byte)(unsafe.Pointer(&s1))
	fmt.Printf("%#v", s)
	return
}

func TestCopy(t *testing.T) {
	to := []int{1, 2, 3, 4}
	fm := make([]int, 6)
	n := copy(to, fm)
	fmt.Println(to, fm, n)

	toA := make([]byte, 3)
	fmA := "abcd"
	n = copy(toA, fmA)
	fmt.Println(toA, fmA, n)
}
