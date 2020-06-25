package application

import (
	"context"
	"time"

	"github.com/pronuu/lincoln/internal/models"
)

// Create ...
func (m *Middleware) Create(ctx context.Context, vote *models.Vote) (output *models.Vote, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Create",
			"input", vote.String(),
			"output", output.String(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Create(ctx, vote)
	return
}
