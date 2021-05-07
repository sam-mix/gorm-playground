package main

import (
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
)

func main() {
	db := v2.Conn()
	t := &model.Task{ID1: 1, ID2: 1}
	db.Table("t_task_100").Save(t)
	t.Version = 1
	db.Table("t_task_100").Save(t)

}
