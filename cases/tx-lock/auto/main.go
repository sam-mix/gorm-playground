package main

import (
	"errors"
	"fmt"
	"log"
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
	tableName := "t_user_888"
	globalDB.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName))
	time.Sleep(time.Second)
	for j := 1; j <= 10; j++ {
		w.Add(1)
		go func(db *gorm.DB, x int) {
			if err := db.Transaction(func(tx *gorm.DB) error {
				defer func() {
					i++
					w.Done()
				}()
				oldVersion := &model.Version{}
				if err := tx.Table(tableName).Where("id = ?", id).Clauses(clause.Locking{Strength: "UPDATE"}).First(oldVersion).Error; err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						oldVersion.ID = id
						oldVersion.Version = i
						return tx.Table(tableName).Error
					} else {
						return err
					}
				}
				if oldVersion.Version != i {
					log.Fatalf("gt: %d, err, old: %v, cur: %v \n", x, oldVersion.Version, i)
					log.Printf("gt: %d, err, old: %v, cur: %v \n", x, oldVersion.Version, i)
					fmt.Printf("gt: %d, err, old: %v, cur: %v \n", x, oldVersion.Version, i)
				} else {
					log.Printf("gt: %d, ok\n", x)
					fmt.Printf("gt: %d, ok\n", x)
				}
				oldVersion.Version = i
				if err := tx.Table(tableName).Save(oldVersion).Error; err != nil {
					return err
				}
				return nil
			}); err != nil {
				fmt.Println(err)
			}
			time.Sleep(50 * time.Microsecond)
		}(globalDB, j)
	}

}
