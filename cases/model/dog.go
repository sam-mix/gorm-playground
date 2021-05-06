package model

import (
	"gorm.io/gorm"

	gormv1 "code.guanmai.cn/back_end/grpckit/database/gorm"
)

type Dog struct {
	gorm.Model
	Name  string
	Color Color
}

type DogV1 struct {
	gormv1.Model
	Name  string
	Color Color
}
