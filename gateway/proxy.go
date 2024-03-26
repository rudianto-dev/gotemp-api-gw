package gateway

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

func (g *Gateway) Proxy(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	serviceName := pathSegments[1]

	r.URL.Path = strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/%s", serviceName))
	serviceURL, ok := g.services[serviceName]
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	proxy := httputil.ReverseProxy{Director: func(r *http.Request) {
		r.URL.Scheme = serviceURL.Scheme
		r.URL.Host = serviceURL.Host
		r.URL.Path = "/external" + serviceURL.Path + r.URL.Path
		r.Host = serviceURL.Host
	}}

	proxy.ServeHTTP(w, r)
}
