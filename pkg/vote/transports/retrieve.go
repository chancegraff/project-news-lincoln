package transports

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pronuu/lincoln/internal/models"
)

// RetrieveRequest ...
type RetrieveRequest struct {
	Vote *models.Vote `json:"vote"`
}

// RetrieveResponse ...
type RetrieveResponse struct {
	Vote *models.Vote `json:"vote"`
	Err  string       `json:"err,omitempty"`
}

// DecodeRetrieveHTTPRequest ...
func DecodeRetrieveHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request RetrieveRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != io.EOF && err != nil {
		return nil, err
	}
	return &request, nil
}

// EncodeRetrieveHTTPResponse ...
func EncodeRetrieveHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
