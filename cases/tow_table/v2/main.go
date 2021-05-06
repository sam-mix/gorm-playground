package main

import (
	"fmt"
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"

	"gorm.io/gorm"
)

func main() {
	db := v2.Conn()
	if err := service(db); err != nil {
		fmt.Printf("Error => %v", err)
	}

	fmt.Printf("db => %#v\n", db)
}

func service(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := c1(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := c2(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := c3(tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func c1(db *gorm.DB) error {
	return db.Create(&model.User{Name: "name test"}).Error
}

func c2(db *gorm.DB) error {
	d := &model.Dog{
		Model: gorm.Model{ID: 1},
		Name:  "name test",
	}
	return db.Create(d).Error
}

func c3(db *gorm.DB) error {
	c := &model.Cat{
		Model: gorm.Model{ID: 1},
		Name:  "name test",
	}
	return db.Create(c).Error
}
