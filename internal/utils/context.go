package utils

import "context"

type contextKey string

func (c contextKey) String() string {
	return "lincoln context key " + string(c)
}

var (
	contextKeyArticlesLimit  = contextKey("articles-limit")
	contextKeyArticlesOffset = contextKey("articles-offset")
)

// GetArticlesLimit ...
func GetArticlesLimit(ctx context.Context) (int, bool) {
	limit, ok := ctx.Value(contextKeyArticlesLimit).(int)
	return limit, ok
}

// SetArticlesLimit ...
func SetArticlesLimit(ctx context.Context, limit int) context.Context {
	return context.WithValue(ctx, contextKeyArticlesLimit, limit)
}

// GetArticlesOffset ...
func GetArticlesOffset(ctx context.Context) (int, bool) {
	offset, ok := ctx.Value(contextKeyArticlesOffset).(int)
	return offset, ok
}

// SetArticlesOffset ...
func SetArticlesOffset(ctx context.Context, offset int) context.Context {
	return context.WithValue(ctx, contextKeyArticlesOffset, offset)
}
