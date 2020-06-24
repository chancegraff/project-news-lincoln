package application

import (
	"context"
	"time"

	"github.com/pronuu/lincoln/internal/models"
)

// Enumerate ...
func (m *Middleware) Enumerate(ctx context.Context, articles models.ArticleArray, offset, limit int) (output models.ArticleArray, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Create",
			"input", articles.String(),
			"output", output.String(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Enumerate(ctx, articles, offset, limit)
	return
}
