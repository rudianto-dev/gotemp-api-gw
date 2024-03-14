package infrastructure

import (
	"net/http"

	"github.com/go-chi/chi"
	chim "github.com/go-chi/chi/middleware"
	"github.com/rudianto-dev/gotemp-api-gw/internal/module"
	"github.com/rudianto-dev/gotemp-sdk/pkg/middleware"
)

func (infra *Infra) CreateAPIService() error {
	r := chi.NewRouter()
	r.Use(
		chim.NoCache,
		chim.RedirectSlashes,
		chim.RequestID,
		chim.Recoverer,
		chim.RealIP,
		chim.Heartbeat("/ping"),
		middleware.RequestLogger(infra.Config.NewLogrus()),
	)

	// load module
	module := module.NewModule(&module.Infra{
		Config:     infra.Config,
		Logger:     infra.Logger,
		Redis:      infra.Redis,
		UserClient: infra.UserClient,
	})
	utilHandlerAPI := module.UtilHandlerAPI()
	userHandlerAPI := module.UserHandlerAPI()

	r.Route("/", func(r chi.Router) {
		r.Route("/health", func(r chi.Router) {
			r.Get("/", utilHandlerAPI.GetHealthStatus)
			r.Get("/{service}", utilHandlerAPI.GetServiceHealth)
		})
		r.Route("/v1", func(r chi.Router) {
			r.Route("/user", func(r chi.Router) {
				r.Get("/{id}", userHandlerAPI.GetProfile)
				r.Put("/{id}", userHandlerAPI.UpdateProfile)
				r.Delete("/{id}", userHandlerAPI.DeleteAccount)
			})
		})
	})

	server := http.Server{
		Addr:    infra.Config.Host.Address,
		Handler: r,
	}
	serverErr := make(chan error, 1)
	go func() {
		infra.Logger.Infof("API GW serving at %s", server.Addr)
		serverErr <- server.ListenAndServe()
	}()

	infra.StopGracefully(&server, serverErr)
	return nil
}
