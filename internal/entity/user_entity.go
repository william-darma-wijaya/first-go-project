package entity

import (
	"time"

	"gorm.io/gorm"
)

// User is a struct that represents a user entity
type User struct {
	Id        string         `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name"`
	Email     string         `gorm:"column:email"`
	Password  string         `gorm:"column:password"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`

	Addresses  []Address      `gorm:"foreignKey:UserId;references:Id"`
	Sicknesses []UserSickness `gorm:"foreignKey:UserId;references:Id"`
}

// To inform gorm the table name in DB
func (u *User) TableName() string {
	return "users"
}
