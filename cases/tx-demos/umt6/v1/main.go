package main

import (
	"fmt"
	"playground/cases/model"
	v1m "playground/cases/util/v1m/conn"

	"code.guanmai.cn/back_end/grpckit/database/gorm"
)

type Impl struct {
	DB *gorm.DB
}

func main() {
	impl := &Impl{
		DB: v1m.Conn(),
	}
	service1(impl)
	service2(impl)
	service3(impl)
}

func service1(impl *Impl) {

	db := impl.DB.Where("name LIKE (?)", fmt.Sprintf("%%%s%%", "sam"))
	catDB := db
	catDB = catDB.Where("eye_color = 3")
	catDB.First(&model.CatV1{})
	dogDB := db
	dogDB = dogDB.Where("color = 4")
	dogDB.First(&model.DogV1{})
	userDB := db
	userDB = userDB.Where("dog_color = 5")
	userDB.First(&model.UserV1{})
}

func service2(impl *Impl) {
	impl.DB.Where("hair_color = ?", model.Black).First(&model.PersonV1{})
}

func service3(impl *Impl) {
	impl.DB.Where("cat_color = ?", model.Black).First(&model.Keeper{})
}
