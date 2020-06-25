package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/lincoln/internal/models"
	"github.com/pronuu/lincoln/pkg/vote/services"
	"github.com/pronuu/lincoln/pkg/vote/transports"
)

// MakeCreateEndpoint ...
func MakeCreateEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*transports.CreateRequest)
		vote, err := svc.Create(ctx, req.Vote)
		if err != nil {
			return &transports.CreateResponse{
				Vote: vote,
				Err:  err.Error(),
			}, nil
		}
		return &transports.CreateResponse{
			Vote: vote,
			Err:  "",
		}, nil
	}
}

// Create ...
func (e *Endpoints) Create(ctx context.Context, vote *models.Vote) (*models.Vote, error) {
	req := &transports.CreateRequest{Vote: vote}
	resp, err := e.CreateEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	createResp := resp.(*transports.CreateResponse)
	if createResp.Err != "" {
		return nil, errors.New(createResp.Err)
	}
	return createResp.Vote, nil
}
