package main

import (
	"errors"
	"fmt"
	"playground/cases/model"
	"playground/cases/util"
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
	globalDB.Exec("DROP TABLE IF EXIST t_user_888")
	for j := 1; j <= 10; j++ {
		w.Add(1)
		go func(db *gorm.DB, x int) {
			tx := util.NewDB(db).Begin()
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
					if err := tx.Table("t_user_888").Save(oldVersion).Error; err != nil {
						fmt.Println(err)
					}
					return
				} else {
					fmt.Println(err)
				}
			}
			if oldVersion.Version != i {
				fmt.Printf("gt: %d, err, old: %v, cur: %v \n", x, oldVersion.Version, i)
			} else {
				fmt.Printf("gt: %d, ok\n", x)
			}
			oldVersion.Version = i
			if err := tx.Table("t_user_888").Save(oldVersion).Error; err != nil {
				fmt.Println(err)
			}
			time.Sleep(50 * time.Microsecond)

		}(globalDB, j)
	}

}
