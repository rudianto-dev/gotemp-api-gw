package user

import (
	"context"

	userDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/user"
	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *UserUseCase) UpdateProfile(ctx context.Context, req userDomain.UpdateProfileRequest) (*userDomain.UpdateProfileResponse, error) {
	response := &userDomain.UpdateProfileResponse{}
	userUpdateRequest, err := s.userService.UpdateUser(ctx, userContract.UpdateUserRequest{
		ID:   req.ID,
		Name: req.Name,
	})
	if err != nil {
		return response, err
	}
	response.User = userDomain.User{
		ID:   userUpdateRequest.User.ID,
		Name: userUpdateRequest.User.Name,
	}

	return response, nil
}
