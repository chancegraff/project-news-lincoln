package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pronuu/lincoln/internal/models"
)

// EnumerateRequest ...
type EnumerateRequest struct {
	Votes  models.VoteArray `json:"votes,omitempty"`
	Offset int              `json:"offset,omitempty"`
	Limit  int              `json:"limit,omitempty"`
}

// EnumerateResponse ...
type EnumerateResponse struct {
	Votes models.VoteArray `json:"votes"`
	Err   string           `json:"err,omitempty"`
}

// DecodeEnumerateHTTPRequest ...
func DecodeEnumerateHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request EnumerateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeEnumerateHTTPResponse ...
func EncodeEnumerateHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
