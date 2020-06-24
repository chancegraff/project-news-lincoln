package services

import (
	"context"
	"time"

	"github.com/pronuu/lincoln/internal/errors"
	"github.com/pronuu/lincoln/internal/models"
)

// Create ...
func (s *services) Create(ctx context.Context, article *models.Article) (output *models.Article, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.NilPointer()
			output = nil
			return
		}
	}()
	if article.Title == "" || article.URL == "" || article.Thumbnail == "" || article.PublishedAt.IsZero() {
		err = errors.MissingKey()
		return
	}
	output = article
	output.CreatedAt = time.Now()
	err = s.Store.Database.Model(article).Create(output).Error
	if err != nil {
		output = nil
		return
	}
	return
}
