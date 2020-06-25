package services

import (
	"context"

	"github.com/pronuu/lincoln/internal/errors"
	"github.com/pronuu/lincoln/internal/models"
)

// Retrieve ...
func (s *services) Retrieve(ctx context.Context, vote *models.Vote) (output *models.Vote, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.NilPointer()
			output = nil
			return
		}
	}()
	if vote.ID == 0 {
		err = errors.MissingKey()
		return
	}
	output = vote
	err = s.Store.Database.Model(vote).Where(vote).First(output).Error
	if err != nil {
		output = nil
		return
	}
	return
}
