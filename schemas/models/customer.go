package models

import (
	"time"

	"github.com/dhyaniarun1993/foody-customer-service/constants"
)

// Customer provides the model definition for Customer
type Customer struct {
	ID          int64     `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// IsActive check if customer status is active
func (customer *Customer) IsActive() bool {
	if customer.Status == constants.CustomerStatusActive {
		return true
	}
	return false
}
