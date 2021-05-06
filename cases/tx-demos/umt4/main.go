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
	pDB := db
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("\n 19 err => %#v\n", err)
			tx.Rollback()
		}
	}()
	pDB = pDB.Where("hair_color = ?", model.Blue)
	if err := pDB.Find(&[]model.Person{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	db = db.Where("cat_color = ?", model.Blue)
	if err := db.Find(&[]*model.Keeper{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("color = ?", model.White).First(&model.Dog{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if err := tx.Model(&model.User{}).Where("dog_color = ?", model.Black).Updates(map[string]interface{}{
		"name": "sam",
	}).Error; err != nil {
		return
	}

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("eye_color = ?", model.White).First(&model.Cat{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
	}

	if err := tx.Model(&model.Dog{}).Where("color = ?", model.Black).Updates(map[string]interface{}{
		"name": "sam's dog",
	}).Error; err != nil {
		tx.Rollback()
	}

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("dog_color = ?", model.White).First(&model.User{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
	}

	if err := tx.Model(&model.Cat{}).Where("eye_color = ?", model.Black).Updates(map[string]interface{}{
		"name": "sam's cat",
	}).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()

	fmt.Printf("db => %#v\n", db)
}
