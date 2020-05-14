package dto

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"

	"github.com/dhyaniarun1993/foody-common/errors"
)

// CreateCustomerRequestBody provides the schema definition for create customer api request body
type CreateCustomerRequestBody struct {
	PhoneNumber string `json:"phone_number" validate:"required,indiaPhoneNumber"`
	Email       string `json:"email" validate:"required,email"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
}

// CreateCustomerRequest provides the schema definition for create customer api request
type CreateCustomerRequest struct {
	Body CreateCustomerRequestBody `json:"body" validate:"required,dive"`
}

// Validate validates CreateProductRequest
func (dto CreateCustomerRequest) Validate(validate *validator.Validate) errors.AppError {
	var errMsg string
	err := validate.Struct(dto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errMsg = fmt.Sprintf("Invalid value for field '%s'", err.Field())
			break
		}
		return errors.NewAppError(errMsg, errors.StatusBadRequest, err)
	}

	return nil
}

// CreateCustomerResponse provides the schema definition for create customer api response
type CreateCustomerResponse struct {
	ID int64 `json:"id"`
}
