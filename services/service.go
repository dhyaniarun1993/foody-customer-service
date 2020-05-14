package services

import (
	"context"

	"github.com/dhyaniarun1993/foody-common/errors"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/dto"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/models"
)

// HealthService provides interface for health service
type HealthService interface {
	HealthCheck(context.Context) errors.AppError
}

// CustomerService provides interface for customer service
type CustomerService interface {
	Create(ctx context.Context, request dto.CreateCustomerRequest) (dto.CreateCustomerResponse, errors.AppError)
	Get(ctx context.Context, request dto.GetCustomerRequest) (models.Customer, errors.AppError)
}
