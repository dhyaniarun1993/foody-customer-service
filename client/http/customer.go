package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/opentracing-contrib/go-stdlib/nethttp"

	"github.com/dhyaniarun1993/foody-common/errors"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/dto"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/models"
)

func (client *httpClient) InternalCreateCustomer(ctx context.Context,
	body dto.CreateCustomerRequestBody) (dto.CreateCustomerResponse, errors.AppError) {

	url := client.config.Endpoint + "/v1/internal/customers"
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(body)
	req, createRequestError := http.NewRequest("POST", url, requestBody)
	if createRequestError != nil {
		return dto.CreateCustomerResponse{}, errors.NewAppError("Something went wrong", http.StatusInternalServerError, createRequestError)
	}

	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)
	newReq, ht := nethttp.TraceRequest(client.tracer, req)
	defer ht.Finish()

	res, clientErr := client.Do(newReq)
	if clientErr != nil {
		return dto.CreateCustomerResponse{}, errors.NewAppError("Something went wrong", http.StatusInternalServerError, createRequestError)
	}
	defer res.Body.Close()

	bodyBytes, bodyReadError := ioutil.ReadAll(res.Body)
	if bodyReadError != nil {
		return dto.CreateCustomerResponse{}, errors.NewAppError("Something went wrong", http.StatusInternalServerError, bodyReadError)
	}

	if res.StatusCode != http.StatusCreated {
		APIErr := client.generateError(bodyBytes, res.StatusCode)
		return dto.CreateCustomerResponse{}, APIErr
	}

	var response dto.CreateCustomerResponse
	unmarshalError := json.Unmarshal(bodyBytes, &response)
	if unmarshalError != nil {
		return dto.CreateCustomerResponse{}, errors.NewAppError("Something went wrong", http.StatusInternalServerError, unmarshalError)
	}
	return response, nil
}

func (client *httpClient) InternalGetCustomer(ctx context.Context,
	query dto.GetCustomerRequestQuery) (models.Customer, errors.AppError) {

	url := client.config.Endpoint + "/v1/internal/customers"
	req, createRequestError := http.NewRequest("GET", url, nil)
	if createRequestError != nil {
		return models.Customer{}, errors.NewAppError("Something went wrong", http.StatusInternalServerError, createRequestError)
	}

	q := req.URL.Query()
	q.Add("phoneNumber", query.PhoneNumber)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)
	req, ht := nethttp.TraceRequest(client.tracer, req)
	defer ht.Finish()

	res, clientErr := client.Do(req)
	if clientErr != nil {
		return models.Customer{}, errors.NewAppError("Something went wrong", http.StatusInternalServerError, createRequestError)
	}
	defer res.Body.Close()

	bodyBytes, bodyReadError := ioutil.ReadAll(res.Body)
	if bodyReadError != nil {
		return models.Customer{}, errors.NewAppError("Something went wrong", http.StatusInternalServerError, bodyReadError)
	}

	if res.StatusCode != http.StatusOK {
		APIErr := client.generateError(bodyBytes, res.StatusCode)
		return models.Customer{}, APIErr
	}

	var response models.Customer
	unmarshalError := json.Unmarshal(bodyBytes, &response)
	if unmarshalError != nil {
		return models.Customer{}, errors.NewAppError("Something went wrong", http.StatusInternalServerError, unmarshalError)
	}
	return response, nil
}
