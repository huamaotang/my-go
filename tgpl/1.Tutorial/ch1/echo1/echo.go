// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package echo

import (
	"fmt"
	"os"
	"strings"
)

func echo() {
	fmt.Println(plusSign(os.Args[1:]))
	fmt.Println(join(os.Args[1:]))
}

func plusSign(arr []string) string {
	var s, sep string
	for i := 0; i < len(arr); i++ {
		s += sep + arr[i]
		sep = " "
	}
	return s
}

func join(arr []string) string {
	return strings.Join(arr, " ")
}

//!-
