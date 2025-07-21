package models

import (
	"dottest/utils"
	"fmt"

	"gorm.io/gorm"
)

type Mapping struct {
	ID     uint   `gorm:"primaryKey"`
	Domain string `gorm:"uniqueIndex;not null"`
	Target string
}

func (m *Mapping) BeforeSave(tx *gorm.DB) (err error) {
	err = utils.GenerateCertificate(fmt.Sprintf("%s.test", m.Domain))
	return err
}
