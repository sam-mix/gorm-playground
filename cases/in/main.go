package main

import (
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
)

func main() {
	names := []string{"123", "321"}
	db := v2.Conn()
	var users []*model.User
	db.Find(&users, "name IN ?", names)
	db.Find(&users, "name IN (?)", names)
	db.Where("name IN (?)", names).Find(&users)
	db.Where("name IN ?", names).Find(&users)
}
