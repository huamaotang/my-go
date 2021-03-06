# Identifier

## 概念
- 对变量、函数、方法等命名时，使用的字符序列，称为标识符
- 自己可以取名的地方都可以叫标识符

## 标识符的命名规则
- 由26个大小写字母、下划线、数字组成
- 数字不可以开头
- 严格区分大小写
- 字符序列中间不能出现空格
- 下划线本身是一个特殊的标识符，称为` 空标识符 `，可以代表任何其他标识符，但是对应的值会被忽略。只能作为` 占位符 `使用，不能作为*标识符*使用
- 不能使用系统保留的关键字作为标识符

## 标识符的例子

```
hello // ok
hello66 // ok
1hello // error，不能以数字开头
a-b // error，不能使用 - 相当于 减号
a bc // error，不能使用含有空格
a_bc // ok
_abc // ok
int // ok，不推荐使用
float64 // ok，不推荐使用
_ // error
Abc // ok
```

## 注意事项
- 包命名，`package`的名字和当前所在目录名保持一致；要有意义、简短、不要与标准库重名

```
比如目录结构：
deme
-----model // 目录
​ ---------utils.go // 包
那么utils.go 文件中包名
package model
```
- 变量、函数、常量：驼峰法

~~~
var studentName string = “fe_cow”
var goodPrice float64 = “666.66”
~~~		
- 首字母大写是公开，小写是私有；如果是变量名、函数名、常量名首字母大写，可以被其他包访问，首字母是小写，只能在本包中访问

## 关键字、预定义标识符
### 25个保留关键字
<table>
	<tr>
		<td>package</td><td>import</td><td>const</td><td>type</td><td>struct</td>
	</tr>
	<tr>
		<td>var</td><td>chan</td><td>map</td><td>func</td><td> defer </td>
	</tr>
	<tr>
		<td>go</td><td>goto</td><td>interface</td><td>range</td><td>select</td>
	</tr>
	<tr>
		<td>if</td><td>else</td><td>for</td><td>continue</td><td> return </td>
	</tr>
	<tr>
		<td>switch</td>	<td>case</td><td>fallthrough</td><td>default</td><td> break </td>
	</tr>
</table>

### 37个预定义标识符
<table>
	<tr>
		<td>内建常量</td><td>true</td><td>false</td><td>iota</td><td>nil</td><td></td>
	</tr>
	<tr>
		<td rowspan=4>内建类型</td><td>int8</td><td>uint8</td><td>byte</td><td>int16</td><td>uint16</td>
	</tr>
	<tr>
		<td> int32 </td><td>uint32</td><td>rune</td><td>int64</td><td>uint64</td>
	</tr>
	<tr>
		<td>int</td><td>uint</td><td>complex</td><td>complex64</td><td>complex128</td>
	</tr>
	<tr>
		<td>float32</td><td>float64</td><td>bool</td><td>string</td><td>error</td>
	</tr>
	<tr>
		<td rowspan="3">内建函数</td><td>make</td><td>append</td><td>new</td><td>len</td><td>cap</td>
	</tr>
	<tr>
		<td>delete</td><td>close</td><td>panic</td><td>recover</td><td>real</td>
	</tr>
	<tr>
		<td>copy</td><td>complex</td><td>real</td><td></td><td></td>
	</tr>
</table>
			
[Reference](https://blog.csdn.net/Fe_cow/article/details/103842774?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromBaidu-1.nonecase&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromBaidu-1.nonecase)				