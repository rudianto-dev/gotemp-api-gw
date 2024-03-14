package user

import (
	"context"

	userDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/user"
	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *UserUseCase) GetProfile(ctx context.Context, req userDomain.GetProfileRequest) (*userDomain.GetProfileResponse, error) {
	response := userDomain.GetProfileResponse{}
	userDetailRequest, err := s.userService.GetDetail(ctx, userContract.GetDetailUserRequest{
		ID: req.ID,
	})
	if err != nil {
		return &response, err
	}
	response.User = userDomain.User{
		ID:   userDetailRequest.User.ID,
		Name: userDetailRequest.User.Name,
	}

	return &response, nil
}
