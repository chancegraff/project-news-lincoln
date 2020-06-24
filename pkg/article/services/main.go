package services

import (
	"context"

	"github.com/pronuu/lincoln/internal/database"
	"github.com/pronuu/lincoln/internal/models"
)

// Services implements the collector interface
type Services interface {
	Create(ctx context.Context, article *models.Article) (*models.Article, error)
	Enumerate(ctx context.Context, articles models.ArticleArray, offset, limit int) (models.ArticleArray, error)
	Retrieve(ctx context.Context, article *models.Article) (*models.Article, error)
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
