package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhyaniarun1993/foody-common/logger"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/dto"
	"github.com/dhyaniarun1993/foody-customer-service/services"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type customerController struct {
	customerService services.CustomerService
	logger          *logger.Logger
	validate        *validator.Validate
}

// NewCustomerController initialize customer endpoint
func NewCustomerController(customerService services.CustomerService,
	logger *logger.Logger, validate *validator.Validate) Controller {

	return &customerController{customerService, logger, validate}
}

func (controller *customerController) LoadRoutes(router *mux.Router) {
	router.HandleFunc("/v1/customers", controller.create).Methods("POST")
}

func (controller *customerController) create(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateCustomerRequest
	var requestBody dto.CreateCustomerRequestBody
	ctx := r.Context()

	logger := controller.logger.WithContext(ctx)
	decodingError := json.NewDecoder(r.Body).Decode(&requestBody)
	if decodingError != nil {
		errorMsg := "Invalid request"
		logger.WithError(decodingError).Error(errorMsg)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"message": %q}`, errorMsg)
		return
	}

	request.Body = requestBody
	validationError := request.Validate(controller.validate)
	if validationError != nil {
		logger.WithError(validationError).Error("Invalid request body")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(validationError.StatusCode())
		fmt.Fprintf(w, `{"message": %q}`, validationError.Error())
		return
	}

	result, serviceError := controller.customerService.Create(ctx, request)
	if serviceError != nil {
		logger.WithError(serviceError).Error("Got Error from Service")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(serviceError.StatusCode())
		fmt.Fprintf(w, `{"message": %q}`, serviceError.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
