package model

type Task struct {
	ID1     uint64 `gorm:"primaryKey;autoIncrement:false"`
	ID2     uint64 `gorm:"primaryKey;autoIncrement:false"`
	Version uint64
}

type TaskV1 struct {
	ID1     uint64 `gorm:"primary_Key;auto_Increment:false"`
	ID2     uint64 `gorm:"primary_Key;auto_Increment:false"`
	Version uint64
}
