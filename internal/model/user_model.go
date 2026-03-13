package model

import "time"

type CreateUserRequest struct {
	Name      string                 `json:"name" validate:"required,min=3,max=100"`
	Email     string                 `json:"email" validate:"required,email"`
	Addresses []CreateAddressRequest `json:"addresses" validate:"required,dive"`
}

type CreateUserResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Addresses []AddressResponse `json:"addresses"`
}

type GetUserRequest struct {
	ID int64 `params:"id" validate:"required"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserWithAddressResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Addresses []AddressResponse `json:"addresses"`
}

type UpdateUserRequest struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"omitempty,min=3,max=100"`
	Email string `json:"email" validate:"omitempty,email"`
}

type UpdateUserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserRequest struct {
	ID string `params:"id" validate:"required"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
