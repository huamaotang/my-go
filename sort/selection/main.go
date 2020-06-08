package main

import "fmt"

func main() {
	s := []int32{6, 5, 4, 4, 3, 1, 7}
	selectionSort(s)
	fmt.Printf("%v", s)
}

func selectionSort(s []int32) {
	// 知识点：一层循环初始化一个变量min，意为最小值的索引值，假设第一个值为最小值；二层循环最小值与其他元素逐一比较，更小的赋值给变量min，得到最小值；
	// 1、从无序数组取出最小值 2、排除掉已经得到的最小值，从无序数组取出最小值直到结束
	length := len(s)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}
}
