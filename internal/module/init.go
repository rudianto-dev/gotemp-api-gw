package module

import (
	"github.com/go-redis/redis"
	"github.com/rudianto-dev/gotemp-api-gw/cmd/configuration"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
	"github.com/rudianto-dev/gotemp-sdk/pkg/transporter"
	userService "github.com/rudianto-dev/gotemp-sdk/service/user"
)

type Module struct {
	srv         *Service
	userService userService.IUserService
}

type Service struct {
	Config     *configuration.ConfigurationSchema
	Logger     *logger.Logger
	Redis      *redis.Client
	UserClient transporter.IHttpTransporter
}

func NewModule(srv *Service) *Module {
	userService := userService.NewService(srv.UserClient)
	module := &Module{
		srv:         srv,
		userService: userService,
	}
	return module
}
