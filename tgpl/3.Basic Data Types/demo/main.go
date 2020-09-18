package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"time"
)

type MY struct {
	name string
}
const timeout = 5*time.Minute
const duration time.Duration = 111
const NameA string = "tang"


func main() {
	var x float64
	fmt.Println(math.Exp(1) , math.IsNaN(x/x), math.IsInf(-1/x, -1))

	name := "汤华茂tang"
	fmt.Println(len(name), string(name[3:]), string(name[2:]))
	for _, v := range name {
		fmt.Println(string(v), v )
	}

	f := strings.Fields("a a bbc. . .")
	fmt.Println(f, len(f), cap(f), len(strings.Fields("   ")))
	fb := bytes.Fields([]byte("aa 汤华 茂 .."))
	fmt.Println(fb, len(fb))

	s := []int{1,2,3}
	var buffer bytes.Buffer
	fmt.Println(buffer)
	buffer.WriteByte('[')
	for i, v := range s {
		if i > 0 {
			buffer.WriteRune('汤')
			buffer.WriteRune('华')
			buffer.WriteString(", ")
		}
		_ , _ = fmt.Fprintf(&buffer, "%d", v)
	}
	buffer.WriteByte(']')
	fmt.Println(buffer.String())

	fmt.Println(NameA)
	NameA := "abc"
	fmt.Println(NameA)
}
