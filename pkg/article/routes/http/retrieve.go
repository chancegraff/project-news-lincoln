package http

import (
	web "net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pronuu/lincoln/internal/utils"
	"github.com/pronuu/lincoln/pkg/article/endpoints"
	"github.com/pronuu/lincoln/pkg/article/transports"
)

// MakeRetrieveRoute ...
func MakeRetrieveRoute(endpoints endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.RetrieveEndpoint,
		transports.DecodeRetrieveHTTPRequest,
		transports.EncodeRetrieveHTTPResponse,
	)
}

// Retrieve ...
func (r *Routes) Retrieve(rw web.ResponseWriter, rq *web.Request) {
	utils.SetCORSHeaders(rw)

	if rq.Method == "OPTIONS" {
		return
	}

	r.RetrieveRoute.ServeHTTP(rw, rq)
}
