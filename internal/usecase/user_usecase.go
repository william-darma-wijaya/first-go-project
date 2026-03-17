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

type UserUseCase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	UserValidator          *CustomValidator.UserValidator
	UserRepository         *repository.UserRepository
	AddressRepository      *repository.AddressRepository
	UserSicknessRepository *repository.UserSicknessRepository
}

func NewUserUseCase(db *gorm.DB, logger *logrus.Logger, validate *CustomValidator.UserValidator,
	userRepository *repository.UserRepository, addressRepository *repository.AddressRepository, userSicknessRepository *repository.UserSicknessRepository) *UserUseCase {
	return &UserUseCase{
		DB:                     db,
		Log:                    logger,
		UserValidator:          validate,
		UserRepository:         userRepository,
		AddressRepository:      addressRepository,
		UserSicknessRepository: userSicknessRepository,
	}
}

func (c *UserUseCase) Create(ctx context.Context, request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.UserValidator.ValidateCreateUser(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	existingUser := new(entity.User)
	err := c.UserRepository.FindByEmail(tx, existingUser, request.Email)
	if err == nil {
		// user ditemukan -> email sudah dipakai
		c.Log.Warnf("Email already exists : %s", request.Email)
		return nil, fiber.ErrConflict
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// error database
		c.Log.Warnf("Failed query user by email : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// dijalankan apabila terdapat error gorm.ErrRecordNotFound
	addresses := make([]entity.Address, len(request.Addresses))
	for i, addr := range request.Addresses {
		addresses[i] = entity.Address{
			ID:         helper.GenerateULID(),
			Address:    addr.Address,
			City:       addr.City,
			PostalCode: addr.PostalCode,
		}
	}
	user := &entity.User{
		Id:        helper.GenerateULID(),
		Name:      request.Name,
		Email:     request.Email,
		Addresses: addresses,
	}

	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToCreateResponse(user), nil
}

func (c *UserUseCase) GetUserAndAddress(ctx context.Context, id string) (*model.UserWithAddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := c.UserRepository.FindByIdWithAddress(tx, user, id); err != nil {
		c.Log.Warnf("Failed get user and address : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserWithAddressToResponse(user), nil
}

func (c *UserUseCase) UpdateUser(ctx context.Context, request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.UserValidator.ValidateUpdateUser(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err := c.UserRepository.FindById(tx, user, request.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Log.Warnf("Failed find user by id: %+v", err)
			return nil, fiber.ErrNotFound
		}
		c.Log.Warnf("Failed to find user by id: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if request.Name != nil {
		user.Name = *request.Name
	}

	if request.Email != nil {
		user.Email = *request.Email
	}

	if err := c.UserRepository.Update(tx, user); err != nil {
		c.Log.Warnf("Failed update user : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.UpdateUserResponse{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (c *UserUseCase) DeleteUser(ctx context.Context, id string) (*model.DeleteUserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	user.Id = id
	if err := c.AddressRepository.DeleteByUserId(tx, user); err != nil {
		c.Log.Warnf("Failed to delete address : %+v", err)
		return nil, fiber.ErrInternalServerError
	}
	if err := c.UserSicknessRepository.DeleteByUserId(tx, user); err != nil {
		c.Log.Warnf("Failed to delete user sickness : %+v", err)
		return nil, fiber.ErrInternalServerError
	}
	if err := c.UserRepository.Delete(tx, user); err != nil {
		c.Log.Warnf("Failed to delete user : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.DeleteUserResponse{Message: "User deleted"}, nil
}

func (c *UserUseCase) FindUserByIdWithSicknesses(ctx context.Context, req *model.GetUserByIdWithSicknessRequest) (*model.GetUserByIdWithSicknessResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := c.UserRepository.FindByIdWithSickness(tx, user, req.Id, req.StartDate, req.EndDate); err != nil {
		c.Log.Warnf("Failed to get User by ID with sicknesses: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToUserWithSicknessesResponse(user), nil
}
