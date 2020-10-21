// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 181.

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"

	"gopl.io/ch7/tempconv"
	"reflect"
)

//!+
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
var d = flag.Int64("b", 1, "xx")

func main() {
	flag.Parse()

	fmt.Println(temp, *temp, temp.String(), reflect.TypeOf(temp), reflect.TypeOf(*temp), reflect.TypeOf(temp.String()))
}

//!-
