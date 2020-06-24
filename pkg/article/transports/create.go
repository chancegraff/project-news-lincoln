package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pronuu/lincoln/internal/models"
)

// CreateRequest ...
type CreateRequest struct {
	Article *models.Article `json:"article"`
}

// CreateResponse ...
type CreateResponse struct {
	Article *models.Article `json:"article"`
	Err     string          `json:"err,omitempty"`
}

// DecodeCreateHTTPRequest ...
func DecodeCreateHTTPRequest(ctx context.Context, rq *http.Request) (interface{}, error) {
	var request CreateRequest
	if err := json.NewDecoder(rq.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeCreateHTTPResponse ...
func EncodeCreateHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
