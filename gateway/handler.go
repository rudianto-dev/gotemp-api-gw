package gateway

import (
	"net/http"
)

func (g *Gateway) Handler(w http.ResponseWriter, r *http.Request) {
	// passing request via proxy
	g.Proxy(w, r)
}
