package module

import (
	userDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/user"
	userHandler "github.com/rudianto-dev/gotemp-api-gw/internal/handler/api/user"
	userUseCase "github.com/rudianto-dev/gotemp-api-gw/internal/usecase/user"
)

func (m *Module) UserHandlerAPI() userDomain.IHandlerAPI {
	userUseCase := userUseCase.NewUseCase(m.infra.Logger, m.userService)
	return userHandler.NewHandler(m.infra.Logger, userUseCase)
}
