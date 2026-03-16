package entity

import (
	"time"

	"gorm.io/gorm"
)

type Sickness struct {
	Id          string         `gorm:"column:id;primaryKey"`
	Name        string         `gorm:"column:name"`
	Description string         `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`

	Users []UserSickness `gorm:"foreignKey:SicknessId;references:Id"`
}

func (s *Sickness) TableName() string {
	return "sickness"
}
