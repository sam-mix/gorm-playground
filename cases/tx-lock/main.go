package main

import (
	"errors"
	"fmt"
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"
	"sync"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	globalDB := v2.Conn()
	id := uint(1)
	var w sync.WaitGroup
	defer w.Wait()
	i := uint(1)
	globalDB.Exec("DROP TABLE IF EXIST t_user_888;")
	for j := 1; j <= 10; j++ {
		w.Add(1)
		go func(db *gorm.DB) {
			tx := db.Begin()
			defer func() {
				tx.Rollback()
				i++
				w.Done()
			}()
			oldVersion := &model.Version{}
			if err := tx.Table("t_user_888").Where("id = ?", id).Clauses(clause.Locking{Strength: "UPDATE"}).First(oldVersion).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					oldVersion.ID = id
					oldVersion.Version = i
					return
				} else {
					panic(err)
				}
			}
			if oldVersion.Version != i {
				fmt.Printf("gt: %d, err, old: %v, cur: %v \n", j, oldVersion.Version, i)
			} else {
				fmt.Printf("gt: %d, ok\n", j)
			}
			oldVersion.Version = i
			tx.Table("t_user_888").Save(oldVersion)

			time.Sleep(50 * time.Microsecond)

		}(globalDB)
	}

}
