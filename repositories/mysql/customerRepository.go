package mysql

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	instrumentedSQL "github.com/dhyaniarun1993/foody-common/datastore/sql"
	"github.com/dhyaniarun1993/foody-common/errors"
	"github.com/dhyaniarun1993/foody-customer-service/repositories"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/models"
	"github.com/go-sql-driver/mysql"
)

type customerRepository struct {
	*instrumentedSQL.DB
}

// NewCustomerRepository creates and returns mysql customer repository
func NewCustomerRepository(db *instrumentedSQL.DB) repositories.CustomerRepository {
	return &customerRepository{db}
}

func (db *customerRepository) Create(ctx context.Context, customer models.Customer) (int64, errors.AppError) {
	timedCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	query := "INSERT INTO customer SET phone_number = ?, email = ?, first_name = ?, last_name = ?, status = ?"
	res, insertErr := db.ExecContext(timedCtx, query, customer.PhoneNumber, customer.Email, customer.FirstName, customer.LastName, customer.Status)
	if insertErr != nil {
		mysqlError, ok := insertErr.(*mysql.MySQLError)
		if ok && mysqlError.Number == 1062 {
			return 0, errors.NewAppError("Phone number or Email already linked to an account", http.StatusUnprocessableEntity, insertErr)
		}
		return 0, errors.NewAppError("Unable to create Customer", http.StatusInternalServerError, insertErr)
	}

	customerID, fetchIDError := res.LastInsertId()
	if fetchIDError != nil {
		return 0, errors.NewAppError("Unable to fetch Id", http.StatusInternalServerError, fetchIDError)
	}
	return customerID, nil
}

func (db *customerRepository) GetByPhoneNumber(ctx context.Context, PhoneNumber string) (models.Customer, errors.AppError) {
	customer := models.Customer{}
	timedCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	query := "SELECT id, first_name, last_name, phone_number, email, status, created_at, updated_at FROM customer WHERE phone_number = ?"
	row := db.QueryRowContext(timedCtx, query, PhoneNumber)

	err := row.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.PhoneNumber, &customer.Email, &customer.Status, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return customer, errors.NewAppError("Something went wrong", http.StatusInternalServerError, err)
	}
	return customer, nil
}

func (db *customerRepository) GetByEmail(ctx context.Context, email string) (models.Customer, errors.AppError) {
	var customer models.Customer
	timedCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	query := "SELECT id, first_name, last_name, phone_number, email, status, created_at, updated_at FROM customer WHERE email = ?"
	row := db.QueryRowContext(timedCtx, query, email)

	err := row.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.PhoneNumber, &customer.Email, &customer.Status, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return customer, errors.NewAppError("Something went wrong", http.StatusInternalServerError, err)
	}
	return customer, nil
}
