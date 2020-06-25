package http

import (
	web "net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pronuu/lincoln/internal/utils"
	"github.com/pronuu/lincoln/pkg/vote/endpoints"
	"github.com/pronuu/lincoln/pkg/vote/transports"
)

// MakeDeleteRoute ...
func MakeDeleteRoute(endpoints endpoints.Endpoints) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.DeleteEndpoint,
		transports.DecodeDeleteHTTPRequest,
		transports.EncodeDeleteHTTPResponse,
	)
}

// Delete ...
func (r *Routes) Delete(rw web.ResponseWriter, rq *web.Request) {
	utils.SetCORSHeaders(rw)

	if rq.Method == "OPTIONS" {
		return
	}

	r.DeleteRoute.ServeHTTP(rw, rq)
}
