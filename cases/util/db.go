package util

import (
	"gorm.io/gorm"
)

func NewDB(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{NewDB: true})
}

func CloneDB(db *gorm.DB) *gorm.DB {
	return db.WithContext(db.Statement.Context)
}

func Rollback(tx *gorm.DB) {
	if err := recover(); err != nil {
		tx.Rollback()
	}
}
