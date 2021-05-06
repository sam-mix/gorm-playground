package main

import (
	"playground/cases/model"
	v1 "playground/cases/util/v1/conn"
)

func main() {
	db := v1.Conn()
	d := &model.Dog{Name: "Darling"}
	db.Table("t_dog_01").Create(d)

	db.Where("id = ", d.ID)
}
