package middlewares

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pronuu/lincoln/pkg/article/endpoints"
	"github.com/pronuu/lincoln/pkg/article/middlewares/application"
	"github.com/pronuu/lincoln/pkg/article/middlewares/transport"
	"github.com/pronuu/lincoln/pkg/article/services"
)

// NewMiddleware ...
func NewMiddleware(logger log.Logger) Middleware {
	logger = level.Info(logger)
	return Middleware{
		ApplicationLogger: application.MakeMiddleware(logger),
		TransportLogger:   transport.MakeMiddleware(logger),
	}
}

// Middleware ...
type Middleware struct {
	ApplicationLogger services.Middleware
	TransportLogger   endpoints.Middleware
}

// BindServices ...
func (mw *Middleware) BindServices(svc services.Services) services.Services {
	return mw.ApplicationLogger(svc)
}

// BindEndpoints ...
func (mw *Middleware) BindEndpoints(end endpoints.Endpoints) endpoints.Endpoints {
	return mw.TransportLogger(end)
}
