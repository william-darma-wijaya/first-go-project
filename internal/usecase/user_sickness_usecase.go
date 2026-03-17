package usecase

import (
	"context"
	"errors"
	"first-go-project/internal/CustomValidator"
	"first-go-project/internal/entity"
	"first-go-project/internal/helper"
	"first-go-project/internal/model"
	"first-go-project/internal/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserSicknessUsecase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	UserSicknessValidator  *CustomValidator.UserSicknessValidator
	UserSicknessRepository *repository.UserSicknessRepository
	SicknessRepository *repository.SicknessRepository
}

func NewUserSicknessUsecase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *CustomValidator.UserSicknessValidator,
	userSicknessRepository *repository.UserSicknessRepository,
	sicknessRepository *repository.SicknessRepository,
) *UserSicknessUsecase {
	return &UserSicknessUsecase{
		DB:                     db,
		Log:                    logger,
		UserSicknessValidator:  validate,
		UserSicknessRepository: userSicknessRepository,
		SicknessRepository: sicknessRepository,
	}
}

func (c *UserSicknessUsecase) CreateUserSickness(ctx context.Context, req *model.CreateUserSicknessRequest) (*model.CreateUserSicknessResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.UserSicknessValidator.ValidateCreateUserSickness(req); err != nil {
		c.Log.Warnf("Failed to validate user sickness body request: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	userSickness := &entity.UserSickness{
		Id:          helper.GenerateULID(),
		UserId:      req.UserId,
		SicknessId:  req.SicknessId,
		DiagnosedAt: time.Now(),
	}

	if err := c.UserSicknessRepository.Create(tx, userSickness); err != nil {
		c.Log.Warnf("(USECASE) Cannot insert User Sickness into database: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("(USECASE) Failed to commit user sickness transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	response := &model.CreateUserSicknessResponse{
		Id:          userSickness.Id,
		UserId:      userSickness.UserId,
		SicknessId:  userSickness.SicknessId,
		DiagnosedAt: userSickness.DiagnosedAt,
	}

	return response, nil
}


func (c *UserSicknessUsecase) CountDiagnosedBySicknessId(ctx context.Context, id string) (*model.CountDiagnosedBySicknessIdResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var userCount int64
	if err := c.UserSicknessRepository.CountDiagnosedBySicknessId(tx, id, &userCount); err != nil {
		c.Log.Warnf("Server failed to count users: %+v", err)
		return  nil, fiber.ErrInternalServerError
	}

	var sickness entity.Sickness
	if err := c.SicknessRepository.FindSicknessById(tx, &sickness, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Log.Warnf("Sickness record not found: %+v", err)
			return nil, fiber.ErrNotFound
		}
		c.Log.Warnf("Server failed to retrieve sickness data: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	response := &model.CountDiagnosedBySicknessIdResponse{
		Id: sickness.Id,
		Name: sickness.Name,
		Counts: userCount,
	}

	return response, nil
}
