package user

import (
	"context"
	"net/http"
)

type IUseCase interface {
	GetProfile(ctx context.Context, req GetProfileRequest) (*GetProfileResponse, error)
	UpdateProfile(ctx context.Context, req UpdateProfileRequest) (*UpdateProfileResponse, error)
	DeleteAccount(ctx context.Context, req DeleteAccountRequest) (*DeleteAccountResponse, error)
}

type IHandlerAPI interface {
	GetProfile(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)
}
