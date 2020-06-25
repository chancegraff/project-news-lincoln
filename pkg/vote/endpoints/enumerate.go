package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/lincoln/internal/models"
	"github.com/pronuu/lincoln/pkg/vote/services"
	"github.com/pronuu/lincoln/pkg/vote/transports"
)

// MakeEnumerateEndpoint ...
func MakeEnumerateEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*transports.EnumerateRequest)
		votes, err := svc.Enumerate(ctx, req.Votes, req.Offset, req.Limit)
		if err != nil {
			return &transports.EnumerateResponse{
				Votes: votes,
				Err:   err.Error(),
			}, nil
		}
		return &transports.EnumerateResponse{
			Votes: votes,
			Err:   "",
		}, nil
	}
}

// Enumerate ...
func (e *Endpoints) Enumerate(ctx context.Context, votes models.VoteArray, offset, limit int) (models.VoteArray, error) {
	req := &transports.EnumerateRequest{Votes: votes, Offset: offset, Limit: limit}
	resp, err := e.EnumerateEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	enumerateResponse := resp.(*transports.EnumerateResponse)
	if enumerateResponse.Err != "" {
		return nil, errors.New(enumerateResponse.Err)
	}
	return enumerateResponse.Votes, nil
}
