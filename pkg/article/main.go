package article

import (
	"github.com/go-kit/kit/log"
	"github.com/pronuu/lincoln/internal/database"
	"github.com/pronuu/lincoln/pkg/article/endpoints"
	"github.com/pronuu/lincoln/pkg/article/middlewares"
	routes "github.com/pronuu/lincoln/pkg/article/routes/http"
	"github.com/pronuu/lincoln/pkg/article/services"
)

// NewArticleService ...
func NewArticleService(str *database.Store, lgr log.Logger) routes.Routes {
	mdl := middlewares.NewMiddleware(lgr)
	svc := services.NewServices(str)
	svc = mdl.BindServices(svc)
	end := endpoints.NewEndpoints(svc)
	end = mdl.BindEndpoints(end)
	rts := routes.NewRoutes(end)
	return rts
}
