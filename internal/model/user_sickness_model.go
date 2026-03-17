package model

import "time"

type CreateUserSicknessRequest struct {
	UserId     string `json:"user_id" validate:"required"`
	SicknessId string `json:"sickness_id" validate:"required"`
}
type CreateUserSicknessResponse struct {
	Id          string    `json:"id"`
	UserId      string    `json:"user_id"`
	SicknessId  string    `json:"sickness_id"`
	DiagnosedAt time.Time `json:"diagnosed_at"`
}

type UserSicknessResponseForSickness struct {
	Id          string    `json:"sickness_id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	DiagnosedAt time.Time `json:"diagnosed_at"`
}
type UserSicknessResponseForUser struct {
	Id          string    `json:"user_id"`
	Name        string    `json:"user_name"`
	Email       string    `json:"email"`
	DiagnosedAt time.Time `json:"diagnosed_at"`
}

type GetUserSicknessByDiagnosedDateRequest struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
type GetUserSicknessByDiagnosedDateResponse struct {
	UserId       string    `json:"user_id"`
	SicknessId   string    `json:"sickness_id"`
	UserName     string    `json:"user_name"`
	Email        string    `json:"email"`
	SicknessName string    `json:"sickness_name"`
	Description  string    `json:"description,omitempty"`
	DiagnosedAt  time.Time `json:"diagnosed_at"`
}

type CountDiagnosedBySicknessIdRequest struct {
	Id string `params:"id"`
}
type CountDiagnosedBySicknessIdResponse struct {
	Id     string `json:"sickness_id"`
	Name   string `json:"name"`
	Counts int64 `json:"diagnosed_count"`
}
