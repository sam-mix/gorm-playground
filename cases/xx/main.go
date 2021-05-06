package main

import "fmt"

type O struct {
	I *I
}

type I struct {
	Name string
}

func main() {
	var uint uint = 1
	println(uint)

	a := &O{}
	b := a
	c := a

	b.I = &I{"xxxx"}
	fmt.Printf("%#v\n", c.I)
}
