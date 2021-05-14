package main

import "fmt"

func main() {
	var buf []byte
	for i := 1; ; i++ {
		buf = make([]byte, 5*1024*1024*1024)
		fmt.Printf("t-%v, x-%v\n", i, buf[0])
	}
}
