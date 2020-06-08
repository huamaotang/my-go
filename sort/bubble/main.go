package main

import "fmt"

func main() {
	// bubble sort
	s := []int32{6,5,4,4,3,1}
	s1 :=bubbleSort(s)
	fmt.Printf("src: %v\ndst: %v", s, s1)
}

func bubbleSort(s []int32) []int32 {
	length := len(s)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if s[j]>s[j+1] {
				s[j], s[j+1] = s[j+1],s[j]
			}
		}
	}
	return s
}