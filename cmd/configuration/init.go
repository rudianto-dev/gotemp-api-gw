package configuration

import (
	"log"

	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

type ConfigurationSchema struct {
	Host     Host      `json:"host"`
	Service  ServiceX  `json:"service"`
	Services []Service `json:"services"`
	Redis    Redis     `json:"redis"`
	Graceful Graceful  `json:"graceful"`
}

type Host struct {
	Address string `json:"address"`
	Debug   bool   `json:"debug"`
}

type Service struct {
	Name string `json:"name"`
	Host string `json:"host"`
}

type ServiceX struct {
	User UserService `json:"user"`
}
type UserService struct {
	URL string `json:"url"`
}

type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	MaxRetry int    `json:"max_retry" mapstructure:"max_retry"`
}

type Graceful struct {
	TimeoutInSecond int64 `json:"timeout_in_second" mapstructure:"timeout_in_second"`
}

func NewConfiguration() *ConfigurationSchema {
	c, err := utils.LoadConfiguration()
	if err != nil {
		log.Panic(err)
	}

	config := ConfigurationSchema{}
	err = c.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	err = c.UnmarshalExact(&config)
	if err != nil {
		log.Panic(err)
	}
	return &config
}
