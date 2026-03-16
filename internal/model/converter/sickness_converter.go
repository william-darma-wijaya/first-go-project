package converter

import (
	"first-go-project/internal/entity"
	"first-go-project/internal/model"
)

func SicknessToCreateResponse(sickness *entity.Sickness) (response *model.CreateSicknessResponse) {
	return &model.CreateSicknessResponse{
		Id: sickness.Id,
		Name: sickness.Name,
		Description: sickness.Description,
	}
}

func SicknessToFindByIdResponse(sickness *entity.Sickness) (response *model.GetSicknessByIdResponse) {
	return &model.GetSicknessByIdResponse{
		Id: sickness.Id,
		Name: sickness.Name,
		Description: sickness.Description,
	}
}

func SicknessToUpdateResponse(sickness *entity.Sickness) (response *model.UpdateSicknessResponse) {
	return &model.UpdateSicknessResponse{
		Id: sickness.Id,
		Name: sickness.Name,
		Description: sickness.Description,
		UpdatedAt: sickness.UpdatedAt,
	}
}