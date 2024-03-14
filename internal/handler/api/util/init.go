package util

import (
	utilDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/util"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type UtilHandler struct {
	logger      logger.ILogger
	utilUseCase utilDomain.IUseCase
}

func NewHandler(logger logger.ILogger, utilUseCase utilDomain.IUseCase) utilDomain.IHandlerAPI {
	return &UtilHandler{
		logger:      logger,
		utilUseCase: utilUseCase,
	}
}
