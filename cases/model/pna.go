package model

type PNA struct {
	PnaID uint64 `gorm:"primaryKey;autoIncrement:false"`
	Name  string
}

type PNAv1 struct {
	PnaID uint64 `gorm:"primary_Key;auto_Increment:false"`
	Name  string
}
