package user

import (
	"context"

	userDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/user"
	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *UserUseCase) DeleteAccount(ctx context.Context, req userDomain.DeleteAccountRequest) (*userDomain.DeleteAccountResponse, error) {
	response := &userDomain.DeleteAccountResponse{}
	userDeleteRequest, err := s.userService.DeleteUser(ctx, userContract.DeleteUserRequest{
		ID: req.ID,
	})
	if err != nil {
		return response, err
	}
	response.User = userDomain.User{
		ID:   userDeleteRequest.User.ID,
		Name: userDeleteRequest.User.Name,
	}

	return response, nil
}
