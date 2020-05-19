package http

import (
	"encoding/json"
	"net/http"

	"github.com/dhyaniarun1993/foody-common/errors"
	"github.com/dhyaniarun1993/foody-customer-service/schemas/dto"
)

func (client *httpClient) generateError(body []byte, statusCode int) errors.AppError {
	var errorResponse dto.ErrorResponse
	unmarshalError := json.Unmarshal(body, &errorResponse)
	if unmarshalError != nil {
		return errors.NewAppError("Something went wrong", http.StatusInternalServerError, unmarshalError)
	}
	return errors.NewAppError(errorResponse.Message, statusCode, nil)
}
