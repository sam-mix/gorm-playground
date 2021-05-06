package main

import (
	"errors"
	"playground/cases/util"
	v2 "playground/cases/util/v2/conn"
	"playground/proto"
	"reflect"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

func main() {
	db := v2.Conn()
	db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')

		tx.Transaction(func(tx1 *gorm.DB) error {
			p1 := &proto.Person{}
			if err := tx.Table("person_1").Clauses(clause.Locking{Strength: "UPDATE"}).First(p1, 102).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			return nil
		})

		tx.Transaction(func(tx2 *gorm.DB) error {
			newPerson := &proto.Person{
				Id: 102,
				// Name: "V2",
			}
			person := util.ConvertStructWithTags(newPerson, map[string]reflect.StructTag{"Id": `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;autoIncrement:false"`})
			if err := tx2.Table("person_1").Create(person).Error; err != nil {
				return err
			}
			return nil
		})

		tx.Transaction(func(tx3 *gorm.DB) error {
			d1 := &proto.Dog{}
			if err := tx3.Table("dog_1").Clauses(clause.Locking{Strength: "UPDATE"}).First(d1, 102).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			return nil
		})

		tx.Transaction(func(tx4 *gorm.DB) error {
			dog := &proto.Dog{
				Id: 102,
			}
			d := util.ConvertStructWithTags(dog, map[string]reflect.StructTag{"Id": `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;autoIncrement:false"`})
			if err := tx4.Table("dog_1").Create(d).Error; err != nil {
				return err
			}
			return nil
		})

		// return nil will commit the whole transaction
		return nil
	})

}
