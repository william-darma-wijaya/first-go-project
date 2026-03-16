package repository

import (
	"first-go-project/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressRepository struct {
	Repository[entity.Address]
	Log *logrus.Logger
}

func NewAddressRepository(log *logrus.Logger) *AddressRepository {
	return &AddressRepository{
		Log: log,
	}
}

func (r *AddressRepository) FindById(tx *gorm.DB, address *entity.Address, id string) error {
	return tx.Where("id = ?", id).First(address).Error
}

func (r *AddressRepository) DeleteByUserId(tx *gorm.DB, user *entity.User) error {
	return tx.Where("user_id = ?", user.Id).Delete(&entity.Address{}).Error
}