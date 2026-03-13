package entity

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID         string         `gorm:"column:id;primaryKey"`
	UserId     string         `gorm:"column:user_id"`
	Address    string         `gorm:"column:address"`
	City       string         `gorm:"column:city"`
	PostalCode string         `gorm:"column:postal_code"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (a *Address) TableName() string {
	return "addresses"
}
