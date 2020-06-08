package insert

import "testing"

func TestSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 2, 1}
	insertSort(arr)
	t.Log(arr)
}

func insertSort(arr []int) {
	// 1、默认第一个元素为有序序列
	// 2、有序序列从后往前推，从第二个元素开始，将之后无序元素一个个加入到有序序列当中
	// 3、无序元素加入有序序列的方法：将无序元素赋值一个变量；将无序元素从后往前比较，当无序元素小于有序序列中的某一个位置的值时，有序序列依次不断后移，直到前面无数据或者无序元素大于有序元素的某个值，记下索引值；
	// 最后，将无序元素插入相应的索引值位置
	var j int
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		for j = i; j > 0 && tmp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}
