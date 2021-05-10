package model

// has not primary key
type NPK struct {
	ID uint64 `gorm:"index"`
	S  string `gorm:"index"`
	L  string `gorm:"index"`
}
