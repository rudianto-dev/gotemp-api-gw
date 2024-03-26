package infrastructure

import (
	"net/http"

	"github.com/rudianto-dev/gotemp-api-gw/gateway"
)

func (srv *Service) RunAPI() {
	router := gateway.NewGateway(srv.Config.ServiceList()).CreateRouting()

	server := &http.Server{Addr: srv.Config.Host.Address, Handler: router}
	srv.Logger.Infof("Gateway service serving at %s", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			srv.Logger.Error(err)
		}
	}()
	srv.StopGracefully(server)
}
