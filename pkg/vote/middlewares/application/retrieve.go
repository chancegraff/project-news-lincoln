package application

import (
	"context"
	"time"

	"github.com/pronuu/lincoln/internal/models"
)

// Retrieve ...
func (m *Middleware) Retrieve(ctx context.Context, vote *models.Vote) (output *models.Vote, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Retrieve",
			"input", vote.String(),
			"output", output.String(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Retrieve(ctx, vote)
	return
}
