package services

import (
	"context"

	"github.com/pronuu/lincoln/internal/errors"
	"github.com/pronuu/lincoln/internal/models"
)

// Retrieve ...
func (s *services) Retrieve(ctx context.Context, article *models.Article) (output *models.Article, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.NilPointer()
			output = nil
			return
		}
	}()
	if article.ID == 0 {
		err = errors.MissingKey()
		return
	}
	output = article
	err = s.Store.Database.Model(article).Where(article).First(output).Error
	if err != nil {
		output = nil
		return
	}
	return
}
