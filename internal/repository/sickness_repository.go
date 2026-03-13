package repository

import (
	"first-go-project/internal/entity"

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

func (c *SicknessRepository) FindSicknessByName(db *gorm.DB, sickness *entity.Sickness, name string) {
	return db.
}