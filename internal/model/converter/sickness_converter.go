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