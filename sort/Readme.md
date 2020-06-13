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
- a

### 动图演示
![排序分类](https://raw.githubusercontent.com/huamaotang/my-images/master/quick-sort.gif)

### 代码实现

```
func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(arr, start, end)
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


