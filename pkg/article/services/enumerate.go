package services

import (
	"context"

	"github.com/pronuu/lincoln/internal/errors"
	"github.com/pronuu/lincoln/internal/models"
	"github.com/pronuu/lincoln/internal/utils"
)

// Enumerate ...
func (s *services) Enumerate(ctx context.Context, articles []*models.Article) (output []*models.Article, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.NilPointer()
			output = nil
			return
		}
	}()
	if len(articles) == 0 {
		err = errors.MissingKey()
		return
	}
	limit, ok := utils.GetArticlesLimit(ctx)
	if !ok {
		limit = 20
	}
	offset, ok := utils.GetArticlesOffset(ctx)
	if !ok {
		offset = 0
	}
	rows, err := s.Store.Database.Model(models.Article{}).Where(articles).Order("published_at desc").Limit(limit).Offset(offset).Rows()
	if err != nil {
		output = nil
		return
	}
	output = make([]*models.Article, len(articles))
	err = s.Store.Database.ScanRows(rows, output)
	if err != nil {
		output = nil
		return
	}
	return
}
