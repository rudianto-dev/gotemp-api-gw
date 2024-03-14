package util

import (
	"context"

	utilDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/util"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UtilUseCase) GetServiceHealth(ctx context.Context, req utilDomain.GetServiceHealthRequest) (*utilDomain.GetServiceHealthResponse, error) {
	response := &utilDomain.GetServiceHealthResponse{}
	switch req.Service {
	case string(utilDomain.USER_SERVICE):
		service, err := s.userService.Health(ctx)
		if err != nil {
			return response, err
		}
		response.Service = "User Service"
		response.Message = service.Message
		response.Status = service.Status
		return response, nil
	default:
		return response, utils.ErrBadRequest
	}
}
