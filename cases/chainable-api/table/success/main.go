package main

import (
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
)

func main() {
	db := v2.Conn()
	var task *model.Task
	db.Table("t_task_999").First(&task)
	db.Table("t_task_999").First(&task, "id1 = ?", 1)
}
