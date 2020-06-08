package merge

import "testing"

func TestMerge(t *testing.T) {
	arr := []int{6, 5, 4, 0, 3, 2, 1}
	t.Log(mergeSort(arr))
}

/* 理解：
1、先折半拆分无序序列至每个序列只有一个元素
2、合并左右已经分别有序的序列，经过不断比较左右两边的值，返回一个有序序列，直到结束
*/

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}
