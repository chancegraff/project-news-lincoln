package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pronuu/lincoln/internal/models"
)

// DeleteRequest ...
type DeleteRequest struct {
	Vote *models.Vote `json:"vote"`
}

// DeleteResponse ...
type DeleteResponse struct {
	Vote *models.Vote `json:"vote"`
	Err  string       `json:"err,omitempty"`
}

// DecodeDeleteHTTPRequest ...
func DecodeDeleteHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeDeleteHTTPResponse ...
func EncodeDeleteHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
