package counting

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{1, 2, 3, 9, -15, 5, 3, 32, 1, 6, 7, 8}
	countingSort(arr)
	fmt.Println(arr)
}

func countingSort(arr []int) {
	var max, min int
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	if min < 0 {
		for k, v := range arr {
			arr[k] = v - min
		}
	}
	var countingArr = make([]int, max-min+1)
	for _, v := range arr {
		countingArr[v] += 1
	}

	var key int
	for k, v := range countingArr {
		for i := 0; i < v; i++ {
			arr[key] = k + min
			key++
		}
	}
	return
}
