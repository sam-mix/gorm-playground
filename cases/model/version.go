package model

import "gorm.io/gorm"

type Version struct {
	gorm.Model
	Version uint
}
