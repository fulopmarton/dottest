package models

type Mapping struct {
	ID     uint   `gorm:"primaryKey"`
	Domain string `gorm:"uniqueIndex;not null"`
	Target string
}
