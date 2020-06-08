package heap

import (
	"fmt"
	"testing"
)

/**
概念：
1、完全二叉树：二叉树中除了最后一层外，其它各层的结点数都达到了最大的个数；且最后一层页结点按照从左往右的顺序连续存在，
只缺最后一层右侧若干结点
2、堆：具有完全二叉树性质，且分为两种，大顶堆、小顶堆
3、大顶堆：每个结点的值都大于或等于孩子结点的值；小顶堆：每个结点的值都小于或等于孩子结点的值

堆排序：将待排序序列构造成一个大顶堆，此时，序列的最大值就是堆顶的根结点。将其与末尾元素交换，末尾就是最大值。
然后将剩余n-1个元素重新构造成一个堆，这样将得到n个元素的次小值。如此反复执行，便能得到一个有序序列

*/

func TestHeapSort(t *testing.T) {
	arr := []int{1,2,3,9,5,6,7,8}
	length := len(arr)

	for i := length/2 - 1; i >= 0; i-- {
		heapAdjust(arr, i, length)
	}

	for i := 0; i < length-1; i++ {
		j := length - 1 - i
		arr[0], arr[j] = arr[j], arr[0]
		heapAdjust(arr, 0, j)
	}
	fmt.Println(arr)
}

func heapAdjust(arr []int, i, len int) {
	temp := arr[i]
	for k := 2*i + 1; k < len; k = 2*k + 1 {
		if k + 1 < len && arr[k] < arr[k+1] {
			k++
		}
		if arr[k] > temp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = temp
}
