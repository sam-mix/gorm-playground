package main

import (
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
)

func main() {
	db := v2.Conn()

	db.Table("t_color_test").Create(&model.Dog{Name: "1", Color: model.Black})
	db.Table("t_color_test").Create(&model.Dog{Name: "1", Color: model.White})
	db.Table("t_color_test").Create(&model.Dog{Name: "1", Color: model.Red})

	var cs []*model.Dog
	db.Table("t_color_test").Where("color = ?", model.Black).Find(&cs)

}
