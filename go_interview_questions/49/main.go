package main

import "fmt"

func main() {
	// 补码：正数的补码为原码；负数的补码为符号位不变，其余位取反，最后加一
	// 取反操作
	// 知识点：1、有符号整数a，简便计算方式: -(a+1);2、数a取反，计算机计算方式：先转为二进制，取a补码；a补码每一位取反，包括符号位；最后取补码，得到原码
	fmt.Printf("%b\n", 3)
	fmt.Printf("%b\n", -3)
	fmt.Printf("%b\n", ^3)
	fmt.Printf("%b\n", ^-3)
	fmt.Printf("%b\n", ^uint8(3))

	var a int8 = 3
	var b int8 = 5
	// 异或操作，对应位相同取0，对应位不同取1
	fmt.Printf("%08b %08b %08b\n", a, b, a^b)
}
