package quick

import (
	"fmt"
	"testing"
)

func TestQuick(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	quickSort(arr, 0, len(arr)-1)
	t.Log(arr)
}

/*
理解：
1、取其中一个元素作为分区点pivot，遍历无序列表中其它数据，将比pivot小的放在左边，大的放在右边，重排无序序列并计算分区点
2、分别递归除分区点的左边、右边无序列表，重排无序序列并计算分区点，直至左右边只剩一个元素，元素也就排序好了

将比pivot小的放在左边，大的放在右边：
1、从i=start+1索引开始，如果值小于分区点，则j=start,j+1与i的值互换，j++
2、将第一个，也就是基准值与第j个元素互换，因为第j个值肯定小于等于第一个值，并返回j
*/

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := partition(arr, start, end)
	fmt.Println(arr, mid)
	return
	quickSort(arr, start, mid-1)
	quickSort(arr, mid+1, end)
}

// 一路快排，还有二路、三路快排
func partition(arr []int, start, end int) int {
	v := arr[start]

	j := start
	for i := start + 1; i <= end; i++ {
		if arr[i] < v {
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j++
		}
	}
	arr[start], arr[j] = arr[j], arr[start]
	return j
}
