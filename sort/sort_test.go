package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, 30)
	for k := range arr {
		arr[k] = rand.Intn(100)
	}
	fmt.Println(arr)
	//shellSort(arr)
	fmt.Println(mergeSort(arr))
}

func radixSort(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	var max int
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	var maxDigit int
	for max > 0 {
		max /= 10
		maxDigit++
	}
	var divisor = 1
	for i := 0; i < maxDigit; i++ {
		bucket := make([][]int, 10)
		for _, v := range arr {
			digit := (v / divisor) % 10
			bucket[digit] = append(bucket[digit], v)
		}

		newK := 0
		for j := 0; j < 10; j++ {
			if len(bucket[j]) == 0 {
				continue
			}
			for _, v := range bucket[j] {
				arr[newK] = v
				newK++
			}
		}
		divisor *= 10
	}
}

func bucketSort(arr []int, size int) {
	var min, max int
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	var bucketArr = make([][]int, (max-min)/size+1)
	for _, v := range arr {
		bucketNo := (v - min) / size
		bucketArr[bucketNo] = append(bucketArr[bucketNo], v)
	}

	var key int
	for _, bucketV := range bucketArr {
		bubbleSort(bucketV)
		for _, v := range bucketV {
			arr[key] = v
			key++
		}
	}
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var res []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		res = append(res, left...)
	} else if len(right) > 0 {
		res = append(res, right...)
	}
	return res
}

func shellSort(arr []int) {
	// 最大增量；2为希尔增量
	maxIncr := 3
	var j int
	length := len(arr)
	// 缩小增量排序；
	for incr := length / maxIncr; incr > 0; incr /= maxIncr {
		for i := incr; i < length; i++ {
			temp := arr[i]
			for j = i; j-incr >= 0 && temp < arr[j-incr]; j -= incr {
				arr[j] = arr[j-incr]
			}
			arr[j] = temp
		}
	}
}

func InsertSort(arr []int) {
	var j int
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		for j = i; j > 0 && tmp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}

func SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
