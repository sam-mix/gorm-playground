package v2

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Conn() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:1qaz@wsx@tcp(127.0.0.1:23306)/x?charset=utf8mb4&parseTime=True&loc=Local"

	zapLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: zapLogger,
	})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	db.Callback().Query().Before("gorm:query").Register("gorm:auto_migrate", migrate)
	db.Callback().Update().Before("gorm:update").Register("gorm:auto_migrate", migrate)
	db.Callback().Create().Before("gorm:create").Register("gorm:auto_migrate", migrate)

	return db

}

func migrate(db *gorm.DB) {
	db.AutoMigrate(db.Statement.Model)
	db.Model(nil)
}
