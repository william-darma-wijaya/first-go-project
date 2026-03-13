package converter

import (
	"first-go-project/internal/entity"
	"first-go-project/internal/model"
)

func AddressToUpdateResponse(address *entity.Address) *model.UpdateAddressResponse {
	return &model.UpdateAddressResponse{
		ID:         address.ID,
		UserID:     address.UserId,
		Address:    address.Address,
		City:       address.City,
		PostalCode: address.PostalCode,
	}
}
