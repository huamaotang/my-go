package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Println(math.Exp(1) , math.IsNaN(x/x), math.IsInf(-1/x, -1))
}
