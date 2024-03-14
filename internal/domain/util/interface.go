package util

import (
	"context"
	"net/http"
)

type IUseCase interface {
	GetServiceHealth(ctx context.Context, req GetServiceHealthRequest) (*GetServiceHealthResponse, error)
}

type IHandlerAPI interface {
	GetHealthStatus(w http.ResponseWriter, r *http.Request)
	GetServiceHealth(w http.ResponseWriter, r *http.Request)
}
