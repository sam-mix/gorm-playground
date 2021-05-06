package main

import (
	"errors"
	"playground/cases/util"
	v1 "playground/cases/util/v1/conn"
	"playground/proto"

	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := v1.Conn()
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()

	newPerson := &proto.Person{
		Id: 101,
		// Name: "v1",
	}
	person := util.ConvertStructWithTags(newPerson, map[string]reflect.StructTag{"Id": `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_Key;auto_Increment:false"`})

	if err := tx.Create(person).Error; err != nil {
		return
	}

	p1 := &proto.Person{}
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(p1, 101).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	p2 := &proto.Person{}
	if err := tx.First(p2, 101).Set("gorm:query_option", "FOR UPDATE").Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	tx = tx.Table("dog_1")
	d1 := &proto.Dog{}
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(d1, 101).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	dog := &proto.Dog{
		Id: 101,
	}
	d := util.ConvertStructWithTags(dog, map[string]reflect.StructTag{"Id": `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;autoIncrement:false"`})
	if err := tx.Create(d).Error; err != nil {
		return
	}
	tx.Commit()
}
