package entity

import (
	"time"

	"gorm.io/gorm"
)

type Sickness struct {
	ID          string         `gorm:"column:id;primaryKey"`
	Name        string         `gorm:"column:name"`
	Description string         `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`

	Users []User `gorm:"many2many:user_sickness"`
}

func (s *Sickness) TableName() string {
	return "sickness"
}
