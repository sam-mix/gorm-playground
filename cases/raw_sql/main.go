package main

import (
	"fmt"
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
)

func main() {
	db := v2.Conn()

	tblName := "t_user_4321"
	u := &model.User{Name: "Test"}
	db.Table(tblName).Create(u)
	var name string
	db.Raw(fmt.Sprintf("SELECT name FROM %s ", tblName)).Scan(&name)
	fmt.Printf("name => %s\n", name)
	var id uint
	db.Raw(fmt.Sprintf("SELECT id FROM %s", tblName)).Scan(&id)
	fmt.Printf("id = %d\n", id)
}
