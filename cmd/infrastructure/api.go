package infrastructure

import (
	"net/http"

	"github.com/go-chi/chi"
	chim "github.com/go-chi/chi/middleware"
	"github.com/rudianto-dev/gotemp-api-gw/internal/module"
	"github.com/rudianto-dev/gotemp-sdk/pkg/middleware"
)

func (srv *Service) RunAPI() {
	router := srv.SetupAPI()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		srv.Logger.Infof("%s %s", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		srv.Logger.Fatal(err)
	}

	server := &http.Server{Addr: srv.Config.Host.Address, Handler: router}
	srv.Logger.Infof("API GW serving at %s", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			srv.Logger.Error(err)
		}
	}()
	srv.StopGracefully(server)
}

func (srv *Service) SetupAPI() *chi.Mux {
	module := module.NewModule(&module.Service{
		Config:     srv.Config,
		Logger:     srv.Logger,
		Redis:      srv.Redis,
		UserClient: srv.UserClient,
	})
	utilHandlerAPI := module.UtilHandlerAPI()
	userHandlerAPI := module.UserHandlerAPI()

	router := chi.NewRouter()
	router.Use(
		chim.NoCache,
		chim.RedirectSlashes,
		chim.RequestID,
		chim.Recoverer,
		chim.RealIP,
		chim.Heartbeat("/ping"),
		middleware.RequestLogger(srv.Config.NewLogrus()),
	)
	router.Route("/", func(r chi.Router) {
		router.Route("/health", func(r chi.Router) {
			router.Get("/", utilHandlerAPI.GetHealthStatus)
			router.Get("/{service}", utilHandlerAPI.GetServiceHealth)
		})
		router.Route("/v1", func(r chi.Router) {
			router.Route("/user", func(r chi.Router) {
				router.Get("/{id}", userHandlerAPI.GetProfile)
				router.Put("/{id}", userHandlerAPI.UpdateProfile)
				router.Delete("/{id}", userHandlerAPI.DeleteAccount)
			})
		})
	})
	return router
}
