package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HealthController provides interface for health controller
type HealthController interface {
	LoadRoutes(*mux.Router)
	HealthCheck(http.ResponseWriter, *http.Request)
}
