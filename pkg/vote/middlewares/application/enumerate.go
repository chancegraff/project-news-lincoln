package application

import (
	"context"
	"time"

	"github.com/pronuu/lincoln/internal/models"
)

// Enumerate ...
func (m *Middleware) Enumerate(ctx context.Context, votes models.VoteArray, offset, limit int) (output models.VoteArray, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Enumerate",
			"input", votes.String(),
			"output", output.String(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Enumerate(ctx, votes, offset, limit)
	return
}
