package main

import (
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
)

func main() {
	db := v2.Conn()
	db = db.Table("t_task_998")
	var task model.Task
	db.First(&task)
	db.First(&task, "id1 = ?", 1)
}
