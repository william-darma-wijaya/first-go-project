package usecase

import (
	"context"
	"first-go-project/internal/CustomValidator"
	"first-go-project/internal/entity"
	"first-go-project/internal/model"
	"first-go-project/internal/model/converter"
	"first-go-project/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	AddressValidator  *CustomValidator.AddressValidator
	AddressRepository *repository.AddressRepository
}

func NewAddressUseCase(db *gorm.DB, logger *logrus.Logger, validate *CustomValidator.AddressValidator,
	addressRepository *repository.AddressRepository) *AddressUseCase {
	return &AddressUseCase{
		DB:                db,
		Log:               logger,
		AddressValidator:  validate,
		AddressRepository: addressRepository,
	}
}

func (c *AddressUseCase) UpdateAddress(ctx context.Context, request *model.UpdateAddressRequest) (*model.UpdateAddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.AddressValidator.ValidateUpdateAddress(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	address := new(entity.Address)
	if err := c.AddressRepository.FindById(tx, address, request.ID); err != nil {
		c.Log.Warnf("Failed find address by id : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if request.Address != "" {
		address.Address = request.Address
	}
	if request.City != "" {
		address.City = request.City
	}
	if request.PostalCode != "" {
		address.PostalCode = request.PostalCode
	}

	if err := c.AddressRepository.Update(tx, address); err != nil {
		c.Log.Warnf("Failed update address : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.AddressToUpdateResponse(address), nil
}

func (c *AddressUseCase) DeleteAddress(ctx context.Context, id string) (*model.DeleteAddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	address := new(entity.Address)
	address.ID = id
	if err := c.AddressRepository.Delete(tx, address); err != nil {
		c.Log.Warnf("Failed to delete address : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.DeleteAddressResponse{Message: "Address deleted"}, nil
}
