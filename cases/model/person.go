package model

import (
	"gorm.io/gorm"

	gormv1 "code.guanmai.cn/back_end/grpckit/database/gorm"
)

type Person struct {
	gorm.Model
	Name      string
	HairColor Color
}

type PersonV1 struct {
	gormv1.Model
	Name      string
	HairColor Color
}
