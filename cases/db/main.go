package main

import (
	"errors"
	"playground/cases/util"
	v2 "playground/cases/util/v2/conn"
	"playground/proto"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	db := v2.Conn()
	tx := db
	p1 := &proto.Person{}
	if err := tx.Table("person_4").Clauses(clause.Locking{Strength: "UPDATE"}).First(p1, 102).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	}

	newPerson := &proto.Person{
		Id: 102,
		// Name: "V2",
	}
	person := util.ConvertStructWithTags(newPerson, map[string]reflect.StructTag{"Id": `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;autoIncrement:false"`})
	if err := tx.Table("person_5").Create(person).Error; err != nil {
	}

	d1 := &proto.Dog{}
	if err := tx.Table("dog_4").Clauses(clause.Locking{Strength: "UPDATE"}).First(d1, 102).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	}

	dog := &proto.Dog{
		Id: 102,
	}
	d := util.ConvertStructWithTags(dog, map[string]reflect.StructTag{"Id": `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;autoIncrement:false"`})
	if err := tx.Table("dog_5").Create(d).Error; err != nil {
	}

}
