package sqlite

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UUID      string
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.UUID = uuid.NewString()
	return nil
}
