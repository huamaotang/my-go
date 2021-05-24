package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for i:=0; i < 10; i++ {
		fmt.Println(i)
		GetHello()
		time.Sleep(10*time.Second)
	}
}

func GetHello()  {
	resp, err := http.Get("http://127.0.0.1:9999/hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
