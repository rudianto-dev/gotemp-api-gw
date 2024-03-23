package infrastructure

import (
	"github.com/go-redis/redis"
	"github.com/rudianto-dev/gotemp-api-gw/cmd/configuration"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type Service struct {
	Config *configuration.ConfigurationSchema
	Logger *logger.Logger
	Redis  *redis.Client
}

func InitInfrastructure(cf *configuration.ConfigurationSchema) *Service {
	// setup logger
	logger := logger.NewLogger(cf.NewLogrus())

	return &Service{
		Config: cf,
		Logger: logger,
		Redis:  cf.NewRedis(),
	}
}
