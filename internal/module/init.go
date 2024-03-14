package module

import (
	"github.com/go-redis/redis"
	"github.com/rudianto-dev/gotemp-api-gw/cmd/configuration"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
	"github.com/rudianto-dev/gotemp-sdk/pkg/transporter"
	userService "github.com/rudianto-dev/gotemp-sdk/service/user"
)

type Module struct {
	infra       *Infra
	userService userService.IUserService
}

type Infra struct {
	Config     *configuration.ConfigurationSchema
	Logger     *logger.Logger
	Redis      *redis.Client
	UserClient transporter.IHttpTransporter
}

func NewModule(infra *Infra) *Module {
	userService := userService.NewService(infra.UserClient)
	module := &Module{
		infra:       infra,
		userService: userService,
	}
	return module
}
