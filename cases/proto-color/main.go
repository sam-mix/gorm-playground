package main

import (
	"fmt"
	v2 "playground/cases/util/v2/conn"
	"playground/proto"
)

func main() {
	db := v2.Conn()
	db.Table("t_proto_color_test").Create(&proto.Dog{
		Name:  "Haele",
		Color: proto.Dog_x,
	})
	db.Table("t_proto_color_test").Create(&proto.Dog{
		Name:  "Ha",
		Color: proto.Dog_xxx,
	})
	db.Table("t_proto_color_test").Create(&proto.Dog{
		Name:  "H",
		Color: proto.Dog_xx,
	})

	var cs []*proto.Dog
	db.Table("t_proto_color_test").Where("color = ?", proto.Dog_x).Find(&cs)
	fmt.Println(len(cs))
}
