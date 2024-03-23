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
	infrastructure := infrastructure.InitInfrastructure(configuration)

	cmd.AddCommand(&cobra.Command{
		Use:   "api",
		Short: "Run API Gateway",
		Run: func(*cobra.Command, []string) {
			infrastructure.RunAPI()
		},
	})

	if err := cmd.Execute(); err != nil {
		log.Panic(err.Error())
	}
}
