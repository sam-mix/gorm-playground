package main

import (
	"fmt"
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
)

func main() {
	db := v2.Conn()
	tx := db.Begin()
	defer tx.Rollback()
	if err := tx.Create(&model.Bee{}).Error; err != nil {
		fmt.Println(err)
	}
	if err := tx.Create(&model.Bee{}).Error; err != nil {
		fmt.Println(err)
	}
	b := &model.Bee{}
	if err := tx.Create(b).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(b.ID)
	b.DukeColor = model.Black
	if err := tx.Save(b).Error; err != nil {
		fmt.Println(err)
	}
	if err := tx.Commit().Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println("end")
}
