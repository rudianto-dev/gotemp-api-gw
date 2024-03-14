package user

import (
	userDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
	userService "github.com/rudianto-dev/gotemp-sdk/service/user"
)

type UserUseCase struct {
	logger      logger.ILogger
	userService userService.IUserService
}

func NewUseCase(logger logger.ILogger, userService userService.IUserService) userDomain.IUseCase {
	return &UserUseCase{
		logger:      logger,
		userService: userService,
	}
}
