package module

import (
	utilDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/util"
	utilHandler "github.com/rudianto-dev/gotemp-api-gw/internal/handler/api/util"
	utilUseCase "github.com/rudianto-dev/gotemp-api-gw/internal/usecase/util"
)

func (m *Module) UtilHandlerAPI() utilDomain.IHandlerAPI {
	utilUseCase := utilUseCase.NewUseCase(m.infra.Logger, m.userService)
	return utilHandler.NewHandler(m.infra.Logger, utilUseCase)
}
