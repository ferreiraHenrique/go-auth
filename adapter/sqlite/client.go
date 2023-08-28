package sqlite

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name      string
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ManagerID uint
	Manager   Manager `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
