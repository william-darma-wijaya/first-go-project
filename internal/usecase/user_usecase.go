package usecase

import (
	"context"
	"errors"
	"first-go-project/internal/CustomValidator"
	"first-go-project/internal/gateway/caching"
	"first-go-project/internal/gateway/messaging"

	// "first-go-project/internal/config"

	// "first-go-project/internal/config"
	"first-go-project/internal/entity"
	"first-go-project/internal/helper"
	"first-go-project/internal/model"
	"first-go-project/internal/model/converter"
	"first-go-project/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
	AuthConfig             *entity.AuthConfig
	UserCache              *caching.UserCache
	UserProducer           *messaging.UserProducer
}

func NewUserUseCase(db *gorm.DB, logger *logrus.Logger, validate *CustomValidator.UserValidator, authConfig *entity.AuthConfig,
	userRepository *repository.UserRepository, addressRepository *repository.AddressRepository, userCache *caching.UserCache,
	userSicknessRepository *repository.UserSicknessRepository, userProducer *messaging.UserProducer) *UserUseCase {
	return &UserUseCase{
		DB:                     db,
		Log:                    logger,
		UserValidator:          validate,
		UserRepository:         userRepository,
		AddressRepository:      addressRepository,
		UserSicknessRepository: userSicknessRepository,
		AuthConfig:             authConfig,
		UserCache:              userCache,
		UserProducer:           userProducer,
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

	hashedPass, err := helper.HashPassword(request.Password)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.User{
		Id:        helper.GenerateULID(),
		Name:      request.Name,
		Password:  hashedPass,
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

func (c *UserUseCase) Login(ctx context.Context, request *model.UserLoginRequest) (*model.UserLoginResponse, error) {

	user := &entity.User{}
	err := c.UserRepository.FindByEmail(c.DB.WithContext(ctx), user, request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Log.Warnf("User email not found : %+v", err)
		}
		c.Log.Errorf("Error from database to retrieve user: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if matchPass := helper.ComparePassAndHashed(request.Password, user.Password); matchPass != true {
		c.Log.Warn("Wrong password")
		return nil, fiber.ErrUnauthorized
	}

	token, refreshToken, err := helper.GenerateJWT(c.AuthConfig, user.Id)
	if err != nil {
		c.Log.Warnf("Error generating JWT token: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// simpan user_id ke redis
	err = c.UserCache.Set(ctx, token, user.Id, c.AuthConfig.MinutesExp)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return &model.UserLoginResponse{Token: token, RefreshToken: refreshToken}, nil
}

func (c *UserUseCase) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {

	token, err := jwt.Parse(request.Token, func(t *jwt.Token) (any, error) {
		return []byte(c.AuthConfig.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, fiber.ErrUnauthorized
	}

	userId, err := c.UserCache.Get(ctx, request.Token)
	if err != nil {
		return nil, fiber.ErrUnauthorized
	}

	return &model.Auth{ID: userId}, nil
}

func (c *UserUseCase) Logout(ctx context.Context, token string) (*model.UserLogoutResponse, error) {
	if err := c.UserCache.Delete(ctx, token); err != nil {
		c.Log.Warnf("Failed to delete token from cache: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.UserLogoutResponse{Message: "User Logged Out"}, nil
}

func (c *UserUseCase) RefreshJWTToken(ctx context.Context, request *model.RefreshTokenRequest) (*model.RefreshTokenResponse, error) {
	refreshToken, err := jwt.Parse(request.RefreshToken, func(t *jwt.Token) (any, error) {
		return []byte(c.AuthConfig.Secret), nil
	})

	if err != nil || !refreshToken.Valid {
		return nil, fiber.ErrUnauthorized
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fiber.ErrUnauthorized
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, fiber.ErrUnauthorized
	}

	token, newRefreshToken, err := helper.GenerateJWT(c.AuthConfig, userID)
	if err != nil {
		c.Log.Warnf("Error generating JWT token: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// simpan user_id ke redis
	err = c.UserCache.Set(ctx, token, userID, c.AuthConfig.MinutesExp)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return &model.RefreshTokenResponse{Token: token, RefreshToken: newRefreshToken}, nil
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
