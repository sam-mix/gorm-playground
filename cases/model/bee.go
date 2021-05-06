package model

import (
	"gorm.io/gorm"

	gormv1 "code.guanmai.cn/back_end/grpckit/database/gorm"
)

type Bee struct {
	gorm.Model
	DukeColor Color
}

type BeeV1 struct {
	gormv1.Model
	DukeColor Color
}
