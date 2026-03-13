package CustomValidator

import (
	"testing"

	"first-go-project/internal/model"

	"github.com/go-playground/validator/v10"
)

func TestUserValidator_ValidateCreateUser(t *testing.T) {
	validate := validator.New()
	userValidator := NewUserValidator(validate)

	tests := []struct {
		name    string
		request *model.CreateUserRequest
		wantErr bool
	}{
		{
			name: "valid request",
			request: &model.CreateUserRequest{
				Name:  "Alice",
				Email: "alice@example.com",
				Addresses: []model.CreateAddressRequest{
					{
						Address:    "123 Main St",
						City:       "Bangkok",
						PostalCode: "10110",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing email",
			request: &model.CreateUserRequest{
				Name: "Alice",
				Addresses: []model.CreateAddressRequest{
					{
						Address:    "123 Main St",
						City:       "Bangkok",
						PostalCode: "10110",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "name too short",
			request: &model.CreateUserRequest{
				Name:  "Al",
				Email: "alice@example.com",
				Addresses: []model.CreateAddressRequest{
					{
						Address:    "123 Main St",
						City:       "Bangkok",
						PostalCode: "10110",
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userValidator.ValidateCreateUser(tt.request)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateCreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
