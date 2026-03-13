package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserSickness struct {
	ID          string         `gorm:"column:id;primaryKey"`
	UserID      string         `gorm:"column:user_id"`
	SicknessID  string         `gorm:"column:sickness_id"`
	DiagnosedAt time.Time      `gorm:"column:diagnosed_at"`

	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	User     User     `gorm:"foreignKey:UserID"`
	Sickness Sickness `gorm:"foreignKey:SicknessID"`
}

func (UserSickness) TableName() string {
	return "user_sickness"
}