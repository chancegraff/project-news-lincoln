package application

import (
	"context"
	"time"

	"github.com/pronuu/lincoln/internal/models"
)

// Delete ...
func (m *Middleware) Delete(ctx context.Context, vote *models.Vote) (output *models.Vote, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Delete",
			"input", vote.String(),
			"output", output.String(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Delete(ctx, vote)
	return
}
