package endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/pronuu/lincoln/internal/models"
	"github.com/pronuu/lincoln/pkg/article/services"
	"github.com/pronuu/lincoln/pkg/article/transports"
)

// MakeEnumerateEndpoint ...
func MakeEnumerateEndpoint(svc services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*transports.EnumerateRequest)
		articles, err := svc.Enumerate(ctx, req.Articles, req.Offset, req.Limit)
		if err != nil {
			return &transports.EnumerateResponse{
				Articles: articles,
				Err:      err.Error(),
			}, nil
		}
		return &transports.EnumerateResponse{
			Articles: articles,
			Err:      "",
		}, nil
	}
}

// Enumerate ...
func (e *Endpoints) Enumerate(ctx context.Context, articles models.ArticleArray, offset, limit int) (models.ArticleArray, error) {
	req := &transports.EnumerateRequest{Articles: articles, Offset: offset, Limit: limit}
	resp, err := e.EnumerateEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	enumerateResponse := resp.(*transports.EnumerateResponse)
	if enumerateResponse.Err != "" {
		return nil, errors.New(enumerateResponse.Err)
	}
	return enumerateResponse.Articles, nil
}
