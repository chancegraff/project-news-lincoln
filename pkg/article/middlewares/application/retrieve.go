package application

import (
	"context"
	"time"

	"github.com/pronuu/lincoln/internal/models"
)

// Retrieve ...
func (m *Middleware) Retrieve(ctx context.Context, article *models.Article) (output *models.Article, err error) {
	defer func(begin time.Time) {
		m.logger.Log(
			"method", "Retrieve",
			"input", article.String(),
			"output", output.String(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = m.next.Retrieve(ctx, article)
	return
}
