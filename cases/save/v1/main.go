package main

import (
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
)

func main() {
	db := v2.Conn()
	t := &model.TaskV1{ID1: 1, ID2: 0}
	db.Save(t)
	t.Version = 1
	db.Save(t)

}
