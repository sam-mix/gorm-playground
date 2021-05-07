package model

type Task struct {
	ID1     uint64 `gorm:"primaryKey;autoIncrement:false"`
	ID2     uint64 `gorm:"primaryKey;autoIncrement:false"`
	Version uint64
}
