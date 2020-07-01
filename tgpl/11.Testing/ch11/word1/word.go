// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 303.
//!+

// Package word provides utilities for word games.
package word

import "fmt"

// IsPalindrome reports whether s reads the same forward and backward.
// (Our first attempt.)
func IsPalindrome(s string) bool {
	for i := range s {
		fmt.Printf("%T\n", s[i])
		fmt.Println(s[i], s[len(s)-1-i], s[i] == s[len(s)-1-i])
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

//!-
