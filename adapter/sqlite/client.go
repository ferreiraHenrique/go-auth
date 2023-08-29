package sqlite

type Client struct {
	Base
	Name        string
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ManagerID   uint
	Manager     Manager      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Permissions []Permission `gorm:"many2many:client_permissions;"`
}
