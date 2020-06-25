package services

import (
	"context"

	"github.com/pronuu/lincoln/internal/errors"
	"github.com/pronuu/lincoln/internal/models"
)

// Enumerate ...
func (s *services) Enumerate(ctx context.Context, votes models.VoteArray, offset, limit int) (output models.VoteArray, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.NilPointer()
			output = nil
			return
		}
	}()
	db := s.Store.Database.Model(models.Vote{})
	if len(votes) > 0 {
		voteIDs := make([]uint, 0)
		for _, vote := range votes {
			voteIDs = append(voteIDs, vote.ID)
		}
		db = db.Where("id IN (?)", voteIDs)
	}
	if limit == 0 {
		limit = 20
	}
	err = db.Where("deleted_at IS nil").Limit(limit).Offset(offset).Order("published_at desc").Scan(output).Error
	if err != nil {
		output = nil
		return
	}
	return
}
