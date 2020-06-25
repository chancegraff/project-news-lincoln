package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/lincoln/internal/models"
	"github.com/pronuu/lincoln/pkg/vote/services"
	"github.com/pronuu/lincoln/pkg/vote/transports"
)

// MakeRetrieveEndpoint ...
func MakeRetrieveEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*transports.RetrieveRequest)
		vote, err := svc.Retrieve(ctx, req.Vote)
		if err != nil {
			return &transports.RetrieveResponse{
				Vote: vote,
				Err:  err.Error(),
			}, nil
		}
		return &transports.RetrieveResponse{
			Vote: vote,
			Err:  "",
		}, nil
	}
}

// Retrieve ...
func (e *Endpoints) Retrieve(ctx context.Context, vote *models.Vote) (*models.Vote, error) {
	req := &transports.RetrieveRequest{Vote: vote}
	resp, err := e.RetrieveEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	retrieveResp := resp.(*transports.RetrieveResponse)
	if retrieveResp.Err != "" {
		return nil, errors.New(retrieveResp.Err)
	}
	return retrieveResp.Vote, nil
}
