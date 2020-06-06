<!-- TOC -->

- [Go中Slice底层实现](#go中slice底层实现)
    - [key point](#key-point)
    - [切片和数组（怎么选择）](#切片和数组怎么选择)
    - [切片的数据结构](#切片的数据结构)
    - [创建切片](#创建切片)
    - [切片扩容](#切片扩容)
    - [切片拷贝](#切片拷贝)

<!-- /TOC -->
# Go中Slice底层实现


## key point
- 切片是Go中的一种数据结构
- 切片本身不是动态数组或者数组指针
- 常见的slice操作：append、copy、reslice

## 切片和数组（怎么选择）
- Go数组是值类型，赋值和函数传参操作都会复制整个数组数据，因此，如果数据量比较大会消耗更多的内存，效率很低。
- 用数组指针、切片效率更高，只进行地址拷贝，避免值拷贝
- 数组指针统一指向**原数组**，**原数组**指针的指向被更改，函数内的指针指向都会更改
- 切片可以更好的处理共享内存的问题，指针地址与数组指针不是同一个

```
func printArray(arr [4]int) {
	fmt.Printf("arr: %p %v\n", &arr, arr)
}
	
func TestArray(t *testing.T) {
	arrA := [4]int{1, 2, 3, 4}
	var arrB [4]int
	
	arrB = arrA
	
	fmt.Printf("arrA: %p %v\n", &arrA, arrA)
	fmt.Printf("arrB: %p %v\n", &arrB, arrB)
	
	printArray(arrA)
}
```

打印结果：

```
arrA: 0xc0000ac040 [1 2 3 4]
arrB: 0xc0000ac060 [1 2 3 4]
arr: 0xc0000ac0a0 [1 2 3 4]
```

- 并非所有时候都适合用切片，底层数组可能在堆上分配内存，或是小数组在栈上拷贝的消耗未必比make消耗大

## 切片的数据结构

- 本身不是动态数组或数组指针。内部实现的数据结构：通过指针引用底层数组
- 一个只读对象
- 工作机制：类似数组指针的一种封装
- 对数组连续片段的引用，由起始和终止索引标识的一些子集，终止索引标识的项不包括在切片里

```
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

## 创建切片
- make和切片字面量

```
s := make([]int, 2, 3)
sA := []int{1, 2}

fmt.Printf("s: %#v, len: %d, cap: %d  \nsA:%#v, len: %d, cap: %d \n", s, len(s), cap(s), sA, len(sA), cap(sA))
```
打印结果：

```
s: []int{0, 0}, len: 2, cap: 3  
sA:[]int{1, 2}, len: 2, cap: 2 
```

- nil和空切片
	- nil切片表示一个不存在的切片，指向的地址为nil
	- 空切片表示一个已经分配内存地址的切片，空集合
	- 区别在于是否创建内存地址
	- 对于append的结果，nil和空切片变量无区别

```
var s []int
sA := make([]int, 0)
fmt.Println(s == nil, sA == nil)
fmt.Printf("s: %#v, ptr: %p, len: %d, cap: %d  \nsA:%#v, ptr: %p, len: %d, cap: %d \n", s, &s, len(s), cap(s), sA, &sA, len(sA), cap(sA))
```

打印结果

```
true false
s: []int(nil), ptr: 0xc000092080, len: 0, cap: 0  
sA:[]int{}, ptr: 0xc0000920a0, len: 0, cap: 0 
```	
## 切片扩容
- 策略
	- 如果cap足够，不重新创建新的底层数组
	- 如果cap容量不足，cap呈2倍的速度增长
	- 判断切片的长度是否超过1024，超过呈1/4数组增长

```
var s = make([]int, 0, 2)
for i := 0; i < 16; i++ {
	s = append(s, 1)
	fmt.Printf("times: %d, ptr: %p\n", i+1, s)
}
```
	
打印结果：

```
times: 1, ptr: 0xc0000801b0
times: 2, ptr: 0xc0000801b0
times: 3, ptr: 0xc0000ac040
times: 4, ptr: 0xc0000ac040
times: 5, ptr: 0xc0000c8000
times: 6, ptr: 0xc0000c8000
times: 7, ptr: 0xc0000c8000
times: 8, ptr: 0xc0000c8000
times: 9, ptr: 0xc0000ce000
times: 10, ptr: 0xc0000ce000
times: 11, ptr: 0xc0000ce000
times: 12, ptr: 0xc0000ce000
times: 13, ptr: 0xc0000ce000
times: 14, ptr: 0xc0000ce000
times: 15, ptr: 0xc0000ce000
times: 16, ptr: 0xc0000ce000

```	
- 一个可能存在的bug
	- 共有的底层数组被引用的某一个切片更改，导致所有引用该数组的切片都被更改

```
s := []int{1, 2, 3, 4}
sA := s[:1:2]
sA = append(sA, 1)
fmt.Printf("s:%v, %p, sA:%v, %p\n", s, s, sA, sA)

```
打印结果：

```
s:[1 1 3 4], 0xc0000a4040, sA:[1 1], 0xc0000a4040

```

## 切片拷贝
- 取决于to、fm中更短的切片


```
to := []int{1, 2, 3, 4}
fm := make([]int, 6)
n := copy(to, fm)
fmt.Println(to, fm, n)
	
toA := make([]byte, 3)
fmA := "abcd"
n = copy(toA, fmA)
fmt.Println(toA, fmA, n)

```

打印结果：

```
[0 0 0 0] [0 0 0 0 0 0] 4
[97 98 99] abcd 3

```

- **for**一个切片，v为值拷贝，且第一次声明变量，后续仅进行值拷贝

```
s := []int{1, 2, 3}
for k, v := range s {
	fmt.Printf("v: %d, %p s:%p\n", v, &v, &s[k])
}

```

打印结果

```
v: 1, 0xc0000821c8 s:0xc0000ac060
v: 2, 0xc0000821c8 s:0xc0000ac068
v: 3, 0xc0000821c8 s:0xc0000ac070
```

[Reference](https://halfrost.com/go_slice/)