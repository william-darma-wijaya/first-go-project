package repository

import (
	"first-go-project/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) FindById(db *gorm.DB, user *entity.User, id string) error {
	return db.Where("id = ?", id).First(user).Error
}

func (r *UserRepository) FindByIdWithAddress(db *gorm.DB, user *entity.User, id string) error {
	return db.Preload("Addresses").Where("id = ?", id).First(user).Error
}

// Check if email already exists
func (r *UserRepository) FindByEmail(db *gorm.DB, user *entity.User, email string) error {
	return db.
		Where("email = ?", email).
		First(user).
		Error
}