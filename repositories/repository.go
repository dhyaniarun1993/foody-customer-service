package repositories

import (
	"context"

	"github.com/dhyaniarun1993/foody-common/errors"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/models"
)

// HealthRepository provides interface for Health repository
type HealthRepository interface {
	HealthCheck(context.Context) errors.AppError
}

// CustomerRepository provides interface for Customer repository
type CustomerRepository interface {
	Create(ctx context.Context, customer models.Customer) (int64, errors.AppError)
	GetByPhoneNumber(ctx context.Context, PhoneNumber string) (models.Customer, errors.AppError)
	GetByEmail(ctx context.Context, email string) (models.Customer, errors.AppError)
}
