package main

import (
	"playground/cases/model"
	v2 "playground/cases/util/v2/conn"

	"gorm.io/gorm"
)

func main() {
	db := v2.Conn()
	tblName := "t_bee_loop_save"
	if err := db.Transaction(func(tx *gorm.DB) error {
		bs := []model.Bee{}
		for i := 1; i < 11; i++ {
			b := model.Bee{DukeColor: model.Black, X: i}
			tx.Table(tblName).Create(&b)
			bs = append(bs, b)
		}
		for i, b := range bs {
			b.X = (i + 1) * (i + 1)
			tx.Table(tblName).Save(b)
		}
		return nil
	}); err != nil {
		panic(err)
	}
}
