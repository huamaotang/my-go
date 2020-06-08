package shell

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{6, 5, 4, 3, 2, 1}
	shellSort(arr)
	t.Log(arr)
}

func shellSort(arr []int) {
	/**
	理解：无序序列转成有序序列；3次循环
	1、增量递减
	2、从增量第一个值开始至最后一个值循环
	3、循环有序序列，将无序序列中的元素插入有序序列
	*/
	n := len(arr)
	for step := n / 2; step > 0; step /= 2 {
		var j int
		for i := step; i < n; i++ {
			temp := arr[i]
			for j = i; j-step >= 0 && temp < arr[j-step]; j -= step {
				arr[j] = arr[j-step]
			}
			arr[j] = temp
			fmt.Println(step, i, j, arr)
		}
	}
}
