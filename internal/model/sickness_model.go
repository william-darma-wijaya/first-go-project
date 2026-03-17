package model

import "time"

type CreateSicknessRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description"`
}
type CreateSicknessResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type GetSicknessByIdRequest struct {
	ID string `params:"id" validate:"required"`
}
type GetSicknessByIdResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateSicknessRequest struct {
	Id          string  `json:"id" validate:"required"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
type UpdateSicknessResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DeleteSicknessRequest struct {
	Id string `params:"id" validate:"required"`
}
type DeleteSicknessResponse struct {
	Message string `json:"message"`
}

type SicknessResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
type GetSicknessesByNameRequest struct {
	Name string `query:"name" validate:"required"`
}
type GetSicknessesByNameResponse struct {
	Sicknesses []SicknessResponse `json:"sicknesses"`
}

type GetSicknessByIdWithUserRequest struct {
	Id        string     `params:"id" validate:"required"`
	StartDate *time.Time `query:"start_date"`
	EndDate   *time.Time `query:"end_date"`
}
type GetSicknessByIdWithUserResponse struct {
	Id           string                        `json:"id"`
	SicknessName string                        `json:"sickness_name"`
	Description  string                        `json:"description"`
	Users        []UserSicknessResponseForUser `json:"users"`
}
