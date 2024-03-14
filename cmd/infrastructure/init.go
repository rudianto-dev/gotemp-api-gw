package infrastructure

import (
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/rudianto-dev/gotemp-api-gw/cmd/configuration"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
	"github.com/rudianto-dev/gotemp-sdk/pkg/transporter"
)

type Infra struct {
	Config     *configuration.ConfigurationSchema
	Logger     *logger.Logger
	Redis      *redis.Client
	UserClient transporter.IHttpTransporter
}

func NewInfrastructure(cf *configuration.ConfigurationSchema) *Infra {
	// setup logger
	logger := logger.NewLogger(cf.NewLogrus())
	// setup http client
	httpClient := http.Client{Timeout: 10 * time.Second}
	// setup user client
	userClient := transporter.NewHTTPTransporter(&httpClient, "api-gw", cf.Service.User.URL)

	return &Infra{
		Config:     cf,
		Logger:     logger,
		Redis:      cf.NewRedis(),
		UserClient: userClient,
	}
}
