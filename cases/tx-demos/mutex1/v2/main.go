package main

import (
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"

	"gorm.io/gorm"
)

func main() {
	db := v2.Conn()
	go loop(db)
}

func loop(db *gorm.DB) {
	for {
		exec(db)
	}
}

func exec(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		var users []*model.User
		tx.Find(&users)
		return nil
	})

}
