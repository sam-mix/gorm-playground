package model

import (
	"gorm.io/gorm"

	gormv1 "code.guanmai.cn/back_end/grpckit/database/gorm"
)

type Keeper struct {
	gorm.Model
	Name     string
	CatColor Color
}
type KeeperV1 struct {
	gormv1.Model
	Name     string
	CatColor Color
}
