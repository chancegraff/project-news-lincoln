package services

import (
	"context"

	"github.com/pronuu/lincoln/internal/errors"
	"github.com/pronuu/lincoln/internal/models"
)

// Enumerate ...
func (s *services) Enumerate(ctx context.Context, articles models.ArticleArray, offset, limit int) (output models.ArticleArray, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.NilPointer()
			output = nil
			return
		}
	}()
	db := s.Store.Database.Model(models.Article{})
	if len(articles) > 0 {
		articleIDs := make([]uint, 0)
		for _, article := range articles {
			articleIDs = append(articleIDs, article.ID)
		}
		db = db.Where("id IN (?)", articleIDs)
	}
	if limit == 0 {
		limit = 20
	}
	rows, err := db.Limit(limit).Offset(offset).Order("published_at desc").Rows()
	if err != nil {
		return
	}
	err = s.Store.Database.ScanRows(rows, output)
	if err != nil {
		output = nil
		return
	}
	return
}
