package repository

import (
	"first-go-project/internal/entity"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserSicknessRepository struct {
	Repository[entity.UserSickness]
	Log *logrus.Logger
}

func NewUserSicknessRepository(log *logrus.Logger) *UserSicknessRepository {
	return &UserSicknessRepository{
		Log: log,
	}
}

func (r *UserSicknessRepository) FindByDiagnosedDate(
	db *gorm.DB,
	startDate time.Time,
	endDate time.Time,
	results *[]entity.UserSickness,
) error {
	return db.
		Preload("User").
		Preload("Sickness").
		Where("diagnosed_at BETWEEN ? AND ?", startDate, endDate).
		Find(results).Error
}

func (r *UserSicknessRepository) CountDiagnosedBySicknessId(db *gorm.DB, id string, userCount *int64) error {
	return db.Where("sickness_id = ?", id).Count(userCount).Error
}

func (r *UserSicknessRepository) DeleteByUserId(tx *gorm.DB, user *entity.User) error {
	return tx.Where("user_id = ?", user.Id).Delete(&entity.UserSickness{}).Error
}

func (r *UserSicknessRepository) DeleteBySicknessId(tx *gorm.DB, sickness *entity.Sickness) error {
	return tx.Where("sickness_id = ?", sickness.Id).Delete(&entity.UserSickness{}).Error
}