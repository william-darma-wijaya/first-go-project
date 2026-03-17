package repository

import (
	"first-go-project/internal/entity"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SicknessRepository struct {
	Repository[entity.Sickness]
	Log *logrus.Logger
}

func NewSicknessRepository(log *logrus.Logger) *SicknessRepository {
	return &SicknessRepository{
		Log: log,
	}
}

func (r *SicknessRepository) FindSicknessById(db *gorm.DB, sickness *entity.Sickness, id string) error {
	return db.Where("id = ?", id).First(sickness).Error
}

func (r *SicknessRepository) FindSicknessByName(db *gorm.DB, sickness *entity.Sickness, name string) error {
	return db.Where("name = ?", name).First(sickness).Error
}

func (r *SicknessRepository) FindSicknessesByName(db *gorm.DB, sicknesses *[]entity.Sickness, name string) error {
	return db.Where("name ILIKE ?", "%"+name+"%").Find(sicknesses).Error
}

func (r *SicknessRepository) FindSicknessByIdWithUser(
	db *gorm.DB, 
	sickness *entity.Sickness, 
	id string,
	startDate *time.Time,
	endDate *time.Time,
) error {
	query := db.
		Where("id = ?", id).
		Preload("Users", func (tx *gorm.DB) *gorm.DB {
			if (startDate != nil)  && (endDate != nil){
				return tx.Where("diagnosed_at BETWEEN ? AND ?", startDate, endDate)
			} else if startDate != nil {
				return tx.Where("diagnosed_at >= ?", startDate)
			} else if endDate != nil {
				return tx.Where("diagnosed_at <= ?", endDate)
			}
			return tx
		}).
		Preload("Users.User")
	return query.First(sickness).Error
}