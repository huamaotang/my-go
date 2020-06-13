<!-- TOC -->

- [经典排序算法](#经典排序算法)
    - [概述](#概述)
        - [排序分类](#排序分类)
        - [算法复杂度](#算法复杂度)
        - [相关概念](#相关概念)
        - [描述](#描述)
    - [bubble sort](#bubble-sort)
        - [算法描述](#算法描述)
        - [动图演示](#动图演示)
        - [代码实现](#代码实现)
    - [select sort](#select-sort)
        - [算法描述](#算法描述-1)
        - [动图演示](#动图演示-1)
        - [代码实现](#代码实现-1)
    - [insert sort](#insert-sort)
        - [算法描述](#算法描述-2)
        - [动图演示](#动图演示-2)
        - [代码实现](#代码实现-2)
    - [shell sort](#shell-sort)
        - [算法描述](#算法描述-3)
        - [动图演示](#动图演示-3)
        - [代码实现](#代码实现-3)
    - [merge sort](#merge-sort)
        - [算法描述](#算法描述-4)
        - [动图演示](#动图演示-4)
        - [代码实现](#代码实现-4)
    - [quick sort](#quick-sort)
        - [算法描述](#算法描述-5)
        - [动图演示](#动图演示-5)
        - [代码实现](#代码实现-5)
    - [heap sort](#heap-sort)
        - [算法描述](#算法描述-6)
        - [动图演示](#动图演示-6)
        - [代码实现](#代码实现-6)
    - [counting sort](#counting-sort)
        - [算法描述](#算法描述-7)
        - [动图演示](#动图演示-7)
        - [代码实现](#代码实现-7)
    - [bucket sort](#bucket-sort)
        - [算法描述](#算法描述-8)
        - [动图演示](#动图演示-8)
        - [代码实现](#代码实现-8)
    - [radix sort](#radix-sort)
        - [算法描述](#算法描述-9)
        - [动图演示](#动图演示-9)
        - [代码实现](#代码实现-9)

<!-- /TOC -->
# 经典排序算法

## 概述
### 排序分类
- 比较类排序：通过比较来决定元素之间的相对次序；因其时间复杂度不能突破O(nlogn)，被称为非线性比较类排序
- 非比较类排序：无需比较元素之间的大小；可以突破比较类的时间下限，以线性时间运行，被称为线性非比较类排序
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/sort-classify.jpg)

### 算法复杂度
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/sort-1.png)
 
### 相关概念
- 稳定：2个相同元素a、b，a在b前，排序之后，a仍在b前
- 不稳定：2个相同元素a、b，a在b前，排序之后，a可能在b之后
- 时间复杂度：排序完成所花的总操作次数。n变化时，操作次数呈现的规律
- 空间复杂度：算法执行时，所需存储空间的度量

###描述
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/sort-2.png)

## bubble sort

### 算法描述
- 比较相邻两个元素：如果第一个比第二个大，则交换他们
- 对每一对相邻的元素进行比较，从开始第一对到最后一对，最后的元素将会是最大值（有序区第一个元素有了）
- 针对无序区做相同比较、交换重复操作（不断新增有序区元素，直到无序区没有元素了）

### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/bubble-sort.gif)

### 代码实现

```
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

```

## select sort

### 算法描述
- 假定第i个元素为无序区最小的，与无序区每个元素比较，找出真正最小值的索引min，则min才是无序区最小元素，交换索引i和min的值
- 第i+1个元素之前此时为有序区
- 不断累加元素至有序区，直到无序区没有元素了

### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/select-sort.gif)

### 代码实现

```
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

```

## insert sort

### 算法描述
- 假定第一个元素为有序区第一个元素
- 取下一个元素A，有序区从后往前扫描每个元素
- 有序区元素若大于无序区元素A，将有序区元素移到下一个位置
- 重复步骤3，直到找到有序区某一元素小于等于元素A的位置
- 将元素A插入到该位置
- 重复2~5

### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/insert-sort.gif)

### 代码实现

```
func insertSort(arr []int) {
	var j int
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		for j = i; j > 0 && tmp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}

```

## shell sort

### 算法描述
- 定义增量的值，序列不断除以增量，直到增量值等于1
- 根据增量的值进行插入排序
- 最后肯定等于1，相当于最后进行一次普通插入排序

### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/shell-sort.gif)

### 代码实现

```
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

```

## merge sort

### 算法描述
- 元素长度为n，若n只有一个元素，则返回这个元素；否则，把长度为n的无序序列分成两个长度为n/2的子序列left、right
- 对这两个无序序列再递归采用步骤1
- 合并并排序left、right，得到一个有序序列

### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/merge-sort.gif)

### 代码实现

```
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
```

## quick sort

### 算法描述
- 假定第一个值索引为中心点pivot，依次与之后的每个元素进行比较，如果中心点之后的元素大于中心值，则不做元素交换
- 否则，交换中心点之后的第一个元素和当前正在比较的元素（有可能两个元素在相同的位置，比如[5 4 3 2 1]）；另起一个变量，从中心点开始，每交换一次加1，直到没有元素了；吐出这个每次加1的变量，也就是中心点
- 将序列按照中心点分成两个序列，对这两个序列（排除掉中心点元素）分别进行分区，不断进行这个操作
- 需要交换的元素越来越少，直到最后


### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/quick-sort.gif)

### 代码实现

```
func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	fmt.Printf("start: %d, end: %d, arr: %v \n", start, end, arr)
	pivot := partition(arr, start, end)
	fmt.Printf("arr: %v, mid: %d\n", arr, pivot)
	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)
}

func partition(arr []int, start, end int) int {
	pivotV := arr[start]
	j := start
	for i := start + 1; i <= end; i++ {
		if arr[i] < pivotV {
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j++
		}
	}
	arr[start], arr[j] = arr[j], arr[start]
	return j
}
```

打印结果（假定arr := []int{4, 5, 6, 1, 2, 3}）：

```
start: 0, end: 5, arr: [4 5 6 1 2 3] 
arr: [3 1 2 4 6 5], mid: 3
start: 0, end: 2, arr: [3 1 2 4 6 5] 
arr: [2 1 3 4 6 5], mid: 2
start: 0, end: 1, arr: [2 1 3 4 6 5] 
arr: [1 2 3 4 6 5], mid: 1
start: 4, end: 5, arr: [1 2 3 4 6 5] 
arr: [1 2 3 4 5 6], mid: 5
```

## heap sort

### 算法描述
- 写一个依据非叶子节点，将其调整成该非叶子节点下是大顶堆
	- 比较节点的左右叶子节点大小，记下更大的
	- 更大的叶子节点与节点比较，如果节点更大，则这小段树枝（即一个节点和最多2个子节点）无需做调整，因为叶子节点以下子节点已经是大顶堆了
	- 如果节点更小，则把叶子节点值赋给节点，并且把叶子节点位置赋给节点
	- 循环下一个叶子节点，重复2~3操作
	- 将节点值赋给新的节点
	- 大顶堆就造好了
- 从最后一个非叶子节点（len(arr)/2 -1）开始造大顶堆
- 交换第一个和最后一个元素，最后一个就是最大的元素
- 对还没确定大小的元素再次进行大顶堆构建（交换之后，只有堆顶是需要调整的，其它的节点已经符合大顶堆的结构）
- 重复3~4步骤	


### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/heap-sort.gif)

### 代码实现

```
func heapSort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeap(arr, i, len(arr))
	}
	for i := 0; i < len(arr)-1; i++ {
		lastIndex := len(arr) - 1 - i
		arr[0], arr[lastIndex] = arr[lastIndex], arr[0]
		maxHeap(arr, 0, lastIndex)
	}
}

func maxHeap(arr []int, i, len int) {
	nodeV := arr[i]
	for j := 2*i + 1; j < len; j = 2*j + 1 {
		if j+1 < len && arr[j+1] > arr[j] {
			j++
		}
		if arr[j] > nodeV {
			arr[i] = arr[j]
			i = j
		} else {
			break
		}
	}
	arr[i] = nodeV
}
```

## counting sort

### 算法描述
- 找出序列中最大值maxV
- 创建maxV长度的切片，把元素放入切片的key中，并不断的累加1
- 最后，将切片循环出来，得到有序序列


### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/counting-sort.gif)

### 代码实现

```
func countingSort(arr []int) {
	var maxV int
	for _, v := range arr {
		if v > maxV {
			maxV = v
		}
	}
	res := make([]int, maxV+1)
	for _, v := range arr {
		res[v]++
	}
	var newK int
	for k, v := range res {
		if v < 1 {
			continue
		}
		for i := 0; i < v; i++ {
			arr[newK] = k
			newK++
		}
	}
}
```

## bucket sort

### 算法描述
- 定义桶大小size，计算出序列中最大值max，最小值min
- 创建一个二维序列，长度为(max-min)/size+1
- 把元素按照(v-min)/size作为key，放入桶中
- 循环出二维序列，把桶中元素进行冒泡排序，依次写入原序列中


### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/bucket-sort.gif)

### 代码实现

```
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
```

## radix sort

### 算法描述
- 计算出最大值max，最大的位数（除以10，看能除几次）
- 按照位数循环（先个位，再十位，再百位，再千位），创建一个二维序列，长度为10（十进制，最多0-9）
- 按照每位的值依次存入新的序列中
- 从0开始，依次把元素取出来，重新存入原序列
- 重复2~4，直到最后位，序列已经有序了


### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/radix-sort.gif)

### 代码实现

```
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
```

![Reference1](https://www.cnblogs.com/onepixel/articles/7674659.html)
![Reference2](https://www.runoob.com/w3cnote/ten-sorting-algorithm.html)