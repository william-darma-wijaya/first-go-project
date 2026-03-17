package converter

import (
	"first-go-project/internal/entity"
	"first-go-project/internal/model"
)

func UserToCreateResponse(user *entity.User) *model.CreateUserResponse {
	addresses := make([]model.AddressResponse, len(user.Addresses))
	for i, addr := range user.Addresses {
		addresses[i] = model.AddressResponse{
			ID:      addr.ID,
			Address: addr.Address,
			City:    addr.City,
			PostalCode: addr.PostalCode,
		}
	}

	return &model.CreateUserResponse{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Addresses: addresses,
	}
}

func UserWithAddressToResponse(user *entity.User) *model.UserWithAddressResponse {
	addresses := make([]model.AddressResponse, len(user.Addresses))
	for i, addr := range user.Addresses {
		addresses[i] = model.AddressResponse{
			ID:      addr.ID,
			Address: addr.Address,
			City:    addr.City,
			PostalCode: addr.PostalCode,
		}
	}

	return &model.UserWithAddressResponse{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Addresses: addresses,
	}
}

func UserToUserWithSicknessesResponse(user *entity.User) *model.GetUserByIdWithSicknessResponse {
	response := &model.GetUserByIdWithSicknessResponse{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
	}

	sicknesses := make([]model.UserSicknessResponseForSickness, 0)
	for _, userSickness := range user.Sicknesses {
		sick := model.UserSicknessResponseForSickness{
			Id: userSickness.Id,
			Name: userSickness.Sickness.Name,
			Description: userSickness.Sickness.Description,
			DiagnosedAt: userSickness.DiagnosedAt,
		}
		sicknesses = append(sicknesses, sick)
	}

	response.Sicknesses = sicknesses

	return response
}