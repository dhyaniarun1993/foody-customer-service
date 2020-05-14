package models

// Customer provides the model definition for Customer
type Customer struct {
	ID          int64  `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}
