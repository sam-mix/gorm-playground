package main

import (
	"fmt"
	v2 "playground/cases/util/v2/conn"
	"playground/proto"
)

func main() {
	db := v2.Conn()
	d := &proto.Dog{Name: "123", Ms: &proto.Dog_Ms{Ms: []string{"123", "456", "789"}}}
	// d := &proto.Dog{}
	db.Create(d)
	d2 := &proto.Dog{}
	db.First(d2, d.Id)
	// fmt.Printf("%#v\n", d2.Ms)
	d2.Ms.Ms = append(d2.Ms.Ms, "xxx")
	fmt.Printf("%#v \n", d2.Ms.Ms)
}
