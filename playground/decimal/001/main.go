package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {

	d, err := decimal.NewFromString("--1")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

}
