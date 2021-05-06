package tests_test

import (
	"playground/proto"
	"testing"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func TestSometging(t *testing.T) {

	t.Run("xx", func(t *testing.T) {
		// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
		dsn := "root:1qaz@wsx@tcp(127.0.0.1:33306)/x?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db = db.Debug()

		db = db.Model(&proto.Person{})

		db = db.Table("person_1")

		// 迁移 schema
		db.AutoMigrate(db.Statement.Model)

		if db.Error != nil {
			t.Log(db.Error)
		}
		newPerson := &proto.Person{
			// Hobbies: &proto.Person_Hobbies{
			// 	Hobbies: []string{"足球", "篮球", "爬山"},
			// 	OuterMap: map[uint64]*proto.Person_Hobbies_Outer{
			// 		1: {
			// 			InnerMap: map[uint64]*proto.Person_Hobbies_Outer_Inner{
			// 				1: {Names: []string{"123"}},
			// 				2: {Names: []string{"345"}},
			// 				3: {Names: []string{"567"}},
			// 			},
			// 		},
			// 	},
			// },
			// Name: "123",
		}

		// Create
		db.Create(newPerson)

		// Rea
		var person proto.Person
		db = db.Debug()
		db.First(&person, newPerson.Id)
		db.First(&person, "name = ?", "123")

		// Update - 将 product 的 price 更新为 200
		db.Model(&person).Update("name", "name")
		// Update - 更新多个字段
		// db.Model(&person).Updates(Product{Price: 200, Code: "F42"})
		// db.Model(&person).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

		// Delete - 删除 product
		db.Delete(&person, newPerson.Id)
	})

}
