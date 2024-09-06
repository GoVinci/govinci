package routing

import (
	"context"
	"net/http"
	"strings"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
	Params  map[string]int
}

type contextKey string

const ParamsKey contextKey = "params"

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Handle(method, path string, handler http.HandlerFunc) {
	params := make(map[string]int)
	for i, segment := range strings.Split(path, "/") {
		if strings.HasPrefix(segment, ":") {
			params[segment[1:]] = i
		}
	}
	r.routes = append(r.routes, Route{method, path, handler, params})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if req.Method == route.Method {
			routeSegments, requestSegments := strings.Split(route.Path, "/"), strings.Split(req.URL.Path, "/")
			if len(routeSegments) == len(requestSegments) {
				params, match := make(map[string]string), true
				for i, segment := range routeSegments {
					if strings.HasPrefix(segment, ":") {
						params[segment[1:]] = requestSegments[i]
					} else if segment != requestSegments[i] {
						match = false
						break
					}
				}
				if match {
					req = req.WithContext(context.WithValue(req.Context(), ParamsKey, params))
					route.Handler(w, req)
					return
				}
			}
		}
	}
	http.NotFound(w, req)
}
