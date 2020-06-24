package http

import (
	web "net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pronuu/lincoln/internal/utils"
	"github.com/pronuu/lincoln/pkg/article/endpoints"
	"github.com/pronuu/lincoln/pkg/article/transports"
)

// MakeEnumerateRoute ...
func MakeEnumerateRoute(endpoints endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.EnumerateEndpoint,
		transports.DecodeEnumerateHTTPRequest,
		transports.EncodeEnumerateHTTPResponse,
	)
}

// Enumerate ...
func (r *Routes) Enumerate(rw web.ResponseWriter, rq *web.Request) {
	utils.SetCORSHeaders(rw)

	if rq.Method == "OPTIONS" {
		return
	}

	r.EnumerateRoute.ServeHTTP(rw, rq)
}
