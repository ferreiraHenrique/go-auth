package sqlite

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name      string
	ManagerID uint
	Manager   Manager `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
