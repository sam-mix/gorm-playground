package v1m

import (
	"code.guanmai.cn/back_end/grpckit/database/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Conn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1qaz@wsx@(localhost:23306)/x?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)

	}
	db = db.Debug()
	db.Callback().Query().Before("gorm:query").Register("gorm:auto_migrate", migrate)
	db.Callback().Update().Before("gorm:update").Register("gorm:auto_migrate", migrate)
	db.Callback().Create().Before("gorm:create").Register("gorm:auto_migrate", migrate)
	return db
}

func migrate(scope *gorm.Scope) {
	scope.DB().AutoMigrate(scope.Value)
}
