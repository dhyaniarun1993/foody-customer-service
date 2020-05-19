package client

import (
	"context"

	"github.com/dhyaniarun1993/foody-common/errors"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/dto"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/models"
)

// Client provides interface definition for Customer Client
type Client interface {
	InternalCreateCustomer(ctx context.Context,
		body dto.CreateCustomerRequestBody) (dto.CreateCustomerResponse, errors.AppError)
	InternalGetCustomer(ctx context.Context,
		query dto.GetCustomerRequestQuery) (models.Customer, errors.AppError)
}
