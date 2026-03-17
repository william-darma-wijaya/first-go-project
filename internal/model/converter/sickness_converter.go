package converter

import (
	"first-go-project/internal/entity"
	"first-go-project/internal/model"
)

func SicknessToCreateResponse(sickness *entity.Sickness) (*model.CreateSicknessResponse) {
	return &model.CreateSicknessResponse{
		Id: sickness.Id,
		Name: sickness.Name,
		Description: sickness.Description,
	}
}

func SicknessToFindByIdResponse(sickness *entity.Sickness) (*model.GetSicknessByIdResponse) {
	return &model.GetSicknessByIdResponse{
		Id: sickness.Id,
		Name: sickness.Name,
		Description: sickness.Description,
	}
}

func SicknessToUpdateResponse(sickness *entity.Sickness) (*model.UpdateSicknessResponse) {
	return &model.UpdateSicknessResponse{
		Id: sickness.Id,
		Name: sickness.Name,
		Description: sickness.Description,
		UpdatedAt: sickness.UpdatedAt,
	}
}

func SicknessesToFindSicknessesResponse(sicknesses *[]entity.Sickness) (*model.GetSicknessesByNameResponse) {
	var responses []model.SicknessResponse

	for _, sickness := range *sicknesses {
		responses = append(responses, model.SicknessResponse{
			Id: sickness.Id,
			Name: sickness.Name,
			Description: sickness.Description,
		})
	}

	return &model.GetSicknessesByNameResponse{
		Sicknesses: responses,
	}
}

func SicknessToSicknessWithUsersResponse(sickness *entity.Sickness) (*model.GetSicknessByIdWithUserResponse) {
	response := &model.GetSicknessByIdWithUserResponse{
		Id: sickness.Id,
		SicknessName: sickness.Name,
		Description: sickness.Description,
	}

	users := make([]model.UserSicknessResponseForUser, 0)
	for _, userSickness := range sickness.Users {
		s := model.UserSicknessResponseForUser{
			Id: userSickness.User.Id,
			Name: userSickness.User.Name,
			Email: userSickness.User.Email,
			DiagnosedAt: userSickness.DiagnosedAt,
		}
		users = append(users, s)
	}

	response.Users = users

	return response
}