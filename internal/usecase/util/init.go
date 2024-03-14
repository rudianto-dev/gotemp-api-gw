package util

import (
	utilDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/util"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
	userService "github.com/rudianto-dev/gotemp-sdk/service/user"
)

type UtilUseCase struct {
	logger      logger.ILogger
	userService userService.IUserService
}

func NewUseCase(logger logger.ILogger, userService userService.IUserService) utilDomain.IUseCase {
	return &UtilUseCase{
		logger:      logger,
		userService: userService,
	}
}
