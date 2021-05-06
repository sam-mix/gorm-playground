package main

import (
	"errors"
	"fmt"
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	db := v2.Conn()
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err := c5(tx); err != nil {
		tx.Rollback()
	}

	if err := c1(tx); err != nil {
		tx.Rollback()
	}

	if err := c4(tx); err != nil {
		tx.Rollback()
	}

	if err := c2(tx); err != nil {
		tx.Rollback()
	}

	if err := c6(tx); err != nil {
		tx.Rollback()
	}

	if err := c3(tx); err != nil {
		tx.Rollback()
	}
	tx.Commit()

	fmt.Printf("db => %#v\n", db)
}

func c1(db *gorm.DB) error {
	return db.Model(&model.User{}).Where("dog_color = ?", model.Black).Updates(map[string]interface{}{
		"name": "sam",
	}).Error
}

func c2(db *gorm.DB) error {
	return db.Model(&model.Dog{}).Where("color = ?", model.Black).Updates(map[string]interface{}{
		"name": "sam's dog",
	}).Error
}

func c3(db *gorm.DB) error {
	return db.Model(&model.Cat{}).Where("eye_color = ?", model.Black).Updates(map[string]interface{}{
		"name": "sam's cat",
	}).Error
}

func c4(db *gorm.DB) error {
	db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	err := db.Where("eye_color = ?", model.White).First(&model.Cat{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}

func c5(db *gorm.DB) error {
	db.Clauses(clause.Locking{Strength: "UPDATE"})
	err := db.Where("color = ?", model.White).First(&model.Dog{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}

func c6(db *gorm.DB) error {
	db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	err := db.Where("dog_color = ?", model.White).First(&model.User{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
