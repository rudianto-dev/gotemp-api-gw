package gateway

import (
	"github.com/go-chi/chi"
)

func (g *Gateway) Routes(mux *chi.Mux) {
	mux.Route("/", func(r chi.Router) {
		// r.Use(GuardAuthentication)
		r.HandleFunc("/*", g.Handler)
	})
}
