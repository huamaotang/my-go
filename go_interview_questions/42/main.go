package main

import "fmt"

type Config struct {
	Name string
}

func (c *Config) String() string {
	//return fmt.Sprintf("print: %v", c) 使用格式化输出导致该函数被递归调用，栈溢出
	return fmt.Sprintf("print: %v", c.Name)
}

func main() {
	x := &Config{Name: "tanghuamao"}
	fmt.Println(x)

	s := []int{1, 2, 3, 4}

	var r []int

	for k, v := range s {
		if k == 0 {
			s = append(s, 5, 6)
			s[3] = 9
		}
		r = append(r, v)
	}
	// for range会使用s的副本s1参与循环，尽管s的长度、值都变了，但s1长度len不变，值不变
	fmt.Println(s, r)
}
