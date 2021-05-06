package model

import (
	"gorm.io/gorm"

	gormv1 "code.guanmai.cn/back_end/grpckit/database/gorm"
)

type Cat struct {
	gorm.Model
	Name     string
	EyeColor Color
}

type CatV1 struct {
	gormv1.Model
	Name     string
	EyeColor Color
}
