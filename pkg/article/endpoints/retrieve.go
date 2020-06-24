package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/lincoln/internal/models"
	"github.com/pronuu/lincoln/pkg/article/services"
	"github.com/pronuu/lincoln/pkg/article/transports"
)

// MakeRetrieveEndpoint ...
func MakeRetrieveEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*transports.RetrieveRequest)
		article, err := svc.Retrieve(ctx, req.Article)
		if err != nil {
			return &transports.RetrieveResponse{
				Article: article,
				Err:     err.Error(),
			}, nil
		}
		return &transports.RetrieveResponse{
			Article: article,
			Err:     "",
		}, nil
	}
}

// Retrieve ...
func (e *Endpoints) Retrieve(ctx context.Context, article *models.Article) (*models.Article, error) {
	req := &transports.RetrieveRequest{Article: article}
	resp, err := e.RetrieveEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	retrieveResp := resp.(*transports.RetrieveResponse)
	if retrieveResp.Err != "" {
		return nil, errors.New(retrieveResp.Err)
	}
	return retrieveResp.Article, nil
}
