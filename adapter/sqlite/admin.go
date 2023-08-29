package sqlite

type Admin struct {
	Base
	Name   string
	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
