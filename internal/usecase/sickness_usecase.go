package usecase

import (
	"context"
	"errors"
	"first-go-project/internal/CustomValidator"
	"first-go-project/internal/entity"
	"first-go-project/internal/helper"
	"first-go-project/internal/model"
	"first-go-project/internal/model/converter"
	"first-go-project/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SicknessUsecase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	SicknessValidator      *CustomValidator.SicknessValidator
	UserSicknessRepository *repository.UserSicknessRepository
	SicknessRepository     *repository.SicknessRepository
}

func NewSicknessUsecase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *CustomValidator.SicknessValidator,
	userSicknessRepository *repository.UserSicknessRepository,
	sicknessRepository *repository.SicknessRepository,
) *SicknessUsecase {
	return &SicknessUsecase{
		DB:                     db,
		Log:                    logger,
		SicknessValidator:      validate,
		UserSicknessRepository: userSicknessRepository,
		SicknessRepository:     sicknessRepository,
	}
}

func (c *SicknessUsecase) CreateSickness(ctx context.Context, request *model.CreateSicknessRequest) (*model.CreateSicknessResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.SicknessValidator.ValidateCreateSickness(request); err != nil {
		c.Log.Warnf("Invalid body request: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	existingSickness := new(entity.Sickness)
	err := c.SicknessRepository.FindSicknessByName(tx, existingSickness, request.Name)
	if err == nil {
		// sickness ditemukan
		c.Log.Warnf("Sickness already exists : %s", request.Name)
		return nil, fiber.ErrConflict
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.Log.Warnf("Failed query sickness by name: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	NewSickness := &entity.Sickness{
		Id:          helper.GenerateULID(),
		Name:        request.Name,
		Description: request.Description,
	}

	if err = c.SicknessRepository.Create(tx, NewSickness); err != nil {
		c.Log.Warnf("Failed to create new sickness: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err = tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.SicknessToCreateResponse(NewSickness), nil
}

func (c *SicknessUsecase) FindById(ctx context.Context, id string) (*model.GetSicknessByIdResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	sickness := new(entity.Sickness)
	if err := c.SicknessRepository.FindSicknessById(tx, sickness, id); err != nil {
		c.Log.Warnf("Failed to get sickness: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.SicknessToFindByIdResponse(sickness), nil
}

func (c *SicknessUsecase) UpdateSickness(ctx context.Context, request *model.UpdateSicknessRequest) (*model.UpdateSicknessResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.SicknessValidator.ValidateUpdateSickness(request); err != nil {
		c.Log.Warnf("Invalid body request: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	sickness := new(entity.Sickness)
	if err := c.SicknessRepository.FindSicknessById(tx, sickness, request.Id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Log.Warnf("Failed find sickness by id: %+v", err)
			return nil, fiber.ErrNotFound
		}
		c.Log.Warnf("Failed to find sickness by id: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if request.Name != nil {
		sickness.Name = *request.Name
	}
	if request.Description != nil {
		sickness.Description = *request.Description
	}

	if err := c.SicknessRepository.Update(tx, sickness); err != nil {
		c.Log.Warnf("Failed to update sickness: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit update transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.SicknessToUpdateResponse(sickness), nil
}

func (c *SicknessUsecase) DeleteSickness(ctx context.Context, id string) (*model.DeleteSicknessResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	sickness := new(entity.Sickness)
	sickness.Id = id
	if err := c.UserSicknessRepository.DeleteBySicknessId(tx, sickness); err != nil {
		c.Log.Warnf("Failed to delete user sickness : %+v", err)
		return nil, fiber.ErrInternalServerError
	}
	if err := c.SicknessRepository.Delete(tx, sickness); err != nil {
		c.Log.Warnf("Failed to delete sickness : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.DeleteSicknessResponse{Message: "Sickness deleted"}, nil
}
