package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/lincoln/internal/models"
	"github.com/pronuu/lincoln/pkg/vote/services"
	"github.com/pronuu/lincoln/pkg/vote/transports"
)

// MakeDeleteEndpoint ...
func MakeDeleteEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*transports.DeleteRequest)
		vote, err := svc.Delete(ctx, req.Vote)
		if err != nil {
			return &transports.DeleteResponse{
				Vote: vote,
				Err:  err.Error(),
			}, nil
		}
		return &transports.DeleteResponse{
			Vote: vote,
			Err:  "",
		}, nil
	}
}

// Delete ...
func (e *Endpoints) Delete(ctx context.Context, vote *models.Vote) (*models.Vote, error) {
	req := &transports.DeleteRequest{Vote: vote}
	resp, err := e.DeleteEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	deleteResp := resp.(*transports.DeleteResponse)
	if deleteResp.Err != "" {
		return nil, errors.New(deleteResp.Err)
	}
	return deleteResp.Vote, nil
}
