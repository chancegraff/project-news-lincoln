package services

import (
	"context"

	"github.com/pronuu/lincoln/internal/database"
	"github.com/pronuu/lincoln/internal/models"
)

// Services implements the collector interface
type Services interface {
	Create(ctx context.Context, vote *models.Vote) (*models.Vote, error)
	Enumerate(ctx context.Context, votes models.VoteArray, offset, limit int) (models.VoteArray, error)
	Retrieve(ctx context.Context, vote *models.Vote) (*models.Vote, error)
	Delete(ctx context.Context, vote *models.Vote) (*models.Vote, error)
}

type services struct {
	Store *database.Store
}

// NewServices instantiates the service with a connection to the database
func NewServices(s *database.Store) Services {
	return &services{
		Store: s,
	}
}

// Middleware is a chainable middleware for Service
type Middleware func(Services) Services
