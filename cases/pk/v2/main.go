package main

import (
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
)

func main() {
	db := v2.Conn()

	npk := &model.NPK{
		ID: 0,
		S:  "S",
		L:  "L",
	}
	db.Table("t_npk_v2").Save(npk)
}
