package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/lincoln/pkg/article/services"
)

// Endpoints ...
type Endpoints struct {
	CreateEndpoint    endpoint.Endpoint
	EnumerateEndpoint endpoint.Endpoint
	RetrieveEndpoint  endpoint.Endpoint
}

// NewEndpoints ...
func NewEndpoints(s services.Services) Endpoints {
	return Endpoints{
		CreateEndpoint:    MakeCreateEndpoint(s),
		EnumerateEndpoint: MakeEnumerateEndpoint(s),
		RetrieveEndpoint:  MakeRetrieveEndpoint(s),
	}
}

// Middleware ...
type Middleware func(Endpoints) Endpoints
