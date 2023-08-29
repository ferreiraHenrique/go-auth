package sqlite

type Permission struct {
	Base
	Name      string
	ManagerID uint
	Manager   Manager `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
