package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserSickness struct {
	Id          string         `gorm:"column:id;primaryKey"`
	UserId      string         `gorm:"column:user_id"`
	SicknessId  string         `gorm:"column:sickness_id"`
	DiagnosedAt time.Time      `gorm:"column:diagnosed_at"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`

	User     User     `gorm:"foreignKey:UserId;references:Id"`
	Sickness Sickness `gorm:"foreignKey:SicknessId;references:Id"`
}

func (UserSickness) TableName() string {
	return "user_sickness"
}
