package services

import (
	"context"
	"time"

	"github.com/pronuu/lincoln/internal/errors"
	"github.com/pronuu/lincoln/internal/models"
)

// Create ...
func (s *services) Create(ctx context.Context, vote *models.Vote) (output *models.Vote, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.NilPointer()
			output = nil
			return
		}
	}()
	if vote.ArticleID == 0 || vote.UserID == 0 {
		err = errors.MissingKey()
		return
	}
	output = vote
	output.CreatedAt = time.Now()
	err = s.Store.Database.Model(vote).Create(output).Error
	if err != nil {
		output = nil
		return
	}
	return
}
