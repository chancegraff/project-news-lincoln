package application

import (
	"github.com/go-kit/kit/log"
	"github.com/pronuu/lincoln/pkg/vote/services"
)

// Middleware ...
type Middleware struct {
	next   services.Services
	logger log.Logger
}

// MakeMiddleware ...
func MakeMiddleware(logger log.Logger) services.Middleware {
	return func(next services.Services) services.Services {
		return &Middleware{
			next,
			log.With(logger, "layer", "application"),
		}
	}
}
