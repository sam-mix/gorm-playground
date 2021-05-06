package main

import (
	"fmt"
	"playground/cases/model"
	v1m "playground/cases/util/v1m/conn"
)

func main() {
	db := v1m.Conn()
	db = db.Where("name LIKE (?)", fmt.Sprintf("%%%s%%", "sam"))
	catDB := db
	catDB = catDB.Where("eye_color = 3")
	catDB.First(&model.Cat{})
	dogDB := db
	dogDB = dogDB.Where("color = 4")
	dogDB.First(&model.Dog{})
	userDB := db
	userDB = userDB.Where("dog_color = 5")
	userDB.First(&model.User{})
}
