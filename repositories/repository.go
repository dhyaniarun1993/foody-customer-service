package repositories

import (
	"context"

	"github.com/dhyaniarun1993/foody-common/errors"
)

// HealthRepository provides interface for Health repositories
type HealthRepository interface {
	HealthCheck(context.Context) errors.AppError
}
