package main

import (
	"fmt"
	"playground/cases/model"
	"playground/cases/util"
	v2 "playground/cases/util/v2/conn"

	"gorm.io/gorm"
)

type Impl struct {
	DB *gorm.DB
}

func main() {
	impl := &Impl{
		DB: v2.Conn(),
	}
	service1(impl)
	service2(impl)
	service3(impl)
}

func service1(impl *Impl) {

	db := impl.DB.Where("name LIKE (?)", fmt.Sprintf("%%%s%%", "sam"))
	catDB := util.CloneDB(db)
	catDB.Where("eye_color = 3").First(&model.Cat{})
	dogDB := util.CloneDB(db)
	dogDB.Where("color = 4").First(&model.Dog{})
	userDB := util.CloneDB(db)
	userDB = userDB.Where("dog_color = 5")
	userDB.First(&model.User{})
}

func service2(impl *Impl) {
	db := util.CloneDB(impl.DB)
	db.Where("hair_color = ?", model.Blue).First(&model.Person{})
}

func service3(impl *Impl) {
	db := util.CloneDB(impl.DB)
	db.Where("cat_color = ?", model.Red).First(&model.Keeper{})
}
