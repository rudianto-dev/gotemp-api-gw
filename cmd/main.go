package main

import (
	"log"

	"github.com/rudianto-dev/gotemp-api-gw/cmd/configuration"
	"github.com/rudianto-dev/gotemp-api-gw/cmd/infrastructure"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{Use: "boilerplate", Short: "Go Boilerplate"}
	configuration := configuration.NewConfiguration()
	infrastructure := infrastructure.NewInfrastructure(configuration)

	cmd.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "Run API Gateway",
		RunE: func(*cobra.Command, []string) error {
			return infrastructure.CreateAPIService()
		},
	})

	if err := cmd.Execute(); err != nil {
		log.Panic(err.Error())
	}
}
