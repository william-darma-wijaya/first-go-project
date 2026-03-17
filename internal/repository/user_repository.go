package repository

import (
	"first-go-project/internal/entity"
	"time"

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

func (r *UserRepository) FindByIdWithSickness(
	db *gorm.DB, 
	user *entity.User, 
	userId string, 
	startDate *time.Time, 
	endDate *time.Time,
) error {
	query := db.
		Where("users.id = ?", userId).
		Preload("Sicknesses", func(tx *gorm.DB) *gorm.DB {
			if (startDate != nil) && (endDate != nil) {
				return tx.Where("diagnosed_at BETWEEN ? AND ?", *startDate, *endDate)
			} else if startDate != nil {
				return tx.Where("diagnosed_at >= ?", *startDate)
			} else if endDate != nil {
				return tx.Where("diagnosed_at <= ?", *endDate)
			}
			return tx
		}).Preload("Sicknesses.Sickness")

	return query.First(user).Error
}