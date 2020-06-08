package main

import "fmt"

type Config struct {
	Name string
}

func work() (string, error) {
	return "work", nil
}

func works() (Config, error) {
	return Config{Name: "work"}, nil
}

func main() {
	// make map可以指定第二个参数，不过会被忽略；cap()适用于slice、channel、array、数组指针，但不适用于map，可以使用len()计算map长度
	m := make(map[string]int, 2)
	fmt.Println(len(m))

	var c Config
	var err error

	//c.Name, err := work()不能使用短变量声明设置结构体的字段值
	c.Name, err = work()
	fmt.Println(c, err)

	c1, err1 := works()
	fmt.Println(c1, err1)

	const x = 1.2

}
