package _switch

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	for i := 0; i < 3; i++ {
		switch {
		case i >= 0:
			fmt.Println("a", i)
		case i >= 1:
			fmt.Println("b", i)
		default:
			fmt.Println("default")
		}
	}
}
