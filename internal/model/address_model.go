package model

type CreateAddressRequest struct {
	Address    string `json:"address" validate:"required"`
	City       string `json:"city"  validate:"required"`
	PostalCode string `json:"postal_code"  validate:"required"`
}

type AddressResponse struct {
	ID         string `json:"id,omitempty"`
	Address    string `json:"address,omitempty"`
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

type UpdateAddressRequest struct {
	ID         string `params:"id" validate:"required"`
	Address    string `json:"address,omitempty"`
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

type UpdateAddressResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

type DeleteAddressRequest struct {
	AddressID string `params:"address_id" validate:"required"`
}

type DeleteAddressResponse struct {
	Message string `json:"message"`
}
