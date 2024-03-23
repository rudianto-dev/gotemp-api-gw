package infrastructure

import (
	"net/http"

	"github.com/rudianto-dev/gotemp-api-gw/gateway"
)

func (srv *Service) RunAPI() {
	router := gateway.NewGateway(srv.Config.ServiceList()).CreateRouting()

	// router := srv.SetupAPI()
	// walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	// 	srv.Logger.Infof("%s %s", method, route)
	// 	return nil
	// }

	// if err := chi.Walk(router, walkFunc); err != nil {
	// 	srv.Logger.Fatal(err)
	// }

	server := &http.Server{Addr: srv.Config.Host.Address, Handler: router}
	srv.Logger.Infof("API GW serving at %s", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			srv.Logger.Error(err)
		}
	}()
	srv.StopGracefully(server)
}
