package store

type Mapping struct {
	ID     uint   `gorm:"primaryKey"`
	domain string `gorm:"uniqueIndex;not null"`
	target string
}
