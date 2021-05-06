package main

import (
	"playground/cases/model"
	"sync"

	"github.com/jinzhu/gorm"
)

func main() {
	// 	db := v1.Conn()
	// 	var mutex *sync.Mutex
	// 	fo
}

func loop(mutex *sync.Mutex, db *gorm.DB) {
	mutex.Lock()
	defer mutex.Unlock()
	exec(db)
}

func exec(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		var users []*model.UserV1
		tx.Find(&users)
		return nil
	})

}
