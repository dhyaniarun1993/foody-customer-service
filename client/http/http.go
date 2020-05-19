package http

import (
	"net/http"
	"time"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"

	"github.com/dhyaniarun1993/foody-customer-service/client"
)

// Configuration provides the configuration parameters for customer http client
type Configuration struct {
	APIEndpoint string        `required:"true" split_words:"true"`
	APITimeout  time.Duration `required:"true" split_words:"true"`
}

type httpClient struct {
	http.Client
	config Configuration
	tracer opentracing.Tracer
}

// NewCustomerClient creates and return http customer client
func NewCustomerClient(config Configuration, tracer opentracing.Tracer) client.Client {
	client := http.Client{
		Timeout:   config.Timeout,
		Transport: &nethttp.Transport{},
	}
	return &httpClient{client, config, tracer}
}
