package sqlite

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name   string
	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
