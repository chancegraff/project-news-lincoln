package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/lincoln/internal/models"
	"github.com/pronuu/lincoln/pkg/article/services"
	"github.com/pronuu/lincoln/pkg/article/transports"
)

// MakeCreateEndpoint ...
func MakeCreateEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*transports.CreateRequest)
		article, err := svc.Create(ctx, req.Article)
		if err != nil {
			return &transports.CreateResponse{
				Article: article,
				Err:     err.Error(),
			}, nil
		}
		return &transports.CreateResponse{
			Article: article,
			Err:     "",
		}, nil
	}
}

// Create ...
func (e *Endpoints) Create(ctx context.Context, article *models.Article) (*models.Article, error) {
	req := &transports.CreateRequest{Article: article}
	resp, err := e.CreateEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	createResp := resp.(*transports.CreateResponse)
	if createResp.Err != "" {
		return nil, errors.New(createResp.Err)
	}
	return createResp.Article, nil
}
