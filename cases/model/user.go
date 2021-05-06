package model

import (
	"gorm.io/gorm"

	gormv1 "code.guanmai.cn/back_end/grpckit/database/gorm"
)

type User struct {
	gorm.Model
	Name     string
	DogColor Color
	Version  uint64
}

type UserV1 struct {
	gormv1.Model
	Name     string
	DogColor Color
	Version  uint64
}
