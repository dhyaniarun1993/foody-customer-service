package services

import (
	"context"
	"reflect"
	"strings"

	"github.com/dhyaniarun1993/foody-common/errors"
	"github.com/dhyaniarun1993/foody-common/logger"
	"github.com/dhyaniarun1993/foody-customer-service/constants"
	"github.com/dhyaniarun1993/foody-customer-service/repositories"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/dto"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/models"
)

type customerService struct {
	customerRepository repositories.CustomerRepository
	logger             *logger.Logger
}

// NewCustomerService creates and return customer service object
func NewCustomerService(customerRepository repositories.CustomerRepository,
	logger *logger.Logger) CustomerService {

	return &customerService{customerRepository, logger}
}

func (service *customerService) Create(ctx context.Context,
	request dto.CreateCustomerRequest) (dto.CreateCustomerResponse, errors.AppError) {

	customer := models.Customer{
		PhoneNumber: request.Body.PhoneNumber,
		Email:       strings.ToLower(request.Body.Email),
		FirstName:   strings.Title(request.Body.FirstName),
		LastName:    strings.Title(request.Body.LastName),
		Status:      constants.CustomerStatusActive,
	}

	customerID, createCustomerError := service.customerRepository.Create(ctx, customer)
	response := dto.CreateCustomerResponse{
		ID: customerID,
	}
	return response, createCustomerError
}

func (service *customerService) Get(ctx context.Context,
	request dto.GetCustomerRequest) (models.Customer, errors.AppError) {

	customer, err := service.customerRepository.GetByPhoneNumber(ctx, request.Query.PhoneNumber)
	if err != nil {
		return models.Customer{}, err
	}

	if reflect.DeepEqual(customer, models.Customer{}) {
		return models.Customer{}, errors.NewAppError("Unable to find customer", errors.StatusNotFound, nil)
	}

	return customer, nil
}
