package infrastructure

import (
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/rudianto-dev/gotemp-api-gw/cmd/configuration"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
	"github.com/rudianto-dev/gotemp-sdk/pkg/transporter"
)

type Service struct {
	Config     *configuration.ConfigurationSchema
	Logger     *logger.Logger
	Redis      *redis.Client
	UserClient transporter.IHttpTransporter
}

func InitInfrastructure(cf *configuration.ConfigurationSchema) *Service {
	// setup logger
	logger := logger.NewLogger(cf.NewLogrus())
	// setup http client
	httpClient := http.Client{Timeout: 10 * time.Second}
	// setup user client
	userClient := transporter.NewHTTPTransporter(&httpClient, "api-gw", cf.Service.User.URL)

	return &Service{
		Config:     cf,
		Logger:     logger,
		Redis:      cf.NewRedis(),
		UserClient: userClient,
	}
}
