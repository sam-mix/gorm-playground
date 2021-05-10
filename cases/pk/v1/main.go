package main

import (
	"playground/cases/model"
	v1m "playground/cases/util/v1m/conn"
)

func main() {
	db := v1m.Conn()

	npk := &model.NPK{
		ID: 0,
		S:  "S",
		L:  "L",
	}
	db.Save(npk)
}
