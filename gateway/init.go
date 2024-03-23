package gateway

import (
	"compress/flate"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Gateway struct {
	services map[string]*url.URL
}

func NewGateway(services map[string]*url.URL) *Gateway {
	g := &Gateway{services: services}
	return g
}

func (g *Gateway) CreateRouting() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(
		middleware.Heartbeat("/ping"),
		middleware.RedirectSlashes,
		middleware.NoCache,
		middleware.Compress(flate.DefaultCompression),
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.DefaultLogger,
	)
	g.Routes(mux)
	return mux
}
