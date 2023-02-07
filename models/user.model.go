package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id,omitempty"`
	Username  string    `gorm:"unique;type:varchar(255);not null" json:"username,omitempty"`
	Email     string    `gorm:"unique;type:varchar(255);not null;primary_key" json:"email,omitempty"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt,omitempty"`
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type CreateUserSchema struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserSchema struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
