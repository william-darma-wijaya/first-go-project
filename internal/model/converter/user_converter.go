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