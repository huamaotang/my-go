package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 2.55
	c := a * 100
	strFloat := strconv.FormatFloat(c, 'f', 0, 64)
	d, _ := strconv.Atoi(strFloat)
	fmt.Println(c, strFloat, d)

	x := 2859.8694
	y := 1.0

	z := fmt.Sprintf("%.2f", x*y)
	fmt.Println(z)
	// 数字修约规则，银行家舍入法（比四舍五入更精确）
	// 四舍六入五考虑，五后非零就进一，五后为零看奇偶，五前为偶应舍去，五前为奇要进一
	fmt.Printf("9.8249	=>	%0.2f(四舍)\n", 9.8249)
	fmt.Printf("9.82671	=>	%0.2f(六入)\n", 9.82671)
	fmt.Printf("9.8351	=>	%0.2f(五后非零就进一)\n", 9.8351)
	fmt.Printf("9.82501	=>	%0.2f(五后非零就进一)\n", 9.82501)
	fmt.Printf("9.8250	=>	%0.2f(五后为零看奇偶，五前为偶应舍去)\n", 9.8250)
	fmt.Printf("9.8350	=>	%0.2f(五后为零看奇偶，五前为奇要进一)\n", 9.8350)

	fmt.Printf("-458.591 => %0.2f \n", -458.591)
}
