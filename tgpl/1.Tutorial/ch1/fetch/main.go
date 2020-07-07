// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type myWriter struct {}

func (myWriter) Write(p []byte) (int, error) {
	//fmt.Println("my write", string(p))
	return len(p), nil
}

var w io.Writer = new(myWriter)

func main() {
	now := time.Now()
	for _, url := range os.Args[1:] {
		nowOne := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			_ , _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		if _, err = io.Copy(ioutil.Discard, resp.Body); err != nil {
			os.Exit(1)
		}

		if err = resp.Body.Close(); err != nil {
			 _ , _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%.2fs %s\n", time.Since(nowOne).Seconds(), url)
	}
	fmt.Printf("%.2fs \n", time.Since(now).Seconds())
}

//!-
